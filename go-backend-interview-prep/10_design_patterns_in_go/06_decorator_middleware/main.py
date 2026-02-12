"""Decorator / middleware pattern -- Python equivalent of the Go example.

Each decorator wraps a sender, adds behavior, delegates to the original.
"""

import time


# --- Base implementation ---

class EmailSender:
    def __init__(self):
        self.fail_count = 0

    def send(self, to: str, message: str) -> None:
        self.fail_count += 1
        if self.fail_count <= 2:
            raise RuntimeError("temporary failure")
        print(f"  [EMAIL] To: {to} | {message}")


class SimpleSender:
    def send(self, to: str, message: str) -> None:
        print(f"  [SEND] {to}: {message}")


# --- Decorator: logging ---

class LoggingSender:
    def __init__(self, next_sender):
        self.next = next_sender

    def send(self, to: str, message: str) -> None:
        print(f"  [LOG] sending to {to} at {time.strftime('%H:%M:%S')}")
        try:
            self.next.send(to, message)
            print("  [LOG] success")
        except Exception as e:
            print(f"  [LOG] error: {e}")
            raise


# --- Decorator: retry ---

class RetrySender:
    def __init__(self, next_sender, max_retries: int = 3):
        self.next = next_sender
        self.max_retries = max_retries

    def send(self, to: str, message: str) -> None:
        last_err = None
        for attempt in range(1, self.max_retries + 1):
            try:
                self.next.send(to, message)
                return
            except Exception as e:
                last_err = e
                print(f"  [RETRY] attempt {attempt}/{self.max_retries} failed: {e}")
        raise RuntimeError(f"all {self.max_retries} retries failed: {last_err}")


# --- Decorator: rate limiter ---

class RateLimitSender:
    def __init__(self, next_sender, delay: float):
        self.next = next_sender
        self.delay = delay
        self.last_sent = 0.0

    def send(self, to: str, message: str) -> None:
        now = time.monotonic()
        elapsed = now - self.last_sent
        if self.last_sent > 0 and elapsed < self.delay:
            wait = self.delay - elapsed
            print(f"  [RATE] throttling {wait*1000:.0f}ms...")
            time.sleep(wait)
        self.last_sent = time.monotonic()
        self.next.send(to, message)


def main() -> None:
    print("=== Decorator / Middleware Pattern ===")
    print()

    # Layer: logging -> retry -> email (fails first 2 times)
    print("--- Logging + Retry (email fails first 2 times) ---")
    base = EmailSender()
    sender = LoggingSender(RetrySender(base, max_retries=3))
    try:
        sender.send("alice@example.com", "Order shipped!")
    except RuntimeError as e:
        print(f"  Final error: {e}")
    print()

    # Rate limited
    print("--- Rate limited sender ---")
    limited = RateLimitSender(LoggingSender(SimpleSender()), delay=0.2)
    for i in range(3):
        limited.send("bob", f"msg-{i}")
    print()

    print("Key: each decorator has .send(), wraps another .send().")
    print("     Compose: LoggingSender(RetrySender(RateLimitSender(base))).")


if __name__ == "__main__":
    main()
