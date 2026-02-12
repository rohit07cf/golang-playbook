"""Context timeouts and cancellation -- Python equivalent of the Go example.

Python has no context.Context. We simulate per-request timeouts using
threading.Timer to abort slow operations.
"""

import json
import threading
import time
import urllib.request
import urllib.error
from http.server import HTTPServer, BaseHTTPRequestHandler


def json_response(handler, status: int, data: dict) -> None:
    handler.send_response(status)
    handler.send_header("Content-Type", "application/json")
    handler.end_headers()
    handler.wfile.write(json.dumps(data).encode())


def simulate_work(duration: float, cancel_event: threading.Event) -> str | None:
    """Simulate work that checks for cancellation."""
    deadline = time.monotonic() + duration
    while time.monotonic() < deadline:
        if cancel_event.is_set():
            return None  # cancelled
        time.sleep(0.05)
    return "done"


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/fast":
            self._with_timeout(work_duration=0.1, timeout=2.0)
        elif self.path == "/slow":
            self._with_timeout(work_duration=5.0, timeout=1.0)
        else:
            json_response(self, 404, {"error": "not found"})

    def _with_timeout(self, work_duration: float, timeout: float):
        """Run simulated work with a timeout (like Go's context.WithTimeout)."""
        cancel = threading.Event()
        result_box: list = []

        def worker():
            r = simulate_work(work_duration, cancel)
            result_box.append(r)

        t = threading.Thread(target=worker)
        t.start()
        t.join(timeout=timeout)

        if t.is_alive():
            cancel.set()  # signal cancellation
            t.join()
            json_response(self, 504, {
                "error": "request timed out",
                "cause": f"exceeded {timeout}s deadline",
            })
        else:
            result = result_box[0] if result_box else None
            if result is None:
                json_response(self, 504, {"error": "cancelled"})
            else:
                json_response(self, 200, {"result": result, "endpoint": self.path})

    def log_message(self, *args):
        pass


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), Handler)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    print(f"Server on {base}")
    print("  GET /fast -> completes within timeout")
    print("  GET /slow -> exceeds timeout (504)")
    threading.Thread(target=server.serve_forever, daemon=True).start()

    print("\n--- Demo ---")

    # Fast: should succeed
    resp = urllib.request.urlopen(f"{base}/fast", timeout=5)
    data = json.loads(resp.read())
    print(f"  /fast -> {resp.status}: {data}")

    # Slow: should timeout (504)
    try:
        resp2 = urllib.request.urlopen(f"{base}/slow", timeout=5)
        data2 = json.loads(resp2.read())
        print(f"  /slow -> {resp2.status}: {data2}")
    except urllib.error.HTTPError as e:
        data2 = json.loads(e.read())
        print(f"  /slow -> {e.code}: {data2}")

    server.shutdown()
    print("\nDone.")


if __name__ == "__main__":
    main()
