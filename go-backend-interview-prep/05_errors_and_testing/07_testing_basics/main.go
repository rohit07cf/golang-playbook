package main

import "fmt"

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Abs returns the absolute value of n.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// IsEven returns true if n is divisible by 2.
func IsEven(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println("Add(2, 3) =", Add(2, 3))
	fmt.Println("Add(-1, 1) =", Add(-1, 1))
	fmt.Println("Abs(-7) =", Abs(-7))
	fmt.Println("Abs(5) =", Abs(5))
	fmt.Println("IsEven(4) =", IsEven(4))
	fmt.Println("IsEven(7) =", IsEven(7))
}
