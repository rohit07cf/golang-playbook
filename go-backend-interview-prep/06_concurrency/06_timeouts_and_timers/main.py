"""Timeouts and timers -- Python equivalent of the Go example.

Python uses asyncio.wait_for for timeouts and asyncio.sleep for timers.
There is no channel-based timer; async sleep is the closest analogy.
"""

import asyncio


async def main() -> None:
    # --- Example 1: timeout with asyncio.wait_for ---
    print("=== asyncio.wait_for timeout ===")

    async def slow_work():
        await asyncio.sleep(0.2)
        return "result"

    try:
        v = await asyncio.wait_for(slow_work(), timeout=0.1)
        print(f"  received: {v}")
    except asyncio.TimeoutError:
        print("  timed out after 100ms")

    # --- Example 2: one-shot timer (asyncio.sleep) ---
    print("\n=== One-shot timer ===")
    await asyncio.sleep(0.1)
    print("  timer fired after 100ms")

    await asyncio.sleep(0.05)
    print("  timer fired again after 50ms (reset)")

    # --- Example 3: periodic ticker (loop + sleep) ---
    print("\n=== Periodic ticker ===")
    count = 0
    interval = 0.08
    deadline = asyncio.get_event_loop().time() + 0.35

    while True:
        await asyncio.sleep(interval)
        now = asyncio.get_event_loop().time()
        if now >= deadline:
            print(f"  stopped after {count} ticks")
            break
        count += 1
        print(f"  tick {count}")


if __name__ == "__main__":
    asyncio.run(main())
