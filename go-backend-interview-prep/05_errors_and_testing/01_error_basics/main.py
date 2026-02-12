"""Error basics -- Python equivalent of the Go example."""


def divide(a: float, b: float) -> float:
    """Divide a by b. Raises ValueError on division by zero."""
    if b == 0:
        raise ValueError("division by zero")
    return a / b


def parse_and_double(s: str) -> int:
    """Parse string to int and double it."""
    try:
        n = int(s)
    except ValueError as e:
        raise ValueError(f"parse_and_double({s!r}): {e}") from e
    return n * 2


def main() -> None:
    # --- Example 1: basic error check ---
    try:
        result = divide(10, 3)
        print(f"10 / 3 = {result:.2f}")
    except ValueError as e:
        print("ERROR:", e)

    try:
        result = divide(10, 0)
        print(f"10 / 0 = {result:.2f}")
    except ValueError as e:
        print("ERROR:", e)

    # --- Example 2: wrapping context ---
    try:
        val = parse_and_double("42")
        print(f'parse_and_double("42") = {val}')
    except ValueError as e:
        print("ERROR:", e)

    try:
        val = parse_and_double("abc")
        print(f'parse_and_double("abc") = {val}')
    except ValueError as e:
        print("ERROR:", e)

    # --- Example 3: creating errors ---
    e1 = ValueError("something went wrong")
    e2 = RuntimeError(f"failed to load config: {e1}")
    print(e2)


if __name__ == "__main__":
    main()
