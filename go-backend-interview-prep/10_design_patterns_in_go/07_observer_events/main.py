"""Observer / events pattern -- Python equivalent of the Go example.

Simple event bus with subscribe/publish.
"""

from dataclasses import dataclass, field


@dataclass
class Event:
    name: str
    data: dict = field(default_factory=dict)


class EventBus:
    def __init__(self):
        self.handlers: dict[str, list] = {}

    def subscribe(self, event_name: str, handler) -> None:
        self.handlers.setdefault(event_name, []).append(handler)

    def publish(self, event: Event) -> None:
        for handler in self.handlers.get(event.name, []):
            handler(event)


class NotificationService:
    def __init__(self, bus: EventBus):
        self.bus = bus

    def send(self, to: str, message: str) -> None:
        print(f"  [SEND] To: {to} | {message}")
        self.bus.publish(Event("notification.sent", {"to": to, "message": message}))

    def fail(self, to: str, reason: str) -> None:
        print(f"  [FAIL] To: {to} | reason: {reason}")
        self.bus.publish(Event("notification.failed", {"to": to, "reason": reason}))


def main() -> None:
    print("=== Observer / Events Pattern ===")
    print()

    bus = EventBus()

    # Subscribe: logger
    bus.subscribe("notification.sent",
                  lambda e: print(f"  [LOG]       Notification sent to {e.data['to']}"))

    # Subscribe: analytics
    bus.subscribe("notification.sent",
                  lambda e: print(f"  [ANALYTICS] Recorded delivery to {e.data['to']}"))

    # Subscribe: alert on failure
    bus.subscribe("notification.failed",
                  lambda e: print(f"  [ALERT]     FAILURE for {e.data['to']}: {e.data['reason']}"))

    svc = NotificationService(bus)

    print("--- Sending notifications ---")
    svc.send("alice@example.com", "Order confirmed")
    print()
    svc.send("bob@example.com", "Payment received")
    print()

    print("--- Simulating failure ---")
    svc.fail("charlie@example.com", "invalid email")
    print()

    print("Key: service publishes events. Subscribers react independently.")
    print("     Adding analytics = one subscribe() call. Zero service changes.")


if __name__ == "__main__":
    main()
