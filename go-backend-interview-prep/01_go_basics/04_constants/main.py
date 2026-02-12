# Python equivalent of Constants (compare with main.go)
# Python has no true const keyword. Convention: UPPER_CASE names.

from enum import IntEnum

# "Constants" by convention (nothing prevents reassignment)
APP_VERSION = "1.0.0"
MAX_CONNECTIONS = 100


# Go iota equivalent: use IntEnum
class Status(IntEnum):
    PENDING = 0
    ACTIVE = 1
    INACTIVE = 2
    DELETED = 3


# Bit-shift constants
KB = 1 << 10
MB = 1 << 20
GB = 1 << 30


# Another enum (iota resets in Go; separate class in Python)
class Color(IntEnum):
    RED = 0
    GREEN = 1
    BLUE = 2


def main():
    # --- Simple constants ---
    print("version:", APP_VERSION)
    print("max connections:", MAX_CONNECTIONS)

    # --- Enum constants ---
    print("--- Status Enum ---")
    print("Pending:", Status.PENDING)
    print("Active:", Status.ACTIVE)
    print("Inactive:", Status.INACTIVE)
    print("Deleted:", Status.DELETED)

    # --- Bit-shift constants ---
    print("--- Sizes ---")
    print("KB:", KB)
    print("MB:", MB)
    print("GB:", GB)

    # --- Color enum ---
    print("--- Colors ---")
    print("Red:", Color.RED)
    print("Green:", Color.GREEN)
    print("Blue:", Color.BLUE)

    # --- Python has no untyped constant flexibility ---
    # In Go, untyped const adapts to context. Python just has one float type.
    PI = 3.14159
    print("PI:", PI)


if __name__ == "__main__":
    main()
