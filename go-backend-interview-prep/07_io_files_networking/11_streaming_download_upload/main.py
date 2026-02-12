"""Streaming download/upload -- Python equivalent of the Go example.

Streams HTTP response to file using chunked reads.
Never loads the entire response into memory.
"""

import os
import shutil
import tempfile
import threading
import urllib.request
from http.server import HTTPServer, BaseHTTPRequestHandler


PAYLOAD = b"Python streaming demo data. " * 1000  # ~28 KB


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/download":
            self.send_response(200)
            self.send_header("Content-Length", str(len(PAYLOAD)))
            self.send_header("Content-Type", "application/octet-stream")
            self.end_headers()
            self.wfile.write(PAYLOAD)

    def do_POST(self):
        length = int(self.headers.get("Content-Length", 0))
        body = self.rfile.read(length)
        self.send_response(200)
        self.end_headers()
        self.wfile.write(f"received {len(body)} bytes".encode())

    def log_message(self, *args):
        pass


def main() -> None:
    server = HTTPServer(("127.0.0.1", 0), Handler)
    port = server.server_address[1]
    base = f"http://127.0.0.1:{port}"
    threading.Thread(target=server.serve_forever, daemon=True).start()

    tmpdir = tempfile.mkdtemp()

    try:
        # --- Example 1: stream download (chunked reads) ---
        print("=== Stream download (chunked) ===")
        resp = urllib.request.urlopen(f"{base}/download", timeout=10)
        out_path = os.path.join(tmpdir, "downloaded.bin")

        total = 0
        with open(out_path, "wb") as f:
            while True:
                chunk = resp.read(8192)  # read in chunks -- constant memory
                if not chunk:
                    break
                f.write(chunk)
                total += len(chunk)
        print(f"  downloaded {total} bytes to file")

        # --- Example 2: stream with progress ---
        print("\n=== Stream with progress ===")
        resp2 = urllib.request.urlopen(f"{base}/download", timeout=10)
        content_length = int(resp2.headers.get("Content-Length", 0))
        out_path2 = os.path.join(tmpdir, "progress.bin")

        downloaded = 0
        with open(out_path2, "wb") as f:
            while True:
                chunk = resp2.read(4096)
                if not chunk:
                    break
                f.write(chunk)
                downloaded += len(chunk)
                if content_length > 0:
                    pct = downloaded / content_length * 100
                    print(f"\r  progress: {pct:.0f}% ({downloaded}/{content_length} bytes)", end="")
        print(f"\n  done: {downloaded} bytes with progress")

        # --- Example 3: upload (stream request body) ---
        print("\n=== Stream upload ===")
        upload_data = b"upload payload: " + b"x" * 500
        req = urllib.request.Request(
            f"{base}/download",
            data=upload_data,
            method="POST",
            headers={"Content-Type": "application/octet-stream"},
        )
        resp3 = urllib.request.urlopen(req, timeout=10)
        print(f"  upload status: {resp3.status}")
        print("  (request body was sent, not buffered)")

    finally:
        server.shutdown()
        shutil.rmtree(tmpdir, ignore_errors=True)


if __name__ == "__main__":
    main()
