# Python equivalent of Defer Deep Dive (compare with main.go)
# Python has no defer. Use try/finally or context managers.


def main():
    # --- try/finally = basic defer ---
    print("--- try/finally ---")
    x = 10
    try:
        x = 999
        print("current x:", x)
    finally:
        # Unlike Go defer, this sees the CURRENT value of x
        print("finally x:", x)  # 999 (Go defer would print 10)

    # --- Context manager (Pythonic defer for resources) ---
    print("\n--- Context manager ---")
    # with open("file.txt") as f:
    #     data = f.read()
    # # f is automatically closed here

    # --- Loop cleanup: context manager per iteration ---
    print("\n--- Loop cleanup ---")
    for i in range(3):
        process_item(i)  # each call has its own try/finally

    # --- No named return + defer pattern in Python ---
    print("\n--- Return enrichment ---")
    # Python has no direct equivalent of Go's named return + defer.
    # Use a wrapper or try/except/finally instead.
    result = enriched_return()
    print("result:", result)

    # --- Key difference ---
    print("\nKey: Go defer captures arg values at defer-time.")
    print("     Python finally sees current values at finally-time.")


def process_item(i: int) -> None:
    try:
        print(f"  processing {i}")
    finally:
        print(f"  cleanup after iteration {i}")


def enriched_return() -> str:
    """Closest equivalent to Go named return + defer."""
    result = "original"
    try:
        return result
    finally:
        # In Python, finally runs but cannot change the return value
        # (unlike Go deferred closures modifying named returns)
        pass
    # To modify return, restructure the code differently


if __name__ == "__main__":
    main()
