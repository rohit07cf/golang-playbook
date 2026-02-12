package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// CountingReader wraps a Reader and counts bytes read.
type CountingReader struct {
	reader    io.Reader
	BytesRead int
}

func (cr *CountingReader) Read(p []byte) (int, error) {
	n, err := cr.reader.Read(p)
	cr.BytesRead += n
	return n, err
}

func main() {
	// --- Example 1: strings.NewReader ---
	fmt.Println("=== strings.NewReader ===")
	r := strings.NewReader("hello, Go IO!")
	buf := make([]byte, 5)

	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("  read %d bytes: %q\n", n, buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("  error:", err)
			break
		}
	}

	// --- Example 2: bytes.Buffer (Reader + Writer) ---
	fmt.Println("\n=== bytes.Buffer ===")
	var b bytes.Buffer
	b.WriteString("hello ")
	b.WriteString("buffer")
	fmt.Println("  buffer contents:", b.String())

	// Read from buffer
	out := make([]byte, 5)
	n, _ := b.Read(out)
	fmt.Printf("  read %d bytes: %q\n", n, out[:n])

	// --- Example 3: io.Copy ---
	fmt.Println("\n=== io.Copy ===")
	src := strings.NewReader("streaming data via io.Copy")
	var dst bytes.Buffer

	written, err := io.Copy(&dst, src)
	if err != nil {
		fmt.Println("  error:", err)
	}
	fmt.Printf("  copied %d bytes: %q\n", written, dst.String())

	// --- Example 4: custom Reader ---
	fmt.Println("\n=== Custom CountingReader ===")
	cr := &CountingReader{reader: strings.NewReader("count these bytes")}
	var result bytes.Buffer
	io.Copy(&result, cr)
	fmt.Printf("  content: %q\n", result.String())
	fmt.Printf("  total bytes read: %d\n", cr.BytesRead)
}
