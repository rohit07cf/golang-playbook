# Python equivalent of Recursion (compare with main.go)
import sys

# Python has a recursion limit (default 1000). Go's stack grows dynamically.
# sys.setrecursionlimit(10000)  # uncomment if needed


def factorial(n: int) -> int:
    if n <= 1:
        return 1
    return n * factorial(n - 1)


def factorial_iter(n: int) -> int:
    result = 1
    for i in range(2, n + 1):
        result *= i
    return result


def fib(n: int) -> int:
    """Naive recursive -- O(2^n). Same as Go version."""
    if n < 2:
        return n
    return fib(n - 1) + fib(n - 2)


def fib_iter(n: int) -> int:
    if n < 2:
        return n
    a, b = 0, 1
    for _ in range(2, n + 1):
        a, b = b, a + b
    return b


def sum_digits(n: int) -> int:
    if n < 10:
        return n
    return n % 10 + sum_digits(n // 10)


def main():
    # --- Factorial ---
    print("--- Factorial ---")
    for n in [0, 1, 5, 10]:
        print(f"factorial({n}) = {factorial(n)}")

    # --- Iterative ---
    print("\n--- Factorial (iterative) ---")
    print("factorial_iter(10):", factorial_iter(10))

    # --- Fibonacci ---
    print("\n--- Fibonacci (recursive) ---")
    for i in range(10):
        print(f"fib({i}) = {fib(i)}")

    # --- Fibonacci (iterative) ---
    print("\n--- Fibonacci (iterative) ---")
    print("fib_iter(30):", fib_iter(30))

    # --- Sum of digits ---
    print("\n--- Sum of digits ---")
    print("sum_digits(12345):", sum_digits(12345))

    # Note: Python default recursion limit is ~1000
    print(f"\nPython recursion limit: {sys.getrecursionlimit()}")


if __name__ == "__main__":
    main()
