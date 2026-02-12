"""CSV basics -- Python equivalent of the Go example."""

import csv
import io
import os
import shutil
import tempfile


def main() -> None:
    # --- Example 1: read CSV from file ---
    print("=== Read sample.csv ===")
    try:
        f = open("07_io_files_networking/07_csv_basics/sample.csv")
    except FileNotFoundError:
        f = open("sample.csv")

    with f:
        reader = csv.reader(f)
        header = next(reader)
        print(f"  header: {header}")
        for row in reader:
            if row:  # skip empty rows
                print(f"  row: name={row[0]} age={row[1]} city={row[2]}")

    # --- Example 1b: DictReader (Python bonus) ---
    print("\n=== DictReader (Python-specific) ===")
    try:
        f = open("07_io_files_networking/07_csv_basics/sample.csv")
    except FileNotFoundError:
        f = open("sample.csv")

    with f:
        for row in csv.DictReader(f):
            print(f"  {row['name']} is {row['age']} in {row['city']}")

    # --- Example 2: read from string ---
    print("\n=== Read CSV from string ===")
    csv_data = 'product,price\n"Widget, Large",9.99\nGadget,4.50'
    reader = csv.reader(io.StringIO(csv_data))
    for row in reader:
        print(f"  {row}")

    # --- Example 3: write CSV ---
    print("\n=== Write CSV ===")
    tmpdir = tempfile.mkdtemp()
    out_path = os.path.join(tmpdir, "output.csv")

    try:
        with open(out_path, "w", newline="") as f:
            writer = csv.writer(f)
            writer.writerow(["name", "score"])
            writer.writerow(["alice", "95"])
            writer.writerow(["bob", "87"])
            writer.writerow(["charlie, jr.", "92"])  # comma -- auto-quoted

        with open(out_path) as f:
            print(f"  written CSV:\n{f.read()}", end="")
    finally:
        shutil.rmtree(tmpdir, ignore_errors=True)


if __name__ == "__main__":
    main()
