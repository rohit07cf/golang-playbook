package main

import (
	"fmt"
	"strings"
)

// ============================================================
// APPROACH 1: INTERFACES (runtime polymorphism, behavior-based)
// ============================================================

// Formatter interface: different types format differently.
type Formatter interface {
	Format() string
}

type PlainText struct{ Content string }

func (p PlainText) Format() string { return p.Content }

type HTMLText struct{ Content string }

func (h HTMLText) Format() string {
	return "<p>" + h.Content + "</p>"
}

// Works with ANY Formatter -- behavior is the contract.
func PrintFormatted(f Formatter) {
	fmt.Println(" ", f.Format())
}

// ============================================================
// APPROACH 2: GENERICS (compile-time, type-based)
// ============================================================

// Reverse works on any slice type -- same logic, different types.
func Reverse[T any](s []T) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}

// Unique removes duplicates -- needs comparable constraint.
func Unique[T comparable](s []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// ============================================================
// WHEN INTERFACE IS BETTER
// ============================================================

// Service layer: behavior-based. Different implementations
// for production, testing, etc.
type UserStore interface {
	GetUser(id int) string
}

type DBStore struct{}

func (d DBStore) GetUser(id int) string {
	return fmt.Sprintf("User-%d (from DB)", id)
}

type MockStore struct{}

func (m MockStore) GetUser(id int) string {
	return fmt.Sprintf("MockUser-%d", id)
}

func greetUser(store UserStore, id int) {
	fmt.Println("  Hello,", store.GetUser(id))
}

func main() {
	// --- Interface approach: behavior polymorphism ---
	fmt.Println("=== INTERFACE: different behavior ===")
	PrintFormatted(PlainText{Content: "hello world"})
	PrintFormatted(HTMLText{Content: "hello world"})

	// --- Generic approach: same logic, different types ---
	fmt.Println("\n=== GENERIC: same logic, different types ===")
	fmt.Println("reversed ints:", Reverse([]int{1, 2, 3, 4, 5}))
	fmt.Println("reversed strs:", Reverse([]string{"a", "b", "c"}))

	fmt.Println("unique ints:", Unique([]int{1, 2, 2, 3, 3, 3}))
	fmt.Println("unique strs:", Unique([]string{"go", "go", "rust"}))

	// --- Interface for service abstraction ---
	fmt.Println("\n=== INTERFACE: service abstraction ===")
	greetUser(DBStore{}, 1)
	greetUser(MockStore{}, 1) // easy to swap for tests

	// --- Decision framework ---
	fmt.Println("\n=== WHEN TO USE WHICH ===")
	decisions := []string{
		"Interface: behavior differs between types (Formatter, Store)",
		"Generic:   logic is identical, types differ (Reverse, Unique)",
		"Interface: dependency injection, mocking, service layers",
		"Generic:   type-safe containers, algorithms, transformations",
		"Default:   start with interface, add generics only if needed",
	}
	for _, d := range decisions {
		fmt.Println(" ", d)
	}
	_ = strings.Builder{} // avoid unused import
}
