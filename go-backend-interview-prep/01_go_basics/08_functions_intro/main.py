# Python equivalent of Functions Intro (compare with main.go)

def greet(name: str) -> str:
    return "hello, " + name


def add(a: int, b: int) -> int:
    return a + b


def log_message(msg: str) -> None:
    print("[LOG]", msg)


def describe(name: str, age: int) -> str:
    return f"{name} is {age} years old"


def apply(value: int, fn) -> int:
    """Takes a function as a parameter (first-class functions)."""
    return fn(value)


def no_change(x: int) -> None:
    x = 999  # rebinds local name; original int is immutable


def main():
    # --- Calling functions ---
    print(greet("world"))
    print("3 + 7 =", add(3, 7))
    log_message("server started")
    print(describe("Alice", 30))

    # --- Functions as values ---
    double = lambda x: x * 2
    print("double(5):", double(5))

    # --- Passing a function as an argument ---
    result = apply(10, double)
    print("apply(10, double):", result)

    # --- Inline anonymous function (lambda) ---
    print("inline:", apply(4, lambda x: x * x))

    # --- Pass-by-value for immutables ---
    original = 42
    no_change(original)
    print("after no_change, original:", original)  # still 42


if __name__ == "__main__":
    main()
