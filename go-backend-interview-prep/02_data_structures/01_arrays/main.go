package main

import "fmt"

func main() {
	// --- Declaration: zero-valued ---
	var nums [5]int
	fmt.Println("zero-valued:", nums) // [0 0 0 0 0]

	// --- Declaration: initialized ---
	colors := [3]string{"red", "green", "blue"}
	fmt.Println("colors:", colors)

	// --- Compiler counts the size ---
	primes := [...]int{2, 3, 5, 7, 11}
	fmt.Println("primes:", primes)
	fmt.Println("length:", len(primes))

	// --- Access and modify ---
	nums[0] = 10
	nums[4] = 50
	fmt.Println("modified:", nums)

	// --- Arrays are VALUE types ---
	// Assigning copies the entire array.
	original := [3]int{1, 2, 3}
	copied := original
	copied[0] = 999

	fmt.Println("original:", original) // [1 2 3] -- unchanged
	fmt.Println("copied:", copied)     // [999 2 3]

	// --- Passing to a function copies too ---
	data := [3]int{10, 20, 30}
	tryToModify(data)
	fmt.Println("after function call:", data) // [10 20 30] -- unchanged

	// --- Iterate with range ---
	fmt.Println("--- iterate ---")
	for i, v := range colors {
		fmt.Printf("  index=%d value=%s\n", i, v)
	}

	// --- 2D array ---
	grid := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D grid:", grid)
}

func tryToModify(arr [3]int) {
	arr[0] = 999 // modifies the local copy, not the original
}
