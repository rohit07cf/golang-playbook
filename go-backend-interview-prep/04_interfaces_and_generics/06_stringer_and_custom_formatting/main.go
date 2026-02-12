package main

import "fmt"

// --- Implement fmt.Stringer ---
type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("%s (age %d)", u.Name, u.Age)
}

// --- GoStringer controls %#v ---
func (u User) GoString() string {
	return fmt.Sprintf("User{Name:%q, Age:%d}", u.Name, u.Age)
}

// --- Another type ---
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// --- Type without Stringer ---
type RawData struct {
	Value int
}

func main() {
	// --- fmt.Stringer in action ---
	fmt.Println("--- fmt.Stringer ---")
	u := User{Name: "Alice", Age: 30}
	fmt.Println(u)              // calls u.String()
	fmt.Printf("user: %v\n", u) // also calls String()
	fmt.Printf("user: %s\n", u) // also calls String()

	// --- GoStringer (%#v) ---
	fmt.Println("\n--- GoStringer ---")
	fmt.Printf("debug: %#v\n", u) // calls u.GoString()

	// --- Custom IP address formatting ---
	fmt.Println("\n--- IPAddr Stringer ---")
	home := IPAddr{127, 0, 0, 1}
	dns := IPAddr{8, 8, 8, 8}
	fmt.Println("home:", home)
	fmt.Println("dns: ", dns)

	// --- Without Stringer: default formatting ---
	fmt.Println("\n--- Without Stringer ---")
	raw := RawData{Value: 42}
	fmt.Println(raw)              // {42}
	fmt.Printf("raw: %v\n", raw) // {42}
	fmt.Printf("raw: %+v\n", raw) // {Value:42}

	// --- Stringer and Sprintf ---
	fmt.Println("\n--- In Sprintf ---")
	msg := fmt.Sprintf("Hello, %s!", u)
	fmt.Println(msg)

	// --- Slice of Stringers ---
	fmt.Println("\n--- Slice of Stringers ---")
	addrs := []IPAddr{{10, 0, 0, 1}, {192, 168, 1, 1}}
	for _, addr := range addrs {
		fmt.Println(" ", addr)
	}
}
