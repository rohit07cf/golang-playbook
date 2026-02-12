"""Wrap errors and unwrap -- Python equivalent of the Go example."""


# --- Custom exception types ---
class NotFoundError(Exception):
    pass


class PermissionError_(Exception):
    """Underscore to avoid shadowing built-in PermissionError."""
    def __init__(self, user: str, action: str):
        self.user = user
        self.action = action
        super().__init__(f"user {user!r} cannot {action}")


# --- Layered functions that chain exceptions ---

def find_record(record_id: int) -> None:
    if record_id <= 0:
        raise NotFoundError("not found")
    if record_id == 99:
        raise PermissionError_(user="guest", action="read")


def get_profile(record_id: int) -> None:
    try:
        find_record(record_id)
    except Exception as e:
        raise RuntimeError(f"get_profile(id={record_id})") from e


def handle_request(record_id: int) -> None:
    try:
        get_profile(record_id)
    except Exception as e:
        raise RuntimeError("handle_request") from e


def main() -> None:
    # --- Example 1: isinstance walks the __cause__ chain ---
    try:
        handle_request(0)
    except Exception as e:
        print("Error:", e)
        # Walk the chain manually
        cause = e.__cause__
        while cause:
            if isinstance(cause, NotFoundError):
                print("Is NotFoundError? True")
                break
            cause = cause.__cause__
        else:
            print("Is NotFoundError? False")

    # --- Example 2: extract typed exception from chain ---
    try:
        handle_request(99)
    except Exception as e:
        print("\nError:", e)
        cause = e.__cause__
        while cause:
            if isinstance(cause, PermissionError_):
                print(f"Permission denied: user={cause.user} action={cause.action}")
                break
            cause = cause.__cause__

    # --- Example 3: raise...from preserves chain ---
    print()
    base = ValueError("disk full")
    try:
        raise RuntimeError("save file") from base
    except RuntimeError as e:
        print("with 'from' preserves cause:", e.__cause__ is base)

    # Without 'from' -- no __cause__
    try:
        raise RuntimeError(f"save file: {base}")
    except RuntimeError as e:
        print("without 'from' loses cause:", e.__cause__ is None)


if __name__ == "__main__":
    main()
