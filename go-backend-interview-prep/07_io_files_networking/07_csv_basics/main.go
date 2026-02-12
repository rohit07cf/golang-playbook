package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// --- Example 1: read CSV from file ---
	fmt.Println("=== Read sample.csv ===")
	f, err := os.Open("07_io_files_networking/07_csv_basics/sample.csv")
	if err != nil {
		f, err = os.Open("sample.csv")
	}
	if err != nil {
		fmt.Println("  could not open sample.csv:", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("  read error:", err)
		return
	}

	header := records[0]
	fmt.Println("  header:", header)
	for _, row := range records[1:] {
		fmt.Printf("  row: name=%s age=%s city=%s\n", row[0], row[1], row[2])
	}

	// --- Example 2: read from string ---
	fmt.Println("\n=== Read CSV from string ===")
	csvData := "product,price\n\"Widget, Large\",9.99\nGadget,4.50"
	r := csv.NewReader(strings.NewReader(csvData))
	for {
		row, err := r.Read()
		if err != nil {
			break
		}
		fmt.Printf("  %v\n", row)
	}

	// --- Example 3: write CSV ---
	fmt.Println("\n=== Write CSV ===")
	dir, _ := os.MkdirTemp("", "csv")
	defer os.RemoveAll(dir)

	outPath := filepath.Join(dir, "output.csv")
	outFile, _ := os.Create(outPath)

	writer := csv.NewWriter(outFile)
	writer.Write([]string{"name", "score"})
	writer.Write([]string{"alice", "95"})
	writer.Write([]string{"bob", "87"})
	writer.Write([]string{"charlie, jr.", "92"}) // comma in field -- auto-quoted
	writer.Flush()
	outFile.Close()

	if err := writer.Error(); err != nil {
		fmt.Println("  write error:", err)
	}

	data, _ := os.ReadFile(outPath)
	fmt.Printf("  written CSV:\n%s", data)
}
