# Python equivalent of Generics Basics (compare with main.go)
# Go generics (compile-time) ~ Python typing.TypeVar (type hints, runtime unchecked)

from typing import TypeVar, Callable

T = TypeVar("T")
U = TypeVar("U")


def map_fn(s: list[T], fn: Callable[[T], U]) -> list[U]:
    """Generic map -- same as Go Map[T, U]."""
    return [fn(v) for v in s]


def filter_fn(s: list[T], predicate: Callable[[T], bool]) -> list[T]:
    return [v for v in s if predicate(v)]


def contains(s: list[T], target: T) -> bool:
    """Python lists support 'in' natively, but this matches Go's Contains."""
    return target in s


class Stack:
    """Python has no generics enforcement at runtime.
    Type hints are for documentation and static checkers only."""

    def __init__(self):
        self.items: list = []

    def push(self, v) -> None:
        self.items.append(v)

    def pop(self) -> tuple:
        if not self.items:
            return None, False
        return self.items.pop(), True

    def __len__(self) -> int:
        return len(self.items)


def main():
    # --- Map ---
    print("--- Map ---")
    nums = [1, 2, 3, 4, 5]
    doubled = map_fn(nums, lambda n: n * 2)
    print("doubled:", doubled)

    strs = map_fn(nums, lambda n: f"#{n}")
    print("as strings:", strs)

    # --- Filter ---
    print("\n--- Filter ---")
    evens = filter_fn(nums, lambda n: n % 2 == 0)
    print("evens:", evens)

    words = ["Go", "Python", "Rust", "Go"]
    long = filter_fn(words, lambda s: len(s) > 2)
    print("long words:", long)

    # --- Contains ---
    print("\n--- Contains ---")
    print("has 3:", contains(nums, 3))
    print("has 9:", contains(nums, 9))
    print("has Go:", contains(words, "Go"))

    # --- Stack ---
    print("\n--- Stack[int] ---")
    int_stack = Stack()
    int_stack.push(10)
    int_stack.push(20)
    int_stack.push(30)
    print("len:", len(int_stack))
    v, ok = int_stack.pop()
    print(f"pop: {v} (ok={ok})")

    print("\n--- Stack[str] ---")
    str_stack = Stack()
    str_stack.push("hello")
    str_stack.push("world")
    s, ok = str_stack.pop()
    print(f"pop: {s!r} (ok={ok})")

    # Key difference:
    # Go generics are enforced at compile time. Type errors are caught before running.
    # Python type hints are NOT enforced at runtime. They are for tools like mypy.


if __name__ == "__main__":
    main()
