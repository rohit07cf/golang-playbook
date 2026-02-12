package main

import (
	"fmt"
	"sync"
	"time"
)

// generate emits integers [lo, hi] on the returned channel.
func generate(lo, hi int) <-chan int {
	out := make(chan int)
	go func() {
		for i := lo; i <= hi; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// square reads ints, squares them, sends to output.
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			time.Sleep(20 * time.Millisecond) // simulate work
			out <- v * v
		}
		close(out)
	}()
	return out
}

// filterEven keeps only even numbers.
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if v%2 == 0 {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

// fanOut spawns n workers that all read from the same input channel.
func fanOut(in <-chan int, n int) []<-chan int {
	outs := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		outs[i] = square(in) // each reads from shared 'in'
	}
	return outs
}

// fanIn merges multiple channels into one.
func fanIn(channels []<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				merged <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	// --- Simple pipeline: generate -> square -> filterEven ---
	fmt.Println("=== Simple pipeline ===")
	nums := generate(1, 10)
	squared := square(nums)
	evens := filterEven(squared)

	for v := range evens {
		fmt.Printf("  %d\n", v)
	}

	// --- Fan-out / Fan-in: 3 workers squaring in parallel ---
	fmt.Println("\n=== Fan-out (3 workers) -> Fan-in ===")
	start := time.Now()
	nums2 := generate(1, 12)
	workers := fanOut(nums2, 3)
	merged := fanIn(workers)

	for v := range merged {
		fmt.Printf("  %d\n", v)
	}
	fmt.Printf("  completed in %v\n", time.Since(start).Round(time.Millisecond))
}
