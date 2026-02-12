package main

import (
	"errors"
	"fmt"
)

// --- Sentinel error ---
var ErrNotFound = errors.New("not found")

// --- Custom error type ---
type PermissionError struct {
	User   string
	Action string
}

func (e *PermissionError) Error() string {
	return fmt.Sprintf("user %q cannot %s", e.User, e.Action)
}

// --- Layered functions that wrap errors ---

func findRecord(id int) error {
	if id <= 0 {
		return ErrNotFound
	}
	if id == 99 {
		return &PermissionError{User: "guest", Action: "read"}
	}
	return nil
}

func getProfile(id int) error {
	err := findRecord(id)
	if err != nil {
		return fmt.Errorf("getProfile(id=%d): %w", id, err)
	}
	return nil
}

func handleRequest(id int) error {
	err := getProfile(id)
	if err != nil {
		return fmt.Errorf("handleRequest: %w", err)
	}
	return nil
}

func main() {
	// --- Example 1: errors.Is walks the chain ---
	err := handleRequest(0)
	fmt.Println("Error:", err)
	fmt.Println("Is ErrNotFound?", errors.Is(err, ErrNotFound))

	// --- Example 2: errors.As extracts a typed error ---
	err = handleRequest(99)
	fmt.Println("\nError:", err)

	var permErr *PermissionError
	if errors.As(err, &permErr) {
		fmt.Printf("Permission denied: user=%s action=%s\n", permErr.User, permErr.Action)
	}

	// --- Example 3: %v vs %w ---
	base := errors.New("disk full")
	withW := fmt.Errorf("save file: %w", base)
	withV := fmt.Errorf("save file: %v", base)

	fmt.Println("\n%w preserves chain:", errors.Is(withW, base))
	fmt.Println("%v breaks chain:   ", errors.Is(withV, base))
}
