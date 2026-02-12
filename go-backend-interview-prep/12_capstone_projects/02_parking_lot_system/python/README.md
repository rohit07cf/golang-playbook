# Parking Lot -- Python Implementation

## Files

| File | Contains |
|------|----------|
| `parkinglot.py` | All domain entities, strategies, repo, service in one file |
| `test_parkinglot.py` | unittest tests mirroring the Go test coverage |

## Go vs Python Comparison

| Concept | Go | Python |
|---------|-----|--------|
| Enum | `const` + `iota` | `IntEnum` |
| Interface | `type PricingStrategy interface` | Duck typing (no explicit interface) |
| Struct methods | `func (s *Spot) CanFit(...)` | `def can_fit(self, ...)` |
| Constructor | `NewVehicle(plate, vType)` | `Vehicle(plate, vType)` |
| Mutex | `sync.Mutex` in repo | Not needed (GIL for simple cases) |
| Project layout | `cmd/` + `internal/` + packages | Single file (for simplicity) |
| Tests | `_test.go` + `testing` package | `unittest.TestCase` |

## Run

```bash
python3 ./12_capstone_projects/02_parking_lot_system/python/parkinglot.py
python3 -m unittest ./12_capstone_projects/02_parking_lot_system/python/test_parkinglot.py -v
```

## TL;DR

- Same design: entities + strategies + repo + service
- Python uses duck typing instead of explicit interfaces
- One file for readability; Go uses packages for encapsulation
- Tests use `unittest` standard library module
