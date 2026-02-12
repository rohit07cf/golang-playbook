"""Sentinel errors vs typed errors -- Python equivalent of the Go example."""


# ---- Sentinel-style exceptions (no fields) ----

class NotFoundError(Exception):
    pass


class UnauthorizedError(Exception):
    pass


# ---- Typed exception (has fields) ----

class RateLimitError(Exception):
    def __init__(self, limit: int, retry_after: int):
        self.limit = limit
        self.retry_after = retry_after
        super().__init__(f"rate limited: max {limit} req/s, retry after {retry_after}s")


# ---- Functions ----

def lookup_user(user_id: int) -> str:
    if user_id == 0:
        raise NotFoundError("not found")
    if user_id == -1:
        raise UnauthorizedError("unauthorized")
    if user_id == 99:
        raise RateLimitError(limit=100, retry_after=30)
    return f"user_{user_id}"


def get_profile(user_id: int) -> str:
    try:
        return lookup_user(user_id)
    except Exception as e:
        raise RuntimeError(f"get_profile({user_id})") from e


def main() -> None:
    ids = [1, 0, -1, 99]

    for uid in ids:
        try:
            name = get_profile(uid)
            print(f"id={uid}  name: {name}")
        except Exception as e:
            print(f"id={uid}  error: {e}")
            cause = e.__cause__

            # Sentinel-style check with isinstance
            if isinstance(cause, NotFoundError):
                print("  -> sentinel: not found")
            elif isinstance(cause, UnauthorizedError):
                print("  -> sentinel: unauthorized")

            # Typed check -- extract fields
            elif isinstance(cause, RateLimitError):
                print(f"  -> typed: limit={cause.limit}, retry_after={cause.retry_after}s")

        print()


if __name__ == "__main__":
    main()
