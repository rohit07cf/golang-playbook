# Python equivalent of Multiple Return Values (compare with main.go)

def divide(a: float, b: float) -> tuple[float, str | None]:
    """Returns (result, error_message). Python uses tuples for multi-return."""
    if b == 0:
        return 0, "cannot divide by zero"
    return a / b, None


def min_max(*nums: int) -> tuple[int, int]:
    """Variadic via *args. Returns (min, max)."""
    lo = hi = nums[0]
    for n in nums[1:]:
        if n < lo:
            lo = n
        if n > hi:
            hi = n
    return lo, hi


def user_info() -> tuple[str, int, bool]:
    return "Alice", 30, True


def main():
    # --- Check error properly ---
    result, err = divide(10, 3)
    if err is not None:
        print("error:", err)
    else:
        print(f"10 / 3 = {result:.2f}")

    # --- Error case ---
    result, err = divide(10, 0)
    if err is not None:
        print("error:", err)

    # --- Discard error with _ ---
    quick, _ = divide(100, 4)
    print("100 / 4 =", quick)

    # --- Two non-error returns ---
    lo, hi = min_max(3, 1, 4, 1, 5, 9, 2, 6)
    print(f"min={lo} max={hi}")

    # --- Three return values ---
    name, age, active = user_info()
    print(f"user: {name}, age: {age}, active: {active}")


if __name__ == "__main__":
    main()
