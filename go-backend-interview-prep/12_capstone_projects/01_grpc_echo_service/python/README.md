# Python gRPC Equivalent -- Conceptual

## Why This Is Conceptual

- Real gRPC in Python requires `grpcio` and `grpcio-tools` (external packages)
- This repo uses **standard library only** for Python
- The files here show **what the code would look like** without actually importing grpcio

---

## Real Python gRPC Setup (for reference)

```bash
pip install grpcio grpcio-tools

# Generate code from proto
python -m grpc_tools.protoc \
  -I. \
  --python_out=. \
  --grpc_python_out=. \
  echo.proto

# This generates:
#   echo_pb2.py        (message classes)
#   echo_pb2_grpc.py   (service stubs)
```

---

## What Each File Shows

| File | Purpose |
|------|---------|
| `server_conceptual.py` | What a gRPC server handler looks like in Python (class + method) |
| `client_minimal.py` | Runnable script that simulates gRPC calls with a mock (no deps) |

---

## Go vs Python gRPC Comparison

| Concept | Go | Python |
|---------|-----|--------|
| Proto compilation | `protoc --go_out` | `python -m grpc_tools.protoc --python_out` |
| Server interface | implement generated interface | subclass generated servicer |
| Client stub | use generated client | use generated stub |
| Deadline | `context.WithTimeout` | `timeout=` param on RPC call |
| Server start | `grpc.NewServer() + Serve()` | `grpc.server(ThreadPoolExecutor)` |

---

## TL;DR

- Python gRPC uses `grpcio` (external) -- not stdlib
- The pattern is the same: proto -> generate -> implement server -> call from client
- Python uses `ThreadPoolExecutor` for the gRPC server thread pool
- Deadlines are passed as `timeout=seconds` on RPC calls
- These files show the structure without requiring any external packages
