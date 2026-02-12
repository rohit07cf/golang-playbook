# Python equivalent of Closures (compare with main.go)

def make_adder(x: int):
    """Returns a closure that adds x to its argument."""
    return lambda y: x + y


def counter_gen():
    """Returns a closure that counts up from 1."""
    n = 0
    def next_val():
        nonlocal n  # required to modify outer variable in Python
        n += 1
        return n
    return next_val


def main():
    # --- Basic closure: captures outer variable ---
    # Python requires 'nonlocal' to modify (Go does not)
    counter = 0
    def increment():
        nonlocal counter
        counter += 1
        return counter

    print("call 1:", increment())
    print("call 2:", increment())
    print("call 3:", increment())
    print("counter:", counter)

    # --- Closure as return value (function factory) ---
    print("\n--- Function factory ---")
    add_five = make_adder(5)
    add_ten = make_adder(10)
    print("add_five(3):", add_five(3))
    print("add_ten(3):", add_ten(3))

    # --- Closure with state ---
    print("\n--- Stateful closure ---")
    next_val = counter_gen()
    print(next_val())
    print(next_val())
    print(next_val())

    # --- Loop variable trap (same issue in Python!) ---
    print("\n--- Loop variable trap ---")
    funcs = []
    for i in range(5):
        funcs.append(lambda: print(i, end=" "))  # captures i by reference

    print("trap: ", end="")
    for fn in funcs:
        fn()  # all print 4
    print()

    # --- Fix: default argument captures current value ---
    funcs2 = []
    for i in range(5):
        funcs2.append(lambda i=i: print(i, end=" "))  # default arg = copy

    print("fixed:", end=" ")
    for fn in funcs2:
        fn()
    print()


if __name__ == "__main__":
    main()
