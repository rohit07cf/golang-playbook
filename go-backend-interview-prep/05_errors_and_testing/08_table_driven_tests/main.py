"""Table-driven tests -- Python equivalent of the Go example."""


def parse_int_safe(s: str, fallback: int = 0) -> tuple[int, str | None]:
    """Parse s as int. Returns (value, error_msg). Uses fallback on failure."""
    try:
        return int(s), None
    except ValueError:
        return fallback, f"parse_int_safe({s!r}): invalid literal"


def clamp(n: int, lo: int, hi: int) -> int:
    """Restrict n to the range [lo, hi]."""
    if n < lo:
        return lo
    if n > hi:
        return hi
    return n


def main() -> None:
    # ParseIntSafe examples
    v, err = parse_int_safe("42", 0)
    print(f'parse_int_safe("42", 0) = {v}, err={err}')

    v, err = parse_int_safe("abc", -1)
    print(f'parse_int_safe("abc", -1) = {v}, err={err}')

    # Clamp examples
    print(f"clamp(5, 0, 10) = {clamp(5, 0, 10)}")
    print(f"clamp(-3, 0, 10) = {clamp(-3, 0, 10)}")
    print(f"clamp(15, 0, 10) = {clamp(15, 0, 10)}")


if __name__ == "__main__":
    main()
