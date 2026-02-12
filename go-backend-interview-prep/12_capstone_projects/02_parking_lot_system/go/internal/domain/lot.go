package domain

// SpotAllocationStrategy decides which spot to assign to a vehicle.
type SpotAllocationStrategy interface {
	Allocate(spots []*Spot, vehicleType VehicleType) *Spot
}

// --- Nearest Available Strategy ---

// NearestAvailable picks the first available spot that fits the vehicle.
// It prefers the smallest fitting spot to conserve larger spots.
type NearestAvailable struct{}

// Allocate returns the first spot that can fit the vehicle type,
// preferring exact-fit spots (e.g., Small for Motorcycle).
func (s *NearestAvailable) Allocate(spots []*Spot, vehicleType VehicleType) *Spot {
	// First pass: look for exact-fit spot
	targetSize := vehicleTypeToMinSpotSize(vehicleType)
	for _, spot := range spots {
		if !spot.Occupied && spot.Size == targetSize {
			return spot
		}
	}
	// Second pass: any spot that fits
	for _, spot := range spots {
		if !spot.Occupied && spot.CanFit(vehicleType) {
			return spot
		}
	}
	return nil
}

// vehicleTypeToMinSpotSize returns the minimum spot size for a vehicle type.
func vehicleTypeToMinSpotSize(vType VehicleType) SpotSize {
	switch vType {
	case Motorcycle:
		return Small
	case Car:
		return Medium
	case Truck:
		return Large
	default:
		return Medium
	}
}

// LotConfig defines the capacity of a parking lot.
type LotConfig struct {
	SmallSpots  int
	MediumSpots int
	LargeSpots  int
}

// CreateSpots generates a slice of spots based on config.
func CreateSpots(config LotConfig) []*Spot {
	spots := make([]*Spot, 0, config.SmallSpots+config.MediumSpots+config.LargeSpots)
	id := 1

	for i := 0; i < config.SmallSpots; i++ {
		spots = append(spots, &Spot{ID: id, Size: Small})
		id++
	}
	for i := 0; i < config.MediumSpots; i++ {
		spots = append(spots, &Spot{ID: id, Size: Medium})
		id++
	}
	for i := 0; i < config.LargeSpots; i++ {
		spots = append(spots, &Spot{ID: id, Size: Large})
		id++
	}
	return spots
}
