# Python equivalent of Generics Constraints (compare with main.go)
# Go constraints ~ Python TypeVar bounds or Protocol constraints

from typing import TypeVar

# --- comparable: Python has __eq__ on everything, so no special constraint ---
T = TypeVar("T")


def index(s: list[T], target: T) -> int:
    """Python has list.index(), but this matches Go's generic Index."""
    for i, v in enumerate(s):
        if v == target:
            return i
    return -1


# --- Number constraint: TypeVar with union ---
# Python TypeVar with constraints restricts to specific types
Num = TypeVar("Num", int, float)


def sum_all(s: list[Num]) -> Num:
    total = type(s[0])(0) if s else 0
    for v in s:
        total += v
    return total


# --- Ordered: Python supports < > on int, float, str natively ---
Ord = TypeVar("Ord", int, float, str)


def max_val(a: Ord, b: Ord) -> Ord:
    return a if a > b else b


def min_val(a: Ord, b: Ord) -> Ord:
    return a if a < b else b


# --- Named types (like Go's type UserID int) ---
class UserID(int):
    """Named type based on int (like Go's ~int with type alias)."""
    pass


def main():
    # --- comparable ---
    print("--- comparable (index) ---")
    nums = [10, 20, 30, 40]
    print("index of 30:", index(nums, 30))
    print("index of 99:", index(nums, 99))

    words = ["go", "python", "rust"]
    print("index of python:", index(words, "python"))

    # --- Number constraint ---
    print("\n--- Number constraint (sum) ---")
    ints = [1, 2, 3, 4, 5]
    print("sum ints:", sum_all(ints))

    floats = [1.1, 2.2, 3.3]
    print("sum floats:", sum_all(floats))

    # --- Named type based on int ---
    # Python: subclass of int works with int operations
    ids = [UserID(1), UserID(2), UserID(3)]
    print("sum UserIDs:", sum(ids))

    # --- Ordered ---
    print("\n--- Ordered (max/min) ---")
    print("max(3, 7):", max_val(3, 7))
    print("min(3, 7):", min_val(3, 7))
    print("max(a, z):", max_val("a", "z"))

    # Key difference:
    # Go constraints are enforced at compile time -- wrong type = compile error.
    # Python TypeVar bounds are hints only -- runtime does not enforce them.


if __name__ == "__main__":
    main()
