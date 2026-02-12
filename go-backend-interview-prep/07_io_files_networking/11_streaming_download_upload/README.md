# Streaming Download / Upload

## What It Is

- **Streaming**: read data in chunks, write it as you go -- never load the whole thing into memory
- **ELI10:** Streaming is drinking from a firehose -- you process water as it flows instead of filling a swimming pool first.
- Go: `io.Copy(dst, resp.Body)` streams an HTTP response directly to a file

## Why It Matters

- Downloading a 2 GB file with `io.ReadAll` would OOM your server
- **ELI10:** Loading a huge file into memory is like trying to drink the ocean in one gulp -- stream it or you'll explode.
- Interviewers test whether you understand memory-efficient data handling

## Syntax Cheat Sheet

```go
// Go: stream HTTP body to file
resp, _ := http.Get(url)
defer resp.Body.Close()
f, _ := os.Create("out.bin")
defer f.Close()
io.Copy(f, resp.Body)  // streams, constant memory
```

```python
# Python: stream response to file
resp = urllib.request.urlopen(url)
with open("out.bin", "wb") as f:
    while chunk := resp.read(8192):
        f.write(chunk)
```

> **Python differs**: no `io.Copy` equivalent that works on HTTP responses.
> Read in a loop with fixed-size chunks.

## Tiny Example

- `main.go` -- local server serves data, client streams to file with io.Copy + progress
- `main.py` -- same with chunked reads from urllib

## Common Interview Traps

- **io.ReadAll loads everything**: never use for large responses -- use io.Copy
- **Forgetting to close response body**: leaks connections
- **No progress tracking**: io.Copy is silent -- wrap with a counting reader for progress
- **Partial downloads**: network failures mid-stream leave partial files
- **Upload streaming**: use `io.Pipe` or pass a reader as request body

## What to Say in Interviews

- "I stream large responses to disk with io.Copy -- constant memory regardless of file size"
- "For progress tracking I wrap the reader with a counting wrapper"
- "For uploads I pass an io.Reader as the request body to stream without buffering"

## Run It

```bash
go run ./07_io_files_networking/11_streaming_download_upload/
python ./07_io_files_networking/11_streaming_download_upload/main.py
```

## TL;DR (Interview Summary)

- `io.Copy(file, resp.Body)` -- streams with constant memory
- Never use `io.ReadAll` for large responses
- Always close `resp.Body` (defer)
- Wrap reader for progress tracking
- For uploads: pass `io.Reader` as request body (streams)
- Handle partial downloads (network failures mid-stream)
- Python: read in chunks with `resp.read(8192)` in a loop
