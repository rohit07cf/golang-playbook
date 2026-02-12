package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // always defer
	fmt.Printf("  worker %d: started\n", id)
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("  worker %d: done\n", id)
}

func main() {
	// --- Example 1: basic WaitGroup ---
	fmt.Println("=== Basic WaitGroup ===")
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Add BEFORE go
		go worker(i, &wg)
	}
	wg.Wait() // blocks until all Done()
	fmt.Println("all workers finished")

	// --- Example 2: WaitGroup + channel for results ---
	fmt.Println("\n=== WaitGroup + channel for results ===")
	results := make(chan int, 5)

	var wg2 sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg2.Add(1)
		go func(n int) {
			defer wg2.Done()
			results <- n * n
		}(i)
	}

	// Close channel after all goroutines finish
	go func() {
		wg2.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("  result: %d\n", r)
	}

	// --- Example 3: Add(n) in bulk ---
	fmt.Println("\n=== Add(n) in bulk ===")
	var wg3 sync.WaitGroup
	n := 3
	wg3.Add(n) // add all at once

	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg3.Done()
			fmt.Printf("  task %d complete\n", id)
		}(i)
	}
	wg3.Wait()
	fmt.Println("all tasks done")
}
