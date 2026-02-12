package main

import "fmt"

// --- Event types ---

type Event struct {
	Name string
	Data map[string]string
}

// --- Event bus ---

type EventBus struct {
	handlers map[string][]func(Event)
}

func NewEventBus() *EventBus {
	return &EventBus{handlers: make(map[string][]func(Event))}
}

func (b *EventBus) Subscribe(eventName string, handler func(Event)) {
	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *EventBus) Publish(event Event) {
	handlers, ok := b.handlers[event.Name]
	if !ok {
		return
	}
	for _, h := range handlers {
		h(event) // in production: run in goroutine with recover
	}
}

// --- Notification service publishes events ---

type NotificationService struct {
	bus *EventBus
}

func NewNotificationService(bus *EventBus) *NotificationService {
	return &NotificationService{bus: bus}
}

func (s *NotificationService) Send(to, message string) {
	fmt.Printf("  [SEND] To: %s | %s\n", to, message)
	s.bus.Publish(Event{
		Name: "notification.sent",
		Data: map[string]string{"to": to, "message": message},
	})
}

func (s *NotificationService) Fail(to, reason string) {
	fmt.Printf("  [FAIL] To: %s | reason: %s\n", to, reason)
	s.bus.Publish(Event{
		Name: "notification.failed",
		Data: map[string]string{"to": to, "reason": reason},
	})
}

func main() {
	fmt.Println("=== Observer / Events Pattern ===")
	fmt.Println()

	bus := NewEventBus()

	// Subscribe: logger
	bus.Subscribe("notification.sent", func(e Event) {
		fmt.Printf("  [LOG]       Notification sent to %s\n", e.Data["to"])
	})

	// Subscribe: analytics
	bus.Subscribe("notification.sent", func(e Event) {
		fmt.Printf("  [ANALYTICS] Recorded delivery to %s\n", e.Data["to"])
	})

	// Subscribe: alert on failure
	bus.Subscribe("notification.failed", func(e Event) {
		fmt.Printf("  [ALERT]     FAILURE for %s: %s\n", e.Data["to"], e.Data["reason"])
	})

	// --- Use the service ---
	svc := NewNotificationService(bus)

	fmt.Println("--- Sending notifications ---")
	svc.Send("alice@example.com", "Order confirmed")
	fmt.Println()
	svc.Send("bob@example.com", "Payment received")
	fmt.Println()

	fmt.Println("--- Simulating failure ---")
	svc.Fail("charlie@example.com", "invalid email")
	fmt.Println()

	fmt.Println("Key: service publishes events. Subscribers react independently.")
	fmt.Println("     Adding analytics = one Subscribe() call. Zero service changes.")
}
