package main

import "fmt"

func main() {
	// --- Example 1: basic send and receive ---
	fmt.Println("=== Basic send/receive ===")
	ch := make(chan string)

	go func() {
		ch <- "hello from goroutine"
	}()

	msg := <-ch // blocks until value is sent
	fmt.Println(msg)

	// --- Example 2: close + range ---
	fmt.Println("\n=== Close + range ===")
	nums := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			nums <- i
		}
		close(nums) // signal: no more values
	}()

	for n := range nums { // loops until channel is closed
		fmt.Printf("  received: %d\n", n)
	}

	// --- Example 3: comma-ok pattern ---
	fmt.Println("\n=== Comma-ok receive ===")
	ch2 := make(chan int, 1)
	ch2 <- 42
	close(ch2)

	v, ok := <-ch2
	fmt.Printf("  v=%d, ok=%v (value available)\n", v, ok)

	v, ok = <-ch2
	fmt.Printf("  v=%d, ok=%v (channel closed, zero value)\n", v, ok)

	// --- Example 4: channel as function return ---
	fmt.Println("\n=== Channel as return value ===")
	result := compute(10, 20)
	fmt.Println("  10 + 20 =", <-result)
}

// compute returns a channel that will receive the sum.
func compute(a, b int) <-chan int {
	ch := make(chan int)
	go func() {
		ch <- a + b
	}()
	return ch
}
