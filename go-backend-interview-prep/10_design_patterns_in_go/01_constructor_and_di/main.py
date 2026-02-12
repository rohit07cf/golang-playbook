"""Constructor and dependency injection -- Python equivalent of the Go example.

Uses duck typing for dependencies (no explicit interface needed).
"""


# --- Concrete implementations ---

class EmailSender:
    def send(self, to: str, message: str) -> None:
        print(f"  [EMAIL] To: {to} | Message: {message}")


class SMSSender:
    def send(self, to: str, message: str) -> None:
        print(f"  [SMS]   To: {to} | Message: {message}")


class ConsoleLogger:
    def log(self, msg: str) -> None:
        print(f"  [LOG]   {msg}")


# --- Service (depends on duck-typed sender and logger) ---

class NotificationService:
    def __init__(self, sender, logger):
        """Constructor injection: dependencies come in as arguments."""
        self.sender = sender
        self.logger = logger

    def notify(self, to: str, message: str) -> None:
        self.logger.log(f"sending notification to {to}")
        self.sender.send(to, message)


# --- Fake for testing ---

class FakeSender:
    def __init__(self):
        self.calls: list[str] = []

    def send(self, to: str, message: str) -> None:
        self.calls.append(f"{to}:{message}")


def main() -> None:
    print("=== Constructor and Dependency Injection ===")
    print()

    logger = ConsoleLogger()

    # Wire with email sender
    print("--- Email notification ---")
    email_svc = NotificationService(EmailSender(), logger)
    email_svc.notify("alice@example.com", "Your order shipped!")
    print()

    # Wire with SMS sender
    print("--- SMS notification ---")
    sms_svc = NotificationService(SMSSender(), logger)
    sms_svc.notify("+1234567890", "Your order shipped!")
    print()

    # Wire with fake for testing
    print("--- Test with fake ---")
    fake = FakeSender()
    test_svc = NotificationService(fake, logger)
    test_svc.notify("bob@test.com", "test message")
    print(f"  Fake recorded {len(fake.calls)} call(s): {fake.calls}")
    print()

    print("Key: same service, different senders. DI makes it testable.")


if __name__ == "__main__":
    main()
