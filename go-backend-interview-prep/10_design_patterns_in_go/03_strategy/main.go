package main

import "fmt"

// --- Strategy interface ---

type RoutingStrategy interface {
	Route(message string, channels []string) string
}

// --- Concrete strategies ---

// PriorityRouting always picks the first channel.
type PriorityRouting struct{}

func (p PriorityRouting) Route(message string, channels []string) string {
	if len(channels) == 0 {
		return ""
	}
	return channels[0] // highest priority
}

// RoundRobinRouting cycles through channels.
type RoundRobinRouting struct {
	counter int
}

func (r *RoundRobinRouting) Route(message string, channels []string) string {
	if len(channels) == 0 {
		return ""
	}
	ch := channels[r.counter%len(channels)]
	r.counter++
	return ch
}

// ContentBasedRouting picks channel based on message content.
type ContentBasedRouting struct{}

func (c ContentBasedRouting) Route(message string, channels []string) string {
	if len(message) > 50 {
		return "email" // long messages go to email
	}
	return "sms" // short messages go to sms
}

// --- Service using strategy ---

type NotificationService struct {
	strategy RoutingStrategy
	channels []string
}

func NewNotificationService(strategy RoutingStrategy, channels []string) *NotificationService {
	return &NotificationService{strategy: strategy, channels: channels}
}

func (n *NotificationService) Send(message string) {
	channel := n.strategy.Route(message, n.channels)
	fmt.Printf("  Sending via %-8s: %s\n", channel, message)
}

// SetStrategy allows swapping at runtime.
func (n *NotificationService) SetStrategy(s RoutingStrategy) {
	n.strategy = s
}

func main() {
	fmt.Println("=== Strategy Pattern ===")
	fmt.Println()

	channels := []string{"email", "sms", "slack"}

	// Priority strategy
	fmt.Println("--- Priority routing (always first channel) ---")
	svc := NewNotificationService(PriorityRouting{}, channels)
	svc.Send("Order confirmed")
	svc.Send("Payment received")
	fmt.Println()

	// Swap to round-robin at runtime
	fmt.Println("--- Round-robin routing ---")
	svc.SetStrategy(&RoundRobinRouting{})
	svc.Send("Order confirmed")
	svc.Send("Payment received")
	svc.Send("Shipped!")
	svc.Send("Delivered!")
	fmt.Println()

	// Swap to content-based
	fmt.Println("--- Content-based routing ---")
	svc.SetStrategy(ContentBasedRouting{})
	svc.Send("short msg")
	svc.Send("This is a much longer message that should be routed to email because it exceeds fifty characters")
	fmt.Println()

	fmt.Println("Key: same service, different routing logic. Swapped at runtime.")
}
