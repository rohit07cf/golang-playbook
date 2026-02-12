"""Channel directions -- Python equivalent of the Go example.

Python has no directional queues. We use naming conventions
(out_q for send, in_q for receive) and documentation to convey intent.
"""

import queue
import threading

SENTINEL = None


def produce(out_q: queue.Queue, count: int) -> None:
    """Send-only by convention: only puts, then signals done."""
    for i in range(1, count + 1):
        out_q.put(i * 10)
    out_q.put(SENTINEL)


def consume(in_q: queue.Queue) -> None:
    """Receive-only by convention: only gets until sentinel."""
    while True:
        val = in_q.get()
        if val is SENTINEL:
            break
        print(f"  received: {val}")


def transform(in_q: queue.Queue, out_q: queue.Queue) -> None:
    """Reads from in_q, doubles, sends to out_q."""
    while True:
        val = in_q.get()
        if val is SENTINEL:
            out_q.put(SENTINEL)
            break
        out_q.put(val * 2)


def generate_nums(lo: int, hi: int) -> queue.Queue:
    """Generator pattern: returns a queue that emits [lo, hi]."""
    q: queue.Queue[int | None] = queue.Queue()

    def _fill():
        for i in range(lo, hi + 1):
            q.put(i)
        q.put(SENTINEL)

    threading.Thread(target=_fill).start()
    return q


def main() -> None:
    # --- Example 1: producer + consumer ---
    print("=== Producer -> Consumer ===")
    ch: queue.Queue = queue.Queue()
    t = threading.Thread(target=produce, args=(ch, 5))
    t.start()
    consume(ch)
    t.join()

    # --- Example 2: pipeline with transform ---
    print("\n=== Producer -> Transform -> Consumer ===")
    raw: queue.Queue = queue.Queue()
    doubled: queue.Queue = queue.Queue()

    t1 = threading.Thread(target=produce, args=(raw, 4))
    t2 = threading.Thread(target=transform, args=(raw, doubled))
    t1.start()
    t2.start()
    consume(doubled)
    t1.join()
    t2.join()

    # --- Example 3: generator pattern ---
    print("\n=== Generator pattern ===")
    nums = generate_nums(1, 5)
    while True:
        n = nums.get()
        if n is SENTINEL:
            break
        print(f"  generated: {n}")


if __name__ == "__main__":
    main()
