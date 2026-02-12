"""Strings, bytes, and builders -- Python equivalent of the Go example.

Compares string building approaches: +=, list+join, io.StringIO.
"""

import io
import time


def concat_plus(n: int) -> str:
    """Slow in general (CPython may optimize via refcount hack)."""
    s = ""
    for i in range(n):
        s += f"item{i},"
    return s


def concat_join(n: int) -> str:
    """Fast: build list, join once at the end."""
    parts = [f"item{i}" for i in range(n)]
    return ",".join(parts)


def concat_stringio(n: int) -> str:
    """io.StringIO: buffer-based, similar to Go's strings.Builder."""
    buf = io.StringIO()
    for i in range(n):
        buf.write(f"item{i},")
    return buf.getvalue()


def bytes_str_conversion(n: int) -> None:
    """Show cost of encoding/decoding between bytes and str."""
    s = "hello " * 100
    start = time.perf_counter()
    for _ in range(n):
        b = s.encode("utf-8")   # str -> bytes
        _ = b.decode("utf-8")   # bytes -> str
    elapsed = time.perf_counter() - start
    print(f"  {n} str<->bytes round-trips: {elapsed:.4f}s")


def main() -> None:
    n = 50_000

    print("=== Strings, Bytes, and Builders ===\n")
    print(f"Building {n}-element string...\n")

    start = time.perf_counter()
    s1 = concat_plus(n)
    t1 = time.perf_counter() - start

    start = time.perf_counter()
    s2 = concat_join(n)
    t2 = time.perf_counter() - start

    start = time.perf_counter()
    s3 = concat_stringio(n)
    t3 = time.perf_counter() - start

    print(f"  += concat:     {t1:.4f}s  (len={len(s1)})")
    print(f"  ''.join():     {t2:.4f}s  (len={len(s2)})")
    print(f"  io.StringIO:   {t3:.4f}s  (len={len(s3)})")
    print()

    # bytes <-> str conversion cost
    print("--- bytes <-> str conversion cost ---")
    bytes_str_conversion(500_000)
    print()

    # f-string vs % formatting vs str()
    print("--- Formatting comparison ---")
    iters = 500_000

    start = time.perf_counter()
    for i in range(iters):
        _ = f"{i}"
    fstring_time = time.perf_counter() - start

    start = time.perf_counter()
    for i in range(iters):
        _ = str(i)
    str_time = time.perf_counter() - start

    print(f"  f-string:  {fstring_time:.4f}s")
    print(f"  str():     {str_time:.4f}s")
    print()

    print("Key: ''.join(list) is Python's equivalent of Go's strings.Builder.")
    print("     f-strings are fast and readable -- use them by default.")


if __name__ == "__main__":
    main()
