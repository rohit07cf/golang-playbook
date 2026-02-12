# Python equivalent of Pointers (compare with main.go)
# Python has no explicit pointers.
# Mutable objects (list, dict, class instances) use reference semantics.
# Immutable objects (int, str, tuple) behave like pass-by-value.


class Config:
    def __init__(self, host: str = "", port: int = 0, verbose: bool = False):
        self.host = host
        self.port = port
        self.verbose = verbose

    def __repr__(self):
        return f"Config(host={self.host!r}, port={self.port}, verbose={self.verbose})"


def increment_immutable(x: int) -> int:
    """Cannot modify caller's int. Must return new value."""
    return x + 1


def enable_verbose(c: Config) -> None:
    """Mutable object: modifies the original (like Go pointer)."""
    c.verbose = True


def no_change(n: int) -> None:
    n = 999  # rebinds local name; original unchanged


def create_config() -> Config:
    """Returning a new object (Python manages memory automatically)."""
    return Config(host="api.example.com", port=443)


def main():
    # --- No explicit & and * in Python ---
    # Python uses id() to see "address"
    x = 42
    print("x:", x)
    print("id(x):", id(x))

    # --- Mutating requires reassignment for immutables ---
    x = increment_immutable(x)
    print("x after increment:", x)

    # --- Mutable objects behave like pointers ---
    print("\n--- Mutable object (like pointer) ---")
    cfg = Config(host="localhost", port=8080)
    print("cfg.host:", cfg.host)
    cfg.port = 9090
    print("cfg.port:", cfg.port)

    # --- Pass mutable object to function ---
    print("\n--- Pass mutable (like *Config) ---")
    enable_verbose(cfg)
    print(f"after enable_verbose: {cfg}")

    # --- Immutable (like Go pass-by-value) ---
    print("\n--- Immutable int (like pass-by-value) ---")
    num = 10
    no_change(num)
    print("num unchanged:", num)

    # --- None = Go's nil ---
    print("\n--- None (like nil) ---")
    ref = None
    print("ref:", ref)
    print("ref is None:", ref is None)
    # Accessing ref.something would raise AttributeError (like nil deref panic)

    # --- Object identity ---
    print("\n--- Identity ---")
    a = [1, 2, 3]
    b = [1, 2, 3]
    print("a == b:", a == b)     # True (same value)
    print("a is b:", a is b)     # False (different objects)

    # --- Factory function ---
    print("\n--- Factory ---")
    cp = create_config()
    print(f"created: {cp}")


if __name__ == "__main__":
    main()
