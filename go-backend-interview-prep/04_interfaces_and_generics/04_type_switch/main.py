# Python equivalent of Type Switch (compare with main.go)
# Go type switch ~ Python match/case with type patterns (3.10+) or isinstance chains


def classify(v: object) -> None:
    """Python 3.10+ structural pattern matching on types."""
    match v:
        case str():
            print(f"  string: {v!r} (len={len(v)})")
        case bool():
            # Must check bool before int (bool is subclass of int)
            print(f"  bool: {v}")
        case int():
            print(f"  int: {v}")
        case float():
            print(f"  float: {v:.2f}")
        case None:
            print("  None")
        case _:
            print(f"  unknown: {v} ({type(v).__name__})")


def is_numeric(v: object) -> bool:
    """Multiple types in one check."""
    return isinstance(v, (int, float)) and not isinstance(v, bool)


def main():
    # --- Basic type dispatch ---
    print("--- Type switch ---")
    values: list[object] = [42, "hello", 3.14, True, None, [1, 2]]
    for v in values:
        classify(v)

    # --- Multiple types check ---
    print("\n--- is_numeric ---")
    print("42:", is_numeric(42))
    print("3.14:", is_numeric(3.14))
    print("hello:", is_numeric("hello"))

    # --- isinstance chains (pre-3.10 style) ---
    print("\n--- isinstance chains ---")
    v: object = "test"
    if isinstance(v, str):
        print(f"  str: {v}")
    elif isinstance(v, int):
        print(f"  int: {v}")
    else:
        print(f"  other: {v}")


if __name__ == "__main__":
    main()
