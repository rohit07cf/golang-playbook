"""SQLite / Repository pattern intro -- Python equivalent of the Go example.

Python has sqlite3 in the stdlib, so we use a real SQLite database (in-memory).
Demonstrates the layered architecture: handler -> service -> repository.
"""

import json
import sqlite3
import threading
import urllib.request
import urllib.error
from http.server import HTTPServer, BaseHTTPRequestHandler


# --- Repository interface (duck typing in Python) ---

class SQLiteRepo:
    """Repository backed by sqlite3 (stdlib). Real database, not a dict."""

    def __init__(self, db_path: str = ":memory:"):
        self.conn = sqlite3.connect(db_path, check_same_thread=False)
        self.conn.row_factory = sqlite3.Row
        self.lock = threading.Lock()
        self.conn.execute("""
            CREATE TABLE IF NOT EXISTS items (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL,
                price REAL NOT NULL
            )
        """)
        self.conn.commit()

    def create(self, name: str, price: float) -> dict:
        with self.lock:
            cur = self.conn.execute(
                "INSERT INTO items (name, price) VALUES (?, ?)",
                (name, price),
            )
            self.conn.commit()
            return {"id": cur.lastrowid, "name": name, "price": price}

    def get_by_id(self, item_id: int) -> dict | None:
        with self.lock:
            row = self.conn.execute(
                "SELECT id, name, price FROM items WHERE id = ?", (item_id,)
            ).fetchone()
        if row is None:
            return None
        return {"id": row["id"], "name": row["name"], "price": row["price"]}

    def list_all(self) -> list[dict]:
        with self.lock:
            rows = self.conn.execute("SELECT id, name, price FROM items").fetchall()
        return [{"id": r["id"], "name": r["name"], "price": r["price"]} for r in rows]

    def delete(self, item_id: int) -> bool:
        with self.lock:
            cur = self.conn.execute("DELETE FROM items WHERE id = ?", (item_id,))
            self.conn.commit()
            return cur.rowcount > 0

    def close(self):
        self.conn.close()


# --- Service layer ---

class ItemService:
    def __init__(self, repo):
        self.repo = repo

    def create_item(self, name: str, price: float) -> dict:
        if not name:
            raise ValueError("name is required")
        if price < 0:
            raise ValueError("price must be non-negative")
        return self.repo.create(name, price)

    def get_item(self, item_id: int) -> dict | None:
        return self.repo.get_by_id(item_id)

    def list_items(self) -> list[dict]:
        return self.repo.list_all()

    def delete_item(self, item_id: int) -> bool:
        return self.repo.delete(item_id)


# --- HTTP Handlers ---

# Module-level service (set in main)
svc: ItemService | None = None


def json_response(handler, status: int, data) -> None:
    handler.send_response(status)
    handler.send_header("Content-Type", "application/json")
    handler.end_headers()
    handler.wfile.write(json.dumps(data).encode())


def read_json(handler) -> dict:
    length = int(handler.headers.get("Content-Length", 0))
    return json.loads(handler.rfile.read(length)) if length else {}


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        import re

        if self.path == "/items":
            items = svc.list_items()
            json_response(self, 200, items)
            return

        m = re.match(r"^/items/(\d+)$", self.path)
        if m:
            item = svc.get_item(int(m.group(1)))
            if item is None:
                json_response(self, 404, {"error": "not found"})
                return
            json_response(self, 200, item)
            return

        json_response(self, 404, {"error": "not found"})

    def do_POST(self):
        if self.path == "/items":
            body = read_json(self)
            try:
                item = svc.create_item(body.get("name", ""), body.get("price", 0))
                json_response(self, 201, item)
            except ValueError as e:
                json_response(self, 400, {"error": str(e)})
            return
        json_response(self, 404, {"error": "not found"})

    def do_DELETE(self):
        import re
        m = re.match(r"^/items/(\d+)$", self.path)
        if m:
            deleted = svc.delete_item(int(m.group(1)))
            if not deleted:
                json_response(self, 404, {"error": "not found"})
                return
            json_response(self, 200, {"deleted": m.group(1)})
            return
        json_response(self, 404, {"error": "not found"})

    def log_message(self, *args):
        pass


def main() -> None:
    global svc

    # Wire up: repo (SQLite) -> service -> handlers
    repo = SQLiteRepo(":memory:")
    svc = ItemService(repo)

    server = HTTPServer(("127.0.0.1", 0), Handler)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    print(f"Repo Pattern API on {base}")
    print("  Architecture: handler -> service -> repository (SQLite)")
    print("  Storage: sqlite3 in-memory (stdlib)")
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

    # Delete
    req5 = urllib.request.Request(f"{base}/items/2", method="DELETE")
    resp5 = urllib.request.urlopen(req5)
    print(f"  DELETE: {resp5.read().decode().strip()}")

    # 404
    try:
        urllib.request.urlopen(f"{base}/items/99")
    except urllib.error.HTTPError as e:
        print(f"  GET 99: status {e.code}")

    server.shutdown()
    repo.close()
    print("\nDone.")


if __name__ == "__main__":
    main()
