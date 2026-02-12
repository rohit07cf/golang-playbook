"""Big-O and hotspots -- Python equivalent of the Go example.

Compares O(n^2) nested-loop duplicate finder vs O(n) set-based approach.
"""

import random
import time


def find_duplicates_slow(nums: list[int]) -> list[int]:
    """O(n^2): nested loop comparison."""
    dups = []
    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):
            if nums[i] == nums[j]:
                dups.append(nums[i])
                break
    return dups


def find_duplicates_fast(nums: list[int]) -> list[int]:
    """O(n): single pass with a set."""
    seen: set[int] = set()
    dups = []
    for n in nums:
        if n in seen:
            dups.append(n)
        seen.add(n)
    return dups


def generate_data(size: int) -> list[int]:
    return [random.randint(0, size // 2) for _ in range(size)]


def main() -> None:
    sizes = [1_000, 5_000, 10_000, 20_000]

    print("=== Big-O: O(n^2) vs O(n) duplicate finder ===")
    print(f"{'Size':<10}  {'O(n^2)':>15}  {'O(n) set':>15}")
    print("-" * 46)

    for size in sizes:
        data = generate_data(size)

        # O(n^2)
        start = time.perf_counter()
        find_duplicates_slow(data)
        slow = time.perf_counter() - start

        # O(n)
        start = time.perf_counter()
        find_duplicates_fast(data)
        fast = time.perf_counter() - start

        print(f"{size:<10}  {slow:>14.4f}s  {fast:>14.6f}s")

    print()
    print("Key takeaway: O(n^2) time grows quadratically;")
    print("O(n) stays roughly linear. The set trades O(n) space for speed.")


if __name__ == "__main__":
    main()
