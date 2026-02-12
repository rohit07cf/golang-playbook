package main

import "fmt"

func classify(v any) {
	switch val := v.(type) {
	case string:
		fmt.Printf("  string: %q (len=%d)\n", val, len(val))
	case int:
		fmt.Printf("  int: %d\n", val)
	case float64:
		fmt.Printf("  float64: %.2f\n", val)
	case bool:
		fmt.Printf("  bool: %t\n", val)
	case nil:
		fmt.Println("  nil")
	default:
		fmt.Printf("  unknown: %v (%T)\n", val, val)
	}
}

// --- Multiple types in one case ---
func isNumeric(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64,
		float32, float64:
		return true
	default:
		return false
	}
}

// --- Interface type switch ---
type Stringer interface {
	String() string
}

type Named struct{ Name string }

func (n Named) String() string { return n.Name }

func describeInterface(v any) {
	switch val := v.(type) {
	case Stringer:
		fmt.Printf("  implements Stringer: %s\n", val.String())
	case error:
		fmt.Printf("  implements error: %s\n", val.Error())
	default:
		fmt.Printf("  no known interface: %T\n", val)
	}
}

func main() {
	// --- Basic type switch ---
	fmt.Println("--- Type switch ---")
	values := []any{42, "hello", 3.14, true, nil, []int{1, 2}}
	for _, v := range values {
		classify(v)
	}

	// --- Multiple types in one case ---
	fmt.Println("\n--- isNumeric ---")
	fmt.Println("42:", isNumeric(42))
	fmt.Println("3.14:", isNumeric(3.14))
	fmt.Println("hello:", isNumeric("hello"))

	// --- Interface type switch ---
	fmt.Println("\n--- Interface type switch ---")
	describeInterface(Named{Name: "Alice"})
	describeInterface(fmt.Errorf("something failed"))
	describeInterface(42)
}
