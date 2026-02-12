"""Paths and directories -- Python equivalent of the Go example."""

import glob
import os
import shutil
import tempfile
from pathlib import Path


def main() -> None:
    # --- Example 1: pathlib.Path / os.path.join ---
    print("=== pathlib.Path ===")
    p = Path("data") / "users" / "profile.json"
    print(f"  joined: {p}")
    print(f"  dir:   {p.parent}")
    print(f"  base:  {p.name}")
    print(f"  ext:   {p.suffix}")

    # --- Example 2: temp directory + create files ---
    print("\n=== Temp directory ===")
    tmpdir = tempfile.mkdtemp(prefix="paths-demo")

    try:
        # Create nested structure
        subdir = Path(tmpdir) / "sub" / "nested"
        subdir.mkdir(parents=True, exist_ok=True)

        # Create a few files
        for name in ["a.txt", "b.go", "c.json"]:
            (Path(tmpdir) / name).write_text(f"content of {name}")
        (subdir / "deep.txt").write_text("deep file")

        print(f"  created temp dir: {tmpdir}")

        # --- Example 3: os.walk ---
        print("\n=== os.walk ===")
        for root, dirs, files in os.walk(tmpdir):
            rel = os.path.relpath(root, tmpdir)
            print(f"  [dir]  {rel}/")
            for fname in files:
                fpath = Path(root) / fname
                size = fpath.stat().st_size
                rel_file = os.path.relpath(fpath, tmpdir)
                print(f"  [file] {rel_file} ({size} bytes)")

        # --- Example 4: glob pattern matching ---
        print("\n=== Glob ===")
        pattern = os.path.join(tmpdir, "*.go")
        for m in glob.glob(pattern):
            print(f"  match: {Path(m).name}")
    finally:
        shutil.rmtree(tmpdir, ignore_errors=True)


if __name__ == "__main__":
    main()
