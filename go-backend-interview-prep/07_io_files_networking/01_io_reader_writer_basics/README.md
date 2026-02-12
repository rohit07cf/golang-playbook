# IO Reader / Writer Basics

## What It Is

- `io.Reader`: `Read(p []byte) (n int, err error)` -- reads bytes into a buffer
- `io.Writer`: `Write(p []byte) (n int, err error)` -- writes bytes from a buffer

## Why It Matters

- These two interfaces underpin **everything** in Go: files, networks, HTTP, compression
- Interviewers expect you to explain why Go's IO is composable

## Syntax Cheat Sheet

```go
// Go: Reader/Writer are interfaces
var r io.Reader = strings.NewReader("hello")
var w io.Writer = &bytes.Buffer{}
n, err := io.Copy(w, r)  // pipe reader into writer
```

```python
# Python: io.StringIO / io.BytesIO
import io, shutil
r = io.StringIO("hello")
w = io.StringIO()
shutil.copyfileobj(r, w)  # pipe reader into writer
```

> **Python differs**: Python has file-like objects with `.read()` / `.write()`
> but no single universal interface like Go's `io.Reader` / `io.Writer`.

## Tiny Example

- `main.go` -- strings.NewReader, bytes.Buffer, io.Copy, custom Reader
- `main.py` -- io.StringIO, io.BytesIO, shutil.copyfileobj

## Common Interview Traps

- **Read returns io.EOF**: this is normal end-of-stream, not an error
- **Read may return fewer bytes**: `Read(p)` can return `n < len(p)` -- call in a loop or use `io.ReadFull`
- **io.Copy is the workhorse**: connects any Reader to any Writer (streaming, no full-memory load)
- **strings.NewReader vs bytes.NewReader**: string input vs byte slice input
- **Closing is not part of Reader/Writer**: close via `io.Closer` or `io.ReadCloser`

## What to Say in Interviews

- "io.Reader and io.Writer are Go's two most important interfaces -- everything composes through them"
- "io.Copy streams data from Reader to Writer without loading it all into memory"
- "I always check for io.EOF to detect end of stream, not treat it as an error"

## Run It

```bash
go run ./07_io_files_networking/01_io_reader_writer_basics/
python ./07_io_files_networking/01_io_reader_writer_basics/main.py
```

## TL;DR (Interview Summary)

- `io.Reader`: one method `Read(p []byte) (n int, err error)`
- `io.Writer`: one method `Write(p []byte) (n int, err error)`
- `io.Copy(dst, src)` streams without full memory load
- `strings.NewReader` / `bytes.NewReader` create readers from strings/bytes
- `bytes.Buffer` is both a Reader and Writer
- `io.EOF` signals end of stream (not an error)
- Python: `io.StringIO`, `io.BytesIO`, `shutil.copyfileobj`
