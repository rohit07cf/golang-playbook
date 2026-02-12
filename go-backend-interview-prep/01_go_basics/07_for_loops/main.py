# Python equivalent of For Loops (compare with main.go)

def main():
    # --- C-style for -> range ---
    print("--- C-style (range) ---")
    for i in range(5):
        print(i, end=" ")
    print()

    # --- While-style ---
    print("--- While-style ---")
    n = 1
    while n < 50:
        n *= 2
    print("n doubled to:", n)

    # --- Infinite loop with break ---
    print("--- Infinite + break ---")
    counter = 0
    while True:
        if counter >= 3:
            break
        print("counter:", counter)
        counter += 1

    # --- Continue (skip iteration) ---
    print("--- Continue (skip even) ---")
    for i in range(6):
        if i % 2 == 0:
            continue
        print(i, end=" ")
    print()

    # --- Iterate over a string (by character, not byte) ---
    print("--- Iterate string ---")
    for idx, ch in enumerate("Go!"):
        print(f"index={idx} char={ch}")

    # --- Ignoring index with _ ---
    print("--- Ignore index ---")
    for _, ch in enumerate("abc"):
        print(ch, end=" ")
    print()
    # Or simply: for ch in "abc":

    # --- Nested loop with break-out ---
    # Python has no labeled break. Use a flag or function.
    print("--- Break from nested ---")
    done = False
    for i in range(3):
        for j in range(3):
            if i == 1 and j == 1:
                done = True
                break
            print(f"({i},{j})", end=" ")
        if done:
            break
    print()


if __name__ == "__main__":
    main()
