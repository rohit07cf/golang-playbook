package main

import (
	"errors"
	"fmt"
	"strconv"
)

// divide returns an error when the divisor is zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// parseAndDouble parses a string to int, doubles it.
func parseAndDouble(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parseAndDouble(%q): %v", s, err)
	}
	return n * 2, nil
}

func main() {
	// --- Example 1: basic error check ---
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// --- Example 2: wrapping context ---
	val, err := parseAndDouble("42")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("parseAndDouble(\"42\") =", val)
	}

	val, err = parseAndDouble("abc")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("parseAndDouble(\"abc\") =", val)
	}

	// --- Example 3: creating errors ---
	e1 := errors.New("something went wrong")
	e2 := fmt.Errorf("failed to load config: %v", e1)
	fmt.Println(e2)
}
