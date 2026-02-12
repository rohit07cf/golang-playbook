package main

import (
	"fmt"
	"time"
)

func main() {
	// --- Example 1: ticker-based rate limiter ---
	fmt.Println("=== Ticker-based (steady rate) ===")
	requests := []int{1, 2, 3, 4, 5}

	limiter := time.NewTicker(100 * time.Millisecond) // 10 req/sec
	defer limiter.Stop()

	start := time.Now()
	for _, req := range requests {
		<-limiter.C // wait for next tick
		fmt.Printf("  request %d at %v\n", req, time.Since(start).Round(time.Millisecond))
	}

	// --- Example 2: token bucket (allows bursts) ---
	fmt.Println("\n=== Token bucket (burst-friendly) ===")
	bucket := make(chan struct{}, 3) // capacity 3 = burst of 3

	// Fill the bucket initially
	for i := 0; i < cap(bucket); i++ {
		bucket <- struct{}{}
	}

	// Refill one token every 200ms
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case bucket <- struct{}{}: // add token if room
			default: // bucket full, discard
			}
		}
	}()

	// Process 8 requests
	start2 := time.Now()
	for i := 1; i <= 8; i++ {
		<-bucket // consume a token (blocks if empty)
		fmt.Printf("  request %d at %v\n", i, time.Since(start2).Round(time.Millisecond))
	}
}
