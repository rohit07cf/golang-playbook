"""URL shortener service -- Python equivalent."""

import json
import random
import string
import threading
import time
from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.request import urlopen, Request
from urllib.error import URLError


# --- Store ---

class MemoryStore:
    def __init__(self):
        self.lock = threading.Lock()
        self.urls = {}  # code -> url

    def save(self, code, url):
        with self.lock:
            self.urls[code] = url

    def load(self, code):
        with self.lock:
            return self.urls.get(code)

    def exists(self, code):
        with self.lock:
            return code in self.urls


# --- Code Generator ---

CHARSET = string.ascii_letters + string.digits


def generate_code(length=6):
    return "".join(random.choices(CHARSET, k=length))


def generate_unique_code(store, length=6):
    for _ in range(10):
        code = generate_code(length)
        if not store.exists(code):
            return code
    return generate_code(length + 2)


# --- Handler ---

store = MemoryStore()


class Handler(BaseHTTPRequestHandler):
    def do_POST(self):
        if self.path == "/shorten":
            length = int(self.headers.get("Content-Length", 0))
            body = json.loads(self.rfile.read(length))
            url = body.get("url", "")
            if not url:
                self.send_response(400)
                self.end_headers()
                self.wfile.write(json.dumps({"error": "provide a valid url"}).encode())
                return

            code = generate_unique_code(store)
            store.save(code, url)

            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps({
                "code": code,
                "short_url": f"http://localhost:9002/r/{code}",
            }).encode())
        else:
            self.send_response(404)
            self.end_headers()

    def do_GET(self):
        if self.path.startswith("/r/"):
            code = self.path[3:]
            url = store.load(code)
            if url is None:
                self.send_response(404)
                self.send_header("Content-Type", "application/json")
                self.end_headers()
                self.wfile.write(json.dumps({"error": "not found"}).encode())
                return

            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(json.dumps({
                "code": code,
                "original_url": url,
            }).encode())
        else:
            self.send_response(404)
            self.end_headers()

    def log_message(self, format, *args):
        pass


# --- Demo ---

def run_demo():
    time.sleep(0.5)
    urls = [
        "https://go.dev/doc/effective_go",
        "https://pkg.go.dev/net/http",
        "https://example.com/very/long/path/to/resource",
    ]

    codes = []
    for u in urls:
        data = json.dumps({"url": u}).encode()
        req = Request("http://localhost:9002/shorten", data=data,
                       headers={"Content-Type": "application/json"})
        resp = urlopen(req, timeout=2)
        result = json.loads(resp.read())
        print(f"POST /shorten {u} -> code={result['code']}")
        codes.append(result["code"])

    print()
    for code in codes:
        resp = urlopen(f"http://localhost:9002/r/{code}", timeout=2)
        result = json.loads(resp.read())
        print(f"GET /r/{code} -> {result['original_url']}")

    print()
    try:
        urlopen("http://localhost:9002/r/missing", timeout=2)
    except URLError as e:
        print(f"GET /r/missing -> status={e.code}")

    print("\ndemo done")
    import os
    os._exit(0)


if __name__ == "__main__":
    print("url shortener on :9002")
    t = threading.Thread(target=run_demo, daemon=True)
    t.start()
    server = HTTPServer(("", 9002), Handler)
    server.serve_forever()
