# Python equivalent of Interfaces Basics (compare with main.go)
# Go interfaces (implicit) ~ Python Protocol (structural typing, 3.8+)
# Python also has ABC, but Protocol is closer to Go's model.

import math
from typing import Protocol


# --- Protocol = structural interface (like Go interface) ---
class Shape(Protocol):
    def area(self) -> float: ...
    def perimeter(self) -> float: ...


# --- Circle satisfies Shape (no explicit declaration) ---
class Circle:
    def __init__(self, radius: float):
        self.radius = radius

    def area(self) -> float:
        return math.pi * self.radius ** 2

    def perimeter(self) -> float:
        return 2 * math.pi * self.radius


# --- Rectangle also satisfies Shape ---
class Rectangle:
    def __init__(self, width: float, height: float):
        self.width = width
        self.height = height

    def area(self) -> float:
        return self.width * self.height

    def perimeter(self) -> float:
        return 2 * (self.width + self.height)


def print_shape(s: Shape) -> None:
    """Accepts anything with area() and perimeter() methods."""
    print(f"  area={s.area():.2f}  perimeter={s.perimeter():.2f}")


def main():
    # --- Implicit satisfaction (duck typing / Protocol) ---
    print("--- Interface basics ---")
    c = Circle(radius=5)
    r = Rectangle(width=10, height=3)

    print_shape(c)
    print_shape(r)

    # --- Variable holding different types ---
    print("\n--- Variable holding different types ---")
    s: Shape = c
    print(f"type holding Circle: area={s.area():.2f}")
    s = r
    print(f"type holding Rect:   area={s.area():.2f}")

    # --- List of shapes ---
    print("\n--- List of shapes ---")
    shapes: list[Shape] = [
        Circle(radius=1),
        Rectangle(width=4, height=5),
        Circle(radius=3),
    ]
    for sh in shapes:
        print_shape(sh)

    # --- None vs nil interface ---
    print("\n--- None comparison ---")
    nil_shape = None
    print("None shape:", nil_shape is None)  # True
    # Python has no nil-pointer-in-interface trap.
    # None is just None. No type-value pair wrapper.


if __name__ == "__main__":
    main()
