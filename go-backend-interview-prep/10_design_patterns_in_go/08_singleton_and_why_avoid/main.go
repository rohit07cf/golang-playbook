package main

import (
	"fmt"
	"sync"
)

// ============================================================
// APPROACH 1: Singleton (the problem)
// ============================================================

type Config struct {
	SMTPHost string
	SMTPPort int
}

var (
	configInstance *Config
	configOnce     sync.Once
)

// GetConfig is a singleton -- returns the same instance every time.
func GetConfig() *Config {
	configOnce.Do(func() {
		fmt.Println("  [SINGLETON] Loading config (runs once)...")
		configInstance = &Config{SMTPHost: "smtp.example.com", SMTPPort: 587}
	})
	return configInstance
}

// singletonService uses the singleton directly -- hidden dependency!
type singletonService struct{}

func (s *singletonService) SendNotification(to, msg string) {
	cfg := GetConfig() // hidden: caller can't see this dependency
	fmt.Printf("  [SINGLETON] Send via %s:%d to=%s msg=%s\n",
		cfg.SMTPHost, cfg.SMTPPort, to, msg)
}

// ============================================================
// APPROACH 2: Dependency Injection (the solution)
// ============================================================

type Sender interface {
	Send(to, message string) error
}

type EmailSender struct {
	Host string
	Port int
}

func (e *EmailSender) Send(to, message string) error {
	fmt.Printf("  [DI]        Send via %s:%d to=%s msg=%s\n",
		e.Host, e.Port, to, message)
	return nil
}

type NotificationService struct {
	sender Sender // explicit dependency -- visible, testable
}

func NewNotificationService(s Sender) *NotificationService {
	return &NotificationService{sender: s}
}

func (n *NotificationService) Notify(to, msg string) {
	n.sender.Send(to, msg)
}

func main() {
	fmt.Println("=== Singleton and Why to Avoid It ===")
	fmt.Println()

	// --- Singleton approach ---
	fmt.Println("--- Singleton (hidden dependency) ---")
	svc1 := &singletonService{}
	svc1.SendNotification("alice@example.com", "Order shipped!")
	svc1.SendNotification("bob@example.com", "Payment received")
	fmt.Println()

	fmt.Println("  Problem: singletonService hides its Config dependency.")
	fmt.Println("  In tests, you can't swap Config without global mutation.")
	fmt.Println()

	// --- DI approach ---
	fmt.Println("--- DI alternative (explicit dependency) ---")
	sender := &EmailSender{Host: "smtp.example.com", Port: 587}
	svc2 := NewNotificationService(sender)
	svc2.Notify("alice@example.com", "Order shipped!")
	svc2.Notify("bob@example.com", "Payment received")
	fmt.Println()

	// In tests: swap with a fake
	fmt.Println("--- In tests: swap with fake ---")
	fakeSender := &FakeSender{}
	testSvc := NewNotificationService(fakeSender)
	testSvc.Notify("test@example.com", "test message")
	fmt.Printf("  Fake recorded: %v\n", fakeSender.calls)
	fmt.Println()

	fmt.Println("Key: DI makes dependencies visible and testable.")
	fmt.Println("     Singleton hides them. Prefer DI in almost all cases.")
}

type FakeSender struct {
	calls []string
}

func (f *FakeSender) Send(to, message string) error {
	f.calls = append(f.calls, to+":"+message)
	return nil
}
