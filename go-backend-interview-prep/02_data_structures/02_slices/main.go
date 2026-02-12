package main

import "fmt"

func main() {
	// --- Slice literal ---
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("nums:", nums)

	// --- Length vs Capacity ---
	s := make([]int, 3, 8)
	fmt.Printf("make([]int, 3, 8): len=%d cap=%d %v\n", len(s), cap(s), s)

	// --- Append (always reassign!) ---
	s = append(s, 99)
	fmt.Printf("after append:      len=%d cap=%d %v\n", len(s), cap(s), s)

	// --- Append past capacity triggers reallocation ---
	small := make([]int, 0, 2)
	fmt.Printf("\nsmall before: len=%d cap=%d\n", len(small), cap(small))
	small = append(small, 1, 2, 3) // exceeds cap=2
	fmt.Printf("small after:  len=%d cap=%d\n", len(small), cap(small))

	// --- Sub-slicing shares the backing array ---
	fmt.Println("\n--- Shared backing array ---")
	original := []int{1, 2, 3, 4, 5}
	sub := original[1:4] // [2, 3, 4]
	fmt.Println("original:", original)
	fmt.Println("sub:     ", sub)

	sub[0] = 999 // also modifies original!
	fmt.Println("after sub[0] = 999:")
	fmt.Println("original:", original) // [1, 999, 3, 4, 5]
	fmt.Println("sub:     ", sub)      // [999, 3, 4]

	// --- Copy to make independent slice ---
	fmt.Println("\n--- Independent copy ---")
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 999
	fmt.Println("src:", src) // unchanged
	fmt.Println("dst:", dst)

	// --- Nil vs empty slice ---
	fmt.Println("\n--- Nil vs empty ---")
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("nil slice:   %v  len=%d  isNil=%t\n", nilSlice, len(nilSlice), nilSlice == nil)
	fmt.Printf("empty slice: %v  len=%d  isNil=%t\n", emptySlice, len(emptySlice), emptySlice == nil)

	// --- Append works on nil slices ---
	var data []int
	data = append(data, 1, 2, 3)
	fmt.Println("\nappend to nil:", data)

	// --- Slice of slice: 3-index syntax limits capacity ---
	fmt.Println("\n--- Three-index slice ---")
	base := []int{0, 1, 2, 3, 4, 5}
	limited := base[1:3:3] // [1, 2] with cap=2 (cannot see beyond index 3)
	fmt.Printf("limited: %v len=%d cap=%d\n", limited, len(limited), cap(limited))
}
