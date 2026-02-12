"""Context cancellation -- Python equivalent of the Go example.

Python has no built-in context.Context. We use:
- asyncio task cancellation (CancelledError) for async code
- threading.Event as a stop flag for threaded code
"""

import asyncio
import threading
import time


# --- Async approach: asyncio task cancellation ---

async def slow_operation_async(name: str) -> None:
    """Respects cancellation via CancelledError."""
    try:
        for i in range(1, 11):
            print(f"  {name}: step {i}")
            await asyncio.sleep(0.05)
        print(f"  {name}: completed all steps")
    except asyncio.CancelledError:
        print(f"  {name}: cancelled at some step")
        raise


# --- Thread approach: Event flag ---

def slow_operation_thread(name: str, stop: threading.Event) -> None:
    """Checks stop event each iteration."""
    for i in range(1, 11):
        if stop.is_set():
            print(f"  {name}: cancelled at step {i}")
            return
        print(f"  {name}: step {i}")
        time.sleep(0.05)
    print(f"  {name}: completed all steps")


async def main() -> None:
    # --- Example 1: asyncio task cancellation (WithCancel) ---
    print("=== asyncio cancel (like WithCancel) ===")
    task = asyncio.create_task(slow_operation_async("job1"))
    await asyncio.sleep(0.15)
    task.cancel()
    try:
        await task
    except asyncio.CancelledError:
        pass

    # --- Example 2: asyncio wait_for (like WithTimeout) ---
    print("\n=== asyncio wait_for (like WithTimeout 200ms) ===")
    try:
        await asyncio.wait_for(slow_operation_async("job2"), timeout=0.2)
    except asyncio.TimeoutError:
        print("  job2: timed out")

    # --- Example 3: threading.Event cancels children ---
    print("\n=== threading.Event cancels children ===")
    stop = threading.Event()

    t1 = threading.Thread(target=slow_operation_thread, args=("child1", stop))
    t2 = threading.Thread(target=slow_operation_thread, args=("child2", stop))
    t1.start()
    t2.start()

    time.sleep(0.12)
    stop.set()  # cancels both children

    t1.join()
    t2.join()
    print("  both children cancelled")


if __name__ == "__main__":
    asyncio.run(main())
