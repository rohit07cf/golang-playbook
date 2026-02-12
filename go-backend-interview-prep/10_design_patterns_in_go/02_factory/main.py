"""Factory pattern -- Python equivalent of the Go example.

Factory function returns a duck-typed sender based on config string.
"""


class EmailSender:
    kind = "email"

    def send(self, to: str, message: str) -> None:
        print(f"  [EMAIL] To: {to} | {message}")


class SMSSender:
    kind = "sms"

    def send(self, to: str, message: str) -> None:
        print(f"  [SMS]   To: {to} | {message}")


class SlackSender:
    kind = "slack"

    def __init__(self, webhook: str = "https://hooks.slack.example/abc"):
        self.webhook = webhook

    def send(self, to: str, message: str) -> None:
        print(f"  [SLACK] Channel: {to} | {message} (webhook: {self.webhook})")


# --- Factory ---

_SENDER_REGISTRY = {
    "email": EmailSender,
    "sms": SMSSender,
    "slack": SlackSender,
}


def new_sender(kind: str):
    """Factory: returns a sender based on config string."""
    cls = _SENDER_REGISTRY.get(kind)
    if cls is None:
        raise ValueError(f"unknown sender type: {kind}")
    return cls()


def main() -> None:
    print("=== Factory Pattern ===")
    print()

    configs = ["email", "sms", "slack", "pigeon"]

    for kind in configs:
        try:
            sender = new_sender(kind)
        except ValueError as e:
            print(f"  factory error: {e}")
            continue

        print(f"  Created sender: kind={sender.kind}")
        sender.send("alice", "Your order shipped!")
        print()

    print("Key: factory returns a duck-typed object. Callers use .send() uniformly.")
    print("     Adding a new sender = register in the dict. Zero caller changes.")


if __name__ == "__main__":
    main()
