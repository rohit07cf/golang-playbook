"""DNS and timeouts -- Python equivalent of the Go example."""

import json
import socket
import threading
import time
import urllib.request
import urllib.error
from http.server import HTTPServer, BaseHTTPRequestHandler


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/fast":
            self.send_response(200)
            self.end_headers()
            self.wfile.write(b"fast response")
        elif self.path == "/slow":
            time.sleep(3)
            self.send_response(200)
            self.end_headers()
            self.wfile.write(b"slow response")

    def log_message(self, *args):
        pass


def main() -> None:
    # --- Example 1: DNS lookup ---
    print("=== DNS Lookup ===")
    hosts = ["localhost", "example.com"]
    for host in hosts:
        try:
            hostname, aliases, addrs = socket.gethostbyname_ex(host)
            print(f"  {host} -> {addrs}")
        except socket.gaierror as e:
            print(f"  {host}: error: {e}")

    # --- Example 2: HTTP client timeout ---
    print("\n=== HTTP Client Timeout ===")

    server = HTTPServer(("127.0.0.1", 0), Handler)
    port = server.server_address[1]
    base = f"http://127.0.0.1:{port}"
    threading.Thread(target=server.serve_forever, daemon=True).start()

    try:
        # Fast request
        resp = urllib.request.urlopen(f"{base}/fast", timeout=2)
        print(f"  fast: {resp.read().decode()}")
    except Exception as e:
        print(f"  fast: error: {e}")

    try:
        # Slow request times out
        urllib.request.urlopen(f"{base}/slow", timeout=2)
    except Exception as e:
        print(f"  slow: timed out (expected): {e}")

    # --- Example 3: socket connect timeout ---
    print("\n=== Socket Connect Timeout ===")
    start = time.perf_counter()
    try:
        # Non-routable address to trigger timeout
        socket.create_connection(("192.0.2.1", 12345), timeout=0.5)
    except (socket.timeout, OSError) as e:
        elapsed = time.perf_counter() - start
        print(f"  dial timeout after {elapsed*1000:.0f}ms: {e}")

    server.shutdown()


if __name__ == "__main__":
    main()
