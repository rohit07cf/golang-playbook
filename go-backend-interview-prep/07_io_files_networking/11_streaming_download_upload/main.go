package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ProgressReader wraps a reader and reports progress.
type ProgressReader struct {
	reader io.Reader
	total  int64
	read   int64
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.read += int64(n)
	if pr.total > 0 {
		pct := float64(pr.read) / float64(pr.total) * 100
		fmt.Printf("\r  progress: %.0f%% (%d/%d bytes)", pct, pr.read, pr.total)
	}
	return n, err
}

func main() {
	// --- Local server that serves a "large" payload ---
	payload := strings.Repeat("Go streaming demo data. ", 1000) // ~24 KB
	mux := http.NewServeMux()
	mux.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(payload)))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write([]byte(payload))
	})

	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	defer ln.Close()
	go http.Serve(ln, mux)
	base := "http://" + ln.Addr().String()

	dir, _ := os.MkdirTemp("", "streaming")
	defer os.RemoveAll(dir)

	// --- Example 1: stream download with io.Copy ---
	fmt.Println("=== Stream download (io.Copy) ===")
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(base + "/download")
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	defer resp.Body.Close()

	outPath := filepath.Join(dir, "downloaded.bin")
	f, _ := os.Create(outPath)

	written, err := io.Copy(f, resp.Body) // streams -- constant memory
	f.Close()
	if err != nil {
		fmt.Println("  copy error:", err)
		return
	}
	fmt.Printf("  downloaded %d bytes to file\n", written)

	// --- Example 2: stream with progress ---
	fmt.Println("\n=== Stream with progress ===")
	resp2, _ := client.Get(base + "/download")
	defer resp2.Body.Close()

	outPath2 := filepath.Join(dir, "progress.bin")
	f2, _ := os.Create(outPath2)

	pr := &ProgressReader{
		reader: resp2.Body,
		total:  resp2.ContentLength,
	}
	written2, _ := io.Copy(f2, pr)
	f2.Close()
	fmt.Printf("\n  done: %d bytes with progress\n", written2)

	// --- Example 3: upload (stream request body) ---
	fmt.Println("\n=== Stream upload ===")
	uploadData := strings.NewReader("upload payload: " + strings.Repeat("x", 500))

	req, _ := http.NewRequest("POST", base+"/download", uploadData)
	req.Header.Set("Content-Type", "application/octet-stream")

	resp3, err := client.Do(req)
	if err != nil {
		fmt.Println("  upload error:", err)
		return
	}
	resp3.Body.Close()
	fmt.Println("  upload status:", resp3.StatusCode)
	fmt.Println("  (request body was streamed, not buffered in memory)")
}
