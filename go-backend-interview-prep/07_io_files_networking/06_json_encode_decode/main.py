"""JSON encode / decode -- Python equivalent of the Go example."""

import json
from dataclasses import dataclass, asdict, field


@dataclass
class User:
    name: str
    age: int = 0
    email: str = ""
    # is_admin is excluded from JSON (like json:"-")
    is_admin: bool = field(default=False, repr=False)

    def to_json_dict(self) -> dict:
        d = asdict(self)
        del d["is_admin"]       # skip like json:"-"
        if d["age"] == 0:
            del d["age"]        # omitempty
        return d


@dataclass
class Profile:
    user: User
    city: str = ""
    country: str = ""

    def to_json_dict(self) -> dict:
        d = {"user": self.user.to_json_dict(), "city": self.city}
        if self.country:
            d["country"] = self.country
        return d


def main() -> None:
    # --- Example 1: Marshal (dumps) ---
    print("=== Marshal ===")
    u = User(name="alice", age=30, email="alice@example.com", is_admin=True)
    data = json.dumps(u.to_json_dict())
    print(f"  json: {data}")

    # --- Example 2: Pretty print ---
    print("\n=== Pretty print ===")
    pretty = json.dumps(u.to_json_dict(), indent=2)
    print(pretty)

    # --- Example 3: Unmarshal (loads) ---
    print("\n=== Unmarshal ===")
    json_str = '{"name":"bob","age":25,"email":"bob@example.com"}'
    parsed = json.loads(json_str)
    u2 = User(**{k: parsed.get(k, d) for k, d in [("name", ""), ("age", 0), ("email", "")]})
    print(f"  parsed: {u2}")
    print(f"  is_admin (skipped): {u2.is_admin}")

    # --- Example 4: omitempty ---
    print("\n=== omitempty ===")
    empty = User(name="charlie", age=0, email="c@x.com")
    data = json.dumps(empty.to_json_dict())
    print(f"  age=0 omitted: {data}")

    # --- Example 5: nested struct ---
    print("\n=== Nested struct ===")
    p = Profile(user=User(name="dana", age=28, email="dana@x.com"), city="NYC")
    data = json.dumps(p.to_json_dict())
    print(f"  json: {data}")

    # --- Example 6: dynamic JSON (dict) ---
    print("\n=== Dynamic JSON (dict) ===")
    dynamic = {"status": "ok", "count": 42, "tags": ["go", "json"]}
    data = json.dumps(dynamic)
    print(f"  json: {data}")

    result = json.loads(data)
    print(f"  parsed dict: {result}")


if __name__ == "__main__":
    main()
