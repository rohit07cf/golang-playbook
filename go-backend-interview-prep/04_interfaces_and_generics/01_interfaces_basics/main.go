package main

import (
	"fmt"
	"math"
)

// --- Interface definition ---
type Shape interface {
	Area() float64
	Perimeter() float64
}

// --- Circle satisfies Shape implicitly ---
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// --- Rectangle also satisfies Shape ---
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// --- Function that accepts the interface ---
func printShape(s Shape) {
	fmt.Printf("  area=%.2f  perimeter=%.2f\n", s.Area(), s.Perimeter())
}

func main() {
	// --- Implicit satisfaction ---
	fmt.Println("--- Interface basics ---")
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 3}

	printShape(c) // Circle satisfies Shape
	printShape(r) // Rectangle satisfies Shape

	// --- Interface variable ---
	fmt.Println("\n--- Interface variable ---")
	var s Shape
	s = c
	fmt.Printf("type holding Circle: area=%.2f\n", s.Area())
	s = r
	fmt.Printf("type holding Rect:   area=%.2f\n", s.Area())

	// --- Slice of interfaces ---
	fmt.Println("\n--- Slice of interfaces ---")
	shapes := []Shape{
		Circle{Radius: 1},
		Rectangle{Width: 4, Height: 5},
		Circle{Radius: 3},
	}
	for _, sh := range shapes {
		printShape(sh)
	}

	// --- Nil interface trap ---
	fmt.Println("\n--- Nil interface trap ---")
	var nilShape Shape
	fmt.Println("nil interface:", nilShape == nil) // true

	var cp *Circle // nil pointer
	nilShape = cp  // storing nil pointer in interface
	fmt.Println("nil pointer in interface:", nilShape == nil) // FALSE!
	// The interface holds (*Circle, nil) -- it is non-nil
}
