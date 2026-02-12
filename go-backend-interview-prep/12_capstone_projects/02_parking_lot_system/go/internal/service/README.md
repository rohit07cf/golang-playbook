# Service Layer

Orchestrates domain entities and repo to fulfill use cases.

## Files

| File | Contains |
|------|----------|
| `parking_service.go` | ParkingService with Park, Unpark, GetOccupancy |

## Key Points

- Accepts interfaces via constructor (repo, allocator, pricing)
- Does not know about storage implementation details
- Contains business rules: allocate spot, issue ticket, calculate fee

## TL;DR

- Service = glue between domain + repo
- All dependencies are injected -- easy to test with fakes
- Business logic lives here, not in entities or repo
