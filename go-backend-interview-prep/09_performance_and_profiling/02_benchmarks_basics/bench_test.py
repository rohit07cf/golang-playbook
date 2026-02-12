"""Benchmark test -- Python equivalent of bench_test.go.

Uses timeit to benchmark string building approaches.
Run: python ./09_performance_and_profiling/02_benchmarks_basics/bench_test.py
"""

import timeit

N = 1000  # string length (matches Go bench)


def bench_concat():
    s = ""
    for _ in range(N):
        s += "a"
    return s


def bench_join():
    parts = []
    for _ in range(N):
        parts.append("a")
    return "".join(parts)


def bench_list_comp_join():
    return "".join("a" for _ in range(N))


def run_bench(name: str, func, iterations: int = 1000) -> None:
    total = timeit.timeit(func, number=iterations)
    per_op = total / iterations
    print(f"  {name:<30}  {iterations} ops  {per_op*1e6:>10.1f} us/op")


def main() -> None:
    print(f"=== Python Benchmarks (string length={N}) ===\n")

    run_bench("concat (+=)", bench_concat)
    run_bench("list + join", bench_join)
    run_bench("genexpr + join", bench_list_comp_join)

    print()
    print("Compare with Go:")
    print("  go test -bench=. -benchmem ./09_performance_and_profiling/02_benchmarks_basics/")


if __name__ == "__main__":
    main()
