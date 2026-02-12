# Python equivalent of Interface Composition (compare with main.go)
# Go interface embedding ~ Python multiple Protocol inheritance

from typing import Protocol


# --- Small, focused protocols ---
class Saver(Protocol):
    def save(self, data: str) -> None: ...

class Loader(Protocol):
    def load(self, id: str) -> str | None: ...

class Deleter(Protocol):
    def delete(self, id: str) -> None: ...


# --- Composed protocol (multiple inheritance) ---
class Storage(Saver, Loader, Deleter, Protocol):
    ...


# --- Concrete type satisfying all ---
class MemoryStore:
    def __init__(self):
        self.data: dict[str, str] = {}

    def save(self, data: str) -> None:
        key = f"key-{len(self.data) + 1}"
        self.data[key] = data
        print(f"  saved: {key} -> {data!r}")

    def load(self, id: str) -> str | None:
        return self.data.get(id)

    def delete(self, id: str) -> None:
        self.data.pop(id, None)
        print(f"  deleted: {id}")


def save_record(s: Saver, data: str) -> None:
    s.save(data)


def load_record(l: Loader, id: str) -> None:
    val = l.load(id)
    if val is None:
        print(f"  error: not found: {id}")
    else:
        print(f"  loaded: {id} -> {val!r}")


def main():
    store = MemoryStore()

    # --- Use as composed interface ---
    print("--- Full Storage interface ---")
    store.save("hello")
    store.save("world")

    # --- Use as Loader only ---
    print("\n--- As Loader only ---")
    load_record(store, "key-1")
    load_record(store, "key-99")

    # --- Use as Saver only ---
    print("\n--- As Saver only ---")
    save_record(store, "another")

    # --- Use as Deleter ---
    print("\n--- As Deleter ---")
    store.delete("key-1")

    print("\n--- Satisfies all sub-protocols ---")
    print("MemoryStore satisfies Saver, Loader, Deleter, and Storage")


if __name__ == "__main__":
    main()
