"""Tests for Parking Lot System -- mirrors Go tests."""

import unittest
from datetime import timedelta

from parkinglot import (
    Vehicle, VehicleType, Spot, SpotSize,
    HourlyPricing, FlatPricing,
    NearestAvailable, InMemoryRepo, ParkingService,
    create_spots,
)


def new_test_service():
    """Create a standard service for tests."""
    spots = create_spots(small=2, medium=2, large=1)
    store = InMemoryRepo(spots)
    pricing = HourlyPricing()
    allocator = NearestAvailable()
    return ParkingService(store, allocator, pricing)


class TestParking(unittest.TestCase):

    def test_park_motorcycle(self):
        svc = new_test_service()
        v = Vehicle("MOTO-1", VehicleType.MOTORCYCLE)
        ticket = svc.park(v)
        self.assertEqual(ticket.vehicle.license_plate, "MOTO-1")
        self.assertTrue(ticket.active)

    def test_park_car(self):
        svc = new_test_service()
        v = Vehicle("CAR-1", VehicleType.CAR)
        ticket = svc.park(v)
        self.assertGreater(ticket.spot_id, 0)

    def test_park_truck(self):
        svc = new_test_service()
        v = Vehicle("TRUCK-1", VehicleType.TRUCK)
        ticket = svc.park(v)
        self.assertEqual(ticket.vehicle.type, VehicleType.TRUCK)


class TestLotFull(unittest.TestCase):

    def test_lot_full_returns_error(self):
        svc = new_test_service()
        svc.park(Vehicle("C1", VehicleType.CAR))
        svc.park(Vehicle("C2", VehicleType.CAR))
        svc.park(Vehicle("T1", VehicleType.TRUCK))

        with self.assertRaises(RuntimeError):
            svc.park(Vehicle("C3", VehicleType.CAR))

    def test_motorcycle_falls_back_to_larger(self):
        svc = new_test_service()
        svc.park(Vehicle("M1", VehicleType.MOTORCYCLE))
        svc.park(Vehicle("M2", VehicleType.MOTORCYCLE))

        # Third motorcycle goes to medium or large
        ticket = svc.park(Vehicle("M3", VehicleType.MOTORCYCLE))
        self.assertGreater(ticket.spot_id, 0)


class TestFeeCalculation(unittest.TestCase):

    def test_hourly_pricing_car(self):
        pricing = HourlyPricing()
        # 2.5 hours -> ceil = 3 * $2 = $6
        fee = pricing.calculate(VehicleType.CAR, timedelta(minutes=150))
        self.assertEqual(fee, 6.0)

    def test_hourly_pricing_motorcycle(self):
        pricing = HourlyPricing()
        # 30 min -> ceil = 1 * $1 = $1
        fee = pricing.calculate(VehicleType.MOTORCYCLE, timedelta(minutes=30))
        self.assertEqual(fee, 1.0)

    def test_hourly_pricing_truck(self):
        pricing = HourlyPricing()
        # 1 hour -> 1 * $3 = $3
        fee = pricing.calculate(VehicleType.TRUCK, timedelta(hours=1))
        self.assertEqual(fee, 3.0)

    def test_flat_pricing(self):
        pricing = FlatPricing(10.0)
        fee = pricing.calculate(VehicleType.CAR, timedelta(hours=3))
        self.assertEqual(fee, 10.0)


class TestUnpark(unittest.TestCase):

    def test_unpark_returns_ticket_with_fee(self):
        svc = new_test_service()
        svc.park(Vehicle("CAR-1", VehicleType.CAR))
        ticket = svc.unpark("CAR-1")
        self.assertFalse(ticket.active)
        self.assertGreater(ticket.fee, 0)

    def test_unpark_nonexistent_vehicle(self):
        svc = new_test_service()
        with self.assertRaises(RuntimeError):
            svc.unpark("GHOST-001")


class TestOccupancy(unittest.TestCase):

    def test_initial_occupancy(self):
        svc = new_test_service()
        occ = svc.get_occupancy()
        self.assertEqual(occ["small"][1], 2)
        self.assertEqual(occ["medium"][1], 2)
        self.assertEqual(occ["large"][1], 1)

    def test_occupancy_after_parking(self):
        svc = new_test_service()
        svc.park(Vehicle("CAR-1", VehicleType.CAR))
        occ = svc.get_occupancy()
        self.assertEqual(occ["medium"][0], 1)


class TestSpotFit(unittest.TestCase):

    def test_motorcycle_fits_everywhere(self):
        for size in [SpotSize.SMALL, SpotSize.MEDIUM, SpotSize.LARGE]:
            spot = Spot(1, size)
            self.assertTrue(spot.can_fit(VehicleType.MOTORCYCLE))

    def test_car_does_not_fit_small(self):
        spot = Spot(1, SpotSize.SMALL)
        self.assertFalse(spot.can_fit(VehicleType.CAR))

    def test_car_fits_medium_and_large(self):
        for size in [SpotSize.MEDIUM, SpotSize.LARGE]:
            spot = Spot(1, size)
            self.assertTrue(spot.can_fit(VehicleType.CAR))

    def test_truck_only_fits_large(self):
        self.assertFalse(Spot(1, SpotSize.SMALL).can_fit(VehicleType.TRUCK))
        self.assertFalse(Spot(1, SpotSize.MEDIUM).can_fit(VehicleType.TRUCK))
        self.assertTrue(Spot(1, SpotSize.LARGE).can_fit(VehicleType.TRUCK))


if __name__ == "__main__":
    unittest.main()
