# Python equivalent of Struct Embedding (compare with main.go)
# Go embedding = composition. Python uses inheritance.
# They look similar but have different semantics.


class Base:
    def __init__(self, id: int = 0):
        self.id = id

    def describe(self) -> str:
        return f"Base(id={self.id})"


class Logger:
    def __init__(self, prefix: str = ""):
        self.prefix = prefix

    def log(self, msg: str) -> None:
        print(f"[{self.prefix}] {msg}")


# --- Python inheritance (similar to Go embedding) ---
class User(Base):
    def __init__(self, id: int, name: str):
        super().__init__(id)
        self.name = name


# --- Multiple inheritance (similar to embedding multiple structs) ---
class Admin(User, Logger):
    def __init__(self, id: int, name: str, prefix: str, level: str):
        User.__init__(self, id, name)
        Logger.__init__(self, prefix)
        self.level = level


# --- Shadowing (override attribute) ---
class Employee(Base):
    def __init__(self, base_id: int, emp_id: str, name: str):
        super().__init__(base_id)
        self.emp_id = emp_id  # separate name to avoid confusion
        self.name = name


def main():
    # --- Basic inheritance ---
    print("--- Basic inheritance ---")
    u = User(id=1, name="Alice")
    print("u.id:", u.id)
    print("u.describe():", u.describe())  # inherited method
    print("u.name:", u.name)

    # --- Multiple inheritance ---
    print("\n--- Multiple inheritance ---")
    a = Admin(id=42, name="Bob", prefix="ADMIN", level="super")
    print("a.id:", a.id)
    print("a.name:", a.name)
    print("a.describe():", a.describe())
    a.log("system started")

    # --- Key difference ---
    print("\n--- Key difference ---")
    # In Python, describe() receives `self` (the Admin instance).
    # In Go, Describe() receives the embedded Base value.
    # Python has true inheritance; Go has composition.
    print("Python: describe() sees full object (inheritance)")
    print("Go: Describe() sees only Base (composition)")

    # --- Employee ---
    print("\n--- Employee ---")
    e = Employee(base_id=100, emp_id="EMP-001", name="Charlie")
    print("e.id:", e.id)          # from Base
    print("e.emp_id:", e.emp_id)  # separate field


if __name__ == "__main__":
    main()
