package main

import (
	"cmp"
	"fmt"
)

// --- comparable: allows == and != ---
func Index[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}

// --- Custom constraint: Number ---
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

// --- ~ matches underlying types ---
type UserID int // underlying type is int

// --- Ordered constraint (uses cmp.Ordered from Go 1.21+) ---
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// --- Constraint with method requirement ---
type Stringer interface {
	String() string
}

func PrintAll[T Stringer](items []T) {
	for _, item := range items {
		fmt.Println(" ", item.String())
	}
}

type Tag struct{ Value string }

func (t Tag) String() string { return "#" + t.Value }

func main() {
	// --- comparable ---
	fmt.Println("--- comparable (Index) ---")
	nums := []int{10, 20, 30, 40}
	fmt.Println("index of 30:", Index(nums, 30))
	fmt.Println("index of 99:", Index(nums, 99))

	words := []string{"go", "python", "rust"}
	fmt.Println("index of python:", Index(words, "python"))

	// --- Custom Number constraint ---
	fmt.Println("\n--- Number constraint (Sum) ---")
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("sum ints:", Sum(ints))

	floats := []float64{1.1, 2.2, 3.3}
	fmt.Println("sum floats:", Sum(floats))

	// --- ~ matches named types ---
	ids := []UserID{1, 2, 3}
	fmt.Println("sum UserIDs:", Sum(ids)) // works because ~int matches UserID

	// --- cmp.Ordered (allows < > comparisons) ---
	fmt.Println("\n--- Ordered (Max/Min) ---")
	fmt.Println("max(3, 7):", Max(3, 7))
	fmt.Println("min(3, 7):", Min(3, 7))
	fmt.Println("max(a, z):", Max("a", "z"))

	// --- Constraint with method ---
	fmt.Println("\n--- Stringer constraint ---")
	tags := []Tag{{Value: "go"}, {Value: "generics"}, {Value: "interview"}}
	PrintAll(tags)
}
