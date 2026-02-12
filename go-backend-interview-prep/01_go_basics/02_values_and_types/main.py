# Python equivalent of Values and Types (compare with main.go)

def main():
    # --- Integers (arbitrary precision in Python) ---
    age: int = 30
    big_num: int = 9_000_000_000
    print("int:", age)
    print("big int:", big_num)

    # --- Floats ---
    pi: float = 3.14159
    print("float:", pi)

    # --- Booleans ---
    active: bool = True
    print("bool:", active)

    # --- Strings ---
    greeting: str = "hello, python"
    print("string:", greeting)
    print("string length:", len(greeting))  # characters, not bytes

    # --- Python has no byte/rune distinction at the type level ---
    ch = ord('A')       # integer value of character
    print(f"ord('A'): {ch}")
    print(f"chr(90):  {chr(90)}")

    # --- Implicit conversion (Python does this; Go does NOT) ---
    result = age + pi  # works in Python, fails in Go
    print("implicit conversion (int + float):", result)

    # --- Type inspection ---
    print(f"type of age: {type(age).__name__}")
    print(f"type of pi: {type(pi).__name__}")
    print(f"type of greeting: {type(greeting).__name__}")


if __name__ == "__main__":
    main()
