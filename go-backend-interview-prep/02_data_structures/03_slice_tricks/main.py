# Python equivalent of Slice Tricks (compare with main.go)
# Python lists have built-in methods for most of these.

def main():
    # === DELETE by index (order preserved) ===
    print("--- Delete (keep order) ---")
    s = [10, 20, 30, 40, 50]
    del s[2]  # or s.pop(2)
    print("after delete index 2:", s)

    # === DELETE (swap-with-last for O(1), order NOT preserved) ===
    print("\n--- Delete (swap-with-last) ---")
    s2 = [10, 20, 30, 40, 50]
    i = 1
    s2[i] = s2[-1]
    s2.pop()
    print("after fast delete index 1:", s2)

    # === INSERT at index ===
    print("\n--- Insert at index ---")
    s3 = [1, 2, 4, 5]
    s3.insert(2, 3)  # Go has no built-in insert
    print("after insert 3 at index 2:", s3)

    # === FILTER ===
    print("\n--- Filter (keep odds) ---")
    data = [1, 2, 3, 4, 5, 6, 7, 8]
    odds = [v for v in data if v % 2 != 0]
    print("odds only:", odds)

    # === PREALLOCATE (not really needed in Python, but shown) ===
    print("\n--- List comprehension (like preallocate) ---")
    inp = [1, 2, 3, 4, 5]
    doubled = [v * 2 for v in inp]
    print(f"doubled: {doubled}")

    # === COPY (independent) ===
    print("\n--- Copy ---")
    src = [10, 20, 30]
    dst = src[:]
    dst[0] = 999
    print("src:", src)
    print("dst:", dst)

    # === REVERSE in-place ===
    print("\n--- Reverse ---")
    r = [1, 2, 3, 4, 5]
    r.reverse()  # or r[::-1] for a new list
    print("reversed:", r)

    # === DEDUPLICATE (preserving order) ===
    print("\n--- Deduplicate ---")
    items = [1, 1, 2, 3, 3, 3, 4]
    seen = set()
    deduped = []
    for x in items:
        if x not in seen:
            seen.add(x)
            deduped.append(x)
    print("deduped:", deduped)


if __name__ == "__main__":
    main()
