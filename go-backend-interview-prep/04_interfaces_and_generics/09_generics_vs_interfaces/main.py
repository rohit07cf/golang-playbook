# Python equivalent of Generics vs Interfaces (compare with main.go)
# In Python, duck typing and ABC/Protocol fill the "interface" role.
# typing module fills the "generics" role (but is not enforced at runtime).

from typing import TypeVar, Protocol

T = TypeVar("T")


# ============================================================
# APPROACH 1: PROTOCOL (structural typing, like Go interface)
# ============================================================

class Formatter(Protocol):
    def format(self) -> str: ...


class PlainText:
    def __init__(self, content: str):
        self.content = content

    def format(self) -> str:
        return self.content


class HTMLText:
    def __init__(self, content: str):
        self.content = content

    def format(self) -> str:
        return f"<p>{self.content}</p>"


def print_formatted(f: Formatter) -> None:
    print(" ", f.format())


# ============================================================
# APPROACH 2: GENERIC-STYLE (same logic, different types)
# ============================================================

def reverse(s: list[T]) -> list[T]:
    return s[::-1]


def unique(s: list) -> list:
    seen = set()
    result = []
    for v in s:
        if v not in seen:
            seen.add(v)
            result.append(v)
    return result


# ============================================================
# PROTOCOL FOR SERVICE ABSTRACTION
# ============================================================

class UserStore(Protocol):
    def get_user(self, id: int) -> str: ...


class DBStore:
    def get_user(self, id: int) -> str:
        return f"User-{id} (from DB)"


class MockStore:
    def get_user(self, id: int) -> str:
        return f"MockUser-{id}"


def greet_user(store: UserStore, id: int) -> None:
    print("  Hello,", store.get_user(id))


def main():
    # --- Protocol approach: behavior polymorphism ---
    print("=== PROTOCOL: different behavior ===")
    print_formatted(PlainText("hello world"))
    print_formatted(HTMLText("hello world"))

    # --- Generic-style: same logic, different types ---
    print("\n=== GENERIC: same logic, different types ===")
    print("reversed ints:", reverse([1, 2, 3, 4, 5]))
    print("reversed strs:", reverse(["a", "b", "c"]))

    print("unique ints:", unique([1, 2, 2, 3, 3, 3]))
    print("unique strs:", unique(["go", "go", "rust"]))

    # --- Protocol for service abstraction ---
    print("\n=== PROTOCOL: service abstraction ===")
    greet_user(DBStore(), 1)
    greet_user(MockStore(), 1)

    # --- Decision framework ---
    print("\n=== WHEN TO USE WHICH ===")
    decisions = [
        "Protocol: behavior differs (Formatter, Store)",
        "Generics: logic identical, types differ (reverse, unique)",
        "Protocol: dependency injection, mocking",
        "Generics: type-safe containers, algorithms",
        "Default:  Python duck typing handles most cases naturally",
    ]
    for d in decisions:
        print(" ", d)


if __name__ == "__main__":
    main()
