# Python equivalent of Maps (compare with main.go)
# Go map -> Python dict

def main():
    # --- Create with literal ---
    ages = {"alice": 30, "bob": 25}
    print("ages:", ages)

    # --- Create empty ---
    scores = {}
    scores["math"] = 95
    scores["english"] = 88
    print("scores:", scores)

    # --- Read: KeyError for missing keys (unlike Go zero value) ---
    print("\n--- Read ---")
    print("alice:", ages["alice"])
    # ages["nobody"]  # would raise KeyError (Go returns zero value)

    # --- .get() is Python's comma-ok equivalent ---
    print("\n--- .get() (like comma-ok) ---")
    val = ages.get("alice", 0)  # default if missing
    print(f"alice: val={val}")

    val = ages.get("nobody", 0)
    print(f"nobody: val={val}")

    # --- Check existence with `in` ---
    print("alice exists:", "alice" in ages)
    print("nobody exists:", "nobody" in ages)

    # --- Write ---
    ages["charlie"] = 35
    print("\nafter add charlie:", ages)

    # --- Update ---
    ages["alice"] = 31
    print("after update alice:", ages)

    # --- Delete ---
    del ages["bob"]
    print("after delete bob:", ages)

    # Delete missing: use pop with default to avoid KeyError
    ages.pop("nobody", None)

    # --- Length ---
    print("length:", len(ages))

    # --- Iteration order (insertion order since Python 3.7) ---
    print("\n--- Iteration (insertion order) ---")
    colors = {"red": "#ff0000", "green": "#00ff00", "blue": "#0000ff"}
    for k, v in colors.items():
        print(f"  {k} -> {v}")
    print("(Python preserves insertion order; Go does NOT)")

    # --- No nil dict trap in Python ---
    # None-type dict just doesn't have methods.
    # But: d = None; d["key"] = 1  -> AttributeError

    # --- Word frequency (same pattern) ---
    print("\n--- Word frequency ---")
    words = ["go", "is", "great", "go", "is", "fast", "go"]
    freq = {}
    for w in words:
        freq[w] = freq.get(w, 0) + 1
    for word, count in freq.items():
        print(f"  {word}: {count}")

    # --- Set ---
    print("\n--- Set ---")
    items = ["a", "b", "a", "c", "b"]
    unique = set(items)
    print("unique count:", len(unique))


if __name__ == "__main__":
    main()
