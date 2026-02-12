"""Hexagonal architecture (ports & adapters) -- Python equivalent.

Core depends on duck-typed ports. Adapters implement ports.
main() wires adapters to core.
"""

from dataclasses import dataclass, field
from datetime import datetime


# ============================================================
# CORE (depends on ports only -- duck typing in Python)
# ============================================================

@dataclass
class Notification:
    id: str
    to: str
    message: str
    sent_at: datetime = field(default_factory=datetime.now)


class NotificationService:
    """Core service -- depends on sender_port and repo_port (duck-typed)."""

    def __init__(self, sender, repo):
        self.sender = sender  # SenderPort
        self.repo = repo      # RepoPort

    def send(self, id: str, to: str, message: str) -> None:
        if not to:
            raise ValueError("recipient required")
        self.sender.send(to, message)
        n = Notification(id=id, to=to, message=message)
        self.repo.save(n)


# ============================================================
# ADAPTERS (implement ports)
# ============================================================

class EmailAdapter:
    """Adapter for SenderPort."""
    def send(self, to: str, message: str) -> None:
        print(f"  [EMAIL ADAPTER] Sent to {to}: {message}")


class MemoryRepoAdapter:
    """Adapter for RepoPort."""
    def __init__(self):
        self._data: dict[str, Notification] = {}

    def save(self, n: Notification) -> None:
        self._data[n.id] = n
        print(f"  [REPO ADAPTER]  Saved notification {n.id}")

    def find_by_id(self, id: str) -> Notification:
        if id not in self._data:
            raise KeyError(f"not found: {id}")
        return self._data[id]


# --- Fakes for testing ---

class FakeSender:
    def __init__(self):
        self.calls: list[str] = []

    def send(self, to: str, message: str) -> None:
        self.calls.append(f"{to}:{message}")


class FakeRepo:
    def __init__(self):
        self.saved: list[Notification] = []

    def save(self, n: Notification) -> None:
        self.saved.append(n)

    def find_by_id(self, id: str) -> Notification:
        for n in self.saved:
            if n.id == id:
                return n
        raise KeyError(f"not found: {id}")


# ============================================================
# MAIN (wires adapters to ports)
# ============================================================

def main() -> None:
    print("=== Hexagonal Architecture (Ports & Adapters) ===")
    print()

    # Production wiring
    print("--- Production: email + memory repo ---")
    svc = NotificationService(EmailAdapter(), MemoryRepoAdapter())
    svc.send("n1", "alice@example.com", "Order confirmed")
    svc.send("n2", "bob@example.com", "Payment received")
    print()

    # Test wiring: fakes
    print("--- Test: fake sender + fake repo ---")
    fake_sender = FakeSender()
    fake_repo = FakeRepo()
    test_svc = NotificationService(fake_sender, fake_repo)
    test_svc.send("t1", "test@test.com", "test message")

    print(f"  Sender calls: {fake_sender.calls}")
    print(f"  Repo saved:   {len(fake_repo.saved)} notification(s)")
    print()

    print("Architecture:")
    print("  Core:     NotificationService (depends on sender + repo duck types)")
    print("  Adapters: EmailAdapter, MemoryRepoAdapter (implement duck types)")
    print("  main():   Wires adapters to core (dependency injection)")
    print("  Tests:    Swap adapters for fakes -- zero core changes")


if __name__ == "__main__":
    main()
