package main

import "fmt"

func main() {
	// --- Basic if/else ---
	temperature := 22

	if temperature > 30 {
		fmt.Println("hot")
	} else if temperature > 15 {
		fmt.Println("comfortable")
	} else {
		fmt.Println("cold")
	}

	// --- If with init statement ---
	// The variable 'length' is scoped to this if/else block only.
	if length := len("backend"); length > 5 {
		fmt.Println("long word, length:", length)
	} else {
		fmt.Println("short word, length:", length)
	}
	// 'length' does NOT exist here -- compile error if used

	// --- Common pattern: check-and-use ---
	if result := multiply(3, 7); result > 20 {
		fmt.Println("product is large:", result)
	}

	// --- No ternary in Go ---
	// In other languages: status = age >= 18 ? "adult" : "minor"
	// In Go, you must use if/else:
	age := 20
	var status string
	if age >= 18 {
		status = "adult"
	} else {
		status = "minor"
	}
	fmt.Println("status:", status)

	// --- Boolean conditions ---
	loggedIn := true
	isAdmin := false

	if loggedIn && isAdmin {
		fmt.Println("admin dashboard")
	} else if loggedIn {
		fmt.Println("user dashboard")
	} else {
		fmt.Println("please log in")
	}
}

func multiply(a, b int) int {
	return a * b
}
