"""WaitGroups -- Python equivalent of the Go example.

Python equivalent: thread.join() or concurrent.futures.
Thread.join() blocks until the thread finishes, like wg.Wait().
"""

import threading
import queue
import time
from concurrent.futures import ThreadPoolExecutor, wait


def worker(wid: int) -> None:
    print(f"  worker {wid}: started")
    time.sleep(0.05)
    print(f"  worker {wid}: done")


def main() -> None:
    # --- Example 1: basic join (WaitGroup equivalent) ---
    print("=== Basic join ===")
    threads = []
    for i in range(1, 6):
        t = threading.Thread(target=worker, args=(i,))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()  # blocks until thread finishes
    print("all workers finished")

    # --- Example 2: threads + queue for results ---
    print("\n=== Threads + queue for results ===")
    results: queue.Queue[int] = queue.Queue()
    threads = []

    for i in range(1, 6):
        def task(n):
            results.put(n * n)
        t = threading.Thread(target=task, args=(i,))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()

    while not results.empty():
        print(f"  result: {results.get()}")

    # --- Example 3: ThreadPoolExecutor (higher-level) ---
    print("\n=== ThreadPoolExecutor ===")
    with ThreadPoolExecutor(max_workers=3) as executor:
        futures = [
            executor.submit(lambda tid: print(f"  task {tid} complete"), i)
            for i in range(3)
        ]
        wait(futures)  # blocks until all complete
    print("all tasks done")


if __name__ == "__main__":
    main()
