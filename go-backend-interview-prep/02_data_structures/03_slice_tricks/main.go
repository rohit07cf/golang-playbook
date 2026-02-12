package main

import "fmt"

func main() {
	// === DELETE (order preserved) ===
	fmt.Println("--- Delete (keep order) ---")
	s := []int{10, 20, 30, 40, 50}
	i := 2 // delete element at index 2 (value 30)
	s = append(s[:i], s[i+1:]...)
	fmt.Println("after delete index 2:", s) // [10 20 40 50]

	// === DELETE (fast, order not preserved) ===
	fmt.Println("\n--- Delete (swap-with-last) ---")
	s2 := []int{10, 20, 30, 40, 50}
	i = 1 // delete index 1 (value 20)
	s2[i] = s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("after fast delete index 1:", s2) // [10 50 30 40]

	// === INSERT at index ===
	fmt.Println("\n--- Insert at index ---")
	s3 := []int{1, 2, 4, 5}
	idx := 2
	val := 3
	// Make room and insert val at idx
	s3 = append(s3[:idx], append([]int{val}, s3[idx:]...)...)
	fmt.Println("after insert 3 at index 2:", s3) // [1 2 3 4 5]

	// === FILTER in-place ===
	fmt.Println("\n--- Filter (keep odds) ---")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := 0
	for _, v := range data {
		if v%2 != 0 {
			data[n] = v
			n++
		}
	}
	data = data[:n]
	fmt.Println("odds only:", data) // [1 3 5 7]

	// === PREALLOCATE to avoid repeated growth ===
	fmt.Println("\n--- Preallocate ---")
	input := []int{1, 2, 3, 4, 5}
	doubled := make([]int, 0, len(input)) // cap = known size
	for _, v := range input {
		doubled = append(doubled, v*2)
	}
	fmt.Printf("doubled: %v (len=%d cap=%d)\n", doubled, len(doubled), cap(doubled))

	// === COPY (independent) ===
	fmt.Println("\n--- Copy (independent) ---")
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 999
	fmt.Println("src:", src) // unchanged
	fmt.Println("dst:", dst)

	// === REVERSE in-place ===
	fmt.Println("\n--- Reverse ---")
	r := []int{1, 2, 3, 4, 5}
	for left, right := 0, len(r)-1; left < right; left, right = left+1, right-1 {
		r[left], r[right] = r[right], r[left]
	}
	fmt.Println("reversed:", r)

	// === DEDUPLICATE (sorted input) ===
	fmt.Println("\n--- Deduplicate (sorted) ---")
	sorted := []int{1, 1, 2, 3, 3, 3, 4}
	j := 0
	for k := 1; k < len(sorted); k++ {
		if sorted[k] != sorted[j] {
			j++
			sorted[j] = sorted[k]
		}
	}
	sorted = sorted[:j+1]
	fmt.Println("deduped:", sorted) // [1 2 3 4]
}
