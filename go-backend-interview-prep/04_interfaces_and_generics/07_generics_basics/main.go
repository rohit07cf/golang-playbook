package main

import "fmt"

// --- Generic function: Map ---
func Map[T any, U any](s []T, fn func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

// --- Generic function: Filter ---
func Filter[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// --- Generic function: Contains (needs comparable) ---
func Contains[T comparable](s []T, target T) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

// --- Generic type: Stack ---
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func main() {
	// --- Generic Map ---
	fmt.Println("--- Map ---")
	nums := []int{1, 2, 3, 4, 5}
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)

	strs := Map(nums, func(n int) string { return fmt.Sprintf("#%d", n) })
	fmt.Println("as strings:", strs)

	// --- Generic Filter ---
	fmt.Println("\n--- Filter ---")
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("evens:", evens)

	words := []string{"Go", "Python", "Rust", "Go"}
	long := Filter(words, func(s string) bool { return len(s) > 2 })
	fmt.Println("long words:", long)

	// --- Generic Contains ---
	fmt.Println("\n--- Contains ---")
	fmt.Println("has 3:", Contains(nums, 3))
	fmt.Println("has 9:", Contains(nums, 9))
	fmt.Println("has Go:", Contains(words, "Go"))

	// --- Generic Stack ---
	fmt.Println("\n--- Stack[int] ---")
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println("len:", intStack.Len())
	v, ok := intStack.Pop()
	fmt.Printf("pop: %d (ok=%t)\n", v, ok)

	fmt.Println("\n--- Stack[string] ---")
	var strStack Stack[string]
	strStack.Push("hello")
	strStack.Push("world")
	s, ok := strStack.Pop()
	fmt.Printf("pop: %q (ok=%t)\n", s, ok)
}
