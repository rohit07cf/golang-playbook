package main

import "fmt"

func main() {
	// --- Basic closure: captures outer variable ---
	counter := 0
	increment := func() int {
		counter++ // modifies the outer variable
		return counter
	}
	fmt.Println("call 1:", increment()) // 1
	fmt.Println("call 2:", increment()) // 2
	fmt.Println("call 3:", increment()) // 3
	fmt.Println("counter:", counter)    // 3

	// --- Closure as return value (function factory) ---
	fmt.Println("\n--- Function factory ---")
	addFive := makeAdder(5)
	addTen := makeAdder(10)
	fmt.Println("addFive(3):", addFive(3))   // 8
	fmt.Println("addTen(3):", addTen(3))     // 13

	// --- Closure with state ---
	fmt.Println("\n--- Stateful closure ---")
	next := counter_gen()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	// --- Loop variable trap (pre-Go 1.22) ---
	fmt.Println("\n--- Loop variable trap ---")
	funcs := make([]func(), 5)
	for i := 0; i < 5; i++ {
		funcs[i] = func() {
			fmt.Print(i, " ") // captures i by reference
		}
	}
	fmt.Print("trap (Go 1.22+ fixes this): ")
	for _, fn := range funcs {
		fn()
	}
	fmt.Println()

	// --- Fix: shadow the variable ---
	fmt.Print("fixed: ")
	funcs2 := make([]func(), 5)
	for i := 0; i < 5; i++ {
		i := i // shadow: creates new variable per iteration
		funcs2[i] = func() {
			fmt.Print(i, " ")
		}
	}
	for _, fn := range funcs2 {
		fn()
	}
	fmt.Println()
}

// makeAdder returns a closure that adds x to its argument.
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// counter_gen returns a closure that counts up from 1.
func counter_gen() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
