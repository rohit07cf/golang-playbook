package main

import (
	"context"
	"fmt"
	"time"
)

// slowOperation simulates work that respects context cancellation.
func slowOperation(ctx context.Context, name string) error {
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  %s: cancelled at step %d (%v)\n", name, i, ctx.Err())
			return ctx.Err()
		default:
			fmt.Printf("  %s: step %d\n", name, i)
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Printf("  %s: completed all steps\n", name)
	return nil
}

func main() {
	// --- Example 1: WithCancel ---
	fmt.Println("=== WithCancel ===")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(150 * time.Millisecond)
		cancel() // cancel after 150ms
	}()

	slowOperation(ctx, "job1")

	// --- Example 2: WithTimeout ---
	fmt.Println("\n=== WithTimeout (200ms) ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel2() // always defer cancel, even with timeout

	slowOperation(ctx2, "job2")

	// --- Example 3: parent cancels children ---
	fmt.Println("\n=== Parent cancels children ===")
	parent, parentCancel := context.WithCancel(context.Background())
	child1, child1Cancel := context.WithCancel(parent)
	child2, child2Cancel := context.WithCancel(parent)
	defer child1Cancel()
	defer child2Cancel()

	done := make(chan struct{})
	go func() {
		slowOperation(child1, "child1")
		done <- struct{}{}
	}()
	go func() {
		slowOperation(child2, "child2")
		done <- struct{}{}
	}()

	time.Sleep(120 * time.Millisecond)
	parentCancel() // cancels both children

	<-done
	<-done
	fmt.Println("  both children cancelled")
}
