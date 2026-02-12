package main

import "fmt"

// --- Interfaces (dependencies) ---

type Sender interface {
	Send(to, message string) error
}

type Logger interface {
	Log(msg string)
}

// --- Concrete implementations ---

type EmailSender struct{}

func (e *EmailSender) Send(to, message string) error {
	fmt.Printf("  [EMAIL] To: %s | Message: %s\n", to, message)
	return nil
}

type SMSSender struct{}

func (s *SMSSender) Send(to, message string) error {
	fmt.Printf("  [SMS]   To: %s | Message: %s\n", to, message)
	return nil
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(msg string) {
	fmt.Printf("  [LOG]   %s\n", msg)
}

// --- Service (depends on interfaces, not concrete types) ---

type NotificationService struct {
	sender Sender
	logger Logger
}

// NewNotificationService is the constructor -- dependencies are injected here.
func NewNotificationService(s Sender, l Logger) *NotificationService {
	return &NotificationService{sender: s, logger: l}
}

func (n *NotificationService) Notify(to, message string) error {
	n.logger.Log("sending notification to " + to)
	return n.sender.Send(to, message)
}

// --- Fake for testing ---

type FakeSender struct {
	Calls []string
}

func (f *FakeSender) Send(to, message string) error {
	f.Calls = append(f.Calls, to+":"+message)
	return nil
}

func main() {
	fmt.Println("=== Constructor and Dependency Injection ===")
	fmt.Println()

	logger := &ConsoleLogger{}

	// Wire with email sender
	fmt.Println("--- Email notification ---")
	emailSvc := NewNotificationService(&EmailSender{}, logger)
	emailSvc.Notify("alice@example.com", "Your order shipped!")
	fmt.Println()

	// Wire with SMS sender -- same service, different behavior
	fmt.Println("--- SMS notification ---")
	smsSvc := NewNotificationService(&SMSSender{}, logger)
	smsSvc.Notify("+1234567890", "Your order shipped!")
	fmt.Println()

	// Wire with fake for testing
	fmt.Println("--- Test with fake ---")
	fake := &FakeSender{}
	testSvc := NewNotificationService(fake, logger)
	testSvc.Notify("bob@test.com", "test message")
	fmt.Printf("  Fake recorded %d call(s): %v\n", len(fake.Calls), fake.Calls)
	fmt.Println()

	fmt.Println("Key: same service, different senders. DI makes it testable.")
}
