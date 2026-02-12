package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- O(n^2): find duplicates with nested loops ---

func findDuplicatesSlow(nums []int) []int {
	var dups []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				dups = append(dups, nums[i])
				break
			}
		}
	}
	return dups
}

// --- O(n): find duplicates with a map ---

func findDuplicatesFast(nums []int) []int {
	seen := make(map[int]bool, len(nums))
	var dups []int
	for _, n := range nums {
		if seen[n] {
			dups = append(dups, n)
		}
		seen[n] = true
	}
	return dups
}

// generateData creates a slice with some duplicates.
func generateData(size int) []int {
	nums := make([]int, size)
	for i := range nums {
		nums[i] = rand.Intn(size / 2) // guarantee duplicates
	}
	return nums
}

func main() {
	sizes := []int{1_000, 5_000, 10_000, 20_000}

	fmt.Println("=== Big-O: O(n^2) vs O(n) duplicate finder ===")
	fmt.Printf("%-10s  %15s  %15s\n", "Size", "O(n^2)", "O(n) map")
	fmt.Println("----------------------------------------------")

	for _, size := range sizes {
		data := generateData(size)

		// O(n^2)
		start := time.Now()
		_ = findDuplicatesSlow(data)
		slow := time.Since(start)

		// O(n)
		start = time.Now()
		_ = findDuplicatesFast(data)
		fast := time.Since(start)

		fmt.Printf("%-10d  %15s  %15s\n", size, slow, fast)
	}

	fmt.Println()
	fmt.Println("Key takeaway: O(n^2) time grows quadratically;")
	fmt.Println("O(n) stays roughly linear. The map trades O(n) space for speed.")
}
