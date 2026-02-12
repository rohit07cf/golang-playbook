"""Rate limiting -- Python equivalent of the Go example.

Python has no channel-based timers. We use:
- time.sleep for fixed interval rate limiting
- queue.Queue as a token bucket
"""

import queue
import threading
import time


def main() -> None:
    # --- Example 1: fixed interval rate limiter ---
    print("=== Fixed interval (steady rate) ===")
    requests = [1, 2, 3, 4, 5]
    interval = 0.1  # 10 req/sec

    start = time.perf_counter()
    for req in requests:
        time.sleep(interval)
        elapsed = time.perf_counter() - start
        print(f"  request {req} at {elapsed*1000:.0f}ms")

    # --- Example 2: token bucket (allows bursts) ---
    print("\n=== Token bucket (burst-friendly) ===")
    bucket: queue.Queue[None] = queue.Queue(maxsize=3)

    # Fill bucket initially
    for _ in range(3):
        bucket.put(None)

    # Refill one token every 200ms in background
    def refiller():
        while True:
            time.sleep(0.2)
            try:
                bucket.put_nowait(None)
            except queue.Full:
                pass  # bucket full, discard

    t = threading.Thread(target=refiller, daemon=True)
    t.start()

    # Process 8 requests
    start2 = time.perf_counter()
    for i in range(1, 9):
        bucket.get()  # consume a token (blocks if empty)
        elapsed = time.perf_counter() - start2
        print(f"  request {i} at {elapsed*1000:.0f}ms")


if __name__ == "__main__":
    main()
