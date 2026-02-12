"""IO Reader/Writer basics -- Python equivalent of the Go example.

Python uses file-like objects with .read() / .write() methods.
io.StringIO and io.BytesIO are the in-memory equivalents.
"""

import io
import shutil


class CountingReader:
    """Wraps a file-like object and counts bytes read."""

    def __init__(self, reader):
        self._reader = reader
        self.bytes_read = 0

    def read(self, n=-1):
        data = self._reader.read(n)
        self.bytes_read += len(data)
        return data


def main() -> None:
    # --- Example 1: io.StringIO (like strings.NewReader) ---
    print("=== io.StringIO ===")
    r = io.StringIO("hello, Go IO!")

    while True:
        chunk = r.read(5)
        if not chunk:
            break
        print(f"  read {len(chunk)} chars: {chunk!r}")

    # --- Example 2: io.StringIO as read+write buffer ---
    print("\n=== io.StringIO (buffer) ===")
    buf = io.StringIO()
    buf.write("hello ")
    buf.write("buffer")
    print("  buffer contents:", buf.getvalue())

    buf.seek(0)
    out = buf.read(5)
    print(f"  read {len(out)} chars: {out!r}")

    # --- Example 3: shutil.copyfileobj (like io.Copy) ---
    print("\n=== shutil.copyfileobj ===")
    src = io.BytesIO(b"streaming data via copyfileobj")
    dst = io.BytesIO()

    shutil.copyfileobj(src, dst)
    content = dst.getvalue()
    print(f"  copied {len(content)} bytes: {content.decode()!r}")

    # --- Example 4: custom CountingReader ---
    print("\n=== Custom CountingReader ===")
    cr = CountingReader(io.StringIO("count these bytes"))
    result = cr.read()
    print(f"  content: {result!r}")
    print(f"  total bytes read: {cr.bytes_read}")


if __name__ == "__main__":
    main()
