"""Escape analysis and pointers -- Python equivalent of the Go example.

Python has NO escape analysis. All objects live on the heap.
This file explains the difference and demonstrates Python's memory model.
"""

import sys
import time


class Point:
    __slots__ = ("x", "y")  # slightly more memory-efficient than __dict__

    def __init__(self, x: float, y: float):
        self.x = x
        self.y = y


def new_point() -> Point:
    """In Python, this ALWAYS allocates on the heap.
    Go equivalent returning by value would stay on the stack."""
    return Point(1.0, 2.0)


def local_only() -> float:
    """Even local variables are heap-allocated in Python.
    Go equivalent keeps Point on the stack (no escape)."""
    p = Point(3.0, 4.0)
    return p.x + p.y


def main() -> None:
    print("=== Escape Analysis -- Python Perspective ===\n")

    n = 1_000_000

    # Time object creation
    start = time.perf_counter()
    for _ in range(n):
        p = new_point()
        _ = p.x
    create_time = time.perf_counter() - start

    start = time.perf_counter()
    for _ in range(n):
        v = local_only()
        _ = v
    local_time = time.perf_counter() - start

    print(f"  new_point() (heap, always): {create_time:.4f}s  ({n} ops)")
    print(f"  local_only() (also heap):   {local_time:.4f}s  ({n} ops)")
    print()

    # Show reference counting
    p = Point(1.0, 2.0)
    print(f"  Reference count of p: {sys.getrefcount(p)}")  # 2 (p + getrefcount arg)
    ref = p
    print(f"  After ref = p:        {sys.getrefcount(p)}")  # 3
    del ref
    print(f"  After del ref:        {sys.getrefcount(p)}")  # 2
    print()

    print("--- Python vs Go memory model ---")
    print("  Python: ALL objects live on the heap")
    print("    - Reference counting frees most objects immediately")
    print("    - Cyclic GC handles reference cycles")
    print("    - No stack vs heap choice -- no escape analysis")
    print()
    print("  Go: compiler decides stack vs heap")
    print("    - Value types stay on stack when possible (fast)")
    print("    - Pointers, interfaces, closures may escape to heap")
    print("    - go build -gcflags=\"-m\" shows escape decisions")
    print()
    print("  Why it matters:")
    print("    - Stack alloc is ~free (just move stack pointer)")
    print("    - Heap alloc needs GC -- reduces allocs = less GC pressure")


if __name__ == "__main__":
    main()
