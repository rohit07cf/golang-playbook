"""TCP / UDP basics -- Python equivalent of the Go example.

Uses socket module for both TCP and UDP.
Server and client run in the same process using threads.
"""

import socket
import threading
import time


def echo_server(server_sock: socket.socket, ready: threading.Event) -> None:
    """Accept one connection and echo lines back."""
    ready.set()
    conn, addr = server_sock.accept()
    with conn:
        while True:
            data = conn.recv(1024)
            if not data:
                break
            line = data.decode().strip()
            if line == "QUIT":
                conn.sendall(b"BYE\n")
                return
            conn.sendall(f"ECHO: {line}\n".encode())


def main() -> None:
    # --- TCP echo server + client ---
    print("=== TCP Echo Server + Client ===")

    server_sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_sock.bind(("127.0.0.1", 0))
    server_sock.listen(1)
    addr = server_sock.getsockname()
    print(f"  server listening on {addr[0]}:{addr[1]}")

    ready = threading.Event()
    t = threading.Thread(target=echo_server, args=(server_sock, ready), daemon=True)
    t.start()
    ready.wait()

    # Client connects
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.settimeout(2.0)
    client.connect(addr)
    print("  client connected")

    messages = ["hello", "world", "Python networking", "QUIT"]
    for msg in messages:
        client.sendall(f"{msg}\n".encode())
        reply = client.recv(1024).decode()
        print(f"  sent: {msg!r} -> got: {reply}", end="")

    client.close()
    server_sock.close()

    # --- UDP example (quick) ---
    print("\n=== UDP (quick demo) ===")
    udp_server = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    udp_server.bind(("127.0.0.1", 0))
    udp_addr = udp_server.getsockname()
    print(f"  udp server on {udp_addr[0]}:{udp_addr[1]}")

    # Client sends to server in another thread
    def udp_client():
        time.sleep(0.05)
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        s.sendto(b"udp hello", udp_addr)
        s.close()

    threading.Thread(target=udp_client, daemon=True).start()

    udp_server.settimeout(2.0)
    try:
        data, remote = udp_server.recvfrom(1024)
        print(f"  received {len(data)} bytes from {remote}: {data.decode()!r}")
    except socket.timeout:
        print("  udp read timed out")

    udp_server.close()


if __name__ == "__main__":
    main()
