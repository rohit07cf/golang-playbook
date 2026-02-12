# 07 -- IO, Files, and Networking

Go's `io.Reader` and `io.Writer` are the two most important interfaces
in the standard library. Everything -- files, network connections,
HTTP bodies, compression streams -- speaks this pair of interfaces.

This module covers reading/writing files, buffered IO, JSON/CSV,
HTTP clients, TCP/UDP basics, DNS lookups, and streaming patterns.

Each example includes `main.go` + Python equivalent `main.py`.

---

## Topics

| # | Folder | What You Learn |
|---|--------|---------------|
| 01 | `01_io_reader_writer_basics` | io.Reader, io.Writer, strings.NewReader, bytes.Buffer, io.Copy |
| 02 | `02_files_read_write` | os.Open, os.Create, read/write/append files |
| 03 | `03_buffered_io` | bufio.Scanner, buffered read/write, line-by-line |
| 04 | `04_paths_and_dirs` | filepath.Join, os.MkdirTemp, WalkDir |
| 05 | `05_env_and_args` | os.Args, flag package, os.Getenv |
| 06 | `06_json_encode_decode` | encoding/json, struct tags, marshal/unmarshal |
| 07 | `07_csv_basics` | encoding/csv reader/writer |
| 08 | `08_http_client_basics` | net/http GET, timeouts, headers |
| 09 | `09_tcp_udp_basics` | net.Listen, net.Dial, echo server |
| 10 | `10_dns_and_timeouts` | net.LookupHost, http.Client timeout |
| 11 | `11_streaming_download_upload` | io.Copy streaming, chunked reads |

---

## 10-Min Revision Path

1. Skim `01_io_reader_writer_basics` -- the two core interfaces
2. Skim `02_files_read_write` -- open, read, write, close patterns
3. Skim `03_buffered_io` -- Scanner for line-by-line, buffered flush
4. Skim `06_json_encode_decode` -- struct tags, marshal/unmarshal
5. Skim `08_http_client_basics` -- always set timeouts
6. Skim `09_tcp_udp_basics` -- connect/accept/read/write loop
7. Skim `_quick_revision/` -- one-screen cheat sheet

---

## Common IO / Networking Mistakes

- Forgetting to close files (use `defer f.Close()`)
- Not checking errors after every IO call
- Reading entire large files into memory (use streaming)
- Forgetting to flush buffered writers (`bufio.Writer.Flush()`)
- HTTP client with no timeout (hangs forever on slow servers)
- Ignoring `io.EOF` -- it's not an error, it signals end of stream
- Hardcoding file paths instead of using `filepath.Join`
- Not handling partial reads (`io.Reader.Read` may return fewer bytes)
- Forgetting `Content-Type` header when sending JSON

---

## TL;DR

- `io.Reader` / `io.Writer` are Go's universal byte abstractions
- Always `defer f.Close()` after opening files
- Use `bufio.Scanner` for line-by-line reading
- `encoding/json`: struct tags control field names; marshal = Go -> JSON, unmarshal = JSON -> Go
- HTTP clients **must** have timeouts (`http.Client{Timeout: 10 * time.Second}`)
- TCP: `net.Listen` + `Accept` loop on server; `net.Dial` on client
- Stream large data with `io.Copy` -- never load entire files into memory
- Python: `open()`, `json`, `csv`, `urllib.request`, `socket`
