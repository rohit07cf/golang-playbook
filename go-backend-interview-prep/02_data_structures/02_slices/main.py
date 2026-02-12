# Python equivalent of Slices (compare with main.go)
# Python lists are dynamic arrays (closest to Go slices).
# Python has no len/cap distinction -- lists manage capacity internally.

def main():
    # --- List literal ---
    nums = [10, 20, 30, 40, 50]
    print("nums:", nums)

    # --- No len/cap distinction in Python ---
    s = [0] * 3
    print(f"[0]*3: len={len(s)} {s}")

    # --- Append ---
    s.append(99)
    print(f"after append: len={len(s)} {s}")

    # --- Append past capacity (Python handles this internally) ---
    small = []
    small.extend([1, 2, 3])
    print(f"small: len={len(small)} {small}")

    # --- Slicing CREATES A NEW LIST (unlike Go!) ---
    print("\n--- Python slicing creates a copy ---")
    original = [1, 2, 3, 4, 5]
    sub = original[1:4]     # [2, 3, 4] -- this is a NEW list
    print("original:", original)
    print("sub:     ", sub)

    sub[0] = 999
    print("after sub[0] = 999:")
    print("original:", original)  # UNCHANGED (unlike Go)
    print("sub:     ", sub)

    # --- Copy ---
    print("\n--- Explicit copy ---")
    src = [10, 20, 30]
    dst = src[:]            # or: dst = list(src)
    dst[0] = 999
    print("src:", src)      # unchanged
    print("dst:", dst)

    # --- None vs empty list ---
    print("\n--- None vs empty ---")
    none_list = None
    empty_list = []
    print(f"None list:  {none_list}  is None={none_list is None}")
    print(f"empty list: {empty_list}  is None={empty_list is None}")

    # --- Append works on empty lists ---
    data = []
    data.extend([1, 2, 3])
    print("\nappend to empty:", data)


if __name__ == "__main__":
    main()
