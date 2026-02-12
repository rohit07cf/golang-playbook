package main

import "fmt"

type Rect struct {
	Width, Height float64
}

// Value receiver: operates on a copy.
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// Value receiver: cannot mutate the original.
func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Pointer receiver: operates on the original.
func (r *Rect) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Pointer receiver: mutates the original.
func (r *Rect) SetWidth(w float64) {
	r.Width = w
}

// --- Named type to add methods to a built-in type ---
type StringSlice []string

func (ss StringSlice) Join(sep string) string {
	result := ""
	for i, s := range ss {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}

func main() {
	// --- Value receiver methods ---
	r := Rect{Width: 10, Height: 5}
	fmt.Println("area:", r.Area())
	fmt.Println("perimeter:", r.Perimeter())

	// --- Pointer receiver: mutates ---
	fmt.Println("\n--- Pointer receiver ---")
	fmt.Println("before scale:", r)
	r.Scale(2)
	fmt.Println("after Scale(2):", r) // {20, 10}

	r.SetWidth(100)
	fmt.Println("after SetWidth(100):", r)

	// --- Value receiver CANNOT mutate ---
	fmt.Println("\n--- Value receiver cannot mutate ---")
	r2 := Rect{Width: 3, Height: 4}
	tryToMutateValue(r2)
	fmt.Println("r2 unchanged:", r2)

	// --- Auto address-taking ---
	// Go lets you call a pointer-receiver method on a value.
	// The compiler automatically takes the address.
	fmt.Println("\n--- Auto address-taking ---")
	val := Rect{Width: 5, Height: 5}
	val.Scale(3) // compiler does (&val).Scale(3)
	fmt.Println("after val.Scale(3):", val)

	// --- Method on named type ---
	fmt.Println("\n--- Named type method ---")
	tags := StringSlice{"go", "python", "rust"}
	fmt.Println("joined:", tags.Join(", "))
}

func tryToMutateValue(r Rect) {
	r.Width = 999 // modifies the copy
}
