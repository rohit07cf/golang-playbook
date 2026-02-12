package domain

// VehicleType represents the category of a vehicle.
type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Truck
)

func (v VehicleType) String() string {
	switch v {
	case Motorcycle:
		return "Motorcycle"
	case Car:
		return "Car"
	case Truck:
		return "Truck"
	default:
		return "Unknown"
	}
}

// Vehicle represents a vehicle entering the lot.
type Vehicle struct {
	LicensePlate string
	Type         VehicleType
}

// NewVehicle creates a Vehicle (factory function).
func NewVehicle(plate string, vType VehicleType) Vehicle {
	return Vehicle{
		LicensePlate: plate,
		Type:         vType,
	}
}
