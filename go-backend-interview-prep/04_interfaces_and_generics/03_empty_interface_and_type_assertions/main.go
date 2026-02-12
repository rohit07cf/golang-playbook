package main

import "fmt"

func main() {
	// --- any accepts anything ---
	fmt.Println("--- any accepts anything ---")
	var v any

	v = 42
	fmt.Printf("int:    %v (type: %T)\n", v, v)

	v = "hello"
	fmt.Printf("string: %v (type: %T)\n", v, v)

	v = true
	fmt.Printf("bool:   %v (type: %T)\n", v, v)

	v = []int{1, 2, 3}
	fmt.Printf("slice:  %v (type: %T)\n", v, v)

	// --- Type assertion: comma-ok form (SAFE) ---
	fmt.Println("\n--- Comma-ok type assertion ---")
	v = "Go is great"

	s, ok := v.(string)
	fmt.Printf("string assertion: val=%q ok=%t\n", s, ok)

	n, ok := v.(int)
	fmt.Printf("int assertion:    val=%d ok=%t\n", n, ok) // ok=false, n=0

	// --- Type assertion: single-value form (DANGEROUS) ---
	fmt.Println("\n--- Single-value assertion ---")
	v = "safe string"
	s2 := v.(string) // works because v IS a string
	fmt.Println("direct assertion:", s2)

	// v.(int) would PANIC here because v holds a string
	// Uncomment to see: _ = v.(int)

	// --- Practical: processing a slice of any ---
	fmt.Println("\n--- Processing []any ---")
	items := []any{42, "hello", 3.14, true, nil}
	for _, item := range items {
		describeItem(item)
	}

	// --- Nil any ---
	fmt.Println("\n--- Nil any ---")
	var nilAny any
	fmt.Printf("nil any: %v (is nil: %t)\n", nilAny, nilAny == nil)
}

func describeItem(v any) {
	if v == nil {
		fmt.Println("  nil")
		return
	}
	if s, ok := v.(string); ok {
		fmt.Printf("  string: %q (len=%d)\n", s, len(s))
		return
	}
	if n, ok := v.(int); ok {
		fmt.Printf("  int: %d\n", n)
		return
	}
	fmt.Printf("  other: %v (%T)\n", v, v)
}
