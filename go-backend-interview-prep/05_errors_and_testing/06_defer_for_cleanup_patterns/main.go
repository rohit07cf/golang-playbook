package main

import (
	"fmt"
	"sync"
)

// --- Pattern 1: LIFO order ---

func demoLIFO() {
	fmt.Println("=== LIFO order ===")
	defer fmt.Println("  deferred 1 (first defer, runs last)")
	defer fmt.Println("  deferred 2")
	defer fmt.Println("  deferred 3 (last defer, runs first)")
	fmt.Println("  function body")
}

// --- Pattern 2: mutex unlock ---

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // guaranteed unlock even if panic occurs
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// --- Pattern 3: defer evaluates args immediately ---

func demoArgEval() {
	fmt.Println("\n=== Defer evaluates args immediately ===")
	x := 10
	defer fmt.Printf("  deferred x = %d (captured at defer line)\n", x)
	x = 20
	fmt.Printf("  current x = %d\n", x)
}

// --- Pattern 4: recover in defer ---

func riskyOperation() {
	panic("something went horribly wrong")
}

func safeWrapper() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  recovered:", r)
		}
	}()
	riskyOperation()
}

// --- Pattern 5: named return + defer for error handling ---

func readData() (data string, err error) {
	fmt.Println("\n=== Named return + defer ===")

	// Simulate opening a resource
	fmt.Println("  opened resource")
	defer func() {
		fmt.Println("  closed resource (via defer)")
		// In real code: err = f.Close() to capture close error
	}()

	data = "hello from resource"
	return data, nil
}

func main() {
	demoLIFO()

	// Mutex pattern
	fmt.Println("\n=== Mutex unlock ===")
	counter := &SafeCounter{}
	for i := 0; i < 5; i++ {
		counter.Increment()
	}
	fmt.Println("  counter:", counter.Value())

	demoArgEval()

	// Recover pattern
	fmt.Println("\n=== Recover in defer ===")
	safeWrapper()
	fmt.Println("  program continues after recover")

	// Named return pattern
	data, err := readData()
	fmt.Printf("  data=%q err=%v\n", data, err)
}
