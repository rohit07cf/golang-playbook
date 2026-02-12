"""Select -- Python equivalent of the Go example.

Python has no native select for queues/channels. We use:
- asyncio for multiplexing async tasks (closest to Go select)
- threading.Event for done/cancellation signaling
"""

import asyncio


async def main() -> None:
    # --- Example 1: wait on two tasks (like select on two channels) ---
    print("=== Select on two tasks ===")

    async def task1():
        await asyncio.sleep(0.05)
        return "one"

    async def task2():
        await asyncio.sleep(0.10)
        return "two"

    t1 = asyncio.create_task(task1())
    t2 = asyncio.create_task(task2())

    # Receive both in completion order
    for coro in asyncio.as_completed([t1, t2]):
        result = await coro
        print(f"  received: {result}")

    # --- Example 2: timeout ---
    print("\n=== Timeout ===")

    async def slow_work():
        await asyncio.sleep(0.5)
        return "finally"

    try:
        result = await asyncio.wait_for(slow_work(), timeout=0.1)
        print(f"  received: {result}")
    except asyncio.TimeoutError:
        print("  timed out after 100ms")

    # --- Example 3: non-blocking check ---
    print("\n=== Non-blocking (queue check) ===")
    q: asyncio.Queue[int] = asyncio.Queue(maxsize=1)

    # Non-blocking get
    try:
        v = q.get_nowait()
        print(f"  received: {v}")
    except asyncio.QueueEmpty:
        print("  nothing ready (non-blocking)")

    # Non-blocking put
    await q.put(42)
    try:
        q.put_nowait(99)
        print("  sent 99")
    except asyncio.QueueFull:
        print("  queue full, skipped send")
    print(f"  value in queue: {await q.get()}")

    # --- Example 4: done signal (cancellation) ---
    print("\n=== Done signal (task cancellation) ===")
    results: asyncio.Queue[int] = asyncio.Queue()

    async def producer():
        i = 1
        try:
            while True:
                await results.put(i)
                i += 1
                await asyncio.sleep(0)  # yield control
        except asyncio.CancelledError:
            print("  producer: received done signal")

    task = asyncio.create_task(producer())

    for _ in range(3):
        val = await results.get()
        print(f"  consumed: {val}")

    task.cancel()
    try:
        await task
    except asyncio.CancelledError:
        pass


if __name__ == "__main__":
    asyncio.run(main())
