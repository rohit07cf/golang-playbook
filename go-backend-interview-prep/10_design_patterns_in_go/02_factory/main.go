package main

import (
	"errors"
	"fmt"
)

// --- Interface ---

type Sender interface {
	Send(to, message string) error
	Kind() string
}

// --- Implementations ---

type EmailSender struct{}

func (e *EmailSender) Send(to, message string) error {
	fmt.Printf("  [EMAIL] To: %s | %s\n", to, message)
	return nil
}
func (e *EmailSender) Kind() string { return "email" }

type SMSSender struct{}

func (s *SMSSender) Send(to, message string) error {
	fmt.Printf("  [SMS]   To: %s | %s\n", to, message)
	return nil
}
func (s *SMSSender) Kind() string { return "sms" }

type SlackSender struct{ webhook string }

func (s *SlackSender) Send(to, message string) error {
	fmt.Printf("  [SLACK] Channel: %s | %s (webhook: %s)\n", to, message, s.webhook)
	return nil
}
func (s *SlackSender) Kind() string { return "slack" }

// --- Factory ---

// NewSender creates a Sender based on a config string.
// Caller gets an interface back -- doesn't know the concrete type.
func NewSender(kind string) (Sender, error) {
	switch kind {
	case "email":
		return &EmailSender{}, nil
	case "sms":
		return &SMSSender{}, nil
	case "slack":
		return &SlackSender{webhook: "https://hooks.slack.example/abc"}, nil
	default:
		return nil, errors.New("unknown sender type: " + kind)
	}
}

func main() {
	fmt.Println("=== Factory Pattern ===")
	fmt.Println()

	// Simulate config-driven sender selection
	configs := []string{"email", "sms", "slack", "pigeon"}

	for _, kind := range configs {
		sender, err := NewSender(kind)
		if err != nil {
			fmt.Printf("  factory error: %v\n", err)
			continue
		}
		fmt.Printf("  Created sender: kind=%s\n", sender.Kind())
		sender.Send("alice", "Your order shipped!")
		fmt.Println()
	}

	fmt.Println("Key: factory returns an interface. Callers don't know the concrete type.")
	fmt.Println("     Adding a new sender = one new case in the factory. Zero caller changes.")
}
