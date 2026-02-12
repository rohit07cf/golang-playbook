# Python equivalent of Init Functions (compare with main.go)
# Python has no init() auto-function.
# Module-level code runs at import time (similar concept).

# --- Module-level code runs when file is imported or executed ---
print("1. module-level code (like Go package var init)")
config = "production"


def setup():
    """Explicit init function (Python convention)."""
    print("2. setup() called (explicit, unlike Go's auto init)")
    print("   config =", config)


def main():
    print("3. main() called")
    print("   config =", config)

    print("\nPython: module-level code runs at import/execution time.")
    print("Go: package vars init, then init(), then main().")
    print("Python has no auto init() -- use explicit setup functions.")


# This runs when the file is executed directly
# (or when imported as a module, the module-level code above runs)
if __name__ == "__main__":
    setup()  # must be called explicitly (Go's init is automatic)
    main()
