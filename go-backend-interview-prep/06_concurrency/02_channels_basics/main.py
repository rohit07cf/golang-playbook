"""Channels basics -- Python equivalent of the Go example.

Python has no built-in channels. queue.Queue is the closest match.
It has no close() or range -- use a sentinel value (None) to signal done.
"""

import queue
import threading

SENTINEL = None  # signals "channel closed"


def main() -> None:
    # --- Example 1: basic send and receive ---
    print("=== Basic send/receive ===")
    q: queue.Queue[str] = queue.Queue()

    threading.Thread(target=lambda: q.put("hello from thread")).start()

    msg = q.get()  # blocks until item available
    print(msg)

    # --- Example 2: close + range (sentinel pattern) ---
    print("\n=== Close + range (sentinel) ===")
    nums: queue.Queue[int | None] = queue.Queue()

    def producer():
        for i in range(1, 6):
            nums.put(i)
        nums.put(SENTINEL)  # signal: no more values

    threading.Thread(target=producer).start()

    while True:
        n = nums.get()
        if n is SENTINEL:
            break
        print(f"  received: {n}")

    # --- Example 3: comma-ok pattern (not native in Python) ---
    print("\n=== Get with timeout (closest to comma-ok) ===")
    q2: queue.Queue[int] = queue.Queue()
    q2.put(42)

    try:
        v = q2.get_nowait()
        print(f"  v={v}, ok=True (value available)")
    except queue.Empty:
        print("  v=0, ok=False (queue empty)")

    try:
        v = q2.get_nowait()
        print(f"  v={v}, ok=True")
    except queue.Empty:
        print("  v=0, ok=False (queue empty)")

    # --- Example 4: queue as function return ---
    print("\n=== Queue as return value ===")
    result = compute(10, 20)
    print("  10 + 20 =", result.get())


def compute(a: int, b: int) -> queue.Queue:
    q: queue.Queue[int] = queue.Queue()
    threading.Thread(target=lambda: q.put(a + b)).start()
    return q


if __name__ == "__main__":
    main()
