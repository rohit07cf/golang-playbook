"""Mocking and dependency injection intro -- Python equivalent of the Go example."""

from typing import Protocol


# ---- Interface (Protocol): the dependency ----

class UserStore(Protocol):
    def get_user(self, user_id: int) -> str: ...


# ---- Real implementation ----

class DBStore:
    def get_user(self, user_id: int) -> str:
        users = {1: "alice", 2: "bob"}
        if user_id not in users:
            raise KeyError(f"user not found: {user_id}")
        return users[user_id]


# ---- Fake implementation for testing ----

class FakeStore:
    def __init__(self, users: dict[int, str] | None = None, err: Exception | None = None):
        self.users = users or {}
        self.err = err

    def get_user(self, user_id: int) -> str:
        if self.err:
            raise self.err
        if user_id not in self.users:
            raise KeyError(f"user not found: {user_id}")
        return self.users[user_id]


# ---- Service that depends on the interface ----

class UserService:
    def __init__(self, store: UserStore):
        self.store = store

    def greet(self, user_id: int) -> str:
        try:
            name = self.store.get_user(user_id)
        except Exception as e:
            raise RuntimeError(f"greet({user_id})") from e
        return f"Hello, {name}!"


def main() -> None:
    # --- Real store ---
    print("=== Real store ===")
    svc = UserService(store=DBStore())
    try:
        print(f"greet(1): {svc.greet(1)!r}, err=None")
    except RuntimeError as e:
        print(f"greet(1): err={e}")

    try:
        print(f"greet(999): {svc.greet(999)!r}, err=None")
    except RuntimeError as e:
        print(f"greet(999): err={e}")

    # --- Fake store (test mode) ---
    print("\n=== Fake store (test mode) ===")
    fake = FakeStore(users={10: "test_user"})
    test_svc = UserService(store=fake)
    print(f"greet(10): {test_svc.greet(10)!r}, err=None")

    # --- Fake store (error mode) ---
    print("\n=== Fake store (error mode) ===")
    err_fake = FakeStore(err=RuntimeError("database down"))
    err_svc = UserService(store=err_fake)
    try:
        err_svc.greet(1)
    except RuntimeError as e:
        print(f"greet(1): err={e}")


if __name__ == "__main__":
    main()
