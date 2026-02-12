"""Memory profiling and allocations -- Python equivalent of the Go example.

Uses tracemalloc (stdlib) to track memory allocations per line.
"""

import tracemalloc


def alloc_heavy(n: int) -> str:
    """Many allocations: new string object per += operation."""
    s = ""
    for _ in range(n):
        s += "x"
    return s


def alloc_light(n: int) -> str:
    """Fewer allocations: list + single join."""
    parts = []
    for _ in range(n):
        parts.append("x")
    return "".join(parts)


def main() -> None:
    n = 50_000

    print("=== Memory Profiling with tracemalloc ===\n")

    # --- Profile alloc_heavy ---
    tracemalloc.start()
    _ = alloc_heavy(n)
    snapshot1 = tracemalloc.take_snapshot()
    tracemalloc.stop()

    print("--- alloc_heavy (+=) top allocations ---")
    for stat in snapshot1.statistics("lineno")[:5]:
        print(f"  {stat}")
    current1, peak1 = tracemalloc.get_traced_memory() if tracemalloc.is_tracing() else (0, 0)
    print()

    # --- Profile alloc_light ---
    tracemalloc.start()
    _ = alloc_light(n)
    snapshot2 = tracemalloc.take_snapshot()
    current2, peak2 = tracemalloc.get_traced_memory()
    tracemalloc.stop()

    print("--- alloc_light (join) top allocations ---")
    for stat in snapshot2.statistics("lineno")[:5]:
        print(f"  {stat}")
    print()

    print(f"  alloc_light peak memory: {peak2 / 1024:.1f} KB")
    print()
    print("--- Key insight ---")
    print("  += creates a new string object each iteration (O(n) allocs)")
    print("  list + join creates one final string (O(1) large allocs)")
    print()
    print("Commands:")
    print("  python -m tracemalloc script.py   # (Python 3.12+)")
    print("  Or start/stop in code as shown above")


if __name__ == "__main__":
    main()
