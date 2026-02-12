"""Worker pools -- Python equivalent of the Go example.

Two approaches shown:
1. ThreadPoolExecutor (high-level, preferred)
2. Manual threads + queue (matches Go pattern exactly)
"""

import queue
import threading
import time
from concurrent.futures import ThreadPoolExecutor


def process_job(job_id: int, value: int) -> tuple[int, int]:
    """Simulate work: square the value."""
    time.sleep(0.05)
    return job_id, value * value


def main() -> None:
    num_jobs = 9

    # --- Approach 1: ThreadPoolExecutor (idiomatic Python) ---
    print("=== ThreadPoolExecutor ===")
    with ThreadPoolExecutor(max_workers=3) as pool:
        futures = [pool.submit(process_job, j, j) for j in range(1, num_jobs + 1)]

    print("\n=== Results (executor) ===")
    for f in futures:
        job_id, output = f.result()
        print(f"  job {job_id} -> {output}")

    # --- Approach 2: manual threads + queue (matches Go) ---
    print("\n=== Manual threads + queue ===")
    SENTINEL = None
    jobs: queue.Queue = queue.Queue()
    results: queue.Queue = queue.Queue()

    def worker(wid: int):
        while True:
            item = jobs.get()
            if item is SENTINEL:
                break
            jid, val = item
            print(f"  worker {wid}: processing job {jid}")
            time.sleep(0.05)
            results.put((jid, val * val))

    # Start workers
    num_workers = 3
    threads = []
    for w in range(1, num_workers + 1):
        t = threading.Thread(target=worker, args=(w,))
        threads.append(t)
        t.start()

    # Send jobs
    for j in range(1, num_jobs + 1):
        jobs.put((j, j))

    # Send sentinel to each worker
    for _ in range(num_workers):
        jobs.put(SENTINEL)

    for t in threads:
        t.join()

    # Collect results
    print("\n=== Results (manual) ===")
    while not results.empty():
        jid, output = results.get()
        print(f"  job {jid} -> {output}")
    print("all done")


if __name__ == "__main__":
    main()
