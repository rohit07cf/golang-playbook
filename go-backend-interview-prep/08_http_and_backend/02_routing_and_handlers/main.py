"""Routing and handlers -- Python equivalent of the Go example.

Manual path parsing with stdlib http.server (no Flask/FastAPI).
"""

import json
import re
import threading
import urllib.request
from http.server import HTTPServer, BaseHTTPRequestHandler

# In-memory store
items: dict[str, str] = {"1": "alpha", "2": "bravo"}
next_id = 3


def json_response(handler, status: int, data: dict) -> None:
    handler.send_response(status)
    handler.send_header("Content-Type", "application/json")
    handler.end_headers()
    handler.wfile.write(json.dumps(data).encode())


class Router(BaseHTTPRequestHandler):
    """Manual routing -- stdlib has no built-in router."""

    def do_GET(self):
        # GET /items -- list all
        if self.path == "/items":
            json_response(self, 200, items)
            return

        # GET /items/{id} -- get one
        m = re.match(r"^/items/(\w+)$", self.path)
        if m:
            item_id = m.group(1)
            if item_id not in items:
                json_response(self, 404, {"error": "not found"})
                return
            json_response(self, 200, {"id": item_id, "name": items[item_id]})
            return

        json_response(self, 404, {"error": "route not found", "path": self.path})

    def do_POST(self):
        if self.path == "/items":
            global next_id
            length = int(self.headers.get("Content-Length", 0))
            body = json.loads(self.rfile.read(length))
            name = body.get("name", "")
            if not name:
                json_response(self, 400, {"error": "name required"})
                return
            item_id = str(next_id)
            next_id += 1
            items[item_id] = name
            json_response(self, 201, {"id": item_id, "name": name})
            return

        json_response(self, 404, {"error": "route not found"})

    def do_DELETE(self):
        m = re.match(r"^/items/(\w+)$", self.path)
        if m:
            item_id = m.group(1)
            if item_id not in items:
                json_response(self, 404, {"error": "not found"})
                return
            del items[item_id]
            json_response(self, 200, {"deleted": item_id})
            return

        json_response(self, 404, {"error": "route not found"})

    def log_message(self, *args):
        pass


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), Router)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    print(f"Server on {base}")
    threading.Thread(target=server.serve_forever, daemon=True).start()

    # --- Demo requests ---
    print("\n--- Demo ---")

    # List items
    resp = urllib.request.urlopen(f"{base}/items")
    print(f"GET /items       -> {resp.read().decode()}")

    # Get one
    resp = urllib.request.urlopen(f"{base}/items/1")
    print(f"GET /items/1     -> {resp.read().decode()}")

    # Create
    req = urllib.request.Request(
        f"{base}/items",
        data=json.dumps({"name": "charlie"}).encode(),
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    resp = urllib.request.urlopen(req)
    print(f"POST /items      -> {resp.read().decode()}")

    # Delete
    req = urllib.request.Request(f"{base}/items/2", method="DELETE")
    resp = urllib.request.urlopen(req)
    print(f"DELETE /items/2   -> {resp.read().decode()}")

    # 404
    try:
        urllib.request.urlopen(f"{base}/nope")
    except urllib.error.HTTPError as e:
        print(f"GET /nope        -> {e.code}: {e.read().decode()}")

    server.shutdown()


if __name__ == "__main__":
    main()
