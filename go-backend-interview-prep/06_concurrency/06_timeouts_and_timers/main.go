package main

import (
	"fmt"
	"time"
)

func main() {
	// --- Example 1: time.After timeout ---
	fmt.Println("=== time.After timeout ===")
	ch := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "result"
	}()

	select {
	case v := <-ch:
		fmt.Println("  received:", v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  timed out after 100ms")
	}

	// --- Example 2: time.NewTimer (reusable) ---
	fmt.Println("\n=== NewTimer ===")
	timer := time.NewTimer(100 * time.Millisecond)

	<-timer.C
	fmt.Println("  timer fired after 100ms")

	// Reset and use again
	timer.Reset(50 * time.Millisecond)
	<-timer.C
	fmt.Println("  timer fired again after 50ms (reset)")

	// --- Example 3: time.NewTicker (periodic) ---
	fmt.Println("\n=== NewTicker (periodic) ===")
	ticker := time.NewTicker(80 * time.Millisecond)
	defer ticker.Stop()

	done := time.After(350 * time.Millisecond)
	count := 0

	for {
		select {
		case t := <-ticker.C:
			count++
			fmt.Printf("  tick %d at %v\n", count, t.Format("15:04:05.000"))
		case <-done:
			fmt.Printf("  stopped after %d ticks\n", count)
			return
		}
	}
}
