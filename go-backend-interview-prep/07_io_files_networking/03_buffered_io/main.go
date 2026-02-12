package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// --- Example 1: Scanner from a string ---
	fmt.Println("=== Scanner from string ===")
	input := "line one\nline two\nline three"
	scanner := bufio.NewScanner(strings.NewReader(input))
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		fmt.Printf("  %d: %s\n", lineNum, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("  scan error:", err)
	}

	// --- Example 2: Scanner from file ---
	fmt.Println("\n=== Scanner from file ===")
	f, err := os.Open("07_io_files_networking/03_buffered_io/sample.txt")
	if err != nil {
		f, err = os.Open("sample.txt")
	}
	if err != nil {
		fmt.Println("  could not open sample.txt:", err)
	} else {
		defer f.Close()
		s := bufio.NewScanner(f)
		for s.Scan() {
			fmt.Println(" ", s.Text())
		}
		if err := s.Err(); err != nil {
			fmt.Println("  scan error:", err)
		}
	}

	// --- Example 3: buffered writer ---
	fmt.Println("\n=== Buffered writer ===")
	dir, _ := os.MkdirTemp("", "bufio")
	defer os.RemoveAll(dir)

	outPath := filepath.Join(dir, "buffered.txt")
	outFile, err := os.Create(outPath)
	if err != nil {
		fmt.Println("  create error:", err)
		return
	}

	w := bufio.NewWriter(outFile)
	w.WriteString("buffered line 1\n")
	w.WriteString("buffered line 2\n")
	w.WriteString("buffered line 3\n")

	fmt.Println("  buffered (not yet flushed), bytes in buffer:", w.Buffered())
	w.Flush() // <-- MUST flush before close
	outFile.Close()

	// Read back
	data, _ := os.ReadFile(outPath)
	fmt.Printf("  file contents:\n%s", data)

	// --- Example 4: word-by-word scanning ---
	fmt.Println("\n=== Word scanner ===")
	ws := bufio.NewScanner(strings.NewReader("the quick brown fox"))
	ws.Split(bufio.ScanWords)
	for ws.Scan() {
		fmt.Printf("  word: %q\n", ws.Text())
	}
}
