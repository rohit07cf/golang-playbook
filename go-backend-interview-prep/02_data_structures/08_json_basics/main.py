# Python equivalent of JSON Basics (compare with main.go)
# Go encoding/json -> Python json module
# Go struct tags -> Python dict keys or dataclass + custom encoder

import json
from dataclasses import dataclass, asdict


@dataclass
class User:
    name: str = ""
    email: str = ""
    age: int = 0
    password: str = ""   # excluded from JSON (like Go's json:"-")
    active: bool = False


def user_to_dict(u: User, omit_empty: bool = False) -> dict:
    """Mimics Go struct tags: exclude password, optionally omit zero values."""
    d = {"name": u.name, "age": u.age}

    if not omit_empty or u.email:
        d["email"] = u.email
    if not omit_empty or u.active:
        d["active"] = u.active
    # password always excluded (like json:"-")
    return d


def main():
    # === MARSHAL: object -> JSON string ===
    print("--- Marshal ---")
    u = User(name="Alice", email="alice@example.com", age=30,
             password="secret123", active=True)

    data = json.dumps(user_to_dict(u))
    print("JSON:", data)

    # === MARSHAL with omitempty ===
    print("\n--- omitempty ---")
    u2 = User(name="Bob", age=25)
    data2 = json.dumps(user_to_dict(u2, omit_empty=True))
    print("JSON:", data2)

    # === PRETTY PRINT ===
    print("\n--- Pretty print ---")
    pretty = json.dumps(user_to_dict(u), indent=2)
    print(pretty)

    # === UNMARSHAL: JSON string -> dict -> object ===
    print("--- Unmarshal ---")
    input_json = '{"name":"Charlie","email":"c@test.com","age":28,"active":true}'
    d = json.loads(input_json)
    u3 = User(**{k: v for k, v in d.items() if k in User.__dataclass_fields__})
    print(f"parsed: {u3}")

    # === Unknown fields silently ignored ===
    print("\n--- Unknown fields ---")
    weird = '{"name":"Dave","unknown_field":"ignored","age":35}'
    d = json.loads(weird)
    u4 = User(name=d.get("name", ""), age=d.get("age", 0))
    print(f"unknown fields ignored: {u4}")

    # === Partial JSON ===
    print("\n--- Partial JSON ---")
    partial = '{"name":"Eve"}'
    d = json.loads(partial)
    u5 = User(name=d.get("name", ""), age=d.get("age", 0))
    print(f"partial: {u5}")

    # === List of dicts ===
    print("\n--- List ---")
    team = [{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}]
    print("team:", json.dumps(team))

    # === Dict -> JSON ===
    print("\n--- Dict ---")
    config = {"host": "localhost", "port": "8080"}
    print("config:", json.dumps(config))


if __name__ == "__main__":
    main()
