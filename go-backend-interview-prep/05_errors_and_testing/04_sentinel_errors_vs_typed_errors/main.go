package main

import (
	"errors"
	"fmt"
)

// ---- Sentinel errors ----

var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

// ---- Typed error ----

type RateLimitError struct {
	Limit      int
	RetryAfter int // seconds
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limited: max %d req/s, retry after %ds", e.Limit, e.RetryAfter)
}

// ---- Functions ----

func lookupUser(id int) (string, error) {
	switch id {
	case 0:
		return "", ErrNotFound
	case -1:
		return "", ErrUnauthorized
	case 99:
		return "", &RateLimitError{Limit: 100, RetryAfter: 30}
	default:
		return fmt.Sprintf("user_%d", id), nil
	}
}

func getProfile(id int) (string, error) {
	name, err := lookupUser(id)
	if err != nil {
		return "", fmt.Errorf("getProfile(%d): %w", id, err)
	}
	return name, nil
}

func main() {
	ids := []int{1, 0, -1, 99}

	for _, id := range ids {
		name, err := getProfile(id)
		if err != nil {
			fmt.Printf("id=%d  error: %v\n", id, err)

			// Sentinel check with errors.Is
			if errors.Is(err, ErrNotFound) {
				fmt.Println("  -> sentinel: not found")
			}
			if errors.Is(err, ErrUnauthorized) {
				fmt.Println("  -> sentinel: unauthorized")
			}

			// Typed check with errors.As
			var rlErr *RateLimitError
			if errors.As(err, &rlErr) {
				fmt.Printf("  -> typed: limit=%d, retry_after=%ds\n", rlErr.Limit, rlErr.RetryAfter)
			}
		} else {
			fmt.Printf("id=%d  name: %s\n", id, name)
		}
		fmt.Println()
	}
}
