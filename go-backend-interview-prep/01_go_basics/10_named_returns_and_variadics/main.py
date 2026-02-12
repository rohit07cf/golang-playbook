# Python equivalent of Named Returns and Variadics (compare with main.go)
# Python has no named returns. You just return values.
# Python has *args for variadics.


def split(total: int) -> tuple[int, int]:
    """Go uses named returns; Python just returns a tuple."""
    half = total // 2
    remainder = total % 2
    return half, remainder


def parse_coords(input_str: str) -> tuple[float, float, str | None]:
    """Named returns in Go serve as docs. In Python, use type hints."""
    lat = 35.6762
    lng = 139.6503
    return lat, lng, None


def sum_all(*nums: int) -> int:
    """Go: func sum(nums ...int). Python: *args."""
    total = 0
    for n in nums:
        total += n
    return total


def greet_all(greeting: str, *names: str) -> None:
    """Go: func greetAll(greeting string, names ...string)."""
    for name in names:
        print(greeting, name)


def main():
    # --- "Named" returns (just tuple unpacking) ---
    h, r = split(17)
    print(f"split(17): half={h} remainder={r}")

    lat, lng, _ = parse_coords("ignored")
    print(f"coords: lat={lat:.4f} lng={lng:.4f}")

    # --- Variadic: individual args ---
    print("sum(1,2,3):", sum_all(1, 2, 3))
    print("sum(10,20):", sum_all(10, 20))
    print("sum():", sum_all())  # valid -- returns 0

    # --- Variadic: spread a list ---
    scores = [90, 85, 92, 78]
    print("sum(*scores):", sum_all(*scores))  # Go: sum(scores...)

    # --- Variadic with fixed param ---
    greet_all("Hello", "Alice", "Bob", "Charlie")


if __name__ == "__main__":
    main()
