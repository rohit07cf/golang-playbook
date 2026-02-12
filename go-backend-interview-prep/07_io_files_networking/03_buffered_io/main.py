"""Buffered IO -- Python equivalent of the Go example.

Python's open() is buffered by default. Iteration over a file object
reads line-by-line efficiently (like bufio.Scanner).
"""

import io
import os
import tempfile


def main() -> None:
    # --- Example 1: line-by-line from string ---
    print("=== Line-by-line from string ===")
    text = "line one\nline two\nline three"
    for i, line in enumerate(io.StringIO(text), 1):
        print(f"  {i}: {line.rstrip()}")

    # --- Example 2: line-by-line from file ---
    print("\n=== Line-by-line from file ===")
    try:
        path = "07_io_files_networking/03_buffered_io/sample.txt"
        f = open(path)
    except FileNotFoundError:
        try:
            f = open("sample.txt")
        except FileNotFoundError:
            f = None
            print("  could not open sample.txt")

    if f:
        with f:
            for line in f:
                print(f"  {line.rstrip()}")

    # --- Example 3: buffered writer ---
    print("\n=== Buffered writer ===")
    tmpdir = tempfile.mkdtemp()
    out_path = os.path.join(tmpdir, "buffered.txt")

    try:
        with open(out_path, "w", buffering=8192) as w:
            w.write("buffered line 1\n")
            w.write("buffered line 2\n")
            w.write("buffered line 3\n")
            # flush happens automatically on close with 'with'

        with open(out_path) as f:
            print(f"  file contents:\n{f.read()}", end="")
    finally:
        import shutil
        shutil.rmtree(tmpdir, ignore_errors=True)

    # --- Example 4: word-by-word ---
    print("\n=== Word scanner ===")
    text = "the quick brown fox"
    for word in text.split():
        print(f'  word: "{word}"')


if __name__ == "__main__":
    main()
