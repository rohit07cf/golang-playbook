"""JSON API CRUD (in-memory) -- Python equivalent of the Go example."""

import json
import re
import threading
import urllib.request
import urllib.error
from http.server import HTTPServer, BaseHTTPRequestHandler

# In-memory store (thread-safe with lock)
store: dict[int, dict] = {}
lock = threading.Lock()
next_id = 1


def json_response(handler, status: int, data) -> None:
    handler.send_response(status)
    handler.send_header("Content-Type", "application/json")
    handler.end_headers()
    handler.wfile.write(json.dumps(data).encode())


def read_json(handler) -> dict:
    length = int(handler.headers.get("Content-Length", 0))
    return json.loads(handler.rfile.read(length)) if length else {}


class CRUDHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        # GET /items -- list
        if self.path == "/items":
            with lock:
                items = list(store.values())
            json_response(self, 200, items)
            return

        # GET /items/{id}
        m = re.match(r"^/items/(\d+)$", self.path)
        if m:
            item_id = int(m.group(1))
            with lock:
                item = store.get(item_id)
            if item is None:
                json_response(self, 404, {"error": "not found"})
                return
            json_response(self, 200, item)
            return

        json_response(self, 404, {"error": "not found"})

    def do_POST(self):
        if self.path == "/items":
            global next_id
            body = read_json(self)
            name = body.get("name", "")
            if not name:
                json_response(self, 400, {"error": "name required"})
                return
            price = body.get("price", 0)
            with lock:
                item = {"id": next_id, "name": name, "price": price}
                store[next_id] = item
                next_id += 1
            json_response(self, 201, item)
            return
        json_response(self, 404, {"error": "not found"})

    def do_PUT(self):
        m = re.match(r"^/items/(\d+)$", self.path)
        if m:
            item_id = int(m.group(1))
            body = read_json(self)
            with lock:
                item = store.get(item_id)
                if item is None:
                    json_response(self, 404, {"error": "not found"})
                    return
                if body.get("name"):
                    item["name"] = body["name"]
                if body.get("price"):
                    item["price"] = body["price"]
                store[item_id] = item
            json_response(self, 200, item)
            return
        json_response(self, 404, {"error": "not found"})

    def do_DELETE(self):
        m = re.match(r"^/items/(\d+)$", self.path)
        if m:
            item_id = int(m.group(1))
            with lock:
                if item_id not in store:
                    json_response(self, 404, {"error": "not found"})
                    return
                del store[item_id]
            json_response(self, 200, {"deleted": str(item_id)})
            return
        json_response(self, 404, {"error": "not found"})

    def log_message(self, *args):
        pass


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), CRUDHandler)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    print(f"CRUD API on {base}")
    threading.Thread(target=server.serve_forever, daemon=True).start()

    print("\n--- Demo ---")

    # Create
    req = urllib.request.Request(
        f"{base}/items",
        data=json.dumps({"name": "widget", "price": 9.99}).encode(),
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    resp = urllib.request.urlopen(req)
    print(f"  CREATE: {resp.read().decode().strip()}")

    # Create another
    req2 = urllib.request.Request(
        f"{base}/items",
        data=json.dumps({"name": "gadget", "price": 19.99}).encode(),
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    urllib.request.urlopen(req2).read()

    # List
    resp3 = urllib.request.urlopen(f"{base}/items")
    items = json.loads(resp3.read())
    print(f"  LIST:   {len(items)} items")

    # Get
    resp4 = urllib.request.urlopen(f"{base}/items/1")
    print(f"  GET:    {resp4.read().decode().strip()}")

    # Update
    req5 = urllib.request.Request(
        f"{base}/items/1",
        data=json.dumps({"name": "super widget", "price": 14.99}).encode(),
        headers={"Content-Type": "application/json"},
        method="PUT",
    )
    resp5 = urllib.request.urlopen(req5)
    print(f"  UPDATE: {resp5.read().decode().strip()}")

    # Delete
    req6 = urllib.request.Request(f"{base}/items/2", method="DELETE")
    resp6 = urllib.request.urlopen(req6)
    print(f"  DELETE: {resp6.read().decode().strip()}")

    server.shutdown()
    print("Done.")


if __name__ == "__main__":
    main()
