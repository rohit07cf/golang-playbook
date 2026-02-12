"""Logging and request ID -- Python equivalent of the Go example.

Python stdlib has no context propagation. We generate a request ID
per request and include it in log output and response headers.
"""

import json
import os
import threading
import time
import urllib.request
from http.server import HTTPServer, BaseHTTPRequestHandler


def generate_id() -> str:
    return os.urandom(8).hex()


def json_response(handler, status: int, data: dict) -> None:
    handler.send_response(status)
    handler.send_header("Content-Type", "application/json")
    handler.end_headers()
    handler.wfile.write(json.dumps(data).encode())


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        start = time.perf_counter()

        # Generate or accept request ID
        request_id = self.headers.get("X-Request-ID", "") or generate_id()

        if self.path == "/hello":
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.send_header("X-Request-ID", request_id)
            self.end_headers()
            body = json.dumps({"message": "hello", "request_id": request_id})
            self.wfile.write(body.encode())
        else:
            self.send_response(404)
            self.send_header("X-Request-ID", request_id)
            self.end_headers()
            self.wfile.write(b"not found")

        elapsed = time.perf_counter() - start
        print(
            f"  [LOG] request_id={request_id} "
            f"method={self.command} path={self.path} "
            f"duration={elapsed*1000:.1f}ms"
        )

    def log_message(self, *args):
        pass


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), Handler)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    print(f"Server on {base}")
    print("  GET /hello -> greeting with request ID")
    threading.Thread(target=server.serve_forever, daemon=True).start()

    print("\n--- Demo ---")

    # Normal request (server generates ID)
    resp = urllib.request.urlopen(f"{base}/hello", timeout=5)
    data = json.loads(resp.read())
    rid = resp.headers.get("X-Request-ID", "")
    print(f"  response: {data}")
    print(f"  X-Request-ID header: {rid}")

    # Request with client-provided ID
    req = urllib.request.Request(
        f"{base}/hello",
        headers={"X-Request-ID": "client-trace-abc123"},
    )
    resp2 = urllib.request.urlopen(req, timeout=5)
    data2 = json.loads(resp2.read())
    print(f"  with client ID: {data2}")

    server.shutdown()
    print("\nDone.")


if __name__ == "__main__":
    main()
