# Python equivalent of Methods and Receivers (compare with main.go)
# Python methods always receive 'self' (a reference), similar to Go pointer receivers.
# There is no "value receiver" concept in Python.


class Rect:
    def __init__(self, width: float = 0, height: float = 0):
        self.width = width
        self.height = height

    def area(self) -> float:
        """All Python methods are like Go pointer receivers (self is a reference)."""
        return self.width * self.height

    def perimeter(self) -> float:
        return 2 * (self.width + self.height)

    def scale(self, factor: float) -> None:
        """Mutates the original (always, in Python)."""
        self.width *= factor
        self.height *= factor

    def set_width(self, w: float) -> None:
        self.width = w

    def __repr__(self):
        return f"Rect(width={self.width}, height={self.height})"


def try_to_mutate_value(r: Rect) -> None:
    """In Python, this WILL mutate the original (unlike Go value receiver)."""
    r.width = 999


def main():
    # --- Methods ---
    r = Rect(width=10, height=5)
    print("area:", r.area())
    print("perimeter:", r.perimeter())

    # --- Mutation (Python always mutates -- no value receiver concept) ---
    print("\n--- Mutation (always pointer-like) ---")
    print("before scale:", r)
    r.scale(2)
    print("after scale(2):", r)

    r.set_width(100)
    print("after set_width(100):", r)

    # --- Python functions ALWAYS get reference to mutable objects ---
    print("\n--- Function receives reference ---")
    r2 = Rect(width=3, height=4)
    try_to_mutate_value(r2)
    print("r2 CHANGED:", r2)  # width=999 (unlike Go value receiver)

    # Key difference:
    # Go value receiver = copy (no mutation)
    # Python self = reference (always mutates)
    print("\nKey: Go value receiver = copy. Python self = reference.")


if __name__ == "__main__":
    main()
