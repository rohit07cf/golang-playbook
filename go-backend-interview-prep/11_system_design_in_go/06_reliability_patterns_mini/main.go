package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// --- Errors ---

var (
	ErrTimeout        = errors.New("timeout")
	ErrCircuitOpen    = errors.New("circuit breaker is open")
	ErrServiceFailure = errors.New("service failure")
)

// --- Retry with Exponential Backoff ---

func retry(ctx context.Context, maxAttempts int, base time.Duration, fn func() error) error {
	var lastErr error
	for i := 0; i < maxAttempts; i++ {
		lastErr = fn()
		if lastErr == nil {
			return nil
		}
		if i < maxAttempts-1 {
			// Exponential backoff with jitter
			backoff := base * time.Duration(1<<uint(i))
			jitter := time.Duration(rand.Int63n(int64(backoff / 2)))
			wait := backoff + jitter
			fmt.Printf("  [retry] attempt %d failed: %v, waiting %v\n", i+1, lastErr, wait)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(wait):
			}
		}
	}
	return fmt.Errorf("all %d attempts failed: %w", maxAttempts, lastErr)
}

// --- Timeout ---

func withTimeout(timeout time.Duration, fn func() error) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() { done <- fn() }()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ErrTimeout
	}
}

// --- Circuit Breaker ---

type CircuitState int

const (
	StateClosed   CircuitState = iota // normal operation
	StateOpen                         // failing, reject calls
	StateHalfOpen                     // probing recovery
)

func (s CircuitState) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateOpen:
		return "OPEN"
	case StateHalfOpen:
		return "HALF-OPEN"
	}
	return "UNKNOWN"
}

type CircuitBreaker struct {
	mu               sync.Mutex
	state            CircuitState
	failures         int
	threshold        int
	cooldown         time.Duration
	lastFailureTime  time.Time
}

func NewCircuitBreaker(threshold int, cooldown time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:     StateClosed,
		threshold: threshold,
		cooldown:  cooldown,
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()

	switch cb.state {
	case StateOpen:
		// Check if cooldown has elapsed
		if time.Since(cb.lastFailureTime) > cb.cooldown {
			cb.state = StateHalfOpen
			fmt.Printf("  [circuit] state: OPEN -> HALF-OPEN (probing)\n")
		} else {
			cb.mu.Unlock()
			return ErrCircuitOpen
		}
	}
	cb.mu.Unlock()

	// Execute the function
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failures++
		cb.lastFailureTime = time.Now()
		if cb.failures >= cb.threshold {
			prev := cb.state
			cb.state = StateOpen
			if prev != StateOpen {
				fmt.Printf("  [circuit] state: %s -> OPEN (failures=%d)\n", prev, cb.failures)
			}
		}
		return err
	}

	// Success -- reset
	if cb.state == StateHalfOpen {
		fmt.Printf("  [circuit] state: HALF-OPEN -> CLOSED (recovered)\n")
	}
	cb.failures = 0
	cb.state = StateClosed
	return nil
}

// --- Flaky Service Simulator ---

type FlakyService struct {
	failCount  int
	callCount  int
	slowAfter  int
}

func (s *FlakyService) Call() error {
	s.callCount++
	if s.callCount <= s.failCount {
		return ErrServiceFailure
	}
	if s.slowAfter > 0 && s.callCount > s.slowAfter {
		time.Sleep(2 * time.Second) // simulate slow response
	}
	return nil
}

// --- Demo ---

func main() {
	fmt.Println("=== 1. Retry with Exponential Backoff ===")
	{
		svc := &FlakyService{failCount: 2} // fails first 2 calls
		ctx := context.Background()
		err := retry(ctx, 5, 100*time.Millisecond, svc.Call)
		if err != nil {
			fmt.Printf("retry result: FAILED -- %v\n", err)
		} else {
			fmt.Printf("retry result: SUCCESS on attempt %d\n", svc.callCount)
		}
	}

	fmt.Print("\n=== 2. Timeout ===\n\n")
	{
		// Fast call -- succeeds
		err := withTimeout(500*time.Millisecond, func() error {
			time.Sleep(50 * time.Millisecond)
			return nil
		})
		fmt.Printf("fast call:  %v\n", err)

		// Slow call -- times out
		err = withTimeout(200*time.Millisecond, func() error {
			time.Sleep(1 * time.Second)
			return nil
		})
		fmt.Printf("slow call:  %v\n", err)
	}

	fmt.Print("\n=== 3. Circuit Breaker ===\n\n")
	{
		cb := NewCircuitBreaker(3, 1*time.Second)
		svc := &FlakyService{failCount: 5} // fails first 5 calls

		for i := 1; i <= 8; i++ {
			err := cb.Call(svc.Call)
			fmt.Printf("call %d: err=%v\n", i, err)
			if errors.Is(err, ErrCircuitOpen) {
				fmt.Println("  (circuit is open, call was rejected)")
			}
			time.Sleep(200 * time.Millisecond)
		}

		// Wait for cooldown
		fmt.Println("\nwaiting 1.2s for cooldown...")
		time.Sleep(1200 * time.Millisecond)

		// Service recovers after 5 failures
		err := cb.Call(svc.Call)
		fmt.Printf("call after cooldown: err=%v\n", err)
	}

	fmt.Println("\ndemo done")
}
