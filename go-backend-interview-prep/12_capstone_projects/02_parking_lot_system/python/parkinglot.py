"""
Parking Lot System -- Python equivalent.

Mirrors the Go implementation:
- Entities: Vehicle, Spot, Ticket
- Strategies: PricingStrategy, SpotAllocationStrategy
- Repository: InMemoryRepo
- Service: ParkingService

Run: python3 parkinglot.py
"""

import math
from datetime import datetime, timedelta
from enum import IntEnum


# --- Domain: Vehicle ---

class VehicleType(IntEnum):
    MOTORCYCLE = 0
    CAR = 1
    TRUCK = 2


class Vehicle:
    def __init__(self, license_plate: str, vehicle_type: VehicleType):
        self.license_plate = license_plate
        self.type = vehicle_type

    def __repr__(self):
        return f"{self.license_plate} ({self.type.name})"


# --- Domain: Spot ---

class SpotSize(IntEnum):
    SMALL = 0
    MEDIUM = 1
    LARGE = 2


class Spot:
    def __init__(self, spot_id: int, size: SpotSize):
        self.id = spot_id
        self.size = size
        self.occupied = False
        self.vehicle = None

    def can_fit(self, vehicle_type: VehicleType) -> bool:
        if vehicle_type == VehicleType.MOTORCYCLE:
            return True
        if vehicle_type == VehicleType.CAR:
            return self.size >= SpotSize.MEDIUM
        if vehicle_type == VehicleType.TRUCK:
            return self.size >= SpotSize.LARGE
        return False

    def park(self, vehicle: Vehicle):
        self.occupied = True
        self.vehicle = vehicle

    def free(self):
        self.occupied = False
        self.vehicle = None


# --- Domain: Ticket ---

class Ticket:
    def __init__(self, ticket_id: int, vehicle: Vehicle, spot_id: int):
        self.id = ticket_id
        self.vehicle = vehicle
        self.spot_id = spot_id
        self.entry_time = datetime.now()
        self.exit_time = None
        self.fee = 0.0
        self.active = True

    def close(self, fee: float):
        self.exit_time = datetime.now()
        self.fee = fee
        self.active = False

    def duration(self) -> timedelta:
        end = self.exit_time if self.exit_time else datetime.now()
        return end - self.entry_time


# --- Strategy: Pricing ---

class HourlyPricing:
    """Charges by the hour, rounded up, per vehicle type."""

    RATES = {
        VehicleType.MOTORCYCLE: 1.0,
        VehicleType.CAR: 2.0,
        VehicleType.TRUCK: 3.0,
    }

    def calculate(self, vehicle_type: VehicleType, duration: timedelta) -> float:
        hours = math.ceil(duration.total_seconds() / 3600)
        if hours < 1:
            hours = 1
        rate = self.RATES.get(vehicle_type, 2.0)
        return hours * rate


class FlatPricing:
    """Charges a flat rate regardless of duration."""

    def __init__(self, rate: float):
        self.rate = rate

    def calculate(self, vehicle_type, duration) -> float:
        return self.rate


# --- Strategy: Spot Allocation ---

VEHICLE_MIN_SPOT = {
    VehicleType.MOTORCYCLE: SpotSize.SMALL,
    VehicleType.CAR: SpotSize.MEDIUM,
    VehicleType.TRUCK: SpotSize.LARGE,
}


class NearestAvailable:
    """Picks the first available spot, preferring exact-fit."""

    def allocate(self, spots: list, vehicle_type: VehicleType):
        target = VEHICLE_MIN_SPOT.get(vehicle_type, SpotSize.MEDIUM)

        # First pass: exact fit
        for spot in spots:
            if not spot.occupied and spot.size == target:
                return spot

        # Second pass: any fit
        for spot in spots:
            if not spot.occupied and spot.can_fit(vehicle_type):
                return spot

        return None


# --- Repository ---

class InMemoryRepo:
    def __init__(self, spots: list):
        self.spots = spots
        self.tickets = {}
        self._ticket_id = 0

    def get_spots(self):
        return self.spots

    def get_spot_by_id(self, spot_id: int):
        for s in self.spots:
            if s.id == spot_id:
                return s
        return None

    def save_ticket(self, ticket: Ticket):
        self.tickets[ticket.id] = ticket

    def get_ticket(self, ticket_id: int):
        return self.tickets.get(ticket_id)

    def get_active_ticket_by_plate(self, plate: str):
        for t in self.tickets.values():
            if t.active and t.vehicle.license_plate == plate:
                return t
        return None

    def next_ticket_id(self) -> int:
        self._ticket_id += 1
        return self._ticket_id


