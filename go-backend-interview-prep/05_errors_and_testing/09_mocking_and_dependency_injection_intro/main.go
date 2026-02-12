package main

import (
	"errors"
	"fmt"
)

// ---- Interface: the dependency ----

type UserStore interface {
	GetUser(id int) (string, error)
}

// ---- Real implementation (would talk to a DB) ----

type DBStore struct{}

func (d DBStore) GetUser(id int) (string, error) {
	// Simulate a database lookup
	users := map[int]string{1: "alice", 2: "bob"}
	name, ok := users[id]
	if !ok {
		return "", errors.New("user not found")
	}
	return name, nil
}

// ---- Fake implementation for testing ----

type FakeStore struct {
	Users map[int]string
	Err   error
}

func (f FakeStore) GetUser(id int) (string, error) {
	if f.Err != nil {
		return "", f.Err
	}
	name, ok := f.Users[id]
	if !ok {
		return "", errors.New("user not found")
	}
	return name, nil
}

// ---- Service that depends on the interface ----

type UserService struct {
	store UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) Greet(id int) (string, error) {
	name, err := s.store.GetUser(id)
	if err != nil {
		return "", fmt.Errorf("Greet(%d): %w", id, err)
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}

func main() {
	// --- Using the real store ---
	fmt.Println("=== Real store ===")
	svc := NewUserService(DBStore{})
	msg, err := svc.Greet(1)
	fmt.Printf("Greet(1): %q, err=%v\n", msg, err)

	msg, err = svc.Greet(999)
	fmt.Printf("Greet(999): %q, err=%v\n", msg, err)

	// --- Using the fake store (like in tests) ---
	fmt.Println("\n=== Fake store (test mode) ===")
	fake := FakeStore{
		Users: map[int]string{10: "test_user"},
	}
	testSvc := NewUserService(fake)
	msg, err = testSvc.Greet(10)
	fmt.Printf("Greet(10): %q, err=%v\n", msg, err)

	// --- Fake that always errors ---
	fmt.Println("\n=== Fake store (error mode) ===")
	errFake := FakeStore{Err: errors.New("database down")}
	errSvc := NewUserService(errFake)
	msg, err = errSvc.Greet(1)
	fmt.Printf("Greet(1): %q, err=%v\n", msg, err)
}
