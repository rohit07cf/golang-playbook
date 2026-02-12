package main

import (
	"fmt"

	"go-backend-interview-prep/12_capstone_projects/02_parking_lot_system/go/internal/domain"
	"go-backend-interview-prep/12_capstone_projects/02_parking_lot_system/go/internal/repo"
	"go-backend-interview-prep/12_capstone_projects/02_parking_lot_system/go/internal/service"
)

func main() {
	// --- Setup: create lot with spots ---
	config := domain.LotConfig{
		SmallSpots:  2,
		MediumSpots: 3,
		LargeSpots:  1,
	}
	spots := domain.CreateSpots(config)
	store := repo.NewInMemoryRepo(spots)
	pricing := domain.NewHourlyPricing()
	allocator := &domain.NearestAvailable{}

	svc := service.NewParkingService(store, allocator, pricing)

	fmt.Println("=== Parking Lot Demo ===")
	fmt.Printf("Lot created: %d small, %d medium, %d large spots\n\n",
		config.SmallSpots, config.MediumSpots, config.LargeSpots)

	// --- Park vehicles ---
	vehicles := []domain.Vehicle{
		domain.NewVehicle("MOTO-001", domain.Motorcycle),
		domain.NewVehicle("CAR-001", domain.Car),
		domain.NewVehicle("CAR-002", domain.Car),
		domain.NewVehicle("TRUCK-01", domain.Truck),
		domain.NewVehicle("MOTO-002", domain.Motorcycle),
	}

	fmt.Println("--- Parking vehicles ---")
	for _, v := range vehicles {
		ticket, err := svc.Park(v)
		if err != nil {
			fmt.Printf("  FAIL: %s (%s) -- %v\n", v.LicensePlate, v.Type, err)
			continue
		}
		fmt.Printf("  OK:   %s (%s) -> spot %d, ticket #%d\n",
			v.LicensePlate, v.Type, ticket.SpotID, ticket.ID)
	}

	// --- Show occupancy ---
	fmt.Printf("\nOccupancy: %s\n", svc.GetOccupancy())

	// --- Try parking when lot is full for a type ---
	fmt.Println("\n--- Attempting to park another truck (should fail) ---")
	extraTruck := domain.NewVehicle("TRUCK-02", domain.Truck)
	_, err := svc.Park(extraTruck)
	if err != nil {
		fmt.Printf("  Expected: %v\n", err)
	}

	// --- Unpark vehicles and show fees ---
	fmt.Println("\n--- Unparking vehicles ---")
	unpark := []string{"CAR-001", "MOTO-001", "TRUCK-01"}
	for _, plate := range unpark {
		ticket, err := svc.Unpark(plate)
		if err != nil {
			fmt.Printf("  FAIL: %s -- %v\n", plate, err)
			continue
		}
		fmt.Printf("  OK:   %s -- fee=$%.2f (parked %v)\n",
			plate, ticket.Fee, ticket.Duration().Round(1000000))
	}

	// --- Final occupancy ---
	fmt.Printf("\nFinal occupancy: %s\n", svc.GetOccupancy())

	fmt.Println("\ndemo done")
}
