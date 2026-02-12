"""Builder vs functional options -- Python equivalent of the Go example.

Python's **kwargs with defaults solves the same problem natively.
"""

import time


class PrintSender:
    def send(self, to: str, message: str) -> None:
        print(f"  [SEND] {to}: {message}")


# ============================================================
# APPROACH 1: kwargs with defaults (Python-idiomatic)
# ============================================================

class NotificationService:
    """Python solves 'many optional params' with kwargs + defaults."""

    def __init__(
        self,
        sender,
        timeout: float = 5.0,
        retries: int = 1,
        verbose: bool = False,
    ):
        self.sender = sender
        self.timeout = timeout
        self.retries = retries
        self.verbose = verbose

    def notify(self, to: str, msg: str) -> None:
        if self.verbose:
            print(f"  [VERBOSE] timeout={self.timeout}s retries={self.retries}")
        print(f"  Sending to {to}: {msg}")


# ============================================================
# APPROACH 2: Builder (verbose in Python, shown for comparison)
# ============================================================

class ServiceBuilder:
    def __init__(self, sender):
        self._sender = sender
        self._timeout = 5.0
        self._retries = 1
        self._verbose = False

    def timeout(self, t: float):
        self._timeout = t
        return self

    def retries(self, n: int):
        self._retries = n
        return self

    def verbose(self):
        self._verbose = True
        return self

    def build(self) -> NotificationService:
        return NotificationService(
            self._sender,
            timeout=self._timeout,
            retries=self._retries,
            verbose=self._verbose,
        )


def main() -> None:
    print("=== Builder vs Functional Options ===")
    print()
    sender = PrintSender()

    # kwargs with defaults (natural)
    print("--- kwargs with defaults (Python-idiomatic) ---")

    svc1 = NotificationService(sender)  # all defaults
    print("  defaults: ", end="")
    svc1.notify("alice", "hello")

    svc2 = NotificationService(sender, timeout=10, retries=3, verbose=True)
    print("  custom:   ", end="")
    svc2.notify("bob", "world")
    print()

    # Builder (for comparison)
    print("--- Builder Pattern (verbose) ---")
    svc3 = (ServiceBuilder(sender)
            .timeout(10)
            .retries(3)
            .verbose()
            .build())
    print("  builder:  ", end="")
    svc3.notify("charlie", "hi")
    print()

    print("Key: Python kwargs with defaults = Go functional options.")
    print("     Both solve 'constructor with many optional params'.")
    print("     Builder is redundant in Python but shown for comparison.")


if __name__ == "__main__":
    main()
