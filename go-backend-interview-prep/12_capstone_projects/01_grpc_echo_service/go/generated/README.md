# Generated Code -- gRPC Echo Service

This folder will contain protobuf-generated Go code after running `protoc`.

**Do not edit generated files manually.**

---

## Prerequisites

Install the protoc compiler and Go plugins:

```bash
# 1. Install protoc (Protocol Buffer compiler)
#    macOS:  brew install protobuf
#    Linux:  apt install -y protobuf-compiler
#    Or download from: https://github.com/protocolbuffers/protobuf/releases

# 2. Install Go protoc plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 3. Ensure $GOPATH/bin is in your PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

## Generate Code

From the repository root (`go-backend-interview-prep/`):

```bash
protoc \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  12_capstone_projects/01_grpc_echo_service/go/generated/echopb/echo.proto
```

Or, copy the proto file first:

```bash
mkdir -p 12_capstone_projects/01_grpc_echo_service/go/generated/echopb
cp 12_capstone_projects/01_grpc_echo_service/proto/echo.proto \
   12_capstone_projects/01_grpc_echo_service/go/generated/echopb/

protoc \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  12_capstone_projects/01_grpc_echo_service/go/generated/echopb/echo.proto
```

---

## What Gets Generated

| File | Contains |
|------|----------|
| `echo.pb.go` | Go structs for `EchoRequest` and `EchoResponse` (message types) |
| `echo_grpc.pb.go` | Go interface for `EchoServiceServer` + client stub `EchoServiceClient` |

- **Server** implements the `EchoServiceServer` interface
- **Client** uses the `EchoServiceClient` stub to make RPC calls
- Both files are auto-generated -- never edit them manually

---

## TL;DR

- Install `protoc` + Go plugins (`protoc-gen-go`, `protoc-gen-go-grpc`)
- Run `protoc` with `--go_out` and `--go-grpc_out` flags
- Two files generated: `.pb.go` (messages) and `_grpc.pb.go` (service stubs)
- Server implements the generated interface, client uses the generated stub
