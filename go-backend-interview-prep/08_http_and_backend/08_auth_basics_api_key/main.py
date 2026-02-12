"""Auth basics (API key) -- Python equivalent of the Go example.

Simple header-based API key check. This is a toy -- production uses JWT/OAuth.
"""

import hmac
import json
import threading
import urllib.request
import urllib.error
from http.server import HTTPServer, BaseHTTPRequestHandler

API_KEY = "secret123"


def json_response(handler, status: int, data: dict) -> None:
    handler.send_response(status)
    handler.send_header("Content-Type", "application/json")
    handler.end_headers()
    handler.wfile.write(json.dumps(data).encode())


def check_api_key(handler) -> bool:
    """Check X-API-Key header. Returns True if valid."""
    key = handler.headers.get("X-API-Key", "")
    # Constant-time comparison (like Go's subtle.ConstantTimeCompare)
    return hmac.compare_digest(key, API_KEY)


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        # Public endpoint -- no auth
        if self.path == "/health":
            json_response(self, 200, {"status": "ok"})
            return

        # Protected endpoint -- requires API key
        if self.path == "/protected":
            if not check_api_key(self):
                json_response(self, 401, {
                    "error": "unauthorized: invalid or missing API key",
                })
                return
            json_response(self, 200, {
                "message": "welcome, authenticated user!",
                "secret": "the answer is 42",
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
    print("  GET /health     -> public (no auth)")
    print("  GET /protected  -> requires X-API-Key header")
    print(f"  API key: {API_KEY}")
    threading.Thread(target=server.serve_forever, daemon=True).start()

    print("\n--- Demo ---")

    # Public endpoint
    resp = urllib.request.urlopen(f"{base}/health", timeout=5)
    print(f"  /health (no key)    -> {resp.status}: {resp.read().decode()}")

    # Protected: no key -> 401
    try:
        urllib.request.urlopen(f"{base}/protected", timeout=5)
    except urllib.error.HTTPError as e:
        print(f"  /protected (no key) -> {e.code}: {e.read().decode()}")

    # Protected: wrong key -> 401
    req = urllib.request.Request(
        f"{base}/protected",
        headers={"X-API-Key": "wrong"},
    )
    try:
        urllib.request.urlopen(req, timeout=5)
    except urllib.error.HTTPError as e:
        print(f"  /protected (wrong)  -> {e.code}: {e.read().decode()}")

    # Protected: valid key -> 200
    req2 = urllib.request.Request(
        f"{base}/protected",
        headers={"X-API-Key": API_KEY},
    )
    resp2 = urllib.request.urlopen(req2, timeout=5)
    print(f"  /protected (valid)  -> {resp2.status}: {resp2.read().decode()}")

    server.shutdown()
    print("\nDone.")


if __name__ == "__main__":
    main()
