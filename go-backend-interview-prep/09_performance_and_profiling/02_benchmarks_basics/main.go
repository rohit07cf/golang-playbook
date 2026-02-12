package main

import (
	"fmt"
	"strings"
	"time"
)

// concatStrings uses += in a loop (slow, many allocations).
func concatStrings(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}

// buildStrings uses strings.Builder (fast, one allocation).
func buildStrings(n int) string {
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < n; i++ {
		b.WriteByte('a')
	}
	return b.String()
}

func main() {
	n := 50_000

	fmt.Println("=== String concat vs strings.Builder ===")
	fmt.Printf("Building a %d-char string...\n\n", n)

	// Concat
	start := time.Now()
	s1 := concatStrings(n)
	concatTime := time.Since(start)

	// Builder
	start = time.Now()
	s2 := buildStrings(n)
	builderTime := time.Since(start)

	fmt.Printf("  += concat:       %s  (len=%d)\n", concatTime, len(s1))
	fmt.Printf("  strings.Builder: %s  (len=%d)\n", builderTime, len(s2))
	fmt.Println()
	fmt.Println("Note: This is a rough timing. For accurate results, use:")
	fmt.Println("  go test -bench=. -benchmem ./09_performance_and_profiling/02_benchmarks_basics/")
	fmt.Println()
	fmt.Println("Benchmark output format:")
	fmt.Println("  BenchmarkConcat-8   1000   150000 ns/op   1234567 B/op   49999 allocs/op")
	fmt.Println("                       ^N     ^time          ^bytes         ^heap allocs")
}
