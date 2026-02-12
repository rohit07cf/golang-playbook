"""Benchmarks basics -- Python equivalent of the Go example.

Uses timeit (stdlib) to compare string building approaches.
"""

import timeit


def concat_strings(n: int) -> str:
    """Slow: += in a loop creates a new string each time."""
    s = ""
    for _ in range(n):
        s += "a"
    return s


def join_strings(n: int) -> str:
    """Fast: list + join allocates once at the end."""
    parts = []
    for _ in range(n):
        parts.append("a")
    return "".join(parts)


def main() -> None:
    n = 50_000

    print("=== String concat vs ''.join() ===")
    print(f"Building a {n}-char string...\n")

    # timeit: run each approach multiple times for stable results
    iters = 10

    concat_time = timeit.timeit(lambda: concat_strings(n), number=iters)
    join_time = timeit.timeit(lambda: join_strings(n), number=iters)

    print(f"  += concat:   {concat_time/iters*1000:.2f} ms/op  ({iters} iterations)")
    print(f"  ''.join():   {join_time/iters*1000:.2f} ms/op  ({iters} iterations)")
    print()
    print("Note: Python's timeit is the equivalent of Go's testing.B")
    print("  - timeit disables GC by default for stable timing")
    print("  - You set iteration count manually (Go auto-adjusts b.N)")
    print()
    print("For structured benchmarks, run:")
    print("  python ./09_performance_and_profiling/02_benchmarks_basics/bench_test.py")


if __name__ == "__main__":
    main()
