package main

import (
	"fmt"
	"strconv"
)

// ParseIntSafe parses s as base-10 int. Returns fallback on failure.
func ParseIntSafe(s string, fallback int) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return fallback, fmt.Errorf("ParseIntSafe(%q): %w", s, err)
	}
	return n, nil
}

// Clamp restricts n to the range [lo, hi].
func Clamp(n, lo, hi int) int {
	if n < lo {
		return lo
	}
	if n > hi {
		return hi
	}
	return n
}

func main() {
	// ParseIntSafe examples
	v, err := ParseIntSafe("42", 0)
	fmt.Printf("ParseIntSafe(\"42\", 0) = %d, err=%v\n", v, err)

	v, err = ParseIntSafe("abc", -1)
	fmt.Printf("ParseIntSafe(\"abc\", -1) = %d, err=%v\n", v, err)

	// Clamp examples
	fmt.Println("Clamp(5, 0, 10) =", Clamp(5, 0, 10))
	fmt.Println("Clamp(-3, 0, 10) =", Clamp(-3, 0, 10))
	fmt.Println("Clamp(15, 0, 10) =", Clamp(15, 0, 10))
}
