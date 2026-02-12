package main

import (
	"errors"
	"fmt"
)

// --- Sentinel error ---
var ErrNotFound = errors.New("not found")

// --- Custom error type ---
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

// --- Functions returning errors ---
func findUser(id int) (string, error) {
	if id <= 0 {
		return "", ErrNotFound
	}
	if id > 100 {
		return "", &ValidationError{Field: "id", Message: "must be <= 100"}
	}
	return "Alice", nil
}

func fetchProfile(id int) (string, error) {
	name, err := findUser(id)
	if err != nil {
		// Wrap with %w to preserve the error chain
		return "", fmt.Errorf("fetchProfile(id=%d): %w", id, err)
	}
	return fmt.Sprintf("Profile: %s", name), nil
}

func main() {
	// --- Basic error handling ---
	fmt.Println("--- Basic error handling ---")
	name, err := findUser(1)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("found:", name)
	}

	// --- Sentinel error ---
	fmt.Println("\n--- Sentinel error ---")
	_, err = findUser(0)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("user not found (sentinel match)")
	}

	// --- Custom error type ---
	fmt.Println("\n--- Custom error type ---")
	_, err = findUser(200)
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("validation: field=%s msg=%s\n", valErr.Field, valErr.Message)
	}

	// --- Wrapped error chain ---
	fmt.Println("\n--- Error wrapping ---")
	_, err = fetchProfile(0)
	fmt.Println("wrapped:", err)

	// errors.Is unwraps the chain
	if errors.Is(err, ErrNotFound) {
		fmt.Println("errors.Is found ErrNotFound in the chain")
	}

	// --- Wrapped custom error ---
	_, err = fetchProfile(200)
	if errors.As(err, &valErr) {
		fmt.Println("errors.As found ValidationError in chain:", valErr.Field)
	}

	// --- error is just an interface ---
	fmt.Println("\n--- error is an interface ---")
	var e error = &ValidationError{Field: "email", Message: "required"}
	fmt.Println("error interface:", e)
	fmt.Printf("concrete type: %T\n", e)
}
