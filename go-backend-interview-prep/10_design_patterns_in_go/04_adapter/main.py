"""Adapter pattern -- Python equivalent of the Go example.

Adapters translate third-party APIs to match your internal interface.
"""


# --- Simulated third-party APIs ---

class TwilioClient:
    def create_message(self, phone_number: str, body: str, from_num: str) -> None:
        print(f"  [TWILIO API] from={from_num} to={phone_number} body={body}")


class SendGridClient:
    def post_mail(self, recipient: str, subject: str, html_body: str) -> None:
        print(f"  [SENDGRID API] to={recipient} subject={subject} body={html_body}")


# --- Adapters: translate to your internal interface ---

class TwilioAdapter:
    def __init__(self, client: TwilioClient, from_num: str):
        self.client = client
        self.from_num = from_num

    def send(self, to: str, message: str) -> None:
        self.client.create_message(to, message, self.from_num)


class SendGridAdapter:
    def __init__(self, client: SendGridClient):
        self.client = client

    def send(self, to: str, message: str) -> None:
        self.client.post_mail(to, "Notification", message)


# --- Service uses YOUR interface ---

class NotificationService:
    def __init__(self, sender):
        self.sender = sender

    def notify(self, to: str, message: str) -> None:
        print(f"  NotificationService.notify -> {to}")
        self.sender.send(to, message)


def main() -> None:
    print("=== Adapter Pattern ===")
    print()

    # Use Twilio via adapter
    print("--- Twilio adapter ---")
    twilio_sender = TwilioAdapter(TwilioClient(), "+1555000000")
    svc1 = NotificationService(twilio_sender)
    svc1.notify("+1234567890", "Your order shipped!")
    print()

    # Use SendGrid via adapter
    print("--- SendGrid adapter ---")
    sg_sender = SendGridAdapter(SendGridClient())
    svc2 = NotificationService(sg_sender)
    svc2.notify("alice@example.com", "Your order shipped!")
    print()

    print("Key: NotificationService depends on .send() duck typing.")
    print("     Adapters translate third-party APIs to match that interface.")
    print("     Swap providers by swapping adapters -- zero service changes.")


if __name__ == "__main__":
    main()
