# Python equivalent of Variables (compare with main.go)

# Module-level variable (like Go package-level var)
app_name: str = "go-basics"


def main():
    # --- Explicit type annotation (optional in Python) ---
    city: str = "Tokyo"
    print("city:", city)

    # --- Type inferred (normal Python style) ---
    population = 14_000_000
    print("population:", population)

    # --- No special short declaration -- just assign ---
    country = "Japan"
    print("country:", country)

    # --- Multiple assignment ---
    x, y = 10, 20
    print("x:", x, "y:", y)

    a, b = "hello", True
    print("a:", a, "b:", b)

    # --- Default values (Python has None, not zero values) ---
    zero_int = 0       # Go: var zeroInt int -> 0
    zero_str = ""      # Go: var zeroStr string -> ""
    zero_bool = False   # Go: var zeroBool bool -> false
    zero_float = 0.0   # Go: var zeroFloat float64 -> 0.0
    print("--- Default Values ---")
    print(f"int:     {zero_int}")
    print(f"string:  {zero_str!r}")
    print(f"bool:    {zero_bool}")
    print(f"float:   {zero_float}")

    # --- Scoping difference ---
    # Python does NOT have block scoping like Go.
    # Variables inside if/for are visible outside.
    score = 100
    print("outer score:", score)
    if True:
        score = 999  # this MODIFIES the outer variable (no new scope)
        print("inner score:", score)
    print("outer score (changed!):", score)

    # --- Module-level variable ---
    print("app:", app_name)

    # Python does NOT enforce unused-variable errors.
    unused_var = 42  # no error in Python; compile error in Go


if __name__ == "__main__":
    main()
