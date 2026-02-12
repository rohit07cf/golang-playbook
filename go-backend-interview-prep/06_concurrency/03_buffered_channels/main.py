"""Buffered channels -- Python equivalent of the Go example.

queue.Queue(maxsize=N) is the Python equivalent of make(chan T, N).
maxsize=0 means unlimited (different from Go where 0 = unbuffered).
"""

import queue
import threading


def main() -> None:
    # --- Example 1: buffered put without blocking ---
    print("=== Buffered queue (maxsize 3) ===")
    q: queue.Queue[str] = queue.Queue(maxsize=3)

    q.put_nowait("a")
    q.put_nowait("b")
    q.put_nowait("c")

    print("qsize:", q.qsize(), "maxsize:", q.maxsize)

    print(q.get())
    print(q.get())
    print(q.get())

    # --- Example 2: producer faster than consumer ---
    print("\n=== Producer/Consumer with buffer ===")
    SENTINEL = None
    jobs: queue.Queue[int | None] = queue.Queue(maxsize=5)

    def producer():
        for i in range(1, 6):
            print(f"  producing {i}")
            jobs.put(i)
        jobs.put(SENTINEL)

    threading.Thread(target=producer).start()

    while True:
        j = jobs.get()
        if j is SENTINEL:
            break
        print(f"  consumed {j}")

    # --- Example 3: buffer size 1 as a signal ---
    print("\n=== Buffer size 1 (signal) ===")
    done: queue.Queue[bool] = queue.Queue(maxsize=1)

    def do_work():
        print("  work complete")
        done.put(True)

    threading.Thread(target=do_work).start()
    done.get()
    print("  received done signal")

    # --- Example 4: semaphore pattern ---
    print("\n=== Semaphore (threading.Semaphore) ===")
    sem = threading.Semaphore(2)  # max 2 concurrent
    threads = []

    for i in range(1, 6):
        def worker(wid):
            sem.acquire()
            try:
                print(f"  worker {wid} running (max 2 at a time)")
            finally:
                sem.release()

        t = threading.Thread(target=worker, args=(i,))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()
    print("  all workers done")


if __name__ == "__main__":
    main()
