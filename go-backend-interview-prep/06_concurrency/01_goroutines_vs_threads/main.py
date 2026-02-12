"""Goroutines vs threads -- Python equivalent of the Go example.

Python threads are real OS threads but the GIL (Global Interpreter Lock)
prevents true CPU parallelism. For I/O-bound work threads are fine;
for CPU-bound work use multiprocessing.
"""

import threading
import time


def work(thread_id: int) -> None:
    print(f"thread {thread_id}: started")
    time.sleep(0.1)  # simulate work
    print(f"thread {thread_id}: done")


def main() -> None:
    # --- Example 1: launching threads ---
    print("=== Launching threads ===")
    threads = []
    for i in range(1, 6):
        t = threading.Thread(target=work, args=(i,))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()  # equivalent of wg.Wait()
    print("all threads finished")

    # --- Example 2: anonymous thread (lambda) ---
    print("\n=== Anonymous thread ===")
    t = threading.Thread(target=lambda: print("hello from anonymous thread"))
    t.start()
    t.join()

    # --- Example 3: thread cost demonstration ---
    # NOTE: 10,000 OS threads is expensive (each ~1 MB stack).
    # Python can do it but it's much slower than Go goroutines.
    print("\n=== Spawning 10,000 threads ===")
    start = time.perf_counter()
    threads = []
    for i in range(10_000):
        t = threading.Thread(target=lambda n: n * n, args=(i,))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()
    elapsed = time.perf_counter() - start
    print(f"10,000 threads completed in {elapsed:.3f}s")
    print("(much slower than Go goroutines due to OS thread overhead)")


if __name__ == "__main__":
    main()
