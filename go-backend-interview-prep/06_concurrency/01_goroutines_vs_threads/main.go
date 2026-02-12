package main

import (
	"fmt"
	"sync"
	"time"
)

// work simulates a task with an ID.
func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("goroutine %d: started\n", id)
	time.Sleep(100 * time.Millisecond) // simulate work
	fmt.Printf("goroutine %d: done\n", id)
}

func main() {
	// --- Example 1: launching goroutines ---
	fmt.Println("=== Launching goroutines ===")
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()
	fmt.Println("all goroutines finished")

	// --- Example 2: anonymous goroutine ---
	fmt.Println("\n=== Anonymous goroutine ===")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello from anonymous goroutine")
	}()
	wg.Wait()

	// --- Example 3: goroutine cost demonstration ---
	fmt.Println("\n=== Spawning 10,000 goroutines ===")
	start := time.Now()
	var wg2 sync.WaitGroup

	for i := 0; i < 10_000; i++ {
		wg2.Add(1)
		go func(n int) {
			defer wg2.Done()
			_ = n * n // trivial work
		}(i)
	}
	wg2.Wait()
	fmt.Printf("10,000 goroutines completed in %v\n", time.Since(start))
}
