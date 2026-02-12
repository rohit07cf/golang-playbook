"""Benchmarks intro -- Python equivalent of the Go example."""


def fib(n: int) -> int:
    """Recursive Fibonacci."""
    if n <= 1:
        return n
    return fib(n - 1) + fib(n - 2)


def fib_iter(n: int) -> int:
    """Iterative Fibonacci."""
    if n <= 1:
        return n
    a, b = 0, 1
    for _ in range(2, n + 1):
        a, b = b, a + b
    return b


def string_concat(n: int) -> str:
    """Build a string by repeated concatenation (slow)."""
    s = ""
    for _ in range(n):
        s += "a"
    return s


def main() -> None:
    print("fib(10) =", fib(10))
    print("fib_iter(10) =", fib_iter(10))
    print("fib(20) =", fib(20))
    print("fib_iter(20) =", fib_iter(20))
    print("len(string_concat(100)) =", len(string_concat(100)))


if __name__ == "__main__":
    main()