# --- Service ---

class ParkingService:
    def __init__(self, repo, allocator, pricing):
        self.repo = repo
        self.allocator = allocator
        self.pricing = pricing

    def park(self, vehicle: Vehicle) -> Ticket:
        spots = self.repo.get_spots()
        spot = self.allocator.allocate(spots, vehicle.type)
        if spot is None:
            raise RuntimeError("no available spot for this vehicle type")

        spot.park(vehicle)
        ticket_id = self.repo.next_ticket_id()
        ticket = Ticket(ticket_id, vehicle, spot.id)
        self.repo.save_ticket(ticket)
        return ticket

    def unpark(self, license_plate: str) -> Ticket:
        ticket = self.repo.get_active_ticket_by_plate(license_plate)
        if ticket is None:
            raise RuntimeError("vehicle is not currently parked")

        spot = self.repo.get_spot_by_id(ticket.spot_id)
        if spot:
            spot.free()

        duration = ticket.duration()
        fee = self.pricing.calculate(ticket.vehicle.type, duration)
        ticket.close(fee)
        self.repo.save_ticket(ticket)
        return ticket

    def get_occupancy(self):
        spots = self.repo.get_spots()
        result = {"small": [0, 0], "medium": [0, 0], "large": [0, 0]}
        for spot in spots:
            key = spot.size.name.lower()
            result[key][1] += 1
            if spot.occupied:
                result[key][0] += 1
        return result


# --- Helper: create spots ---

def create_spots(small=0, medium=0, large=0):
    spots = []
    spot_id = 1
    for _ in range(small):
        spots.append(Spot(spot_id, SpotSize.SMALL))
        spot_id += 1
    for _ in range(medium):
        spots.append(Spot(spot_id, SpotSize.MEDIUM))
        spot_id += 1
    for _ in range(large):
        spots.append(Spot(spot_id, SpotSize.LARGE))
        spot_id += 1
    return spots


# --- CLI Demo ---

def main():
    spots = create_spots(small=2, medium=3, large=1)
    store = InMemoryRepo(spots)
    pricing = HourlyPricing()
    allocator = NearestAvailable()
    svc = ParkingService(store, allocator, pricing)

    print("=== Parking Lot Demo ===")
    print("Lot created: 2 small, 3 medium, 1 large spots\n")

    vehicles = [
        Vehicle("MOTO-001", VehicleType.MOTORCYCLE),
        Vehicle("CAR-001", VehicleType.CAR),
        Vehicle("CAR-002", VehicleType.CAR),
        Vehicle("TRUCK-01", VehicleType.TRUCK),
        Vehicle("MOTO-002", VehicleType.MOTORCYCLE),
    ]

    print("--- Parking vehicles ---")
    for v in vehicles:
        try:
            ticket = svc.park(v)
            print(f"  OK:   {v.license_plate} ({v.type.name}) "
                  f"-> spot {ticket.spot_id}, ticket #{ticket.id}")
        except RuntimeError as e:
            print(f"  FAIL: {v.license_plate} ({v.type.name}) -- {e}")

    occ = svc.get_occupancy()
    print(f"\nOccupancy: Small: {occ['small'][0]}/{occ['small'][1]} | "
          f"Medium: {occ['medium'][0]}/{occ['medium'][1]} | "
          f"Large: {occ['large'][0]}/{occ['large'][1]}")

    print("\n--- Attempting to park another truck (should fail) ---")
    try:
        svc.park(Vehicle("TRUCK-02", VehicleType.TRUCK))
    except RuntimeError as e:
        print(f"  Expected: {e}")

    print("\n--- Unparking vehicles ---")
    for plate in ["CAR-001", "MOTO-001", "TRUCK-01"]:
        try:
            ticket = svc.unpark(plate)
            print(f"  OK:   {plate} -- fee=${ticket.fee:.2f} "
                  f"(parked {ticket.duration()})")
        except RuntimeError as e:
            print(f"  FAIL: {plate} -- {e}")

    occ = svc.get_occupancy()
    print(f"\nFinal occupancy: Small: {occ['small'][0]}/{occ['small'][1]} | "
          f"Medium: {occ['medium'][0]}/{occ['medium'][1]} | "
          f"Large: {occ['large'][0]}/{occ['large'][1]}")

    print("\ndemo done")


if __name__ == "__main__":
    main()
