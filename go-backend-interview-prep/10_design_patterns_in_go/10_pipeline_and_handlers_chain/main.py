"""Pipeline and handlers chain -- Python equivalent of the Go example.

Message flows through ordered handlers: validate -> enrich -> log -> send.
"""

from datetime import datetime


class Message:
    def __init__(self, to: str, body: str):
        self.to = to
        self.body = body
        self.timestamp: datetime | None = None
        self.metadata: dict[str, str] = {}


class Pipeline:
    def __init__(self, *handlers):
        self.handlers = handlers

    def run(self, msg: Message) -> str | None:
        """Returns error string on failure, None on success."""
        for h in self.handlers:
            err = h.handle(msg)
            if err:
                return err  # short-circuit
        return None


# --- Concrete handlers ---

class ValidateHandler:
    def handle(self, msg: Message) -> str | None:
        print("  [VALIDATE] checking message...")
        if not msg.to:
            return "validation: 'to' is required"
        if not msg.body:
            return "validation: 'body' is required"
        if "@" not in msg.to:
            return "validation: 'to' must be an email"
        print("  [VALIDATE] OK")
        return None


class EnrichHandler:
    def handle(self, msg: Message) -> str | None:
        msg.timestamp = datetime.now()
        msg.metadata["enriched"] = "true"
        msg.metadata["priority"] = "normal"
        print("  [ENRICH]   Added timestamp + metadata")
        return None


class LogHandler:
    def handle(self, msg: Message) -> str | None:
        print(f"  [LOG]      to={msg.to} body={msg.body!r} metadata={msg.metadata}")
        return None


class SendHandler:
    def handle(self, msg: Message) -> str | None:
        ts = msg.timestamp.strftime("%H:%M:%S") if msg.timestamp else "?"
        print(f"  [SEND]     Delivered to {msg.to} at {ts}")
        return None


def main() -> None:
    print("=== Pipeline and Handlers Chain ===")
    print()

    pipeline = Pipeline(
        ValidateHandler(),
        EnrichHandler(),
        LogHandler(),
        SendHandler(),
    )

    # Valid message
    print("--- Valid message ---")
    err = pipeline.run(Message("alice@example.com", "Your order shipped!"))
    if err:
        print(f"  PIPELINE ERROR: {err}")
    print()

    # Invalid: no @
    print("--- Invalid message (no @) ---")
    err = pipeline.run(Message("bob", "test"))
    if err:
        print(f"  PIPELINE ERROR: {err}")
    print()

    # Invalid: empty body
    print("--- Invalid message (empty body) ---")
    err = pipeline.run(Message("charlie@example.com", ""))
    if err:
        print(f"  PIPELINE ERROR: {err}")
    print()

    print("Key: validate -> enrich -> log -> send. Error = short-circuit.")
    print("     Each handler does one thing. Add/remove steps freely.")


if __name__ == "__main__":
    main()
