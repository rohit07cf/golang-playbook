package main

import (
	"errors"
	"fmt"
)

// --- The classic (value, error) pattern ---
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// --- Returning two non-error values ---
func minMax(nums ...int) (int, int) {
	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// --- Returning three values ---
func userInfo() (string, int, bool) {
	return "Alice", 30, true
}

func main() {
	// --- Check error properly ---
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// --- Error case ---
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	}

	// --- Discard error with _ (not recommended in production) ---
	quick, _ := divide(100, 4)
	fmt.Println("100 / 4 =", quick)

	// --- Two non-error returns ---
	lo, hi := minMax(3, 1, 4, 1, 5, 9, 2, 6)
	fmt.Printf("min=%d max=%d\n", lo, hi)

	// --- Three return values ---
	name, age, active := userInfo()
	fmt.Printf("user: %s, age: %d, active: %t\n", name, age, active)
}
