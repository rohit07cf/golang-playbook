"""Strategy pattern -- Python equivalent of the Go example.

Strategies are duck-typed objects with a route() method.
"""


class PriorityRouting:
    """Always picks the first (highest priority) channel."""
    def route(self, message: str, channels: list[str]) -> str:
        return channels[0] if channels else ""


class RoundRobinRouting:
    """Cycles through channels."""
    def __init__(self):
        self.counter = 0

    def route(self, message: str, channels: list[str]) -> str:
        if not channels:
            return ""
        ch = channels[self.counter % len(channels)]
        self.counter += 1
        return ch


class ContentBasedRouting:
    """Picks channel based on message length."""
    def route(self, message: str, channels: list[str]) -> str:
        return "email" if len(message) > 50 else "sms"


class NotificationService:
    def __init__(self, strategy, channels: list[str]):
        self.strategy = strategy
        self.channels = channels

    def send(self, message: str) -> None:
        channel = self.strategy.route(message, self.channels)
        print(f"  Sending via {channel:<8}: {message}")

    def set_strategy(self, strategy) -> None:
        """Swap strategy at runtime."""
        self.strategy = strategy


def main() -> None:
    print("=== Strategy Pattern ===")
    print()

    channels = ["email", "sms", "slack"]

    # Priority strategy
    print("--- Priority routing (always first channel) ---")
    svc = NotificationService(PriorityRouting(), channels)
    svc.send("Order confirmed")
    svc.send("Payment received")
    print()

    # Swap to round-robin
    print("--- Round-robin routing ---")
    svc.set_strategy(RoundRobinRouting())
    svc.send("Order confirmed")
    svc.send("Payment received")
    svc.send("Shipped!")
    svc.send("Delivered!")
    print()

    # Swap to content-based
    print("--- Content-based routing ---")
    svc.set_strategy(ContentBasedRouting())
    svc.send("short msg")
    svc.send("This is a much longer message that should be routed to email because it exceeds fifty characters")
    print()

    print("Key: same service, different routing logic. Swapped at runtime.")


if __name__ == "__main__":
    main()
