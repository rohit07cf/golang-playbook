package tests

import (
	"testing"
	"time"

	"go-backend-interview-prep/12_capstone_projects/02_parking_lot_system/go/internal/domain"
	"go-backend-interview-prep/12_capstone_projects/02_parking_lot_system/go/internal/repo"
	"go-backend-interview-prep/12_capstone_projects/02_parking_lot_system/go/internal/service"
)

// helper: create a standard parking service for tests.
func newTestService() *service.ParkingService {
	spots := domain.CreateSpots(domain.LotConfig{
		SmallSpots:  2,
		MediumSpots: 2,
		LargeSpots:  1,
	})
	store := repo.NewInMemoryRepo(spots)
	pricing := domain.NewHourlyPricing()
	allocator := &domain.NearestAvailable{}
	return service.NewParkingService(store, allocator, pricing)
}

// --- Allocation Tests ---

func TestParkMotorcycle(t *testing.T) {
	svc := newTestService()
	v := domain.NewVehicle("MOTO-1", domain.Motorcycle)

	ticket, err := svc.Park(v)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if ticket.Vehicle.LicensePlate != "MOTO-1" {
		t.Errorf("expected plate MOTO-1, got %s", ticket.Vehicle.LicensePlate)
	}
	if !ticket.Active {
		t.Error("ticket should be active")
	}
}

func TestParkCar(t *testing.T) {
	svc := newTestService()
	v := domain.NewVehicle("CAR-1", domain.Car)

	ticket, err := svc.Park(v)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if ticket.SpotID == 0 {
		t.Error("spot ID should be assigned")
	}
}

func TestParkTruck(t *testing.T) {
	svc := newTestService()
	v := domain.NewVehicle("TRUCK-1", domain.Truck)

	ticket, err := svc.Park(v)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if ticket.Vehicle.Type != domain.Truck {
		t.Errorf("expected truck, got %v", ticket.Vehicle.Type)
	}
}

// --- Full Lot Tests ---

func TestLotFullReturnsError(t *testing.T) {
	svc := newTestService()

	// Park 2 cars (fills both medium spots)
	svc.Park(domain.NewVehicle("CAR-1", domain.Car))
	svc.Park(domain.NewVehicle("CAR-2", domain.Car))

	// Park 1 truck (fills the large spot)
	svc.Park(domain.NewVehicle("TRUCK-1", domain.Truck))

	// Next car should fail (no medium or large spots left)
	_, err := svc.Park(domain.NewVehicle("CAR-3", domain.Car))
	if err == nil {
		t.Fatal("expected lot full error, got nil")
	}
}

func TestMotorcycleFallsBackToLargerSpot(t *testing.T) {
	svc := newTestService()

	// Fill both small spots
	svc.Park(domain.NewVehicle("M1", domain.Motorcycle))
	svc.Park(domain.NewVehicle("M2", domain.Motorcycle))

	// Third motorcycle should fall back to medium or large
	ticket, err := svc.Park(domain.NewVehicle("M3", domain.Motorcycle))
	if err != nil {
		t.Fatalf("motorcycle should fit in larger spot, got %v", err)
	}
	if ticket.SpotID == 0 {
		t.Error("expected valid spot ID")
	}
}

// --- Fee Calculation Tests ---

func TestFeeCalculationHourly(t *testing.T) {
	pricing := domain.NewHourlyPricing()

	// Car for 2.5 hours -> ceil(2.5) = 3 hours * $2 = $6
	fee := pricing.Calculate(domain.Car, 150*time.Minute)
	if fee != 6.0 {
		t.Errorf("expected $6.00, got $%.2f", fee)
	}

	// Motorcycle for 30 minutes -> ceil(0.5) = 1 hour * $1 = $1
	fee = pricing.Calculate(domain.Motorcycle, 30*time.Minute)
	if fee != 1.0 {
		t.Errorf("expected $1.00, got $%.2f", fee)
	}

	// Truck for 1 hour -> 1 * $3 = $3
	fee = pricing.Calculate(domain.Truck, 60*time.Minute)
	if fee != 3.0 {
		t.Errorf("expected $3.00, got $%.2f", fee)
	}
}

func TestFlatPricing(t *testing.T) {
	pricing := &domain.FlatPricing{Rate: 10.0}
	fee := pricing.Calculate(domain.Car, 3*time.Hour)
	if fee != 10.0 {
		t.Errorf("expected $10.00, got $%.2f", fee)
	}
}

// --- Unpark Tests ---

func TestUnparkReturnsTicketWithFee(t *testing.T) {
	svc := newTestService()
	v := domain.NewVehicle("CAR-1", domain.Car)

	svc.Park(v)
	ticket, err := svc.Unpark("CAR-1")
	if err != nil {
		t.Fatalf("unpark failed: %v", err)
	}
	if ticket.Active {
		t.Error("ticket should be inactive after unpark")
	}
	if ticket.Fee <= 0 {
		t.Error("fee should be positive")
	}
}

func TestUnparkNonexistentVehicle(t *testing.T) {
	svc := newTestService()
	_, err := svc.Unpark("GHOST-001")
	if err == nil {
		t.Fatal("expected error for non-parked vehicle")
	}
}

// --- Occupancy Tests ---

func TestOccupancy(t *testing.T) {
	svc := newTestService()

	occ := svc.GetOccupancy()
	if occ.SmallTotal != 2 || occ.MediumTotal != 2 || occ.LargeTotal != 1 {
		t.Errorf("unexpected totals: %v", occ)
	}

	svc.Park(domain.NewVehicle("CAR-1", domain.Car))
	occ = svc.GetOccupancy()
	if occ.MediumUsed != 1 {
		t.Errorf("expected 1 medium used, got %d", occ.MediumUsed)
	}
}

// --- Spot Fit Tests ---

func TestSpotCanFit(t *testing.T) {
	small := &domain.Spot{ID: 1, Size: domain.Small}
	medium := &domain.Spot{ID: 2, Size: domain.Medium}
	large := &domain.Spot{ID: 3, Size: domain.Large}

	// Motorcycle fits anywhere
	if !small.CanFit(domain.Motorcycle) {
		t.Error("motorcycle should fit in small")
	}
	if !large.CanFit(domain.Motorcycle) {
		t.Error("motorcycle should fit in large")
	}

	// Car does not fit in small
	if small.CanFit(domain.Car) {
		t.Error("car should not fit in small")
	}
	if !medium.CanFit(domain.Car) {
		t.Error("car should fit in medium")
	}

	// Truck only fits in large
	if medium.CanFit(domain.Truck) {
		t.Error("truck should not fit in medium")
	}
	if !large.CanFit(domain.Truck) {
		t.Error("truck should fit in large")
	}
}
