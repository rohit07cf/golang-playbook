"""Graceful shutdown -- Python equivalent of the Go example.

Python's HTTPServer.shutdown() stops serve_forever() but doesn't drain
in-flight requests as gracefully as Go's Shutdown(ctx). We demonstrate
signal handling and the best stdlib approach.
"""

import json
import signal
import threading
import time
import urllib.request
from http.server import HTTPServer, BaseHTTPRequestHandler


def json_response(handler, status: int, data: dict) -> None:
    handler.send_response(status)
    handler.send_header("Content-Type", "application/json")
    handler.end_headers()
    handler.wfile.write(json.dumps(data).encode())


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/health":
            json_response(self, 200, {"status": "ok"})
        elif self.path == "/slow":
            print("  [handler] slow request started, sleeping 2s...")
            time.sleep(2)
            json_response(self, 200, {"message": "slow request completed"})
            print("  [handler] slow request done")
        else:
            json_response(self, 404, {"error": "not found"})

    def log_message(self, *args):
        pass


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), Handler)
    host, port = server.server_address
    base = f"http://{host}:{port}"

    # --- Signal handler for graceful shutdown ---
    def on_signal(signum, frame):
        name = signal.Signals(signum).name
        print(f"\n  [shutdown] received signal: {name}")
        print("  [shutdown] stopping server...")
        # shutdown() must run in a separate thread (it blocks until serve_forever stops)
        threading.Thread(target=server.shutdown).start()

    signal.signal(signal.SIGINT, on_signal)
    signal.signal(signal.SIGTERM, on_signal)

    print(f"Server on {base}")
    print("  GET /health -> health check")
    print("  GET /slow   -> 2s slow request")
    print("  Press Ctrl+C to trigger graceful shutdown")

    # Self-test demo
    def demo():
        time.sleep(0.2)
        print("\n--- Demo ---")

        # Start slow request in background
        result = [None]

        def slow_req():
            try:
                resp = urllib.request.urlopen(f"{base}/slow", timeout=10)
                result[0] = (resp.status, json.loads(resp.read()))
            except Exception as e:
                result[0] = ("error", str(e))

        t = threading.Thread(target=slow_req)
        t.start()

        # While slow request is in-flight, trigger shutdown
        time.sleep(0.5)
        print("  [demo] triggering shutdown while slow request is in-flight...")
        threading.Thread(target=server.shutdown).start()

        # Wait for slow request
        t.join()
        if result[0]:
            print(f"  slow request result: {result[0]}")

    threading.Thread(target=demo, daemon=True).start()

    # serve_forever blocks until shutdown() is called
    server.serve_forever()
    print("  [shutdown] server stopped")
    print("\nServer exited.")


if __name__ == "__main__":
    main()
