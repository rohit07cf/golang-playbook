# Python equivalent of Arrays (compare with main.go)
# Python has no fixed-size array type. Lists are dynamic.
# For fixed-size, use tuples (immutable) or array module.

def main():
    # --- Python list (closest to Go array for demo) ---
    nums = [0, 0, 0, 0, 0]
    print("zero-valued:", nums)

    # --- Initialized ---
    colors = ["red", "green", "blue"]
    print("colors:", colors)

    primes = [2, 3, 5, 7, 11]
    print("primes:", primes)
    print("length:", len(primes))

    # --- Access and modify ---
    nums[0] = 10
    nums[4] = 50
    print("modified:", nums)

    # --- Lists are REFERENCE types (unlike Go arrays!) ---
    # Assigning does NOT copy in Python.
    original = [1, 2, 3]
    copied = original       # same object!
    copied[0] = 999
    print("original:", original)  # [999, 2, 3] -- CHANGED (unlike Go)
    print("copied:", copied)

    # --- To actually copy, use [:] or list() ---
    original2 = [1, 2, 3]
    copied2 = original2[:]  # shallow copy
    copied2[0] = 999
    print("original2:", original2)  # [1, 2, 3] -- unchanged
    print("copied2:", copied2)

    # --- Passing to a function ---
    data = [10, 20, 30]
    try_to_modify(data)
    print("after function call:", data)  # [999, 20, 30] -- CHANGED (unlike Go)

    # --- Iterate ---
    print("--- iterate ---")
    for i, v in enumerate(colors):
        print(f"  index={i} value={v}")

    # --- 2D list ---
    grid = [[1, 2, 3], [4, 5, 6]]
    print("2D grid:", grid)


def try_to_modify(lst):
    lst[0] = 999  # modifies the original (Python lists are reference types)


if __name__ == "__main__":
    main()
