# Parking Lot -- Go Implementation

## Project Structure

```
go/
  cmd/parkinglot/main.go     -- CLI demo entry point
  internal/
    domain/
      vehicle.go              -- Vehicle entity + VehicleType enum
      spot.go                 -- Spot entity + SpotSize enum + CanFit logic
      ticket.go               -- Ticket entity + Duration/Close
      pricing.go              -- PricingStrategy interface + HourlyPricing + FlatPricing
      lot.go                  -- SpotAllocationStrategy + NearestAvailable + CreateSpots
    service/
      parking_service.go      -- ParkingService (Park, Unpark, GetOccupancy)
    repo/
      in_memory_repo.go       -- ParkingRepo interface + InMemoryRepo (mutex-protected)
  tests/
    parkinglot_test.go        -- All tests
```

## Key Design Decisions

- **`internal/`** prevents external imports -- enforces encapsulation
- **`cmd/`** holds the executable entry point -- standard Go project layout
- **Interfaces defined where consumed**: PricingStrategy and SpotAllocationStrategy in domain, ParkingRepo in repo
- **Mutex in repo layer**: protects shared state without polluting domain logic
- **Factory functions** (NewVehicle, NewTicket) instead of constructors -- Go-idiomatic

## Run

```bash
go run ./12_capstone_projects/02_parking_lot_system/go/cmd/parkinglot
go test ./12_capstone_projects/02_parking_lot_system/go/tests/ -v
```

## TL;DR

- Standard Go layout: `cmd/` for binaries, `internal/` for private packages
- Domain layer is pure logic -- no I/O, no frameworks
- Service layer orchestrates domain objects
- Repo layer abstracts persistence behind an interface
- Tests exercise all layers via the service
