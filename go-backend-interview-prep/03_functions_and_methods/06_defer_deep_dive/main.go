package main

import "fmt"

func main() {
	// --- Argument evaluation timing ---
	fmt.Println("--- Arg evaluation ---")
	x := 10
	defer fmt.Println("deferred x:", x) // captures x=10 NOW
	x = 999
	fmt.Println("current x:", x)
	// Output: current x: 999, then deferred x: 10

	// --- Closure captures current value at EXIT ---
	fmt.Println("\n--- Closure vs direct arg ---")
	y := "hello"
	defer func() { fmt.Println("closure y:", y) }() // captures reference
	defer fmt.Println("direct y:", y)                // captures value NOW
	y = "goodbye"
	// direct prints "hello", closure prints "goodbye"

	// --- Defer in loop: THE TRAP ---
	fmt.Println("\n--- Defer in loop (trap) ---")
	deferInLoopTrap()

	// --- Defer in loop: THE FIX ---
	fmt.Println("\n--- Defer in loop (fix) ---")
	deferInLoopFix()

	// --- Named return + defer ---
	fmt.Println("\n--- Named return + defer ---")
	result := enrichedReturn()
	fmt.Println("result:", result)
}

func deferInLoopTrap() {
	// BAD: all defers pile up until function exits
	for i := 0; i < 3; i++ {
		// Imagine: f, _ := os.Open(files[i])
		// defer f.Close()  // NOT released until function returns!
		defer fmt.Printf("  loop defer (all at end): %d\n", i)
	}
	fmt.Println("  after loop (defers haven't run yet)")
}

func deferInLoopFix() {
	// GOOD: extract body into helper so defer runs per iteration
	for i := 0; i < 3; i++ {
		processItem(i)
	}
}

func processItem(i int) {
	defer fmt.Printf("  defer runs immediately after iteration %d\n", i)
	fmt.Printf("  processing %d\n", i)
}

// Named return values can be modified by deferred closures.
func enrichedReturn() (result string) {
	result = "original"
	defer func() {
		result = result + " (enriched by defer)"
	}()
	return result
}
