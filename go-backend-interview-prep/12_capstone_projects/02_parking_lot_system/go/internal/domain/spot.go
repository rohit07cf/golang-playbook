package domain

// SpotSize represents the physical size of a parking spot.
type SpotSize int

const (
	Small  SpotSize = iota // fits Motorcycle
	Medium                 // fits Car
	Large                  // fits Truck (also fits smaller vehicles)
)

func (s SpotSize) String() string {
	switch s {
	case Small:
		return "Small"
	case Medium:
		return "Medium"
	case Large:
		return "Large"
	default:
		return "Unknown"
	}
}

// Spot represents a single parking spot.
type Spot struct {
	ID       int
	Size     SpotSize
	Occupied bool
	Vehicle  *Vehicle // nil if empty
}

// CanFit checks if a vehicle type can fit in this spot size.
func (s *Spot) CanFit(vType VehicleType) bool {
	switch vType {
	case Motorcycle:
		return true // fits any spot
	case Car:
		return s.Size >= Medium
	case Truck:
		return s.Size >= Large
	default:
		return false
	}
}

// Park places a vehicle in this spot.
func (s *Spot) Park(v *Vehicle) {
	s.Occupied = true
	s.Vehicle = v
}

// Free removes the vehicle from this spot.
func (s *Spot) Free() {
	s.Occupied = false
	s.Vehicle = nil
}
