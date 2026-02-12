package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Use a temp directory so we don't pollute the source tree
	dir, err := os.MkdirTemp("", "fileio")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer os.RemoveAll(dir)

	// --- Example 1: read entire file ---
	fmt.Println("=== Read sample.txt ===")
	data, err := os.ReadFile("07_io_files_networking/02_files_read_write/sample.txt")
	if err != nil {
		// Fallback: try running from repo root or topic folder
		data, err = os.ReadFile("sample.txt")
	}
	if err != nil {
		fmt.Println("  could not read sample.txt:", err)
	} else {
		fmt.Printf("  %d bytes:\n%s", len(data), data)
	}

	// --- Example 2: write entire file ---
	fmt.Println("\n=== Write new file ===")
	outPath := filepath.Join(dir, "output.txt")
	content := []byte("first line\nsecond line\n")
	err = os.WriteFile(outPath, content, 0644)
	if err != nil {
		fmt.Println("  write error:", err)
		return
	}
	fmt.Println("  wrote", len(content), "bytes to", outPath)

	// --- Example 3: append to file ---
	fmt.Println("\n=== Append to file ===")
	f, err := os.OpenFile(outPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("  open error:", err)
		return
	}
	_, err = f.WriteString("appended third line\n")
	f.Close()
	if err != nil {
		fmt.Println("  write error:", err)
		return
	}
	fmt.Println("  appended one line")

	// --- Example 4: read back ---
	fmt.Println("\n=== Read back ===")
	result, _ := os.ReadFile(outPath)
	fmt.Print(string(result))

	// --- Example 5: streaming read with os.Open ---
	fmt.Println("\n=== Streaming read (os.Open) ===")
	f2, err := os.Open(outPath)
	if err != nil {
		fmt.Println("  open error:", err)
		return
	}
	defer f2.Close()

	buf := make([]byte, 16)
	for {
		n, err := f2.Read(buf)
		if n > 0 {
			fmt.Printf("  chunk (%d bytes): %q\n", n, buf[:n])
		}
		if err != nil {
			break
		}
	}
}
