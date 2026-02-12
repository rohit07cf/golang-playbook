package main

import "fmt"

// produce sends values to a send-only channel and closes it.
func produce(out chan<- int, count int) {
	for i := 1; i <= count; i++ {
		out <- i * 10
	}
	close(out) // sender closes
}

// consume reads from a receive-only channel until closed.
func consume(in <-chan int) {
	for val := range in {
		fmt.Printf("  received: %d\n", val)
	}
}

// transform reads from in, doubles values, sends to out.
func transform(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * 2
	}
	close(out)
}

func main() {
	// --- Example 1: producer + consumer ---
	fmt.Println("=== Producer -> Consumer ===")
	ch := make(chan int) // bidirectional, converts to directional

	go produce(ch, 5) // ch converts to chan<- int
	consume(ch)        // ch converts to <-chan int

	// --- Example 2: pipeline with transform ---
	fmt.Println("\n=== Producer -> Transform -> Consumer ===")
	raw := make(chan int)
	doubled := make(chan int)

	go produce(raw, 4)
	go transform(raw, doubled)
	consume(doubled)

	// --- Example 3: return receive-only channel ---
	fmt.Println("\n=== Generator pattern ===")
	nums := generateNums(1, 5)
	for n := range nums {
		fmt.Printf("  generated: %d\n", n)
	}
}

// generateNums returns a receive-only channel that emits values [lo, hi].
func generateNums(lo, hi int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := lo; i <= hi; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
