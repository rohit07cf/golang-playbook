# Python equivalent of Strings and Runes (compare with main.go)
# Python str is Unicode (not bytes). len() returns characters.

def main():
    # --- Length = character count (unlike Go!) ---
    s = "Hello, world"
    print("string:", s)
    print("len (characters):", len(s))  # Go: len() returns bytes
    print("len (bytes):", len(s.encode("utf-8")))

    # --- Multi-byte characters ---
    print("\n--- Multi-byte ---")
    jp = "Go\u8a00\u8a9e"  # "Go言語"
    print("string:", jp)
    print("len (characters):", len(jp))              # 4
    print("len (bytes):", len(jp.encode("utf-8")))   # 8

    # --- Indexing gives a character (unlike Go which gives a byte) ---
    print("\n--- Character indexing ---")
    print(f"jp[0] = {jp[0]}")   # 'G'
    print(f"jp[2] = {jp[2]}")   # '言' (full character, unlike Go)

    # --- Iterate by character (default in Python) ---
    print("\n--- Iterate (by character) ---")
    for i, ch in enumerate(jp):
        print(f"  index={i} char={ch} (U+{ord(ch):04X})")

    # --- Strings are immutable (same as Go) ---
    # s[0] = 'h'  # TypeError

    # --- Modify: convert to list of chars ---
    print("\n--- Modify via list ---")
    chars = list(jp)
    chars[2] = 'X'
    print("modified:", "".join(chars))

    # --- bytes vs str ---
    print("\n--- bytes vs str ---")
    b = b"hello"      # bytes literal
    print("bytes:", b)
    print("decoded:", b.decode("utf-8"))

    # --- Building strings efficiently ---
    print("\n--- String building ---")
    parts = []
    for i in range(5):
        parts.append(f"item{i}")
    print("built:", " ".join(parts))

    # --- Useful string methods ---
    print("\n--- String methods ---")
    msg = "  Python is Great  "
    print("strip:", repr(msg.strip()))
    print("contains:", "Great" in msg)
    print("split:", "a,b,c".split(","))
    print("join:", "-".join(["x", "y", "z"]))


if __name__ == "__main__":
    main()
