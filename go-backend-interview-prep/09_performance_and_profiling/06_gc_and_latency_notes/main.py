"""GC and latency notes -- Python equivalent of the Go example.

Python uses reference counting + cyclic garbage collector.
"""

import gc
import sys
import time


class Node:
    """Simple object for allocation demos."""
    __slots__ = ("value", "next")

    def __init__(self, value: int):
        self.value = value
        self.next = None


def main() -> None:
    print("=== GC and Latency Notes (Python) ===\n")

    # Show GC generation stats
    gc.collect()
    stats = gc.get_stats()
    for i, gen in enumerate(stats):
        print(f"  Generation {i}: {gen}")
    print()

    # Allocate many objects
    print("  Allocating 500K objects...")
    start = time.perf_counter()
    objects = []
    for i in range(500_000):
        objects.append(Node(i))
    alloc_time = time.perf_counter() - start
    print(f"  Alloc time: {alloc_time:.4f}s")
    print()

    # Create reference cycles (to trigger cyclic GC)
    print("  Creating 10K reference cycles...")
    cycles = []
    for i in range(10_000):
        a = Node(i)
        b = Node(i + 1)
        a.next = b
        b.next = a  # cycle!
        cycles.append(a)
    print()

    # Release references
    del objects
    del cycles

    # Force cyclic GC and time it
    gc.collect()  # clear any pending
    start = time.perf_counter()
    collected = gc.collect()
    gc_time = time.perf_counter() - start
    print(f"  gc.collect(): freed {collected} cyclic objects in {gc_time*1000:.2f}ms")
    print()

    # Show reference counting
    obj = Node(42)
    print(f"  Reference counting demo:")
    print(f"    refcount(obj):          {sys.getrefcount(obj)}")  # 2 (obj + arg)
    ref = obj
    print(f"    After ref=obj:          {sys.getrefcount(obj)}")  # 3
    del ref
    print(f"    After del ref:          {sys.getrefcount(obj)}")  # 2
    print(f"    When refcount -> 0: immediately freed (no GC needed)")
    print()

    print("--- Python GC model ---")
    print("  1. Reference counting (primary)")
    print("     - Every object has a refcount")
    print("     - Freed immediately when refcount reaches 0")
    print("     - No pause -- deterministic lifetime")
    print()
    print("  2. Cyclic GC (supplementary)")
    print("     - Handles reference cycles (A->B->A)")
    print("     - Runs periodically based on allocation/deallocation thresholds")
    print("     - 3 generations: gen0 (young), gen1, gen2 (old)")
    print("     - Can cause small pauses")
    print()
    print("  vs Go GC:")
    print("     - Go: concurrent mark-and-sweep, sub-ms STW pauses")
    print("     - Python: refcount (instant) + cyclic GC (periodic)")
    print("     - Go: tunable via GOGC; Python: tunable via gc.set_threshold()")


if __name__ == "__main__":
    main()
