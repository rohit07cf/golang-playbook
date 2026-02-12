"""Maps and slices perf gotchas -- Python equivalent of the Go example.

Compares list/dict with and without preallocation.
"""

import time


def list_no_prealloc(n: int) -> list[int]:
    """Append to empty list (grows dynamically)."""
    s = []
    for i in range(n):
        s.append(i)
    return s


def list_prealloc(n: int) -> list[int]:
    """Preallocate with known size."""
    s = [0] * n
    for i in range(n):
        s[i] = i
    return s


def list_comprehension(n: int) -> list[int]:
    """List comprehension (fastest in CPython)."""
    return [i for i in range(n)]


def dict_no_hint(n: int) -> dict[str, int]:
    """Build dict entry by entry."""
    d = {}
    for i in range(n):
        d[str(i)] = i
    return d


def dict_comprehension(n: int) -> dict[str, int]:
    """Dict comprehension (often slightly faster)."""
    return {str(i): i for i in range(n)}


def main() -> None:
    print("=== Maps and Slices -- Performance Gotchas ===\n")

    # --- List benchmark ---
    sizes = [10_000, 100_000, 1_000_000]

    print("--- List: append vs preallocated vs comprehension ---")
    print(f"{'Size':<12}  {'append':>12}  {'preallocated':>12}  {'comprehension':>13}")
    print("-" * 56)

    for n in sizes:
        start = time.perf_counter()
        list_no_prealloc(n)
        t1 = time.perf_counter() - start

        start = time.perf_counter()
        list_prealloc(n)
        t2 = time.perf_counter() - start

        start = time.perf_counter()
        list_comprehension(n)
        t3 = time.perf_counter() - start

        print(f"{n:<12}  {t1:>11.4f}s  {t2:>11.4f}s  {t3:>12.4f}s")
    print()

    # --- Dict benchmark ---
    dict_sizes = [10_000, 100_000]

    print("--- Dict: loop vs comprehension ---")
    print(f"{'Size':<12}  {'loop':>12}  {'comprehension':>13}")
    print("-" * 40)

    for n in dict_sizes:
        start = time.perf_counter()
        dict_no_hint(n)
        t1 = time.perf_counter() - start

        start = time.perf_counter()
        dict_comprehension(n)
        t2 = time.perf_counter() - start

        print(f"{n:<12}  {t1:>11.4f}s  {t2:>12.4f}s")
    print()

    # --- List growth pattern ---
    print("--- List internal size growth ---")
    import sys
    s: list = []
    prev_size = sys.getsizeof(s)
    for i in range(20):
        s.append(i)
        new_size = sys.getsizeof(s)
        if new_size != prev_size:
            print(f"  len={len(s):<3} size={new_size} bytes (grew from {prev_size})")
            prev_size = new_size
    print()

    print("Key: list comprehension is fastest in Python.")
    print("     Go's make([]T, 0, n) has no exact Python equivalent,")
    print("     but [None]*n + index assignment is closest.")


if __name__ == "__main__":
    main()
