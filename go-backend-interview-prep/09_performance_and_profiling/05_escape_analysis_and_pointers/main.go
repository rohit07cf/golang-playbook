package main

import (
	"fmt"
	"time"
)

// Point is a small struct for demonstration.
type Point struct {
	X, Y float64
}

// --- Does NOT escape: returns value ---
func newPointValue() Point {
	return Point{X: 1.0, Y: 2.0} // stays on stack (returned by value)
}

// --- ESCAPES: returns pointer ---
func newPointPointer() *Point {
	p := Point{X: 1.0, Y: 2.0}
	return &p // escapes to heap (pointer outlives function)
}

// --- ESCAPES: assigned to interface ---
func toInterface(p Point) any {
	return p // escapes: value copied to heap for interface
}

// --- Does NOT escape: used only locally ---
func localOnly() float64 {
	p := Point{X: 3.0, Y: 4.0} // stays on stack
	return p.X + p.Y
}

func main() {
	fmt.Println("=== Escape Analysis and Pointers ===")
	fmt.Println()

	n := 10_000_000

	// Benchmark: value return vs pointer return
	start := time.Now()
	for i := 0; i < n; i++ {
		p := newPointValue()
		_ = p.X
	}
	valueTime := time.Since(start)

	start = time.Now()
	for i := 0; i < n; i++ {
		p := newPointPointer()
		_ = p.X
	}
	ptrTime := time.Since(start)

	start = time.Now()
	for i := 0; i < n; i++ {
		v := localOnly()
		_ = v
	}
	localTime := time.Since(start)

	fmt.Printf("  newPointValue()   (stack):  %s\n", valueTime)
	fmt.Printf("  newPointPointer() (heap):   %s\n", ptrTime)
	fmt.Printf("  localOnly()       (stack):  %s\n", localTime)
	fmt.Println()

	fmt.Println("--- Escape analysis rules ---")
	fmt.Println("  1. Returning a pointer -> escapes to heap")
	fmt.Println("  2. Returning a value   -> stays on stack")
	fmt.Println("  3. Assigning to interface{} -> escapes")
	fmt.Println("  4. Closure capturing a var  -> may escape")
	fmt.Println("  5. Slice growing beyond cap  -> new array on heap")
	fmt.Println()
	fmt.Println("--- Check it yourself ---")
	fmt.Println("  go build -gcflags=\"-m\" ./09_performance_and_profiling/05_escape_analysis_and_pointers/")
	fmt.Println()
	fmt.Println("You'll see output like:")
	fmt.Println("  ./main.go:XX: &p escapes to heap")
	fmt.Println("  ./main.go:XX: p does not escape")
}
