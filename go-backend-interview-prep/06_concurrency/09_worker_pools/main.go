package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a unit of work.
type Job struct {
	ID    int
	Value int
}

// Result holds the outcome of processing a Job.
type Result struct {
	JobID  int
	Output int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("  worker %d: processing job %d\n", id, job.ID)
		time.Sleep(50 * time.Millisecond) // simulate work
		results <- Result{JobID: job.ID, Output: job.Value * job.Value}
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j}
	}
	close(jobs) // no more jobs

	// Close results after all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	fmt.Println("\n=== Results ===")
	for r := range results {
		fmt.Printf("  job %d -> %d\n", r.JobID, r.Output)
	}
	fmt.Println("all done")
}
