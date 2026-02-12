# Python equivalent of Stringer and Custom Formatting (compare with main.go)
# Go fmt.Stringer ~ Python __str__
# Go GoStringer ~ Python __repr__


class User:
    def __init__(self, name: str, age: int):
        self.name = name
        self.age = age

    def __str__(self) -> str:
        """Like Go's String() method (fmt.Stringer)."""
        return f"{self.name} (age {self.age})"

    def __repr__(self) -> str:
        """Like Go's GoString() method (%#v)."""
        return f"User(name={self.name!r}, age={self.age})"


class IPAddr:
    def __init__(self, a: int, b: int, c: int, d: int):
        self.octets = (a, b, c, d)

    def __str__(self) -> str:
        return ".".join(str(o) for o in self.octets)

    def __repr__(self) -> str:
        return f"IPAddr{self.octets}"


class RawData:
    """No __str__ -- uses default repr."""
    def __init__(self, value: int):
        self.value = value


def main():
    # --- __str__ in action ---
    print("--- __str__ (like Stringer) ---")
    u = User(name="Alice", age=30)
    print(u)                    # calls __str__
    print(f"user: {u}")         # also calls __str__

    # --- __repr__ ---
    print("\n--- __repr__ (like GoStringer) ---")
    print(f"debug: {u!r}")      # calls __repr__

    # --- IPAddr ---
    print("\n--- IPAddr ---")
    home = IPAddr(127, 0, 0, 1)
    dns = IPAddr(8, 8, 8, 8)
    print("home:", home)
    print("dns: ", dns)

    # --- Without __str__: default output ---
    print("\n--- Without __str__ ---")
    raw = RawData(value=42)
    print(raw)                  # shows <__main__.RawData object at 0x...>

    # --- In f-string ---
    print("\n--- In f-string ---")
    msg = f"Hello, {u}!"
    print(msg)

    # --- List of objects ---
    print("\n--- List ---")
    addrs = [IPAddr(10, 0, 0, 1), IPAddr(192, 168, 1, 1)]
    for addr in addrs:
        print(" ", addr)


if __name__ == "__main__":
    main()
