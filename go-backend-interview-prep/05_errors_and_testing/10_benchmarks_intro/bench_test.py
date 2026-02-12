"""Benchmarks intro -- Python equivalent of bench_test.go using timeit."""

import timeit
from main import fib, fib_iter, string_concat

RUNS = 1000


def bench(name: str, func, number: int = RUNS) -> None:
    elapsed = timeit.timeit(func, number=number)
    per_op = elapsed / number * 1e6  # microseconds
    print(f"{name:30s}  {number} runs  {elapsed:.4f}s total  {per_op:.1f} us/op")


def main() -> None:
    print("Python benchmarks (timeit)\n")
    bench("fib(20) recursive", lambda: fib(20))
    bench("fib_iter(20) iterative", lambda: fib_iter(20))
    bench("fib(10) recursive", lambda: fib(10))
    bench("fib_iter(10) iterative", lambda: fib_iter(10))
    bench("string_concat(100)", lambda: string_concat(100))


if __name__ == "__main__":
    main()
