package main

import (
	"errors"
	"fmt"
	"time"
)

// --- Domain ---

type Notification struct {
	ID      string
	To      string
	Message string
	SentAt  time.Time
}

// --- Repository interface ---

type NotificationRepo interface {
	Save(n Notification) error
	FindByID(id string) (Notification, error)
	List() ([]Notification, error)
}

// --- In-memory implementation ---

type MemoryRepo struct {
	data map[string]Notification
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{data: make(map[string]Notification)}
}

func (r *MemoryRepo) Save(n Notification) error {
	r.data[n.ID] = n
	return nil
}

func (r *MemoryRepo) FindByID(id string) (Notification, error) {
	n, ok := r.data[id]
	if !ok {
		return Notification{}, errors.New("not found: " + id)
	}
	return n, nil
}

func (r *MemoryRepo) List() ([]Notification, error) {
	result := make([]Notification, 0, len(r.data))
	for _, n := range r.data {
		result = append(result, n)
	}
	return result, nil
}

// --- Service (depends on interface, not concrete repo) ---

type NotificationService struct {
	repo NotificationRepo
}

func NewNotificationService(repo NotificationRepo) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) Send(id, to, message string) error {
	n := Notification{
		ID:      id,
		To:      to,
		Message: message,
		SentAt:  time.Now(),
	}
	fmt.Printf("  Sending: id=%s to=%s msg=%s\n", id, to, message)
	return s.repo.Save(n)
}

func (s *NotificationService) GetHistory() ([]Notification, error) {
	return s.repo.List()
}

func main() {
	fmt.Println("=== Repository Pattern ===")
	fmt.Println()

	// Wire with in-memory repo
	repo := NewMemoryRepo()
	svc := NewNotificationService(repo)

	// Send some notifications
	svc.Send("n1", "alice@example.com", "Order confirmed")
	svc.Send("n2", "bob@example.com", "Payment received")
	svc.Send("n3", "alice@example.com", "Order shipped")
	fmt.Println()

	// Query via repository
	fmt.Println("--- Find by ID ---")
	n, err := repo.FindByID("n2")
	if err != nil {
		fmt.Println("  error:", err)
	} else {
		fmt.Printf("  Found: id=%s to=%s msg=%s\n", n.ID, n.To, n.Message)
	}

	_, err = repo.FindByID("n99")
	fmt.Printf("  Not found: %v\n", err)
	fmt.Println()

	// List all
	fmt.Println("--- List all ---")
	all, _ := svc.GetHistory()
	for _, n := range all {
		fmt.Printf("  %s -> %s: %s\n", n.ID, n.To, n.Message)
	}
	fmt.Println()

	fmt.Println("Key: service calls repo.Save() / repo.FindByID().")
	fmt.Println("     Swap MemoryRepo for SQLRepo -- zero service changes.")
}
