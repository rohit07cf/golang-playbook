# Capstone Projects -- Quick Revision

One-screen cheat sheet for the two capstone projects.

---

## gRPC Echo Service Flow

```
proto file  -->  protoc  -->  generated stubs  -->  server implements interface  -->  client calls stubs
```

1. Define messages + service in `.proto`
2. Run `protoc` to generate Go/Python stubs
3. Server: implement the generated `EchoServiceServer` interface
4. Client: create connection, use generated `EchoServiceClient`
5. Deadline: `context.WithTimeout` on client side, propagates to server

---

## Parking Lot System Flow

```
entities  -->  strategies  -->  repo (interface)  -->  service  -->  CLI  -->  tests
```

1. **Entities**: Vehicle (type + plate), Spot (size + occupied), Ticket (entry/exit/fee)
2. **Strategies**: PricingStrategy (hourly/flat), SpotAllocationStrategy (nearest fit)
3. **Repo**: ParkingRepo interface + InMemoryRepo (map + mutex)
4. **Service**: ParkingService (Park, Unpark, GetOccupancy) -- DI via constructor
5. **Tests**: allocation, full lot, fee calculation, unpark, occupancy, spot fit rules

---

## 8 Interview One-Liners

| # | One-Liner |
|---|-----------|
| 1 | "gRPC uses **protobuf** for schema-first design -- the proto file is the contract between client and server." |
| 2 | "Deadlines propagate automatically in gRPC -- `context.WithTimeout` on the client side is enforced on the server." |
| 3 | "I used the **strategy pattern** for pricing and spot allocation so I can swap algorithms without changing the service." |
| 4 | "The **repository interface** decouples business logic from storage -- I can swap in-memory for a database without touching the service." |
| 5 | "All dependencies are **injected via constructor** -- no global state, easy to test with fakes." |
| 6 | "The domain layer is **pure logic** -- no I/O, no frameworks, no external dependencies." |
| 7 | "Go's `internal/` package prevents external imports -- it enforces encapsulation at the compiler level." |
| 8 | "Tests exercise the full stack through the service layer -- if the tests pass, the business logic is correct." |

---

## TL;DR

- gRPC: proto -> generate -> implement -> call -> deadline propagation
- Parking lot: entities -> strategies -> repo -> service -> tests
- Strategy pattern = swap algorithms without changing consumers
- Repository pattern = swap storage without changing business logic
- Constructor injection = testable, no global state
- Always be ready to explain **why** you chose each pattern
