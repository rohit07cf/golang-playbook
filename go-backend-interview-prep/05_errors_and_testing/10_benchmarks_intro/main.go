package main

import "fmt"

// Fib computes the n-th Fibonacci number recursively.
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// FibIter computes the n-th Fibonacci number iteratively.
func FibIter(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// StringConcat builds a string by concatenation (slow).
func StringConcat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "a"
	}
	return s
}

func main() {
	fmt.Println("Fib(10) =", Fib(10))
	fmt.Println("FibIter(10) =", FibIter(10))
	fmt.Println("Fib(20) =", Fib(20))
	fmt.Println("FibIter(20) =", FibIter(20))
	fmt.Println("len(StringConcat(100)) =", len(StringConcat(100)))
}
