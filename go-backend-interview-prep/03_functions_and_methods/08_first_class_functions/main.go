package main

import "fmt"

// --- Named function type ---
type Transform func(int) int

// --- Apply a function to every element ---
func apply(nums []int, fn Transform) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}

// --- Filter elements ---
func filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// --- Return a function (factory) ---
func multiplier(factor int) Transform {
	return func(n int) int {
		return n * factor
	}
}

// --- Strategy pattern with function types ---
type SortStrategy func(a, b int) bool

func bubbleSort(nums []int, less SortStrategy) []int {
	result := make([]int, len(nums))
	copy(result, nums)
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if less(result[j], result[i]) {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// --- Pass function as argument ---
	fmt.Println("--- Apply (map) ---")
	doubled := apply(nums, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)

	squared := apply(nums, func(n int) int { return n * n })
	fmt.Println("squared:", squared)

	// --- Function factory ---
	fmt.Println("\n--- Multiplier factory ---")
	triple := multiplier(3)
	fmt.Println("tripled:", apply(nums, triple))

	// --- Filter ---
	fmt.Println("\n--- Filter ---")
	evens := filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("evens:", evens)

	// --- Strategy pattern ---
	fmt.Println("\n--- Strategy pattern (sort) ---")
	data := []int{5, 3, 1, 4, 2}
	ascending := bubbleSort(data, func(a, b int) bool { return a < b })
	descending := bubbleSort(data, func(a, b int) bool { return a > b })
	fmt.Println("ascending:", ascending)
	fmt.Println("descending:", descending)

	// --- Storing function in a variable ---
	fmt.Println("\n--- Function variable ---")
	var fn Transform
	fmt.Println("nil function:", fn == nil) // true
	fn = func(n int) int { return n + 100 }
	fmt.Println("fn(5):", fn(5))
}
