package main

import "fmt"

// --- Interface ---
type Speaker interface {
	Speak() string
}

// --- Dog: value receiver ---
type Dog struct{ Name string }

func (d Dog) Speak() string {
	return d.Name + " says Woof"
}

// --- Cat: pointer receiver ---
type Cat struct{ Name string }

func (c *Cat) Speak() string {
	return c.Name + " says Meow"
}

// --- Mutable type with both receiver types ---
type Counter struct{ N int }

func (c Counter) Value() int { return c.N }  // value receiver
func (c *Counter) Increment()  { c.N++ }     // pointer receiver

type Getter interface{ Value() int }
type Incrementer interface{ Increment() }

func main() {
	// --- Value receiver: both T and *T satisfy the interface ---
	fmt.Println("--- Dog (value receiver) ---")
	var s Speaker

	s = Dog{Name: "Rex"}    // OK: Dog value has Speak()
	fmt.Println(s.Speak())

	s = &Dog{Name: "Buddy"} // OK: *Dog also has Speak()
	fmt.Println(s.Speak())

	// --- Pointer receiver: only *T satisfies the interface ---
	fmt.Println("\n--- Cat (pointer receiver) ---")
	s = &Cat{Name: "Whiskers"} // OK: *Cat has Speak()
	fmt.Println(s.Speak())

	// s = Cat{Name: "Fluffy"}  // COMPILE ERROR:
	// Cat does not implement Speaker (Speak has pointer receiver)

	// --- Demonstrating concrete variable auto-addressing ---
	fmt.Println("\n--- Auto-addressing on concrete variable ---")
	cat := Cat{Name: "Luna"}
	fmt.Println(cat.Speak()) // works: compiler does (&cat).Speak()
	// But you cannot do: Speaker(Cat{Name: "Luna"})

	// --- Method sets and interfaces ---
	fmt.Println("\n--- Counter method sets ---")
	c := Counter{N: 0}

	var g Getter = c     // OK: Counter has Value() (value receiver)
	fmt.Println("value:", g.Value())

	var inc Incrementer = &c // OK: *Counter has Increment() (pointer receiver)
	inc.Increment()
	inc.Increment()
	fmt.Println("after 2 increments:", c.N)

	// var inc2 Incrementer = c // COMPILE ERROR: Counter lacks Increment()
}
