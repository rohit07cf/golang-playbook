"""Testing basics -- Python equivalent of the Go example."""


def add(a: int, b: int) -> int:
    return a + b


def abs_val(n: int) -> int:
    return -n if n < 0 else n


def is_even(n: int) -> bool:
    return n % 2 == 0


def main() -> None:
    print("add(2, 3) =", add(2, 3))
    print("add(-1, 1) =", add(-1, 1))
    print("abs_val(-7) =", abs_val(-7))
    print("abs_val(5) =", abs_val(5))
    print("is_even(4) =", is_even(4))
    print("is_even(7) =", is_even(7))


if __name__ == "__main__":
    main()
