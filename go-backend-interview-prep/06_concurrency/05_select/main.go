package main

import (
	"fmt"
	"time"
)

func main() {
	// --- Example 1: select on two channels ---
	fmt.Println("=== Select on two channels ===")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "two"
	}()

	// Receive both (ch1 should arrive first)
	for i := 0; i < 2; i++ {
		select {
		case v := <-ch1:
			fmt.Println("  from ch1:", v)
		case v := <-ch2:
			fmt.Println("  from ch2:", v)
		}
	}

	// --- Example 2: timeout with time.After ---
	fmt.Println("\n=== Timeout ===")
	slow := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond) // too slow
		slow <- "finally"
	}()

	select {
	case v := <-slow:
		fmt.Println("  received:", v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  timed out after 100ms")
	}

	// --- Example 3: non-blocking with default ---
	fmt.Println("\n=== Non-blocking (default) ===")
	ch := make(chan int, 1)

	// Non-blocking receive
	select {
	case v := <-ch:
		fmt.Println("  received:", v)
	default:
		fmt.Println("  nothing ready (non-blocking)")
	}

	// Non-blocking send
	ch <- 42
	select {
	case ch <- 99:
		fmt.Println("  sent 99")
	default:
		fmt.Println("  channel full, skipped send")
	}
	fmt.Println("  value in ch:", <-ch)

	// --- Example 4: done channel pattern ---
	fmt.Println("\n=== Done channel ===")
	done := make(chan struct{})
	data := make(chan int)

	go func() {
		defer close(data)
		for i := 1; ; i++ {
			select {
			case data <- i:
			case <-done:
				fmt.Println("  producer: received done signal")
				return
			}
		}
	}()

	// Take 3 values then signal done
	for i := 0; i < 3; i++ {
		fmt.Println("  consumed:", <-data)
	}
	close(done) // signal producer to stop

	time.Sleep(50 * time.Millisecond) // let producer print
}
