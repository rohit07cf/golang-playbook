"""Singleton and why to avoid it -- Python equivalent of the Go example.

Shows module-level singleton vs dependency injection alternative.
"""


# ============================================================
# APPROACH 1: Singleton (the problem)
# ============================================================

_config = None


def get_config():
    """Module-level singleton -- returns same dict every time."""
    global _config
    if _config is None:
        print("  [SINGLETON] Loading config (runs once)...")
        _config = {"smtp_host": "smtp.example.com", "smtp_port": 587}
    return _config


class SingletonService:
    """Uses the singleton directly -- hidden dependency!"""

    def send_notification(self, to: str, msg: str) -> None:
        cfg = get_config()  # hidden: caller can't see this dependency
        print(f"  [SINGLETON] Send via {cfg['smtp_host']}:{cfg['smtp_port']} to={to} msg={msg}")


# ============================================================
# APPROACH 2: Dependency Injection (the solution)
# ============================================================

class EmailSender:
    def __init__(self, host: str, port: int):
        self.host = host
        self.port = port

    def send(self, to: str, message: str) -> None:
        print(f"  [DI]        Send via {self.host}:{self.port} to={to} msg={message}")


class NotificationService:
    def __init__(self, sender):
        self.sender = sender  # explicit dependency

    def notify(self, to: str, msg: str) -> None:
        self.sender.send(to, msg)


class FakeSender:
    def __init__(self):
        self.calls: list[str] = []

    def send(self, to: str, message: str) -> None:
        self.calls.append(f"{to}:{message}")


def main() -> None:
    print("=== Singleton and Why to Avoid It ===")
    print()

    # Singleton approach
    print("--- Singleton (hidden dependency) ---")
    svc1 = SingletonService()
    svc1.send_notification("alice@example.com", "Order shipped!")
    svc1.send_notification("bob@example.com", "Payment received")
    print()

    print("  Problem: SingletonService hides its config dependency.")
    print("  In tests, you can't swap config without global mutation.")
    print()

    # DI approach
    print("--- DI alternative (explicit dependency) ---")
    sender = EmailSender("smtp.example.com", 587)
    svc2 = NotificationService(sender)
    svc2.notify("alice@example.com", "Order shipped!")
    svc2.notify("bob@example.com", "Payment received")
    print()

    # In tests: swap with fake
    print("--- In tests: swap with fake ---")
    fake = FakeSender()
    test_svc = NotificationService(fake)
    test_svc.notify("test@example.com", "test message")
    print(f"  Fake recorded: {fake.calls}")
    print()

    print("Key: DI makes dependencies visible and testable.")
    print("     Singleton hides them. Prefer DI in almost all cases.")


if __name__ == "__main__":
    main()
