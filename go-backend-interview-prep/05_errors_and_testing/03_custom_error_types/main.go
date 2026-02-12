package main

import (
	"errors"
	"fmt"
)

// --- Custom error: validation ---

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s -- %s", e.Field, e.Message)
}

// --- Custom error: HTTP-like ---

type HTTPError struct {
	Code    int
	Status  string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Status)
}

// --- Functions that return custom errors ---

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistic value"}
	}
	return nil
}

func fetchUser(id int) error {
	if id <= 0 {
		return &HTTPError{Code: 400, Status: "bad request"}
	}
	if id == 999 {
		return &HTTPError{Code: 404, Status: "not found"}
	}
	return nil
}

func processRequest(id int, age int) error {
	if err := fetchUser(id); err != nil {
		return fmt.Errorf("processRequest: %w", err)
	}
	if err := validateAge(age); err != nil {
		return fmt.Errorf("processRequest: %w", err)
	}
	return nil
}

func main() {
	// --- Check ValidationError with errors.As ---
	err := processRequest(1, -5)
	fmt.Println("Error:", err)

	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("  Field: %s, Message: %s\n", valErr.Field, valErr.Message)
	}

	// --- Check HTTPError with errors.As ---
	err = processRequest(999, 25)
	fmt.Println("\nError:", err)

	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		fmt.Printf("  Code: %d, Status: %s\n", httpErr.Code, httpErr.Status)
	}

	// --- Happy path ---
	err = processRequest(1, 25)
	fmt.Println("\nNo error:", err)
}
