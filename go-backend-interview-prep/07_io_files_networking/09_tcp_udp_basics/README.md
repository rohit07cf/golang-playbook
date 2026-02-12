# TCP / UDP Basics

## What It Is

- **TCP**: `net.Listen("tcp", addr)` on server, `net.Dial("tcp", addr)` on client
- **UDP**: same API but connectionless -- `net.ListenPacket` / `net.DialUDP`

## Why It Matters

- Understanding TCP/UDP is essential for backend system design
- Interviewers ask about connection-oriented vs connectionless, read loops, graceful close

## Syntax Cheat Sheet

```go
// Go TCP: server
ln, _ := net.Listen("tcp", ":0")
conn, _ := ln.Accept()
io.Copy(conn, conn)  // echo

// Go TCP: client
conn, _ := net.Dial("tcp", addr)
conn.Write([]byte("hello"))
```

```python
# Python TCP: server
import socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(("", 0)); s.listen(1)
conn, _ = s.accept()
conn.sendall(conn.recv(1024))  # echo
```

> **Python differs**: uses the lower-level `socket` module directly.
> Go's `net` package wraps sockets with higher-level Read/Write methods.

## Tiny Example

- `main.go` -- TCP echo server + client in the same program
- `main.py` -- same pattern with Python sockets

## Common Interview Traps

- **Read may return partial data**: TCP is a byte stream, not message-based
- **Forgetting to close connections**: leaked connections exhaust file descriptors
- **UDP is unreliable**: packets can be lost, duplicated, or reordered
- **Blocking Accept**: runs in a goroutine/thread to avoid blocking main
- **Buffer size**: too small = truncated reads; too large = wasted memory

## What to Say in Interviews

- "TCP is a reliable byte stream; UDP is unreliable datagrams -- I choose based on requirements"
- "I always handle partial reads in a loop and close connections when done"
- "For production servers I'd use goroutine-per-connection with graceful shutdown"

## Run It

```bash
go run ./07_io_files_networking/09_tcp_udp_basics/
python ./07_io_files_networking/09_tcp_udp_basics/main.py
```

## TL;DR (Interview Summary)

- TCP: `net.Listen` + `Accept` loop (server); `net.Dial` (client)
- Connection is `io.ReadWriteCloser` -- use `Read`/`Write`/`Close`
- Always close connections; always handle partial reads
- UDP: connectionless, unreliable, use `ListenPacket` / `ReadFrom` / `WriteTo`
- Goroutine-per-connection is the standard Go server pattern
- Python: `socket` module, `AF_INET` + `SOCK_STREAM` (TCP) / `SOCK_DGRAM` (UDP)
