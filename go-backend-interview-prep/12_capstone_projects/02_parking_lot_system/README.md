# 02 -- Parking Lot System (Capstone)

## What We Are Building

- A classic LLD (Low-Level Design) interview problem: parking lot with multiple vehicle types
- **ELI10:** The parking lot problem is just controlled chaos with rules -- vehicles come in, spots get assigned, money gets collected.
- Clean architecture with domain entities, strategy pattern, repository pattern, and DI

## Requirements

**Functional:**
- Support vehicle types: Motorcycle, Car, Truck
- Spot types: Small, Medium, Large (with fitting rules)
- Park a vehicle: allocate spot, issue ticket
- Unpark a vehicle: calculate fee by duration and type, free spot
- Show lot occupancy summary

**Non-functional:**
- Thread-safety in the repo layer (mutex in Go)
- Testable via dependency injection
- Configurable pricing and allocation strategies

## High-Level Design

```
+--------+     +---------+     +---------+     +-----------+
|  CLI   | --> | Parking | --> | Repo    | --> | In-Memory |
| (main) |     | Service |     | (iface) |     | Store     |
+--------+     +---------+     +---------+     +-----------+
                  |    |
           +------+    +-------+
           |                   |
     +-----------+      +----------+
     | Pricing   |      | Allocator|
     | Strategy  |      | Strategy |
     +-----------+      +----------+
```

## Design Patterns Used

- **ELI10:** Strategy pattern here is like having multiple pricing plans -- same parking lot, different rates depending on the plan.
- **Strategy**: PricingStrategy (hourly vs flat), SpotAllocationStrategy (nearest available)
- **Repository**: ParkingRepo interface decouples service from storage
- **Factory**: NewVehicle, NewTicket, CreateSpots -- factory functions for entities
- **Dependency Injection**: ParkingService accepts interfaces via constructor
- **SOLID**:
  - S: Each entity has one responsibility (Vehicle, Spot, Ticket)
  - O: New pricing/allocation strategies without modifying service
  - L: Any PricingStrategy implementation works in the service
  - I: Small interfaces (PricingStrategy: 1 method, SpotAllocationStrategy: 1 method)
  - D: Service depends on interfaces, not concrete implementations

## What to Say in Interviews

- "I separated domain entities from the service layer for testability"
- "Pricing and allocation use the **strategy pattern** -- I can swap algorithms without changing the service"
- "The **repository interface** means I can replace in-memory storage with a database without touching business logic"
- "Dependencies are **injected via constructor** -- no global state, easy to test"
- "Thread safety is handled at the repo layer with a mutex"

## Common Traps

- Making Vehicle a class with Park/Unpark methods (vehicles do not park themselves)
- Using inheritance instead of composition (Go does not have inheritance)
- Hardcoding pricing logic inside the service (violates Open/Closed)
- Forgetting thread safety on the shared spot map
- Not testing edge cases: full lot, invalid unpark, motorcycle in large spot
- Over-engineering with event systems or notification patterns

## Run It

```bash
# Go CLI demo
go run ./12_capstone_projects/02_parking_lot_system/go/cmd/parkinglot

# Go tests
go test ./12_capstone_projects/02_parking_lot_system/go/tests/ -v

# Python CLI demo
python3 ./12_capstone_projects/02_parking_lot_system/python/parkinglot.py

# Python tests
python3 -m unittest 12_capstone_projects/02_parking_lot_system/python/test_parkinglot.py -v
```

## TL;DR

- Entities: Vehicle, Spot, Ticket -- each with a single responsibility
- Strategy pattern for pricing (hourly/flat) and allocation (nearest available)
- Repository interface abstracts storage -- swap in-memory for DB later
- Constructor injection makes everything testable
- Go: `cmd/` + `internal/` structure; Python: single file with classes
- Tests cover: allocation, full lot, fee calculation, unpark, occupancy
- In interviews: lead with the **patterns**, then show the code
