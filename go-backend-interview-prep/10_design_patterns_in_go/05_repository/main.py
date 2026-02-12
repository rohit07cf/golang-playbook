"""Repository pattern -- Python equivalent of the Go example.

Dict-backed repository implementing save/find/list.
"""

from dataclasses import dataclass, field
from datetime import datetime


@dataclass
class Notification:
    id: str
    to: str
    message: str
    sent_at: datetime = field(default_factory=datetime.now)


class MemoryRepo:
    """In-memory repository -- swap for SQLRepo in production."""

    def __init__(self):
        self._data: dict[str, Notification] = {}

    def save(self, n: Notification) -> None:
        self._data[n.id] = n

    def find_by_id(self, id: str) -> Notification:
        if id not in self._data:
            raise KeyError(f"not found: {id}")
        return self._data[id]

    def list_all(self) -> list[Notification]:
        return list(self._data.values())


class NotificationService:
    """Service depends on repo (duck-typed), not on storage details."""

    def __init__(self, repo):
        self.repo = repo

    def send(self, id: str, to: str, message: str) -> None:
        n = Notification(id=id, to=to, message=message)
        print(f"  Sending: id={id} to={to} msg={message}")
        self.repo.save(n)

    def get_history(self) -> list[Notification]:
        return self.repo.list_all()


def main() -> None:
    print("=== Repository Pattern ===")
    print()

    repo = MemoryRepo()
    svc = NotificationService(repo)

    svc.send("n1", "alice@example.com", "Order confirmed")
    svc.send("n2", "bob@example.com", "Payment received")
    svc.send("n3", "alice@example.com", "Order shipped")
    print()

    # Find by ID
    print("--- Find by ID ---")
    n = repo.find_by_id("n2")
    print(f"  Found: id={n.id} to={n.to} msg={n.message}")

    try:
        repo.find_by_id("n99")
    except KeyError as e:
        print(f"  Not found: {e}")
    print()

    # List all
    print("--- List all ---")
    for n in svc.get_history():
        print(f"  {n.id} -> {n.to}: {n.message}")
    print()

    print("Key: service calls repo.save() / repo.find_by_id().")
    print("     Swap MemoryRepo for SQLRepo -- zero service changes.")


if __name__ == "__main__":
    main()
