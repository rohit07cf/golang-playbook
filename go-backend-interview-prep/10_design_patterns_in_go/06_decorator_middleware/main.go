package main

import (
	"errors"
	"fmt"
	"time"
)

// --- Core interface ---

type Sender interface {
	Send(to, message string) error
}

// --- Base implementation ---

type EmailSender struct {
	failCount int // simulate failures for retry demo
}

func (e *EmailSender) Send(to, message string) error {
	e.failCount++
	if e.failCount <= 2 {
		return errors.New("temporary failure")
	}
	fmt.Printf("  [EMAIL] To: %s | %s\n", to, message)
	return nil
}

// --- Decorator: logging ---

type loggingSender struct {
	next Sender
}

func WithLogging(next Sender) Sender {
	return &loggingSender{next: next}
}

func (l *loggingSender) Send(to, message string) error {
	fmt.Printf("  [LOG] sending to %s at %s\n", to, time.Now().Format("15:04:05"))
	err := l.next.Send(to, message)
	if err != nil {
		fmt.Printf("  [LOG] error: %v\n", err)
	} else {
		fmt.Printf("  [LOG] success\n")
	}
	return err
}

// --- Decorator: retry ---

type retrySender struct {
	next       Sender
	maxRetries int
}

func WithRetry(next Sender, maxRetries int) Sender {
	return &retrySender{next: next, maxRetries: maxRetries}
}

func (r *retrySender) Send(to, message string) error {
	var err error
	for attempt := 1; attempt <= r.maxRetries; attempt++ {
		err = r.next.Send(to, message)
		if err == nil {
			return nil
		}
		fmt.Printf("  [RETRY] attempt %d/%d failed: %v\n", attempt, r.maxRetries, err)
	}
	return fmt.Errorf("all %d retries failed: %w", r.maxRetries, err)
}

// --- Decorator: rate limiter (simple) ---

type rateLimitSender struct {
	next    Sender
	delay   time.Duration
	lastSent time.Time
}

func WithRateLimit(next Sender, delay time.Duration) Sender {
	return &rateLimitSender{next: next, delay: delay}
}

func (r *rateLimitSender) Send(to, message string) error {
	since := time.Since(r.lastSent)
	if since < r.delay {
		wait := r.delay - since
		fmt.Printf("  [RATE] throttling %s...\n", wait.Round(time.Millisecond))
		time.Sleep(wait)
	}
	r.lastSent = time.Now()
	return r.next.Send(to, message)
}

func main() {
	fmt.Println("=== Decorator / Middleware Pattern ===")
	fmt.Println()

	// Layer decorators: logging -> retry -> email sender
	fmt.Println("--- Logging + Retry (email fails first 2 times) ---")
	base := &EmailSender{}
	sender := WithLogging(WithRetry(base, 3))
	err := sender.Send("alice@example.com", "Order shipped!")
	if err != nil {
		fmt.Printf("  Final error: %v\n", err)
	}
	fmt.Println()

	// Demonstrate rate limiting
	fmt.Println("--- Rate limited sender ---")
	simpleSender := WithRateLimit(
		WithLogging(&simplePrintSender{}),
		200*time.Millisecond,
	)
	for i := 0; i < 3; i++ {
		simpleSender.Send("bob", fmt.Sprintf("msg-%d", i))
	}
	fmt.Println()

	fmt.Println("Key: each decorator implements Sender, wraps another Sender.")
	fmt.Println("     Compose: WithLogging(WithRetry(WithRateLimit(base))).")
	fmt.Println("     Same pattern as Go HTTP middleware.")
}

type simplePrintSender struct{}

func (s *simplePrintSender) Send(to, message string) error {
	fmt.Printf("  [SEND] %s: %s\n", to, message)
	return nil
}
