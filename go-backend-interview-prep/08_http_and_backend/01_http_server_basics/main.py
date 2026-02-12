"""HTTP server basics -- Python equivalent of the Go example.

Uses only the stdlib http.server module (no Flask/FastAPI).
"""

import json
import threading
import time
import urllib.request
from http.server import HTTPServer, BaseHTTPRequestHandler


class MyHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/hello":
            self.send_response(200)
            self.send_header("Content-Type", "text/plain")
            self.end_headers()
            self.wfile.write(b"Hello from Python HTTP server!\n")

        elif self.path == "/health":
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            body = json.dumps({
                "status": "ok",
                "time": time.strftime("%Y-%m-%dT%H:%M:%S%z"),
            })
            self.wfile.write(body.encode())

        else:
            self.send_response(404)
            self.end_headers()
            self.wfile.write(b"not found\n")

    def log_message(self, *args):
        pass  # silence default stderr logging


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), MyHandler)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    print(f"Server running on {base}")
    print("  GET /hello  -> greeting")
    print("  GET /health -> JSON health check")

    # Run server in background thread
    threading.Thread(target=server.serve_forever, daemon=True).start()

    # --- Demo: hit the endpoints ---
    print("\n--- Demo requests ---")

    resp1 = urllib.request.urlopen(f"{base}/hello", timeout=5)
    print(f"GET /hello  -> {resp1.status}: {resp1.read().decode().strip()}")

    resp2 = urllib.request.urlopen(f"{base}/health", timeout=5)
    data = json.loads(resp2.read().decode())
    print(f"GET /health -> {resp2.status}: {data}")

    server.shutdown()
    print("\nServer stopped.")


if __name__ == "__main__":
    main()
