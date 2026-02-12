package main

import (
	"fmt"
	"strconv"
	"time"
)

func sliceNoPrealloc(n int) []int {
	var s []int // nil slice, cap=0
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}

func slicePrealloc(n int) []int {
	s := make([]int, 0, n) // cap=n
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}

func mapNoHint(n int) map[string]int {
	m := map[string]int{}
	for i := 0; i < n; i++ {
		m[strconv.Itoa(i)] = i
	}
	return m
}

func mapWithHint(n int) map[string]int {
	m := make(map[string]int, n) // hint
	for i := 0; i < n; i++ {
		m[strconv.Itoa(i)] = i
	}
	return m
}

func main() {
	fmt.Println("=== Maps and Slices -- Performance Gotchas ===")
	fmt.Println()

	// --- Slice benchmark ---
	sizes := []int{10_000, 100_000, 1_000_000}

	fmt.Println("--- Slice: no prealloc vs make([]T, 0, n) ---")
	fmt.Printf("%-12s  %15s  %15s\n", "Size", "No prealloc", "Preallocated")
	fmt.Println(repeat("-", 46))

	for _, n := range sizes {
		start := time.Now()
		_ = sliceNoPrealloc(n)
		t1 := time.Since(start)

		start = time.Now()
		_ = slicePrealloc(n)
		t2 := time.Since(start)

		fmt.Printf("%-12d  %15s  %15s\n", n, t1, t2)
	}
	fmt.Println()

	// --- Map benchmark ---
	mapSizes := []int{10_000, 100_000}

	fmt.Println("--- Map: no hint vs make(map, n) ---")
	fmt.Printf("%-12s  %15s  %15s\n", "Size", "No hint", "With hint")
	fmt.Println(repeat("-", 46))

	for _, n := range mapSizes {
		start := time.Now()
		_ = mapNoHint(n)
		t1 := time.Since(start)

		start = time.Now()
		_ = mapWithHint(n)
		t2 := time.Since(start)

		fmt.Printf("%-12d  %15s  %15s\n", n, t1, t2)
	}
	fmt.Println()

	// --- Slice growth demo ---
	fmt.Println("--- Slice growth pattern ---")
	s := make([]int, 0)
	prevCap := 0
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if cap(s) != prevCap {
			fmt.Printf("  len=%-3d cap=%-5d (grew from %d)\n", len(s), cap(s), prevCap)
			prevCap = cap(s)
		}
	}
	fmt.Println()

	fmt.Println("Key: preallocating avoids repeated grow + copy.")
	fmt.Println("     Map hint avoids rehashing during growth.")
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
