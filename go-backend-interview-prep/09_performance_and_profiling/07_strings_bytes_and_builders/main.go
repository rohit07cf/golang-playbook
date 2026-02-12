package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func concatPlus(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "item" + strconv.Itoa(i) + ","
	}
	return s
}

func concatBuilder(n int) string {
	var b strings.Builder
	b.Grow(n * 10) // estimate capacity
	for i := 0; i < n; i++ {
		b.WriteString("item")
		b.WriteString(strconv.Itoa(i))
		b.WriteByte(',')
	}
	return b.String()
}

func concatJoin(n int) string {
	parts := make([]string, n)
	for i := 0; i < n; i++ {
		parts[i] = "item" + strconv.Itoa(i)
	}
	return strings.Join(parts, ",")
}

// byteStringConversion shows the cost of converting between types.
func byteStringConversion(n int) {
	s := strings.Repeat("hello ", 100)
	start := time.Now()
	for i := 0; i < n; i++ {
		b := []byte(s)   // copy 1: string -> []byte
		_ = string(b)    // copy 2: []byte -> string
	}
	fmt.Printf("  %d string<->[]byte round-trips: %s\n", n, time.Since(start))
}

func main() {
	n := 50_000

	fmt.Println("=== Strings, Bytes, and Builders ===")
	fmt.Println()
	fmt.Printf("Building %d-element string...\n\n", n)

	start := time.Now()
	s1 := concatPlus(n)
	t1 := time.Since(start)

	start = time.Now()
	s2 := concatBuilder(n)
	t2 := time.Since(start)

	start = time.Now()
	s3 := concatJoin(n)
	t3 := time.Since(start)

	fmt.Printf("  += concat:        %s  (len=%d)\n", t1, len(s1))
	fmt.Printf("  strings.Builder:  %s  (len=%d)\n", t2, len(s2))
	fmt.Printf("  strings.Join:     %s  (len=%d)\n", t3, len(s3))
	fmt.Println()

	// Show byte<->string conversion cost
	fmt.Println("--- []byte <-> string conversion cost ---")
	byteStringConversion(1_000_000)
	fmt.Println()

	// fmt.Sprintf vs strconv
	fmt.Println("--- fmt.Sprintf vs strconv ---")
	iters := 1_000_000

	start = time.Now()
	for i := 0; i < iters; i++ {
		_ = fmt.Sprintf("%d", i)
	}
	sprintfTime := time.Since(start)

	start = time.Now()
	for i := 0; i < iters; i++ {
		_ = strconv.Itoa(i)
	}
	strconvTime := time.Since(start)

	fmt.Printf("  fmt.Sprintf: %s\n", sprintfTime)
	fmt.Printf("  strconv.Itoa: %s\n", strconvTime)
	fmt.Println()

	fmt.Println("Key: strings.Builder is 10-100x faster than += in loops.")
	fmt.Println("     strconv is 2-5x faster than fmt.Sprintf for numbers.")
}
