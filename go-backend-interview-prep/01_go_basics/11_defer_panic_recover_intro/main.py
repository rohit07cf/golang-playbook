# Python equivalent of Defer, Panic, Recover (compare with main.go)
# Go defer  -> Python try/finally or context managers
# Go panic  -> Python raise
# Go recover -> Python except


def main():
    # --- Defer equivalent: try/finally ---
    print("--- try/finally (like defer) ---")
    print("start")
    try:
        print("end")
    finally:
        # Runs when leaving the block, like defer at function exit
        print("finally block (like deferred)")

    # --- LIFO order: multiple finally blocks via nested try ---
    # In practice, Python uses context managers for cleanup.

    # --- Context manager (Pythonic defer) ---
    print("\n--- Context manager (Pythonic defer) ---")
    # with open("file.txt") as f:
    #     data = f.read()
    # # f.Close() equivalent is automatic

    # --- Cleanup pattern ---
    print("\n--- Cleanup pattern ---")
    do_work()

    # --- Panic + Recover = raise + except ---
    print("\n--- raise + except (like panic/recover) ---")
    safe_call()
    print("program continues after caught exception")


def do_work():
    print("  opening resource")
    try:
        print("  doing work...")
        print("  work done")
    finally:
        print("  closing resource (finally)")


def safe_call():
    try:
        print("  about to raise")
        raise RuntimeError("something broke")
        # Code below never executes
    except RuntimeError as e:
        print("  recovered from exception:", e)


if __name__ == "__main__":
    main()
