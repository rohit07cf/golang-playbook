"""Custom error types -- Python equivalent of the Go example."""


# --- Custom exception: validation ---

class ValidationError(Exception):
    def __init__(self, field: str, message: str):
        self.field = field
        self.message = message
        super().__init__(f"validation: {field} -- {message}")


# --- Custom exception: HTTP-like ---

class HTTPError(Exception):
    def __init__(self, code: int, status: str):
        self.code = code
        self.status = status
        super().__init__(f"HTTP {code}: {status}")


# --- Functions that raise custom exceptions ---

def validate_age(age: int) -> None:
    if age < 0:
        raise ValidationError(field="age", message="must be non-negative")
    if age > 150:
        raise ValidationError(field="age", message="unrealistic value")


def fetch_user(user_id: int) -> None:
    if user_id <= 0:
        raise HTTPError(code=400, status="bad request")
    if user_id == 999:
        raise HTTPError(code=404, status="not found")


def process_request(user_id: int, age: int) -> None:
    try:
        fetch_user(user_id)
    except Exception as e:
        raise RuntimeError("process_request") from e
    try:
        validate_age(age)
    except Exception as e:
        raise RuntimeError("process_request") from e


def main() -> None:
    # --- Check ValidationError with isinstance ---
    try:
        process_request(1, -5)
    except Exception as e:
        print("Error:", e)
        cause = e.__cause__
        if isinstance(cause, ValidationError):
            print(f"  Field: {cause.field}, Message: {cause.message}")

    # --- Check HTTPError with isinstance ---
    try:
        process_request(999, 25)
    except Exception as e:
        print("\nError:", e)
        cause = e.__cause__
        if isinstance(cause, HTTPError):
            print(f"  Code: {cause.code}, Status: {cause.status}")

    # --- Happy path ---
    try:
        process_request(1, 25)
        print("\nNo error: None")
    except Exception as e:
        print("Error:", e)


if __name__ == "__main__":
    main()
