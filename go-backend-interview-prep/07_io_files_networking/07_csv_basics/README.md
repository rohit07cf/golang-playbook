# CSV Basics

## What It Is

- Go: `encoding/csv` provides `csv.NewReader` / `csv.NewWriter`
- Reads/writes CSV with proper quoting, escaping, and delimiters

## Why It Matters

- CSV is still the most common data exchange format for batch processing
- Interviewers test whether you handle headers, quoting, and edge cases

## Syntax Cheat Sheet

```go
// Go: csv reader/writer
r := csv.NewReader(file)
records, _ := r.ReadAll()       // [][]string
w := csv.NewWriter(file)
w.Write([]string{"a", "b"})
w.Flush()
```

```python
# Python: csv module
import csv
reader = csv.reader(open("data.csv"))
for row in reader: print(row)
writer = csv.writer(open("out.csv", "w"))
writer.writerow(["a", "b"])
```

> **Python differs**: `csv.DictReader` / `csv.DictWriter` give dict-based
> access by column name -- very convenient, no Go equivalent in stdlib.

## Tiny Example

- `main.go` -- reads sample.csv, prints rows, writes a new CSV
- `main.py` -- same with csv.reader/writer and DictReader

## Common Interview Traps

- **Forgetting csv.Writer.Flush()**: data stays in buffer without it
- **Fields with commas**: the CSV library handles quoting automatically
- **Header row**: first ReadAll/Read row is usually headers -- handle it
- **Newline at end of file**: trailing newline may produce an empty last row

## What to Say in Interviews

- "encoding/csv handles quoting and escaping automatically"
- "I always call csv.Writer.Flush after writing and check csv.Writer.Error"
- "Python's DictReader is more ergonomic; Go uses slice-of-strings"

## Run It

```bash
go run ./07_io_files_networking/07_csv_basics/
python ./07_io_files_networking/07_csv_basics/main.py
```

## TL;DR (Interview Summary)

- `csv.NewReader(r)` reads from any `io.Reader`
- `ReadAll()` returns `[][]string`; `Read()` returns one row at a time
- `csv.NewWriter(w)` + `Write(row)` + `Flush()` -- always flush
- Quoting/escaping is automatic
- Python: `csv.reader`, `csv.writer`, `csv.DictReader`
