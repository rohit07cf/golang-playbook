# Python equivalent of Structs (compare with main.go)
# Go struct -> Python dataclass (or simple class)

from dataclasses import dataclass


@dataclass
class User:
    name: str = ""
    email: str = ""
    age: int = 0


def new_user(name: str, email: str, age: int) -> User:
    """Factory function (same convention as Go's NewUser)."""
    return User(name=name, email=email, age=age)


def try_to_modify(u: User) -> None:
    # Dataclasses are mutable by default -- this WILL modify the original
    # (unlike Go structs which are value types and get copied)
    u.name = "CHANGED"


def main():
    # --- Create with field names ---
    u1 = User(name="Alice", email="alice@example.com", age=30)
    print("u1:", u1)

    # --- Zero-value equivalent ---
    u2 = User()
    print(f"u2 (default): {u2}")

    # --- Partial initialization ---
    u3 = User(name="Bob")
    print(f"u3 (partial): {u3}")

    # --- Factory function ---
    u4 = new_user("Charlie", "charlie@example.com", 25)
    print("u4:", u4)

    # --- Access and modify ---
    u1.age = 31
    print("u1 after birthday:", u1.age)

    # --- Python dataclasses are REFERENCE types (unlike Go structs!) ---
    print("\n--- Reference semantics ---")
    original = User(name="Dave", age=40)
    copied = original            # same object!
    copied.name = "Modified"
    print("original:", original.name)  # "Modified" -- CHANGED (unlike Go)

    # --- To copy, use dataclasses.replace or copy ---
    from copy import copy
    original2 = User(name="Dave", age=40)
    copied2 = copy(original2)
    copied2.name = "Modified"
    print("original2:", original2.name)  # "Dave" -- unchanged
    print("copied2:", copied2.name)

    # --- Passing to function mutates original ---
    print("\n--- Pass by reference ---")
    test_user = User(name="Eve")
    try_to_modify(test_user)
    print("after function:", test_user.name)  # "CHANGED" (unlike Go)

    # --- Comparison ---
    print("\n--- Comparison ---")
    a = User(name="Eve", age=28)
    b = User(name="Eve", age=28)
    print("a == b:", a == b)  # True (dataclass __eq__)


if __name__ == "__main__":
    main()
