"""Defer for cleanup patterns -- Python equivalent of the Go example."""

import threading


# --- Pattern 1: LIFO order (simulated with try/finally nesting) ---

def demo_lifo() -> None:
    print("=== LIFO order (try/finally nesting) ===")
    try:
        try:
            try:
                print("  function body")
            finally:
                print("  cleanup 3 (innermost finally, runs first)")
        finally:
            print("  cleanup 2")
    finally:
        print("  cleanup 1 (outermost finally, runs last)")


# --- Pattern 2: lock unlock (with statement) ---

class SafeCounter:
    def __init__(self) -> None:
        self._lock = threading.Lock()
        self._count = 0

    def increment(self) -> None:
        with self._lock:  # auto-releases lock (like defer mu.Unlock())
            self._count += 1

    def value(self) -> int:
        with self._lock:
            return self._count


# --- Pattern 3: args evaluated eagerly (Python closures capture by reference) ---

def demo_arg_eval() -> None:
    print("\n=== Python closures capture by reference ===")
    x = 10
    # To capture current value, use default arg
    try:
        x = 20
        print(f"  current x = {x}")
    finally:
        # In Python, finally sees the latest value (unlike Go defer)
        print(f"  finally x = {x} (Python captures by reference, not value)")


# --- Pattern 4: recover from exception ---

def risky_operation() -> None:
    raise RuntimeError("something went horribly wrong")


def safe_wrapper() -> None:
    try:
        risky_operation()
    except RuntimeError as e:
        print(f"  recovered: {e}")


# --- Pattern 5: context manager for resource cleanup ---

class Resource:
    def __enter__(self):
        print("  opened resource")
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        print("  closed resource (via __exit__)")
        return False

    def read(self) -> str:
        return "hello from resource"


def read_data() -> str:
    print("\n=== Context manager (with statement) ===")
    with Resource() as r:
        return r.read()


def main() -> None:
    demo_lifo()

    # Lock pattern
    print("\n=== Lock unlock ===")
    counter = SafeCounter()
    for _ in range(5):
        counter.increment()
    print(f"  counter: {counter.value()}")

    demo_arg_eval()

    # Recover pattern
    print("\n=== Recover from exception ===")
    safe_wrapper()
    print("  program continues after recovery")

    # Context manager pattern
    data = read_data()
    print(f"  data={data!r} err=None")


if __name__ == "__main__":
    main()
