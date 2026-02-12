"""Job queue worker -- Python equivalent using queue.Queue + ThreadPoolExecutor."""

import queue
import random
import threading
import time
from concurrent.futures import ThreadPoolExecutor


# --- Job ---

class Job:
    def __init__(self, job_id, payload):
        self.id = job_id
        self.payload = payload


# --- Worker ---

def worker(worker_id, job_queue, shutdown_event):
    """Process jobs from the queue until shutdown."""
    while True:
        try:
            job = job_queue.get(timeout=0.5)
        except queue.Empty:
            if shutdown_event.is_set():
                break
            continue

        duration = random.uniform(0.1, 0.4)
        print(f"[worker {worker_id}] processing job {job.id}: {job.payload}")
        time.sleep(duration)
        print(f"[worker {worker_id}] completed  job {job.id} (took {duration:.2f}s)")
        job_queue.task_done()

    print(f"[worker {worker_id}] shutting down")


# --- Demo ---

def main():
    num_jobs = 12
    num_workers = 3
    queue_size = 4  # bounded queue -- backpressure when full

    job_queue = queue.Queue(maxsize=queue_size)
    shutdown_event = threading.Event()

    print(f"starting {num_workers} workers (queue capacity: {queue_size})\n")

    # Start worker threads
    threads = []
    for i in range(1, num_workers + 1):
        t = threading.Thread(target=worker, args=(i, job_queue, shutdown_event))
        t.start()
        threads.append(t)

    # Producer: enqueue jobs
    print(f"producing {num_jobs} jobs...\n")
    for i in range(1, num_jobs + 1):
        job = Job(i, f"send-email-{i}")
        print(f"[producer] queuing job {i} (queue size: ~{job_queue.qsize()}/{queue_size})")
        job_queue.put(job)  # blocks if queue is full (backpressure)

    # Wait for all jobs to be processed
    job_queue.join()
    print("\n[producer] all jobs processed")

    # Signal shutdown
    shutdown_event.set()
    for t in threads:
        t.join()

    print("\nall workers done -- exiting")


if __name__ == "__main__":
    main()
