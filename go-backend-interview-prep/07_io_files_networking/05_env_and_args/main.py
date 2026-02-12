"""Environment variables and command-line arguments -- Python equivalent."""

import argparse
import os
import sys


def main() -> None:
    # --- Example 1: sys.argv ---
    print("=== sys.argv ===")
    print(f"  program: {sys.argv[0]}")
    print(f"  all args: {sys.argv[1:]}")

    # --- Example 2: argparse ---
    print("\n=== argparse ===")
    parser = argparse.ArgumentParser(description="Greeting demo")
    parser.add_argument("--name", default="world", help="who to greet")
    parser.add_argument("--count", type=int, default=1, help="how many times")
    parser.add_argument("-v", action="store_true", help="verbose output")
    args, remaining = parser.parse_known_args()

    print(f"  name={args.name!r} count={args.count} verbose={args.v}")
    print(f"  remaining args: {remaining}")

    for _ in range(args.count):
        print(f"  hello, {args.name}!")

    # --- Example 3: environment variables ---
    print("\n=== Environment variables ===")

    # os.environ.get returns default if unset
    home = os.environ.get("HOME", "")
    print(f"  HOME: {home}")

    # Check existence explicitly
    secret = os.environ.get("MY_APP_SECRET")
    if secret is not None:
        print(f"  MY_APP_SECRET: {secret}")
    else:
        print("  MY_APP_SECRET: (not set)")

    # List env vars matching a prefix
    print("\n=== Env vars starting with 'GO' ===")
    for key, val in sorted(os.environ.items()):
        if key.startswith("GO"):
            print(f"  {key}={val}")


if __name__ == "__main__":
    main()
