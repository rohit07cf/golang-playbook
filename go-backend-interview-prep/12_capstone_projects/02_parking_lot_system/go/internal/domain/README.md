# Domain Layer

Pure business entities and interfaces. No I/O, no external dependencies.

## Files

| File | Contains |
|------|----------|
| `vehicle.go` | VehicleType enum + Vehicle struct + NewVehicle factory |
| `spot.go` | SpotSize enum + Spot struct + CanFit/Park/Free methods |
| `ticket.go` | Ticket struct + NewTicket factory + Close/Duration methods |
| `pricing.go` | PricingStrategy interface + HourlyPricing + FlatPricing |
| `lot.go` | SpotAllocationStrategy interface + NearestAvailable + CreateSpots |

## TL;DR

- Entities own their data and basic behavior
- Strategy interfaces defined here, implementations are pluggable
- No imports from service or repo layers (dependency direction: inward)
