# 01 -- gRPC Echo Service (Capstone)

## What We Are Building

- A minimal gRPC service with one RPC: `Echo(EchoRequest) returns (EchoResponse)`
- **ELI10:** gRPC is a strict contract phone call -- both sides agree on exactly what words mean before anyone speaks.
- Server + client in Go, conceptual Python equivalent

## Requirements

**Functional:**
- Client sends a message + request_id
- Server echoes the message back with server_time and request_id
- Client demonstrates deadline/timeout on one call

**Non-functional:**
- Protobuf for serialization (schema-first API design)
- gRPC for transport (HTTP/2, bidirectional streaming capable)
- Request tracing via request_id field

## High-Level Design

```
+--------+   gRPC/HTTP2   +--------+
| Client | -------------> | Server |
| (Go)   | <------------- | (Go)   |
+--------+  EchoResponse  +--------+
     |                         |
  sends:                   returns:
  - message                - message (echoed)
  - request_id             - request_id
                           - server_time
```

## Key Go Building Blocks

- `google.golang.org/grpc` -- gRPC framework
- `google.golang.org/protobuf` -- protobuf runtime
- `protoc` -- code generator (proto -> Go stubs)
- `context.WithTimeout` -- client-side deadline
- `grpc.NewServer()` -- create server
- `grpc.Dial()` -- create client connection

## What to Say in Interviews

- "gRPC uses **protobuf** for schema-first API design -- the proto file is the contract"
- "The server implements a generated interface, the client calls generated stubs"
- "Deadlines propagate from client to server via gRPC metadata -- no manual plumbing"
- "gRPC runs over **HTTP/2** which gives us multiplexing, header compression, and streaming"
- "I chose gRPC over REST because of strong typing, code generation, and streaming support"

## Common Traps

- Forgetting to install protoc plugins (`protoc-gen-go`, `protoc-gen-go-grpc`)
- Not understanding the difference between generated `.pb.go` (messages) and `_grpc.pb.go` (service stubs)
- Confusing deadline (absolute time) with timeout (relative duration)
- Not closing the client connection (resource leak)
- Assuming gRPC works in browsers (it does not without grpc-web proxy)

## Trade-Offs

- **gRPC vs REST** -- gRPC is faster and typed, but harder to debug (binary protocol)
- **Unary vs streaming** -- unary is simpler; streaming adds complexity but handles real-time data
- **Protobuf vs JSON** -- protobuf is smaller/faster but not human-readable
- **ELI10:** Protobuf vs JSON is like Morse code vs English -- one is compact and fast, the other is human-friendly.
- **Code generation** -- ensures type safety but adds a build step

## Run It

```bash
# See go/generated/README.md for protoc setup and code generation

# Terminal 1: start server
go run ./12_capstone_projects/01_grpc_echo_service/go/server

# Terminal 2: run client
go run ./12_capstone_projects/01_grpc_echo_service/go/client

# Python conceptual equivalent
python3 ./12_capstone_projects/01_grpc_echo_service/python/client_minimal.py
```

## TL;DR

- Proto file defines the contract: service + messages
- `protoc` generates Go stubs (messages + service interface)
- Server implements the generated interface
- Client uses generated stubs to call RPCs
- Deadlines/timeouts propagate automatically via context
- gRPC = HTTP/2 + protobuf + code generation + streaming
- In interviews: emphasize **schema-first design** and **type safety**
