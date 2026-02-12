"""Files: read and write -- Python equivalent of the Go example."""

import os
import tempfile


def main() -> None:
    # Use a temp directory
    tmpdir = tempfile.mkdtemp()

    try:
        # --- Example 1: read entire file ---
        print("=== Read sample.txt ===")
        try:
            with open("07_io_files_networking/02_files_read_write/sample.txt") as f:
                data = f.read()
        except FileNotFoundError:
            try:
                with open("sample.txt") as f:
                    data = f.read()
            except FileNotFoundError:
                data = None
                print("  could not read sample.txt")

        if data:
            print(f"  {len(data)} chars:")
            print(data, end="")

        # --- Example 2: write entire file ---
        print("\n=== Write new file ===")
        out_path = os.path.join(tmpdir, "output.txt")
        content = "first line\nsecond line\n"
        with open(out_path, "w") as f:
            f.write(content)
        print(f"  wrote {len(content)} chars to {out_path}")

        # --- Example 3: append to file ---
        print("\n=== Append to file ===")
        with open(out_path, "a") as f:
            f.write("appended third line\n")
        print("  appended one line")

        # --- Example 4: read back ---
        print("\n=== Read back ===")
        with open(out_path) as f:
            print(f.read(), end="")

        # --- Example 5: streaming read (chunked) ---
        print("\n=== Streaming read (chunked) ===")
        with open(out_path, "rb") as f:
            while True:
                chunk = f.read(16)
                if not chunk:
                    break
                print(f"  chunk ({len(chunk)} bytes): {chunk!r}")
    finally:
        # Cleanup
        import shutil
        shutil.rmtree(tmpdir, ignore_errors=True)


if __name__ == "__main__":
    main()
