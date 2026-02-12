package main

import "fmt"

// --- Basic function ---
func greet(name string) string {
	return "hello, " + name
}

// --- Same-type parameter shorthand ---
func add(a, b int) int {
	return a + b
}

// --- No return value ---
func logMessage(msg string) {
	fmt.Println("[LOG]", msg)
}

// --- Multiple types ---
func describe(name string, age int) string {
	return fmt.Sprintf("%s is %d years old", name, age)
}

func main() {
	// --- Calling functions ---
	fmt.Println(greet("world"))
	fmt.Println("3 + 7 =", add(3, 7))
	logMessage("server started")
	fmt.Println(describe("Alice", 30))

	// --- Functions as values ---
	double := func(x int) int {
		return x * 2
	}
	fmt.Println("double(5):", double(5))

	// --- Passing a function as an argument ---
	result := apply(10, double)
	fmt.Println("apply(10, double):", result)

	// --- Inline anonymous function ---
	fmt.Println("inline:", apply(4, func(x int) int {
		return x * x
	}))

	// --- Pass-by-value demonstration ---
	original := 42
	noChange(original)
	fmt.Println("after noChange, original:", original) // still 42
}

// --- Function that takes a function as a parameter ---
func apply(value int, fn func(int) int) int {
	return fn(value)
}

// --- Demonstrates pass-by-value ---
func noChange(x int) {
	x = 999 // modifies the copy, not the original
}
