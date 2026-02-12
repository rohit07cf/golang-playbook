"""Observability basics -- Python equivalent with request ID, logging, metrics."""

import json
import os
import random
import string
import threading
import time
from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.request import urlopen


# --- Request ID ---

def generate_request_id():
    return "".join(random.choices(string.hexdigits[:16], k=16))


# --- Metrics ---

class Metrics:
    def __init__(self):
        self.lock = threading.Lock()
        self.total_requests = 0
        self.total_errors = 0
        self.total_latency_ms = 0.0

    def record(self, status_code, latency_ms):
        with self.lock:
            self.total_requests += 1
            self.total_latency_ms += latency_ms
            if status_code >= 400:
                self.total_errors += 1

    def snapshot(self):
        with self.lock:
            avg = 0.0
            if self.total_requests > 0:
                avg = self.total_latency_ms / self.total_requests
            return {
                "requests_total": self.total_requests,
                "errors_total": self.total_errors,
                "latency_total_ms": round(self.total_latency_ms, 2),
                "latency_avg_ms": round(avg, 2),
            }


metrics = Metrics()


# --- Handler ---

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        start = time.time()
        request_id = generate_request_id()

        if self.path == "/hello":
            body = {"message": "hello", "request_id": request_id}
            status = 200
        elif self.path == "/error":
            body = {"error": "something went wrong"}
            status = 500
        elif self.path == "/metrics":
            snap = metrics.snapshot()
            text = "\n".join(f"{k} {v}" for k, v in snap.items())
            self.send_response(200)
            self.send_header("Content-Type", "text/plain")
            self.end_headers()
            self.wfile.write(text.encode())
            return
        else:
            self.send_response(404)
            self.end_headers()
            return

        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("X-Request-ID", request_id)
        self.end_headers()
        self.wfile.write(json.dumps(body).encode())

        latency_ms = (time.time() - start) * 1000
        metrics.record(status, latency_ms)

        # Structured log
        print(f"request_id={request_id} method={self.command} "
              f"path={self.path} status={status} latency={latency_ms:.1f}ms")

    def log_message(self, format, *args):
        pass


# --- Demo ---

def run_demo():
    time.sleep(0.5)

    print("--- sending requests ---\n")
    for i in range(5):
        resp = urlopen("http://localhost:9005/hello", timeout=2)
        body = json.loads(resp.read())
        print(f"GET /hello -> request_id={body['request_id']}")

    resp = urlopen("http://localhost:9005/error", timeout=2)
    print(f"GET /error -> status={resp.status}")

    print("\n--- GET /metrics ---\n")
    resp = urlopen("http://localhost:9005/metrics", timeout=2)
    print(resp.read().decode())

    print("demo done")
    os._exit(0)


if __name__ == "__main__":
    print("observability demo on :9005")
    t = threading.Thread(target=run_demo, daemon=True)
    t.start()
    server = HTTPServer(("", 9005), Handler)
    server.serve_forever()
