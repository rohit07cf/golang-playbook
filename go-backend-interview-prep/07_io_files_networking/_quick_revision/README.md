# IO, Files, Networking -- Quick Revision

> One-screen cheat sheet. Skim before interviews.

---

## 1. Reader / Writer Mental Model

```go
// Go: everything speaks io.Reader / io.Writer
var r io.Reader = strings.NewReader("hello")
var w io.Writer = os.Stdout
io.Copy(w, r)  // pipe: reader -> writer
```

```python
# Python: file-like objects with .read() / .write()
import io, shutil
r = io.StringIO("hello")
shutil.copyfileobj(r, sys.stdout)
```

## 2. Buffered IO vs Direct

```go
// Go: bufio wraps for efficiency
scanner := bufio.NewScanner(file)
for scanner.Scan() { fmt.Println(scanner.Text()) }
w := bufio.NewWriter(f); w.WriteString("x"); w.Flush()
```

```python
# Python: open() is buffered by default
for line in open("f.txt"): print(line, end="")
```

## 3. Files: Permissions + Close + Flush

```go
f, _ := os.Create("out.txt")   // perm 0666 (umask applied)
defer f.Close()                  // ALWAYS defer close
os.WriteFile("f.txt", data, 0644)
```

```python
with open("out.txt", "w") as f:  # auto-close
    f.write("data")
```

## 4. JSON / CSV Quick Rules

```go
// JSON: struct tags control keys
type U struct { Name string `json:"name,omitempty"` }
json.Marshal(u); json.Unmarshal(data, &u)

// CSV: always Flush
w := csv.NewWriter(f); w.Write(row); w.Flush()
```

```python
json.dumps(obj); json.loads(s)
csv.writer(f).writerow(row)
```

## 5. HTTP Client: Timeouts ALWAYS

```go
client := &http.Client{Timeout: 5 * time.Second}
resp, _ := client.Get(url)
defer resp.Body.Close()
```

```python
urllib.request.urlopen(url, timeout=5)
```

## 6. TCP: Connect / Read / Write Loop

```go
ln, _ := net.Listen("tcp", ":8080")
conn, _ := ln.Accept()
defer conn.Close()
io.Copy(conn, conn)  // echo server
```

```python
s = socket.socket(); s.bind(("", 8080)); s.listen(1)
conn, _ = s.accept()
conn.sendall(conn.recv(1024))
```

## 7. Streaming: Don't Load Everything

```go
// Stream HTTP body to file -- constant memory
io.Copy(file, resp.Body)
```

```python
while chunk := resp.read(8192):
    f.write(chunk)
```

## 8. DNS + Env + Args

```go
addrs, _ := net.LookupHost("example.com")
val := os.Getenv("KEY")
name := flag.String("name", "default", "usage")
```

```python
socket.gethostbyname_ex("example.com")
os.environ.get("KEY", "")
argparse.ArgumentParser().add_argument("--name")
```

---

## Interview One-Liners

1. "io.Reader and io.Writer are Go's two most important interfaces"
2. "Always defer Close after opening files and HTTP response bodies"
3. "bufio.Scanner for line-by-line; always check scanner.Err()"
4. "JSON struct tags control field names -- unexported fields are invisible"
5. "HTTP clients must have explicit timeouts -- the default has none"
6. "io.Copy streams data with constant memory -- never ReadAll for large data"
7. "TCP is a byte stream, not message-based -- handle partial reads"
8. "Use filepath.Join for paths, os.MkdirTemp for temp dirs"

---

## TL;DR

- `io.Reader` + `io.Writer` = universal byte abstractions
- `defer f.Close()` -- always, immediately after open
- `bufio.Scanner` for lines; `bufio.Writer` needs `Flush()`
- JSON: `Marshal`/`Unmarshal` + struct tags + `omitempty`
- HTTP: always set `Timeout`; always close `resp.Body`
- TCP: `Listen` + `Accept` + `Read`/`Write` loop
- Streaming: `io.Copy` -- constant memory for any size
- Python: `open()`, `json`, `csv`, `urllib.request`, `socket`
