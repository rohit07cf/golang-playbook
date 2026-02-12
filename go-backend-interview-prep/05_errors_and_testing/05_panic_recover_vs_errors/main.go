package main

import (
	"errors"
	"fmt"
)

// safeDivide recovers from a panic inside a deferred function.
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered panic: %v", r)
		}
	}()
	// integer division by zero panics in Go
	return a / b, nil
}

// mustPositive panics on programmer error (precondition violation).
func mustPositive(n int) int {
	if n <= 0 {
		panic(fmt.Sprintf("mustPositive: got %d, want > 0", n))
	}
	return n
}

// parseAge returns an error for expected bad input.
func parseAge(s string) (int, error) {
	var age int
	_, err := fmt.Sscanf(s, "%d", &age)
	if err != nil {
		return 0, fmt.Errorf("parseAge(%q): %w", s, err)
	}
	if age < 0 || age > 150 {
		return 0, errors.New("age out of range")
	}
	return age, nil
}

func main() {
	// --- Example 1: recover from panic ---
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("safeDivide(10, 0):", err)
	} else {
		fmt.Println("safeDivide(10, 0):", result)
	}

	result, err = safeDivide(10, 3)
	fmt.Println("safeDivide(10, 3):", result)

	// --- Example 2: panic for programmer bugs ---
	fmt.Println("\nmustPositive(5):", mustPositive(5))

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("caught panic:", r)
			}
		}()
		mustPositive(-1) // this panics
	}()

	// --- Example 3: errors for expected failures ---
	age, err := parseAge("25")
	fmt.Println("\nparseAge(\"25\"):", age, err)

	age, err = parseAge("abc")
	fmt.Println("parseAge(\"abc\"):", age, err)

	age, err = parseAge("999")
	fmt.Println("parseAge(\"999\"):", age, err)
}
