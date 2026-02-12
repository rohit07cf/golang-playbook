"""Fan-out, fan-in pipeline -- Python equivalent of the Go example.

Uses asyncio queues to build a pipeline with fan-out/fan-in.
"""

import asyncio
import time

SENTINEL = None


async def generate(lo: int, hi: int) -> asyncio.Queue:
    """Emit integers [lo, hi] into a queue."""
    out: asyncio.Queue = asyncio.Queue()
    for i in range(lo, hi + 1):
        await out.put(i)
    await out.put(SENTINEL)
    return out


async def square_worker(wid: int, in_q: asyncio.Queue, out_q: asyncio.Queue) -> None:
    """Read ints, square them, send to output."""
    while True:
        v = await in_q.get()
        if v is SENTINEL:
            await in_q.put(SENTINEL)  # re-post so other workers see it
            break
        await asyncio.sleep(0.02)  # simulate work
        await out_q.put(v * v)


async def filter_even(in_q: asyncio.Queue) -> asyncio.Queue:
    """Keep only even numbers."""
    out: asyncio.Queue = asyncio.Queue()
    while True:
        v = await in_q.get()
        if v is SENTINEL:
            await out.put(SENTINEL)
            break
        if v % 2 == 0:
            await out.put(v)
    return out


async def drain(q: asyncio.Queue) -> list[int]:
    """Collect all values from queue until sentinel."""
    results = []
    while True:
        v = await q.get()
        if v is SENTINEL:
            break
        results.append(v)
    return results


async def main() -> None:
    # --- Simple pipeline: generate -> square -> filter_even ---
    print("=== Simple pipeline ===")
    nums_q = await generate(1, 10)
    squared_q: asyncio.Queue = asyncio.Queue()

    # Single square worker
    await square_worker(0, nums_q, squared_q)
    await squared_q.put(SENTINEL)

    evens_q = await filter_even(squared_q)
    for v in await drain(evens_q):
        print(f"  {v}")

    # --- Fan-out (3 workers) -> Fan-in ---
    print("\n=== Fan-out (3 workers) -> Fan-in ===")
    start = time.perf_counter()

    nums_q2 = await generate(1, 12)
    merged_q: asyncio.Queue = asyncio.Queue()

    # Fan-out: 3 workers read from same queue
    workers = [
        asyncio.create_task(square_worker(w, nums_q2, merged_q))
        for w in range(3)
    ]
    await asyncio.gather(*workers)
    await merged_q.put(SENTINEL)

    # Fan-in: drain merged queue
    for v in await drain(merged_q):
        print(f"  {v}")

    elapsed = time.perf_counter() - start
    print(f"  completed in {elapsed*1000:.0f}ms")


if __name__ == "__main__":
    asyncio.run(main())
