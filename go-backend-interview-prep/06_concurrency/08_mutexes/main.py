"""Mutexes -- Python equivalent of the Go example.

Python's GIL protects some built-in operations, but compound
operations (read-modify-write) still need threading.Lock.
"""

import threading


# --- Example 1: safe counter with Lock ---

class SafeCounter:
    def __init__(self) -> None:
        self._lock = threading.Lock()
        self._value = 0

    def inc(self) -> None:
        with self._lock:  # auto lock/unlock (like defer Unlock)
            self._value += 1

    def value(self) -> int:
        with self._lock:
            return self._value


# --- Example 2: read-heavy cache with RLock ---
# Python has no RWLock in stdlib. threading.RLock is reentrant,
# not a read-write lock. We use a regular Lock here.

class Cache:
    def __init__(self) -> None:
        self._lock = threading.Lock()
        self._data: dict[str, str] = {}

    def get(self, key: str) -> str | None:
        with self._lock:
            return self._data.get(key)

    def set(self, key: str, val: str) -> None:
        with self._lock:
            self._data[key] = val


def main() -> None:
    # --- Safe counter ---
    print("=== Safe Counter (Lock) ===")
    counter = SafeCounter()
    threads = []

    for _ in range(1000):
        t = threading.Thread(target=counter.inc)
        threads.append(t)
        t.start()

    for t in threads:
        t.join()
    print(f"  counter: {counter.value()}")  # always 1000

    # --- Cache ---
    print("\n=== Cache (Lock) ===")
    cache = Cache()
    cache.set("name", "alice")
    cache.set("role", "engineer")

    threads = []
    for i in range(5):
        def reader(rid):
            val = cache.get("name")
            if val:
                print(f"  reader {rid}: name={val}")

        t = threading.Thread(target=reader, args=(i,))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()

    cache.set("name", "bob")
    print("  after write: name =", cache.get("name"))


if __name__ == "__main__":
    main()
