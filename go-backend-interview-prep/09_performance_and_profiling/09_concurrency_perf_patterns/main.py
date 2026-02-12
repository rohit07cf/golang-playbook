"""Concurrency performance patterns -- Python equivalent of the Go example.

Compares unbounded threading vs ThreadPoolExecutor.
"""

import os
import time
import threading
from concurrent.futures import ThreadPoolExecutor


def simulate_work(task_id: int) -> int:
    """Small CPU-ish work unit."""
    total = 0
    for i in range(1000):
        total += i * task_id
    return total


def unbounded_threads(tasks: int) -> None:
    """Spawn one thread per task (dangerous at scale)."""
    threads = []
    for i in range(tasks):
        t = threading.Thread(target=simulate_work, args=(i,))
        t.start()
        threads.append(t)
    for t in threads:
        t.join()


def thread_pool(tasks: int, workers: int) -> None:
    """Fixed thread pool with bounded concurrency."""
    with ThreadPoolExecutor(max_workers=workers) as pool:
        list(pool.map(simulate_work, range(tasks)))


def main() -> None:
    num_cpu = os.cpu_count() or 4
    print("=== Concurrency Performance Patterns ===\n")
    print(f"  CPUs available: {num_cpu}\n")

    # Smaller task counts for Python (threads are heavier)
    task_counts = [100, 1_000, 5_000]

    print("--- Unbounded threads vs ThreadPoolExecutor ---")
    print(f"{'Tasks':<10}  {'Unbounded':>15}  {'Pool':>15}  Workers")
    print("-" * 55)

    for tasks in task_counts:
        start = time.perf_counter()
        unbounded_threads(tasks)
        t1 = time.perf_counter() - start

        start = time.perf_counter()
        thread_pool(tasks, num_cpu)
        t2 = time.perf_counter() - start

        print(f"{tasks:<10}  {t1:>14.4f}s  {t2:>14.4f}s  {num_cpu}")
    print()

    # Queue-based backpressure example
    import queue

    print("--- Queue backpressure demo ---")
    q: queue.Queue = queue.Queue(maxsize=10)  # buffer of 10
    produced = 0
    consumed = 0

    def producer():
        nonlocal produced
        for i in range(50):
            q.put(i)  # blocks if queue is full (backpressure!)
            produced += 1

    def consumer():
        nonlocal consumed
        while True:
            try:
                _ = q.get(timeout=0.1)
                consumed += 1
                time.sleep(0.01)  # simulate slow consumer
                q.task_done()
            except queue.Empty:
                break

    # 1 producer, 2 consumers
    threads = [
        threading.Thread(target=producer),
        threading.Thread(target=consumer),
        threading.Thread(target=consumer),
    ]
    for t in threads:
        t.start()
    for t in threads:
        t.join()

    print(f"  Produced: {produced}, Consumed: {consumed}")
    print(f"  Queue maxsize=10 provided backpressure")
    print()

    print("Key: ThreadPoolExecutor is Python's worker pool.")
    print("     queue.Queue(maxsize=N) provides backpressure.")
    print("     Python threads are ~8MB stack vs Go's ~2KB goroutine stack.")


if __name__ == "__main__":
    main()
