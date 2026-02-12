package main

import "fmt"

func main() {
	// --- Example 1: buffered send without blocking ---
	fmt.Println("=== Buffered channel (cap 3) ===")
	ch := make(chan string, 3)

	// These don't block because buffer has room
	ch <- "a"
	ch <- "b"
	ch <- "c"

	fmt.Println("len:", len(ch), "cap:", cap(ch))

	// Drain
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// --- Example 2: producer faster than consumer ---
	fmt.Println("\n=== Producer/Consumer with buffer ===")
	jobs := make(chan int, 5)

	// Producer: fills buffer quickly
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("  producing %d\n", i)
			jobs <- i
		}
		close(jobs)
	}()

	// Consumer: drains at its own pace
	for j := range jobs {
		fmt.Printf("  consumed %d\n", j)
	}

	// --- Example 3: buffer size 1 as a signal ---
	fmt.Println("\n=== Buffer size 1 (signal) ===")
	done := make(chan bool, 1)

	go func() {
		fmt.Println("  work complete")
		done <- true // non-blocking (buffer has room)
	}()

	<-done
	fmt.Println("  received done signal")

	// --- Example 4: semaphore pattern ---
	fmt.Println("\n=== Semaphore (buffered channel) ===")
	sem := make(chan struct{}, 2) // max 2 concurrent

	for i := 1; i <= 5; i++ {
		sem <- struct{}{} // acquire
		go func(id int) {
			defer func() { <-sem }() // release
			fmt.Printf("  worker %d running (max 2 at a time)\n", id)
		}(i)
	}
	// Wait for all to release
	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
	fmt.Println("  all workers done")
}
