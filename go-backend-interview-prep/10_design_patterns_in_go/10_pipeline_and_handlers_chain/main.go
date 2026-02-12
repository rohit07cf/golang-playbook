package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// --- Message ---

type Message struct {
	To        string
	Body      string
	Timestamp time.Time
	Metadata  map[string]string
}

// --- Handler interface ---

type Handler interface {
	Handle(msg *Message) error
}

// --- Pipeline chains handlers ---

type Pipeline struct {
	handlers []Handler
}

func NewPipeline(handlers ...Handler) *Pipeline {
	return &Pipeline{handlers: handlers}
}

func (p *Pipeline) Run(msg *Message) error {
	for _, h := range p.handlers {
		if err := h.Handle(msg); err != nil {
			return err // short-circuit on error
		}
	}
	return nil
}

// --- Concrete handlers ---

// ValidateHandler checks required fields.
type ValidateHandler struct{}

func (v *ValidateHandler) Handle(msg *Message) error {
	fmt.Println("  [VALIDATE] checking message...")
	if msg.To == "" {
		return errors.New("validation: 'to' is required")
	}
	if msg.Body == "" {
		return errors.New("validation: 'body' is required")
	}
	if !strings.Contains(msg.To, "@") {
		return errors.New("validation: 'to' must be an email")
	}
	fmt.Println("  [VALIDATE] OK")
	return nil
}

// EnrichHandler adds metadata.
type EnrichHandler struct{}

func (e *EnrichHandler) Handle(msg *Message) error {
	msg.Timestamp = time.Now()
	if msg.Metadata == nil {
		msg.Metadata = make(map[string]string)
	}
	msg.Metadata["enriched"] = "true"
	msg.Metadata["priority"] = "normal"
	fmt.Printf("  [ENRICH]   Added timestamp + metadata\n")
	return nil
}

// LogHandler logs the message.
type LogHandler struct{}

func (l *LogHandler) Handle(msg *Message) error {
	fmt.Printf("  [LOG]      to=%s body=%q metadata=%v\n",
		msg.To, msg.Body, msg.Metadata)
	return nil
}

// SendHandler delivers the message.
type SendHandler struct{}

func (s *SendHandler) Handle(msg *Message) error {
	fmt.Printf("  [SEND]     Delivered to %s at %s\n",
		msg.To, msg.Timestamp.Format("15:04:05"))
	return nil
}

func main() {
	fmt.Println("=== Pipeline and Handlers Chain ===")
	fmt.Println()

	pipeline := NewPipeline(
		&ValidateHandler{},
		&EnrichHandler{},
		&LogHandler{},
		&SendHandler{},
	)

	// Valid message
	fmt.Println("--- Valid message ---")
	msg1 := &Message{To: "alice@example.com", Body: "Your order shipped!"}
	if err := pipeline.Run(msg1); err != nil {
		fmt.Printf("  PIPELINE ERROR: %v\n", err)
	}
	fmt.Println()

	// Invalid: missing email
	fmt.Println("--- Invalid message (no @) ---")
	msg2 := &Message{To: "bob", Body: "test"}
	if err := pipeline.Run(msg2); err != nil {
		fmt.Printf("  PIPELINE ERROR: %v\n", err)
	}
	fmt.Println()

	// Invalid: empty body
	fmt.Println("--- Invalid message (empty body) ---")
	msg3 := &Message{To: "charlie@example.com", Body: ""}
	if err := pipeline.Run(msg3); err != nil {
		fmt.Printf("  PIPELINE ERROR: %v\n", err)
	}
	fmt.Println()

	fmt.Println("Key: validate -> enrich -> log -> send. Error = short-circuit.")
	fmt.Println("     Each handler does one thing. Add/remove steps freely.")
}
