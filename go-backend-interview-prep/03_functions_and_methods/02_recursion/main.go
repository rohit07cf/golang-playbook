package main

import "fmt"

// --- Factorial (recursive) ---
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// --- Factorial (iterative, preferred for large n) ---
func factorialIter(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// --- Fibonacci (naive recursive, O(2^n)) ---
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// --- Fibonacci (iterative, O(n)) ---
func fibIter(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// --- Sum of digits (recursive) ---
func sumDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumDigits(n/10)
}

func main() {
	// --- Factorial ---
	fmt.Println("--- Factorial ---")
	for _, n := range []int{0, 1, 5, 10} {
		fmt.Printf("factorial(%d) = %d\n", n, factorial(n))
	}

	// --- Iterative comparison ---
	fmt.Println("\n--- Factorial (iterative) ---")
	fmt.Println("factorialIter(10):", factorialIter(10))

	// --- Fibonacci ---
	fmt.Println("\n--- Fibonacci (recursive) ---")
	for i := 0; i < 10; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fib(i))
	}

	// --- Fibonacci (iterative) ---
	fmt.Println("\n--- Fibonacci (iterative) ---")
	fmt.Println("fibIter(30):", fibIter(30))

	// --- Sum of digits ---
	fmt.Println("\n--- Sum of digits ---")
	fmt.Println("sumDigits(12345):", sumDigits(12345)) // 15
}
