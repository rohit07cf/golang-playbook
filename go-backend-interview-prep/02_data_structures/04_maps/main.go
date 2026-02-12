package main

import "fmt"

func main() {
	// --- Create with literal ---
	ages := map[string]int{
		"alice": 30,
		"bob":   25,
	}
	fmt.Println("ages:", ages)

	// --- Create with make ---
	scores := make(map[string]int)
	scores["math"] = 95
	scores["english"] = 88
	fmt.Println("scores:", scores)

	// --- Read: zero value for missing keys ---
	fmt.Println("\n--- Read ---")
	fmt.Println("alice:", ages["alice"])     // 30
	fmt.Println("missing:", ages["nobody"])  // 0 (zero value, no panic)

	// --- Comma-ok idiom ---
	fmt.Println("\n--- Comma-ok ---")
	val, ok := ages["alice"]
	fmt.Printf("alice: val=%d exists=%t\n", val, ok)

	val, ok = ages["nobody"]
	fmt.Printf("nobody: val=%d exists=%t\n", val, ok)

	// --- Write ---
	ages["charlie"] = 35
	fmt.Println("\nafter add charlie:", ages)

	// --- Update ---
	ages["alice"] = 31
	fmt.Println("after update alice:", ages)

	// --- Delete ---
	delete(ages, "bob")
	fmt.Println("after delete bob:", ages)

	// Delete a missing key: no panic, no-op
	delete(ages, "nobody")

	// --- Length ---
	fmt.Println("length:", len(ages))

	// --- Iteration order is RANDOM ---
	fmt.Println("\n--- Iteration (random order) ---")
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}
	for k, v := range colors {
		fmt.Printf("  %s -> %s\n", k, v)
	}
	fmt.Println("(run again -- order may differ)")

	// --- Nil map trap ---
	fmt.Println("\n--- Nil map ---")
	var nilMap map[string]int
	fmt.Println("nil map read:", nilMap["key"]) // 0, no panic
	fmt.Println("nil map len:", len(nilMap))     // 0
	// nilMap["key"] = 1  // PANIC: assignment to entry in nil map

	// --- Count word frequency (common pattern) ---
	fmt.Println("\n--- Word frequency ---")
	words := []string{"go", "is", "great", "go", "is", "fast", "go"}
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++ // zero value of int is 0, so this works cleanly
	}
	for word, count := range freq {
		fmt.Printf("  %s: %d\n", word, count)
	}

	// --- Set pattern (map[T]bool or map[T]struct{}) ---
	fmt.Println("\n--- Set pattern ---")
	seen := make(map[string]struct{})
	items := []string{"a", "b", "a", "c", "b"}
	for _, item := range items {
		seen[item] = struct{}{}
	}
	fmt.Println("unique count:", len(seen))
}
