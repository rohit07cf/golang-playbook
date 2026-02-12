# Files: Read and Write

## What It Is

- `os.ReadFile` / `os.WriteFile` for small files (loads entire content)
- **ELI10:** Reading a file is opening a book. Writing a file is scribbling in a notebook. Always close the cover when you're done.
- `os.Open` + `os.Create` for streaming or controlled access

## Why It Matters

- File IO is a basic backend skill -- config files, logs, data import/export
- **ELI10:** Leaving a file open is like leaving the fridge door open -- eventually you run out of resources and everything breaks.
- Interviewers test whether you always close files and check errors

## Syntax Cheat Sheet

```go
// Go: read/write files
data, err := os.ReadFile("sample.txt")       // read all
os.WriteFile("out.txt", data, 0644)           // write all
f, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)
defer f.Close()
```

```python
# Python: open() with mode
data = open("sample.txt").read()              # read all
open("out.txt", "w").write(data)              # write all
f = open("log.txt", "a")                      # append
f.close()
```

> **Python differs**: `with open(...) as f:` is the idiomatic close pattern.
> Go uses `defer f.Close()`.

## Tiny Example

- `main.go` -- reads sample.txt, writes a new file, appends to it, reads back
- `main.py` -- same flow using `open()` and `with` statement

## Common Interview Traps

- **Forgetting defer Close()**: leaked file descriptors exhaust OS limits
- **os.ReadFile loads everything into memory**: don't use for large files
- **File permissions**: `0644` = owner rw, group r, others r (know this)
- **Append vs overwrite**: `os.O_APPEND` for append; `os.Create` truncates
- **Writing without flush**: buffered writers need explicit `Flush()` before close

## What to Say in Interviews

- "I use os.ReadFile for small files, os.Open + streaming for large files"
- "I always defer Close immediately after a successful open"
- "For append I use os.OpenFile with O_APPEND|O_WRONLY flags"

## Run It

```bash
go run ./07_io_files_networking/02_files_read_write/
python ./07_io_files_networking/02_files_read_write/main.py
```

## TL;DR (Interview Summary)

- `os.ReadFile("path")` -- read all bytes (small files only)
- `os.WriteFile("path", data, perm)` -- write all bytes
- `os.Open` for reading; `os.Create` for writing (truncates)
- `os.OpenFile` with flags for append, read-write, etc.
- Always `defer f.Close()` after open
- File permissions: `0644` (rw-r--r--)
- Python: `open(path, mode)` with `with` statement for auto-close
