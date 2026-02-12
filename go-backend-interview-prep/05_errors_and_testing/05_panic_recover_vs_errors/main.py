"""Panic, recover vs errors -- Python equivalent of the Go example."""


def safe_divide(a: int, b: int) -> int:
    """Catches ZeroDivisionError (like Go recover from panic)."""
    try:
        return a // b
    except ZeroDivisionError as e:
        raise ValueError(f"recovered: {e}") from e


def must_positive(n: int) -> int:
    """Panics (asserts) on programmer error."""
    assert n > 0, f"must_positive: got {n}, want > 0"
    return n


def parse_age(s: str) -> int:
    """Returns age or raises ValueError for bad input."""
    try:
        age = int(s)
    except ValueError:
        raise ValueError(f"parse_age({s!r}): not a number")
    if age < 0 or age > 150:
        raise ValueError("age out of range")
    return age


def main() -> None:
    # --- Example 1: catch crash-level error ---
    try:
        result = safe_divide(10, 0)
        print(f"safe_divide(10, 0): {result}")
    except ValueError as e:
        print(f"safe_divide(10, 0): {e}")

    print(f"safe_divide(10, 3): {safe_divide(10, 3)}")

    # --- Example 2: assertion for programmer bugs ---
    print(f"\nmust_positive(5): {must_positive(5)}")

    try:
        must_positive(-1)
    except AssertionError as e:
        print(f"caught assertion: {e}")

    # --- Example 3: exceptions for expected failures ---
    try:
        age = parse_age("25")
        print(f'\nparse_age("25"): {age} None')
    except ValueError as e:
        print(f'\nparse_age("25"): {e}')

    try:
        age = parse_age("abc")
        print(f'parse_age("abc"): {age} None')
    except ValueError as e:
        print(f'parse_age("abc"): 0 {e}')

    try:
        age = parse_age("999")
        print(f'parse_age("999"): {age} None')
    except ValueError as e:
        print(f'parse_age("999"): 0 {e}')


if __name__ == "__main__":
    main()
