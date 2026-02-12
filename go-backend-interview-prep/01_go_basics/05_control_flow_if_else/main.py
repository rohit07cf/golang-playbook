# Python equivalent of Control Flow: If/Else (compare with main.go)

def multiply(a: int, b: int) -> int:
    return a * b


def main():
    # --- Basic if/else ---
    temperature = 22

    if temperature > 30:
        print("hot")
    elif temperature > 15:
        print("comfortable")
    else:
        print("cold")

    # --- Python has no if-init statement ---
    # Go:  if length := len("backend"); length > 5 { ... }
    # Python: assign first, then check
    length = len("backend")
    if length > 5:
        print("long word, length:", length)
    else:
        print("short word, length:", length)

    # --- Walrus operator (:=) is the closest equivalent (Python 3.8+) ---
    if (result := multiply(3, 7)) > 20:
        print("product is large:", result)

    # --- No ternary? Python HAS one (Go does not) ---
    age = 20
    status = "adult" if age >= 18 else "minor"
    print("status:", status)

    # --- Boolean conditions ---
    logged_in = True
    is_admin = False

    if logged_in and is_admin:
        print("admin dashboard")
    elif logged_in:
        print("user dashboard")
    else:
        print("please log in")


if __name__ == "__main__":
    main()
