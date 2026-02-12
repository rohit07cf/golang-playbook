# Python equivalent of Errors as Interfaces (compare with main.go)
# Go error interface ~ Python Exception hierarchy
# Go errors.Is ~ Python isinstance or direct comparison
# Go errors.As ~ Python except with specific exception type
# Go error wrapping ~ Python exception chaining (raise ... from ...)


# --- Custom exception types ---
class NotFoundError(Exception):
    """Sentinel-style exception."""
    pass


class ValidationError(Exception):
    """Structured exception with extra fields."""
    def __init__(self, field: str, message: str):
        self.field = field
        self.message = message
        super().__init__(f"validation error: {field} - {message}")


def find_user(id: int) -> str:
    """Raises instead of returning (value, error) tuple."""
    if id <= 0:
        raise NotFoundError("not found")
    if id > 100:
        raise ValidationError(field="id", message="must be <= 100")
    return "Alice"


def fetch_profile(id: int) -> str:
    """Wraps exceptions with chaining (like Go's %w)."""
    try:
        name = find_user(id)
        return f"Profile: {name}"
    except Exception as e:
        raise RuntimeError(f"fetch_profile(id={id})") from e


def main():
    # --- Basic error handling ---
    print("--- Basic error handling ---")
    try:
        name = find_user(1)
        print("found:", name)
    except Exception as e:
        print("error:", e)

    # --- Sentinel-style exception ---
    print("\n--- Sentinel exception ---")
    try:
        find_user(0)
    except NotFoundError:
        print("user not found (specific exception match)")

    # --- Custom exception type ---
    print("\n--- Custom exception type ---")
    try:
        find_user(200)
    except ValidationError as e:
        print(f"validation: field={e.field} msg={e.message}")

    # --- Exception chaining (like Go error wrapping) ---
    print("\n--- Exception chaining ---")
    try:
        fetch_profile(0)
    except RuntimeError as e:
        print("wrapped:", e)
        # __cause__ is the chained exception (like Go's Unwrap)
        if isinstance(e.__cause__, NotFoundError):
            print("found NotFoundError in chain")

    # --- Wrapped custom exception ---
    try:
        fetch_profile(200)
    except RuntimeError as e:
        if isinstance(e.__cause__, ValidationError):
            print("found ValidationError in chain:", e.__cause__.field)


if __name__ == "__main__":
    main()
