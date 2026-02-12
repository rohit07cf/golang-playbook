"""HTTP client basics -- Python equivalent of the Go example.

Uses http.server for a local test server and urllib.request for client.
"""

import json
import threading
import time
import urllib.request
import urllib.error
from http.server import HTTPServer, BaseHTTPRequestHandler


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/hello":
            self.send_response(200)
            self.send_header("X-Custom", "demo-value")
            self.end_headers()
            self.wfile.write(f"Hello from local server! Method=GET".encode())
        elif self.path == "/json":
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps({"status": "ok", "count": 42}).encode())
        elif self.path == "/slow":
            time.sleep(3)
            self.send_response(200)
            self.end_headers()
            self.wfile.write(b"slow response")
        else:
            self.send_response(404)
            self.end_headers()

    def log_message(self, format, *args):
        pass  # suppress logs


def main() -> None:
    # --- Spin up local server ---
    server = HTTPServer(("127.0.0.1", 0), Handler)
    port = server.server_address[1]
    base_url = f"http://127.0.0.1:{port}"

    t = threading.Thread(target=server.serve_forever, daemon=True)
    t.start()
    print(f"local server at {base_url}")

    try:
        # --- Example 1: GET with timeout ---
        print("\n=== GET /hello ===")
        resp = urllib.request.urlopen(f"{base_url}/hello", timeout=5)
        body = resp.read().decode()
        print(f"  status: {resp.status}")
        print(f"  body: {body}")
        print(f"  X-Custom: {resp.headers.get('X-Custom')}")

        # --- Example 2: GET JSON ---
        print("\n=== GET /json ===")
        resp = urllib.request.urlopen(f"{base_url}/json", timeout=5)
        body = resp.read().decode()
        print(f"  Content-Type: {resp.headers.get('Content-Type')}")
        print(f"  body: {body}")

        # --- Example 3: custom request with headers ---
        print("\n=== Custom request with headers ===")
        req = urllib.request.Request(
            f"{base_url}/hello",
            headers={
                "Accept": "text/plain",
                "User-Agent": "py-interview-prep/1.0",
            },
        )
        resp = urllib.request.urlopen(req, timeout=5)
        print(f"  body: {resp.read().decode()}")

        # --- Example 4: timeout on slow endpoint ---
        print("\n=== Timeout (1s on /slow) ===")
        try:
            urllib.request.urlopen(f"{base_url}/slow", timeout=1)
        except Exception as e:
            print(f"  expected timeout: {e}")

    finally:
        server.shutdown()


if __name__ == "__main__":
    main()
