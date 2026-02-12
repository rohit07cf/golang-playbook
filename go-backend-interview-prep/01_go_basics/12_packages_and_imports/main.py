# Python equivalent of Packages and Imports (compare with main.go)
# Python uses modules and packages (directories with __init__.py).

import math
import string

def main():
    # --- Using the math module ---
    print("--- math module ---")
    print("pi:", math.pi)
    print("sqrt(16):", math.sqrt(16))

    # --- Using the string module ---
    print("\n--- string module ---")
    msg = "Hello, Python World"
    print("upper:", msg.upper())       # method on str, not a module function
    print("lower:", msg.lower())
    print("contains 'Python':", "Python" in msg)
    print("replace:", msg.replace("Python", "Pythonista"))
    print("split:", msg.split(" "))

    # --- Visibility rules ---
    # Python convention: _prefix means private, no prefix means public.
    # Go: uppercase = exported, lowercase = unexported.
    print("\n--- Visibility ---")
    print("public function:", exported_greet("Alice"))
    print("private function:", _unexported_helper())

    # --- Module name = file name ---
    # Go: package name = last segment of import path
    # Python: module name = file name (without .py)
    print("\nPython module names = file names (not path segments).")


# Public function (no underscore prefix)
def exported_greet(name: str) -> str:
    return "Hello, " + name


# Private function (underscore prefix, convention only)
def _unexported_helper() -> str:
    return "I am module-private (by convention)"


if __name__ == "__main__":
    main()
