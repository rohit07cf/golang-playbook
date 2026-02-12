package main

import "fmt"

// --- Named return values ---
func split(total int) (half, remainder int) {
	half = total / 2
	remainder = total % 2
	return // naked return: returns half and remainder
}

// --- Named returns for documentation ---
func parseCoords(input string) (lat, lng float64, err error) {
	// In real code you would parse the string.
	// Named returns make the API self-documenting.
	lat = 35.6762
	lng = 139.6503
	return
}

// --- Variadic function ---
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// --- Variadic with a fixed first param ---
func greetAll(greeting string, names ...string) {
	for _, name := range names {
		fmt.Println(greeting, name)
	}
}

func main() {
	// --- Named returns ---
	h, r := split(17)
	fmt.Printf("split(17): half=%d remainder=%d\n", h, r)

	lat, lng, _ := parseCoords("ignored")
	fmt.Printf("coords: lat=%.4f lng=%.4f\n", lat, lng)

	// --- Variadic: individual args ---
	fmt.Println("sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("sum(10,20):", sum(10, 20))
	fmt.Println("sum():", sum()) // valid -- returns 0

	// --- Variadic: spread a slice ---
	scores := []int{90, 85, 92, 78}
	fmt.Println("sum(scores...):", sum(scores...))

	// --- Variadic with fixed param ---
	greetAll("Hello", "Alice", "Bob", "Charlie")
}
