package main

import (
	"fmt"
	"time"
)

// ============================================================
// APPROACH 1: Functional Options (Go-idiomatic)
// ============================================================

type Sender interface {
	Send(to, message string) error
}

type NotificationService struct {
	sender  Sender
	timeout time.Duration
	retries int
	verbose bool
}

// Option is a function that configures the service.
type Option func(*NotificationService)

func WithTimeout(d time.Duration) Option {
	return func(s *NotificationService) { s.timeout = d }
}

func WithRetries(n int) Option {
	return func(s *NotificationService) { s.retries = n }
}

func WithVerbose(v bool) Option {
	return func(s *NotificationService) { s.verbose = v }
}

// NewNotificationService uses functional options for optional config.
func NewNotificationService(sender Sender, opts ...Option) *NotificationService {
	// Defaults
	svc := &NotificationService{
		sender:  sender,
		timeout: 5 * time.Second,
		retries: 1,
		verbose: false,
	}
	// Apply options
	for _, opt := range opts {
		opt(svc)
	}
	return svc
}

func (s *NotificationService) Notify(to, msg string) {
	if s.verbose {
		fmt.Printf("  [VERBOSE] timeout=%s retries=%d\n", s.timeout, s.retries)
	}
	fmt.Printf("  Sending to %s: %s\n", to, msg)
}

// ============================================================
// APPROACH 2: Builder (less idiomatic in Go, shown for comparison)
// ============================================================

type ServiceBuilder struct {
	sender  Sender
	timeout time.Duration
	retries int
	verbose bool
}

func NewServiceBuilder(sender Sender) *ServiceBuilder {
	return &ServiceBuilder{
		sender:  sender,
		timeout: 5 * time.Second,
		retries: 1,
	}
}

func (b *ServiceBuilder) Timeout(d time.Duration) *ServiceBuilder {
	b.timeout = d
	return b
}

func (b *ServiceBuilder) Retries(n int) *ServiceBuilder {
	b.retries = n
	return b
}

func (b *ServiceBuilder) Verbose() *ServiceBuilder {
	b.verbose = true
	return b
}

func (b *ServiceBuilder) Build() *NotificationService {
	return &NotificationService{
		sender:  b.sender,
		timeout: b.timeout,
		retries: b.retries,
		verbose: b.verbose,
	}
}

// --- Simple sender for demo ---

type printSender struct{}

func (p *printSender) Send(to, message string) error {
	fmt.Printf("  [SEND] %s: %s\n", to, message)
	return nil
}

func main() {
	fmt.Println("=== Builder vs Functional Options ===")
	fmt.Println()
	sender := &printSender{}

	// Functional options (preferred)
	fmt.Println("--- Functional Options (Go-idiomatic) ---")

	svc1 := NewNotificationService(sender) // all defaults
	fmt.Print("  defaults: ")
	svc1.Notify("alice", "hello")

	svc2 := NewNotificationService(sender,
		WithTimeout(10*time.Second),
		WithRetries(3),
		WithVerbose(true),
	)
	fmt.Print("  custom:   ")
	svc2.Notify("bob", "world")
	fmt.Println()

	// Builder (for comparison)
	fmt.Println("--- Builder Pattern (less idiomatic) ---")
	svc3 := NewServiceBuilder(sender).
		Timeout(10 * time.Second).
		Retries(3).
		Verbose().
		Build()
	fmt.Print("  builder:  ")
	svc3.Notify("charlie", "hi")
	fmt.Println()

	fmt.Println("Key: functional options are Go-idiomatic for optional config.")
	fmt.Println("     Builder works but adds more ceremony with chained methods.")
}
