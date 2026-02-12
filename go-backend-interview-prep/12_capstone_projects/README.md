# 12 -- Capstone Projects

Two complete mini-projects that combine everything from modules 01-11.

These are **interview-grade implementations** -- small enough to build in 45 minutes,
but structured well enough to demonstrate real engineering skill.

Each capstone includes Go code + Python equivalent for side-by-side learning.

---

## Capstone Projects

| # | Project | Skills Covered |
|---|---------|---------------|
| 01 | **gRPC Echo Service** | Protobuf, gRPC server/client, deadlines, request tracing, code generation |
| 02 | **Parking Lot System** | OOP in Go, SOLID, Strategy pattern, Repository pattern, DI, testing, CLI |

---

## How to Run

**gRPC Echo Service:**
```bash
# 1. Generate code (requires protoc + Go gRPC plugins)
#    See 01_grpc_echo_service/go/generated/README.md for setup

# 2. Server
go run ./12_capstone_projects/01_grpc_echo_service/go/server

# 3. Client (separate terminal)
go run ./12_capstone_projects/01_grpc_echo_service/go/client

# Python conceptual equivalent
python3 ./12_capstone_projects/01_grpc_echo_service/python/client_minimal.py
```

**Parking Lot System:**
```bash
# Go CLI demo
go run ./12_capstone_projects/02_parking_lot_system/go/cmd/parkinglot

# Go tests
go test ./12_capstone_projects/02_parking_lot_system/go/tests/ -v

# Python equivalent
python3 ./12_capstone_projects/02_parking_lot_system/python/parkinglot.py

# Python tests
python3 -m pytest ./12_capstone_projects/02_parking_lot_system/python/test_parkinglot.py -v
# or: python3 -m unittest ./12_capstone_projects/02_parking_lot_system/python/test_parkinglot.py -v
```

---

## Interview Story: How to Present These

- "I built a **gRPC echo service** to learn protobuf, code generation, and deadline propagation."
- "I implemented a **parking lot system** as an LLD exercise -- it uses strategy pattern for pricing and spot allocation, repository pattern for storage, and constructor injection for testability."
- "Both projects have **tests** and follow **clean architecture** with separated layers."
- Mention trade-offs you made (in-memory vs DB, simplified vs production-ready)
- Show you can discuss **what you would change** for production

---

## Common Traps

- Building too much -- keep it interview-sized (30-45 min)
- Not being able to explain your own design decisions
- Skipping tests -- interviewers check for testing instinct
- Over-engineering with patterns you cannot explain
- Forgetting to mention trade-offs and limitations

---

## TL;DR

- Capstones prove you can **combine concepts** into a working system
- gRPC: proto definition -> code generation -> server/client -> deadlines
- Parking lot: entities -> strategies -> repo -> service -> CLI -> tests
- Both follow **clean architecture**: separated layers, interfaces, DI
- Keep implementations small but demonstrate **real patterns**
- Always be ready to explain **why** you chose each pattern
- Tests are not optional -- they show engineering maturity
