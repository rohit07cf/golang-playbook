"""Rate limit basics -- Python equivalent of the Go example.

Token bucket rate limiter with threading.Lock for concurrency safety.
"""

import json
import threading
import time
import urllib.request
import urllib.error
from http.server import HTTPServer, BaseHTTPRequestHandler


class TokenBucket:
    def __init__(self, max_tokens: int, refill_interval: float):
        self.max_tokens = max_tokens
        self.tokens = max_tokens
        self.refill_interval = refill_interval
        self.lock = threading.Lock()
        # Refill tokens periodically
        self._start_refill()

    def _start_refill(self):
        def refill():
            while True:
                time.sleep(self.refill_interval)
                with self.lock:
                    self.tokens = self.max_tokens
        t = threading.Thread(target=refill, daemon=True)
        t.start()

    def allow(self) -> bool:
        with self.lock:
            if self.tokens > 0:
                self.tokens -= 1
                return True
            return False

    def remaining(self) -> int:
        with self.lock:
            return self.tokens


# Global limiter: 5 requests per 3 seconds
limiter = TokenBucket(max_tokens=5, refill_interval=3.0)


def json_response(handler, status: int, data: dict, extra_headers=None) -> None:
    handler.send_response(status)
    handler.send_header("Content-Type", "application/json")
    if extra_headers:
        for k, v in extra_headers.items():
            handler.send_header(k, v)
    handler.end_headers()
    handler.wfile.write(json.dumps(data).encode())


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/api":
            if not limiter.allow():
                json_response(self, 429, {
                    "error": "rate limit exceeded",
                    "retry_after": f"{limiter.refill_interval:.0f}s",
                }, extra_headers={
                    "Retry-After": str(int(limiter.refill_interval)),
                })
                return

            json_response(self, 200, {"message": "API response"}, extra_headers={
                "X-RateLimit-Remaining": str(limiter.remaining()),
            })
            return

        json_response(self, 404, {"error": "not found"})

    def log_message(self, *args):
        pass


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), Handler)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    print(f"Server on {base}")
    print("  GET /api -> rate limited (5 req / 3s)")
    threading.Thread(target=server.serve_forever, daemon=True).start()

    print("\n--- Demo: 8 rapid requests ---")
    for i in range(1, 9):
        try:
            resp = urllib.request.urlopen(f"{base}/api", timeout=5)
            data = json.loads(resp.read())
            remaining = resp.headers.get("X-RateLimit-Remaining", "?")
            print(f"  request {i}: status={resp.status} remaining={remaining} msg={data.get('message', '')}")
        except urllib.error.HTTPError as e:
            data = json.loads(e.read())
            print(f"  request {i}: status={e.code} msg={data.get('error', '')}")

    print(f"\n  waiting 3s for token refill...")
    time.sleep(3)

    resp = urllib.request.urlopen(f"{base}/api", timeout=5)
    print(f"  after refill: status={resp.status}")

    server.shutdown()
    print("\nDone.")


if __name__ == "__main__":
    main()
