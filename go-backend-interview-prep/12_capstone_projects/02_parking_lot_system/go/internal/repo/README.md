# Repository Layer

Abstracts data persistence behind an interface.

## Files

| File | Contains |
|------|----------|
| `in_memory_repo.go` | ParkingRepo interface + InMemoryRepo (map-based, mutex-protected) |

## Key Points

- **ParkingRepo interface**: GetSpots, GetSpotByID, SaveTicket, GetTicket, GetActiveTicketByPlate, NextTicketID
- **InMemoryRepo**: stores spots as slice, tickets as map, protected by `sync.Mutex`
- Swap in a database implementation later without changing the service layer

## TL;DR

- Interface defined here, consumed by service
- Mutex protects concurrent access to shared state
- In-memory for demos and tests; DB for production
