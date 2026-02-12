"""API gateway basics -- Python equivalent with auth, rate limit, routing."""

import json
import os
import threading
import time
import uuid
from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.request import urlopen, Request
from urllib.error import URLError


# --- Rate Limiter (fixed window) ---

class RateLimiter:
    def __init__(self, limit, window_sec):
        self.lock = threading.Lock()
        self.counts = {}
        self.limit = limit
        self.window = window_sec
        self.last_reset = time.time()

    def allow(self, key):
        with self.lock:
            if time.time() - self.last_reset > self.window:
                self.counts = {}
                self.last_reset = time.time()
            self.counts[key] = self.counts.get(key, 0) + 1
            return self.counts[key] <= self.limit


# --- Config ---

VALID_KEYS = {"key-abc-123", "key-xyz-789"}
rate_limiter = RateLimiter(limit=5, window_sec=10)


# --- Gateway Handler ---

class GatewayHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        request_id = uuid.uuid4().hex[:16]

        # Auth check
        api_key = self.headers.get("X-API-Key", "")
        if api_key not in VALID_KEYS:
            self._respond(401, {"error": "invalid or missing API key"}, request_id)
            return

        # Rate limit
        client_key = self.client_address[0]
        if not rate_limiter.allow(client_key):
            self._respond(429, {"error": "rate limit exceeded"}, request_id)
            return

        # Route
        if self.path.startswith("/api/users"):
            body = {
                "service": "users",
                "path": self.path,
                "request_id": request_id,
                "data": "user list placeholder",
            }
            self._respond(200, body, request_id)
        elif self.path.startswith("/api/orders"):
            body = {
                "service": "orders",
                "path": self.path,
                "request_id": request_id,
                "data": "order list placeholder",
            }
            self._respond(200, body, request_id)
        else:
            self._respond(404, {"error": "route not found"}, request_id)

    def _respond(self, status, body, request_id):
        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("X-Request-ID", request_id)
        self.end_headers()
        self.wfile.write(json.dumps(body).encode())

    def log_message(self, format, *args):
        pass


# --- Demo ---

def run_demo():
    time.sleep(0.5)
    print("--- API Gateway Demo ---\n")

    def fetch(path, api_key=None):
        req = Request(f"http://localhost:9007{path}")
        if api_key:
            req.add_header("X-API-Key", api_key)
        try:
            resp = urlopen(req, timeout=2)
            body = json.loads(resp.read())
            return resp.status, body
        except URLError as e:
            if hasattr(e, "code"):
                try:
                    body = json.loads(e.read())
                except Exception:
                    body = {}
                return e.code, body
            return 0, {"error": str(e)}

    # No API key
    status, body = fetch("/api/users")
    print(f"GET /api/users (no key): status={status}")

    # Valid key -- users
    status, body = fetch("/api/users", "key-abc-123")
    print(f"GET /api/users (valid key): status={status} service={body.get('service')} req_id={body.get('request_id')}")

    # Valid key -- orders
    status, body = fetch("/api/orders", "key-abc-123")
    print(f"GET /api/orders (valid key): status={status} service={body.get('service')}")

    # Unknown route
    status, body = fetch("/api/unknown", "key-abc-123")
    print(f"GET /api/unknown: status={status}")

    # Rate limit test
    print("\n--- Rate limit test (limit=5/10s) ---")
    for i in range(1, 8):
        status, _ = fetch("/api/users", "key-xyz-789")
        print(f"req {i}: status={status}")

    print("\ndemo done")
    os._exit(0)


if __name__ == "__main__":
    print("api gateway on :9007")
    t = threading.Thread(target=run_demo, daemon=True)
    t.start()
    server = HTTPServer(("", 9007), GatewayHandler)
    server.serve_forever()
