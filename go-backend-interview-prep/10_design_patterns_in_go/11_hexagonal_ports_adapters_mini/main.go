package main

import (
	"errors"
	"fmt"
	"time"
)

// ============================================================
// CORE (depends on ports only -- zero external imports)
// ============================================================

// --- Domain ---

type Notification struct {
	ID      string
	To      string
	Message string
	SentAt  time.Time
}

// --- Ports (interfaces defined by core) ---

type SenderPort interface {
	Send(to, message string) error
}

type RepoPort interface {
	Save(n Notification) error
	FindByID(id string) (Notification, error)
}

// --- Core service (depends on ports, not adapters) ---

type NotificationService struct {
	sender SenderPort
	repo   RepoPort
}

func NewNotificationService(sender SenderPort, repo RepoPort) *NotificationService {
	return &NotificationService{sender: sender, repo: repo}
}

func (s *NotificationService) Send(id, to, message string) error {
	// Business logic lives here -- no knowledge of email/DB specifics
	if to == "" {
		return errors.New("recipient required")
	}
	if err := s.sender.Send(to, message); err != nil {
		return fmt.Errorf("send failed: %w", err)
	}
	n := Notification{ID: id, To: to, Message: message, SentAt: time.Now()}
	return s.repo.Save(n)
}

// ============================================================
// ADAPTERS (implement ports -- live outside core)
// ============================================================

// --- Email adapter (implements SenderPort) ---

type EmailAdapter struct{}

func (e *EmailAdapter) Send(to, message string) error {
	fmt.Printf("  [EMAIL ADAPTER] Sent to %s: %s\n", to, message)
	return nil
}

// --- In-memory repo adapter (implements RepoPort) ---

type MemoryRepoAdapter struct {
	data map[string]Notification
}

func NewMemoryRepoAdapter() *MemoryRepoAdapter {
	return &MemoryRepoAdapter{data: make(map[string]Notification)}
}

func (r *MemoryRepoAdapter) Save(n Notification) error {
	r.data[n.ID] = n
	fmt.Printf("  [REPO ADAPTER]  Saved notification %s\n", n.ID)
	return nil
}

func (r *MemoryRepoAdapter) FindByID(id string) (Notification, error) {
	n, ok := r.data[id]
	if !ok {
		return Notification{}, errors.New("not found: " + id)
	}
	return n, nil
}

// --- Fake adapters for testing ---

type FakeSender struct{ Calls []string }

func (f *FakeSender) Send(to, msg string) error {
	f.Calls = append(f.Calls, to+":"+msg)
	return nil
}

type FakeRepo struct{ Saved []Notification }

func (f *FakeRepo) Save(n Notification) error {
	f.Saved = append(f.Saved, n)
	return nil
}
func (f *FakeRepo) FindByID(id string) (Notification, error) {
	for _, n := range f.Saved {
		if n.ID == id {
			return n, nil
		}
	}
	return Notification{}, errors.New("not found")
}

// ============================================================
// MAIN (wires adapters to ports)
// ============================================================

func main() {
	fmt.Println("=== Hexagonal Architecture (Ports & Adapters) ===")
	fmt.Println()

	// --- Production wiring ---
	fmt.Println("--- Production: email + memory repo ---")
	svc := NewNotificationService(
		&EmailAdapter{},        // adapter for SenderPort
		NewMemoryRepoAdapter(), // adapter for RepoPort
	)
	svc.Send("n1", "alice@example.com", "Order confirmed")
	svc.Send("n2", "bob@example.com", "Payment received")
	fmt.Println()

	// --- Test wiring: fakes ---
	fmt.Println("--- Test: fake sender + fake repo ---")
	fakeSender := &FakeSender{}
	fakeRepo := &FakeRepo{}
	testSvc := NewNotificationService(fakeSender, fakeRepo)
	testSvc.Send("t1", "test@test.com", "test message")

	fmt.Printf("  Sender calls: %v\n", fakeSender.Calls)
	fmt.Printf("  Repo saved:   %d notification(s)\n", len(fakeRepo.Saved))
	fmt.Println()

	fmt.Println("Architecture:")
	fmt.Println("  Core:     NotificationService (depends on SenderPort + RepoPort)")
	fmt.Println("  Adapters: EmailAdapter, MemoryRepoAdapter (implement ports)")
	fmt.Println("  main():   Wires adapters to ports (dependency injection)")
	fmt.Println("  Tests:    Swap adapters for fakes -- zero core changes")
}
