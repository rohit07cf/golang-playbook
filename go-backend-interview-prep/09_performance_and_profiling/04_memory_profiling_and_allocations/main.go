package main

import (
	"fmt"
	"runtime"
	"strings"
)

// allocHeavy creates many small allocations (new string per iteration).
func allocHeavy(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "x"
	}
	return s
}

// allocLight preallocates and does one allocation via Builder.
func allocLight(n int) string {
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < n; i++ {
		b.WriteByte('x')
	}
	return b.String()
}

func printMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("  [%s]\n", label)
	fmt.Printf("    HeapAlloc:  %6d KB  (current live heap)\n", m.HeapAlloc/1024)
	fmt.Printf("    TotalAlloc: %6d KB  (cumulative allocated)\n", m.TotalAlloc/1024)
	fmt.Printf("    NumGC:      %6d     (garbage collections)\n", m.NumGC)
	fmt.Printf("    HeapObjects:%6d     (live objects on heap)\n", m.HeapObjects)
	fmt.Println()
}

func main() {
	n := 100_000

	fmt.Println("=== Memory Profiling and Allocations ===")
	fmt.Println()

	// Force GC and take baseline
	runtime.GC()
	printMemStats("baseline after GC")

	// Allocation-heavy approach
	_ = allocHeavy(n)
	printMemStats("after allocHeavy (string += loop)")

	runtime.GC()
	printMemStats("after GC")

	// Allocation-light approach
	_ = allocLight(n)
	printMemStats("after allocLight (strings.Builder)")

	runtime.GC()
	printMemStats("final after GC")

	fmt.Println("--- Key insight ---")
	fmt.Println("  allocHeavy: ~N allocations (new string each +=)")
	fmt.Println("  allocLight: ~1 allocation  (Builder.Grow preallocates)")
	fmt.Println()
	fmt.Println("  Check allocs/op with:")
	fmt.Println("    go test -bench=. -benchmem")
	fmt.Println()
	fmt.Println("  Heap profile:")
	fmt.Println("    go test -bench=. -memprofile=mem.prof")
	fmt.Println("    go tool pprof mem.prof")
	fmt.Println("    > top10 -cum")
}
