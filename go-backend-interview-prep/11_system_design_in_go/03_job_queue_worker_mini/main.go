package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// --- Job ---

type Job struct {
	ID      int
	Payload string
}

// --- Worker ---

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("[worker %d] processing job %d: %s\n", id, job.ID, job.Payload)
		// Simulate variable work duration
		duration := time.Duration(100+rand.Intn(300)) * time.Millisecond
		time.Sleep(duration)
		fmt.Printf("[worker %d] completed  job %d (took %v)\n", id, job.ID, duration)
	}
	fmt.Printf("[worker %d] shutting down (channel closed)\n", id)
}

// --- Demo ---

func main() {
	const (
		numJobs    = 12
		numWorkers = 3
		queueSize  = 4 // bounded queue -- backpressure when full
	)

	jobs := make(chan Job, queueSize)
	var wg sync.WaitGroup

	// Start worker pool
	fmt.Printf("starting %d workers (queue capacity: %d)\n\n", numWorkers, queueSize)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Producer: enqueue jobs
	fmt.Printf("producing %d jobs...\n\n", numJobs)
	for i := 1; i <= numJobs; i++ {
		job := Job{ID: i, Payload: fmt.Sprintf("send-email-%d", i)}
		fmt.Printf("[producer] queuing job %d (queue len: %d/%d)\n", i, len(jobs), queueSize)
		jobs <- job // blocks if queue is full (backpressure)
	}

	// Close channel -- workers will drain remaining jobs then exit
	close(jobs)
	fmt.Println("\n[producer] all jobs queued, channel closed")
	fmt.Println("[main]     waiting for workers to finish...")

	wg.Wait()
	fmt.Println("\nall workers done -- exiting")
}
