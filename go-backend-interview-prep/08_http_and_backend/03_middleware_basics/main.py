"""Middleware basics -- Python equivalent of the Go example.

Python stdlib has no middleware chain. We simulate it by wrapping
handler methods with decorators / manual composition.
"""

import json
import threading
import time
import traceback
import urllib.request
from http.server import HTTPServer, BaseHTTPRequestHandler


# --- Middleware-style wrappers ---
# In Python stdlib, middleware = wrapping handler logic manually.

class MyHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        # --- Recovery middleware (try/except around everything) ---
        try:
            start = time.perf_counter()
            self._route()
            elapsed = time.perf_counter() - start
            # --- Logging middleware ---
            print(f"  [LOG] {self.command} {self.path}  {elapsed*1000:.1f}ms")
        except Exception:
            print(f"  [RECOVER] panic: {traceback.format_exc().splitlines()[-1]}")
            self.send_response(500)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps({"error": "internal server error"}).encode())

    def _route(self):
        if self.path == "/hello":
            self.send_response(200)
            self.send_header("Content-Type", "text/plain")
            self.end_headers()
            self.wfile.write(b"Hello from middleware demo!\n")

        elif self.path == "/panic":
            raise RuntimeError("something went wrong!")

        else:
            self.send_response(404)
            self.end_headers()
            self.wfile.write(b"not found\n")

    def log_message(self, *args):
        pass


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), MyHandler)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    print(f"Server on {base}")
    print("  GET /hello  -> greeting (logged + timed)")
    print("  GET /panic  -> triggers recovery middleware")
    threading.Thread(target=server.serve_forever, daemon=True).start()

    # --- Demo requests ---
    print("\n--- Demo requests ---")

    resp = urllib.request.urlopen(f"{base}/hello", timeout=5)
    print(f"  /hello -> status {resp.status}")

    try:
        urllib.request.urlopen(f"{base}/panic", timeout=5)
    except urllib.error.HTTPError as e:
        print(f"  /panic -> status {e.code} (recovered)")

    server.shutdown()
    print("\nDone.")


if __name__ == "__main__":
    main()
