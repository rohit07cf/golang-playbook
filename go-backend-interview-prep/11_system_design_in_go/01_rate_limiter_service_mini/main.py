"""Rate limiter service -- Python equivalent using token bucket."""

import json
import threading
import time
from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.request import urlopen, Request
from urllib.error import URLError


# --- Token Bucket ---

class Bucket:
    def __init__(self, max_tokens, refill_rate):
        self.tokens = max_tokens
        self.max_tokens = max_tokens
        self.refill_rate = refill_rate  # tokens per second
        self.last_refill = time.time()

    def allow(self):
        now = time.time()
        elapsed = now - self.last_refill
        self.tokens += elapsed * self.refill_rate
        if self.tokens > self.max_tokens:
            self.tokens = self.max_tokens
        self.last_refill = now

        if self.tokens >= 1:
            self.tokens -= 1
            return True
        return False


# --- Rate Limiter ---

class RateLimiter:
    def __init__(self, refill_rate, max_tokens):
        self.lock = threading.Lock()
        self.buckets = {}
        self.max_tokens = max_tokens
        self.refill_rate = refill_rate

    def allow(self, key):
        with self.lock:
            if key not in self.buckets:
                self.buckets[key] = Bucket(self.max_tokens, self.refill_rate)
            return self.buckets[key].allow()


# --- Handler with rate limit ---

limiter = RateLimiter(refill_rate=2, max_tokens=5)


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/ping":
            key = self.client_address[0]
            if not limiter.allow(key):
                self.send_response(429)
                self.send_header("Content-Type", "application/json")
                self.send_header("Retry-After", "1")
                self.end_headers()
                self.wfile.write(json.dumps({"error": "rate limit exceeded"}).encode())
                return

            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps({"message": "pong"}).encode())
        else:
            self.send_response(404)
            self.end_headers()

    def log_message(self, format, *args):
        pass  # suppress default logs


# --- Demo ---

def run_demo():
    time.sleep(0.5)
    for i in range(1, 9):
        try:
            req = Request("http://localhost:9001/ping")
            resp = urlopen(req, timeout=2)
            body = json.loads(resp.read())
            print(f"req {i}: status={resp.status} body={body}")
        except URLError as e:
            if hasattr(e, "code"):
                print(f"req {i}: status={e.code} body=rate limit exceeded")
            else:
                print(f"req {i}: error: {e}")

    print("\n--- waiting 2s for token refill ---")
    time.sleep(2)

    try:
        resp = urlopen("http://localhost:9001/ping", timeout=2)
        body = json.loads(resp.read())
        print(f"req after wait: status={resp.status} body={body}")
    except URLError as e:
        if hasattr(e, "code"):
            print(f"req after wait: status={e.code}")
        else:
            print(f"req after wait: error: {e}")

    print("\ndemo done")
    import os
    os._exit(0)


if __name__ == "__main__":
    print("rate limiter server on :9001 (2 tokens/sec, burst 5)")
    t = threading.Thread(target=run_demo, daemon=True)
    t.start()
    server = HTTPServer(("", 9001), Handler)
    server.serve_forever()
