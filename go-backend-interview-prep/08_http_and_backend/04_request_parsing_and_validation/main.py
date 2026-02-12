"""Request parsing and validation -- Python equivalent of the Go example."""

import json
import threading
import urllib.request
import urllib.error
from http.server import HTTPServer, BaseHTTPRequestHandler


def json_response(handler, status: int, data: dict) -> None:
    handler.send_response(status)
    handler.send_header("Content-Type", "application/json")
    handler.end_headers()
    handler.wfile.write(json.dumps(data).encode())


class Handler(BaseHTTPRequestHandler):
    def do_POST(self):
        if self.path == "/users":
            self._create_user()
            return
        json_response(self, 404, {"error": "not found"})

    def do_GET(self):
        if self.path.startswith("/search"):
            self._search()
            return
        json_response(self, 404, {"error": "not found"})

    def _create_user(self):
        # Check Content-Type
        ct = self.headers.get("Content-Type", "")
        if "application/json" not in ct:
            json_response(self, 400, {"error": "Content-Type must be application/json"})
            return

        # Decode JSON
        length = int(self.headers.get("Content-Length", 0))
        if length > 1_048_576:  # 1 MB limit
            json_response(self, 400, {"error": "body too large"})
            return

        try:
            body = json.loads(self.rfile.read(length))
        except json.JSONDecodeError as e:
            json_response(self, 400, {"error": f"invalid JSON: {e}"})
            return

        # Validate fields
        errors = []
        name = body.get("name", "")
        email = body.get("email", "")
        if not name:
            errors.append("name is required")
        if not email:
            errors.append("email is required")
        elif "@" not in email:
            errors.append("email must contain @")

        if errors:
            json_response(self, 400, {"error": "validation failed", "fields": errors})
            return

        json_response(self, 201, {"message": "user created", "name": name, "email": email})

    def _search(self):
        from urllib.parse import urlparse, parse_qs
        qs = parse_qs(urlparse(self.path).query)
        page = qs.get("page", ["1"])[0]
        limit = qs.get("limit", ["10"])[0]
        json_response(self, 200, {"page": page, "limit": limit})

    def log_message(self, *args):
        pass


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), Handler)
    host, port = server.server_address
    base = f"http://{host}:{port}"
    print(f"Server on {base}")
    threading.Thread(target=server.serve_forever, daemon=True).start()

    print("\n--- Demo ---")

    # Valid request
    req = urllib.request.Request(
        f"{base}/users",
        data=json.dumps({"name": "Alice", "email": "alice@example.com"}).encode(),
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    resp = urllib.request.urlopen(req)
    print(f"  valid:     {resp.read().decode()}")

    # Missing fields
    req2 = urllib.request.Request(
        f"{base}/users",
        data=b"{}",
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    try:
        urllib.request.urlopen(req2)
    except urllib.error.HTTPError as e:
        print(f"  invalid:   {e.read().decode()}")

    # Bad email
    req3 = urllib.request.Request(
        f"{base}/users",
        data=json.dumps({"name": "Bob", "email": "nope"}).encode(),
        headers={"Content-Type": "application/json"},
        method="POST",
    )
    try:
        urllib.request.urlopen(req3)
    except urllib.error.HTTPError as e:
        print(f"  bad email: {e.read().decode()}")

    # Query params
    resp4 = urllib.request.urlopen(f"{base}/search?page=2&limit=25")
    print(f"  query:     {resp4.read().decode()}")

    server.shutdown()
    print("Done.")


if __name__ == "__main__":
    main()
