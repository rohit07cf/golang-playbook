# Python equivalent of Empty Interface and Type Assertions (compare with main.go)
# Go any ~ Python object (everything is an object)
# Go type assertion ~ Python isinstance()


def describe_item(v: object) -> None:
    """isinstance = Python's type assertion."""
    if v is None:
        print("  None")
    elif isinstance(v, str):
        print(f"  string: {v!r} (len={len(v)})")
    elif isinstance(v, int) and not isinstance(v, bool):
        # bool is a subclass of int in Python, so check bool first or exclude it
        print(f"  int: {v}")
    else:
        print(f"  other: {v} ({type(v).__name__})")


def main():
    # --- object accepts anything (like Go's any) ---
    print("--- object accepts anything ---")
    v: object

    v = 42
    print(f"int:    {v} (type: {type(v).__name__})")

    v = "hello"
    print(f"string: {v} (type: {type(v).__name__})")

    v = True
    print(f"bool:   {v} (type: {type(v).__name__})")

    v = [1, 2, 3]
    print(f"list:   {v} (type: {type(v).__name__})")

    # --- isinstance = safe type assertion (like comma-ok) ---
    print("\n--- isinstance (safe assertion) ---")
    v = "Go is great"

    if isinstance(v, str):
        print(f"string assertion: val={v!r} ok=True")

    if isinstance(v, int):
        print(f"int assertion: val={v} ok=True")
    else:
        print(f"int assertion: val=0 ok=False")

    # --- No panic equivalent ---
    # Python has no "single-value assertion that panics."
    # The closest is: cast without check -> AttributeError later.
    print("\n--- Direct use (no assertion needed in Python) ---")
    v = "safe string"
    print("direct use:", v)  # Python is dynamically typed -- just use it

    # --- Processing a list of mixed types ---
    print("\n--- Processing list[object] ---")
    items: list[object] = [42, "hello", 3.14, True, None]
    for item in items:
        describe_item(item)

    # --- None ---
    print("\n--- None ---")
    nil_any = None
    print(f"None: {nil_any} (is None: {nil_any is None})")


if __name__ == "__main__":
    main()
