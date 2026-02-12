"""HTTP performance basics -- Python equivalent of the Go example.

Demonstrates urllib with timeouts. Python's urllib has limited connection reuse.
For proper pooling, you'd use requests.Session (third-party, not shown).
"""

import json
import threading
import time
import urllib.request
from http.server import HTTPServer, BaseHTTPRequestHandler


class PingHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-Type", "application/json")
        self.end_headers()
        self.wfile.write(json.dumps({"status": "pong"}).encode())

    def log_message(self, *args):
        pass


def main() -> None:
    print("=== HTTP Performance Basics (Python) ===\n")

    # Start local test server
    server = HTTPServer(("127.0.0.1", 0), PingHandler)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    url = f"{base}/ping"
    threading.Thread(target=server.serve_forever, daemon=True).start()

    n = 100
    print(f"  Making {n} requests...\n")

    # --- urllib: each request opens a new connection ---
    start = time.perf_counter()
    for _ in range(n):
        resp = urllib.request.urlopen(url, timeout=5)
        resp.read()
        resp.close()
    urllib_time = time.perf_counter() - start

    print(f"  urllib (no reuse):   {urllib_time:.4f}s  ({n} requests)")
    print(f"  Per request:         {urllib_time/n*1000:.2f}ms")
    print()

    # --- Show connection reuse concept ---
    print("--- Connection reuse in Python ---")
    print("  urllib.request.urlopen: opens a new connection each time")
    print("  http.client.HTTPConnection: can reuse with keep-alive")
    print("  requests.Session(): proper connection pooling (third-party)")
    print()

    # Demo: http.client with keep-alive
    import http.client

    start = time.perf_counter()
    conn = http.client.HTTPConnection(host, port)
    for _ in range(n):
        conn.request("GET", "/ping")
        resp = conn.getresponse()
        resp.read()
    conn.close()
    keepalive_time = time.perf_counter() - start

    print(f"  http.client (reused):  {keepalive_time:.4f}s  ({n} requests)")
    print(f"  Per request:           {keepalive_time/n*1000:.2f}ms")
    print(f"  Speedup:               {urllib_time/keepalive_time:.1f}x")
    print()

    # --- Timeout importance ---
    print("--- Timeouts ---")
    print("  Always set timeouts on HTTP clients:")
    print("    urllib.request.urlopen(url, timeout=10)")
    print("    http.client.HTTPConnection(host, timeout=10)")
    print("    requests.get(url, timeout=10)  # third-party")
    print()
    print("  Without timeouts, a stuck server hangs your thread forever.")
    print()

    print("--- Go vs Python ---")
    print("  Go: http.Client reuses TCP connections by default (keep-alive)")
    print("  Python urllib: no connection reuse (new TCP each time)")
    print("  Python http.client: can reuse a single connection")
    print("  Python requests.Session: proper pool (like Go's Transport)")

    server.shutdown()


if __name__ == "__main__":
    main()
