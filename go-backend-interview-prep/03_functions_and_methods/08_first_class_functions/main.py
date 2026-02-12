# Python equivalent of First-Class Functions (compare with main.go)
from typing import Callable


# Type alias (like Go's named function type)
Transform = Callable[[int], int]


def apply(nums: list[int], fn: Transform) -> list[int]:
    """Apply a function to every element (like Go apply)."""
    return [fn(v) for v in nums]


def filter_nums(nums: list[int], predicate: Callable[[int], bool]) -> list[int]:
    return [v for v in nums if predicate(v)]


def multiplier(factor: int) -> Transform:
    """Function factory (returns a closure)."""
    return lambda n: n * factor


def bubble_sort(nums: list[int], less: Callable[[int, int], bool]) -> list[int]:
    """Strategy pattern: sort behavior determined by 'less' function."""
    result = nums[:]
    for i in range(len(result)):
        for j in range(i + 1, len(result)):
            if less(result[j], result[i]):
                result[i], result[j] = result[j], result[i]
    return result


def main():
    nums = [1, 2, 3, 4, 5]

    # --- Pass function as argument ---
    print("--- Apply (map) ---")
    doubled = apply(nums, lambda n: n * 2)
    print("doubled:", doubled)

    squared = apply(nums, lambda n: n * n)
    print("squared:", squared)

    # Python built-in: list(map(lambda n: n*2, nums))

    # --- Function factory ---
    print("\n--- Multiplier factory ---")
    triple = multiplier(3)
    print("tripled:", apply(nums, triple))

    # --- Filter ---
    print("\n--- Filter ---")
    evens = filter_nums(nums, lambda n: n % 2 == 0)
    print("evens:", evens)

    # Python built-in: list(filter(lambda n: n%2==0, nums))

    # --- Strategy pattern ---
    print("\n--- Strategy pattern (sort) ---")
    data = [5, 3, 1, 4, 2]
    ascending = bubble_sort(data, lambda a, b: a < b)
    descending = bubble_sort(data, lambda a, b: a > b)
    print("ascending:", ascending)
    print("descending:", descending)

    # --- Function variable ---
    print("\n--- Function variable ---")
    fn: Transform | None = None
    print("None function:", fn is None)
    fn = lambda n: n + 100
    print("fn(5):", fn(5))


if __name__ == "__main__":
    main()
