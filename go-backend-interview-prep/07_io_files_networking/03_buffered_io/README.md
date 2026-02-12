# Buffered IO

## What It Is

- `bufio.Scanner` reads input line-by-line (or by custom split function)
- **ELI10:** Buffered I/O is like loading a shopping cart instead of carrying items one by one to the checkout.
- `bufio.NewReader` / `bufio.NewWriter` wrap readers/writers with a buffer

## Why It Matters

- Line-by-line reading is the most common file processing pattern
- **ELI10:** Without buffering, every tiny read or write is a separate trip to the OS -- like driving to the store each time you need one egg.
- Buffered writes reduce syscalls -- critical for performance

## Syntax Cheat Sheet

```go
// Go: Scanner for line-by-line
scanner := bufio.NewScanner(file)
for scanner.Scan() { fmt.Println(scanner.Text()) }

// Buffered writer (must flush!)
w := bufio.NewWriter(file)
w.WriteString("data")
w.Flush()  // <-- don't forget this
```

```python
# Python: iterate file object for lines
for line in open("file.txt"):
    print(line, end="")
# Buffered by default; flush with f.flush()
```

> **Python differs**: `open()` is buffered by default. Go's `os.Open`
> is unbuffered -- you wrap it with `bufio` for buffering.

## Tiny Example

- `main.go` -- reads sample.txt line-by-line with Scanner, writes with buffered writer
- `main.py` -- same patterns with file iteration and explicit flush

## Common Interview Traps

- **Forgetting Flush()**: buffered writer holds data in memory until flushed
- **Scanner default limit**: bufio.Scanner has a 64 KB line limit (use `scanner.Buffer()` to increase)
- **scanner.Err()**: always check `scanner.Err()` after the loop for IO errors
- **Scan() vs ReadString**: Scanner is simpler; ReadString includes the delimiter

## What to Say in Interviews

- "I use bufio.Scanner for line-by-line reading and always check scanner.Err() after the loop"
- "Buffered writers reduce syscalls but need explicit Flush before closing"
- "In Python, open() is already buffered, but in Go I wrap with bufio explicitly"

## Run It

```bash
go run ./07_io_files_networking/03_buffered_io/
python ./07_io_files_networking/03_buffered_io/main.py
```

## TL;DR (Interview Summary)

- `bufio.Scanner`: line-by-line with `Scan()` + `Text()`
- Always check `scanner.Err()` after the scan loop
- Default Scanner line limit = 64 KB (increase with `scanner.Buffer()`)
- `bufio.NewWriter` buffers writes; **must** call `Flush()` before close
- Buffering reduces syscalls = better performance
- Python: `open()` is buffered by default; iterate with `for line in f`
