package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// simulateWork simulates a small unit of CPU work.
func simulateWork(id int) int {
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += i * id
	}
	return sum
}

// --- Unbounded: one goroutine per task ---

func unbounded(tasks int) {
	var wg sync.WaitGroup
	for i := 0; i < tasks; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			simulateWork(id)
		}(i)
	}
	wg.Wait()
}

// --- Worker pool: fixed goroutines ---

func workerPool(tasks, workers int) {
	jobs := make(chan int, workers*2) // buffered for backpressure
	var wg sync.WaitGroup

	// Start workers
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range jobs {
				simulateWork(id)
			}
		}()
	}

	// Send tasks
	for i := 0; i < tasks; i++ {
		jobs <- i
	}
	close(jobs) // signal workers to exit
	wg.Wait()
}

func main() {
	fmt.Println("=== Concurrency Performance Patterns ===")
	fmt.Println()

	numCPU := runtime.NumCPU()
	fmt.Printf("  CPUs available: %d\n\n", numCPU)

	taskCounts := []int{1_000, 10_000, 100_000}

	fmt.Println("--- Unbounded goroutines vs Worker pool ---")
	fmt.Printf("%-10s  %15s  %15s  %s\n", "Tasks", "Unbounded", "Pool (numCPU)", "Pool workers")
	fmt.Println(repeat('-', 62))

	for _, tasks := range taskCounts {
		// Unbounded
		start := time.Now()
		unbounded(tasks)
		t1 := time.Since(start)

		// Worker pool with numCPU workers
		start = time.Now()
		workerPool(tasks, numCPU)
		t2 := time.Since(start)

		fmt.Printf("%-10d  %15s  %15s  %d\n", tasks, t1, t2, numCPU)
	}
	fmt.Println()

	// Show goroutine overhead
	fmt.Println("--- Goroutine memory overhead ---")
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	baseline := m.HeapAlloc

	// Spawn 10K goroutines that block on a channel
	ch := make(chan struct{})
	for i := 0; i < 10_000; i++ {
		go func() { <-ch }()
	}
	runtime.ReadMemStats(&m)
	perGoroutine := (m.HeapAlloc - baseline) / 10_000
	fmt.Printf("  ~%d bytes per goroutine (rough estimate)\n", perGoroutine)
	close(ch) // release them
	time.Sleep(50 * time.Millisecond)

	fmt.Println()
	fmt.Println("Key: worker pools bound memory + CPU, handle backpressure.")
	fmt.Println("     Unbounded spawning works for small N but risks OOM at scale.")
}

func repeat(b byte, n int) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = b
	}
	return string(s)
}
