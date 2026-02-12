package main

import "fmt"

// --- Your internal interface ---

type Sender interface {
	Send(to, message string) error
}

// --- Simulated third-party APIs (incompatible signatures) ---

// TwilioClient simulates Twilio's API -- different method names and signatures.
type TwilioClient struct{}

func (t *TwilioClient) CreateMessage(phoneNumber string, body string, from string) error {
	fmt.Printf("  [TWILIO API] from=%s to=%s body=%s\n", from, phoneNumber, body)
	return nil
}

// SendGridClient simulates SendGrid's API -- completely different shape.
type SendGridClient struct{}

func (s *SendGridClient) PostMail(recipient, subject, htmlBody string) error {
	fmt.Printf("  [SENDGRID API] to=%s subject=%s body=%s\n", recipient, subject, htmlBody)
	return nil
}

// --- Adapters: translate third-party APIs to your Sender interface ---

type TwilioAdapter struct {
	client   *TwilioClient
	fromNum  string
}

func (a *TwilioAdapter) Send(to, message string) error {
	// Translate: your simple interface -> Twilio's complex API
	return a.client.CreateMessage(to, message, a.fromNum)
}

type SendGridAdapter struct {
	client *SendGridClient
}

func (a *SendGridAdapter) Send(to, message string) error {
	// Translate: your simple interface -> SendGrid's API
	return a.client.PostMail(to, "Notification", message)
}

// --- Service uses YOUR interface, not third-party types ---

type NotificationService struct {
	sender Sender
}

func (n *NotificationService) Notify(to, message string) {
	fmt.Printf("  NotificationService.Notify -> %s\n", to)
	n.sender.Send(to, message)
}

func main() {
	fmt.Println("=== Adapter Pattern ===")
	fmt.Println()

	// Use Twilio via adapter
	fmt.Println("--- Twilio adapter ---")
	twilioSender := &TwilioAdapter{
		client:  &TwilioClient{},
		fromNum: "+1555000000",
	}
	svc1 := &NotificationService{sender: twilioSender}
	svc1.Notify("+1234567890", "Your order shipped!")
	fmt.Println()

	// Use SendGrid via adapter -- same service, same interface
	fmt.Println("--- SendGrid adapter ---")
	sgSender := &SendGridAdapter{
		client: &SendGridClient{},
	}
	svc2 := &NotificationService{sender: sgSender}
	svc2.Notify("alice@example.com", "Your order shipped!")
	fmt.Println()

	fmt.Println("Key: NotificationService depends on Sender interface.")
	fmt.Println("     Adapters translate third-party APIs to match that interface.")
	fmt.Println("     Swap providers by swapping adapters -- zero service changes.")
}
