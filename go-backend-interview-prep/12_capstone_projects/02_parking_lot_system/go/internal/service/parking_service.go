package service

import (
	"errors"
	"fmt"

	"go-backend-interview-prep/12_capstone_projects/02_parking_lot_system/go/internal/domain"
	"go-backend-interview-prep/12_capstone_projects/02_parking_lot_system/go/internal/repo"
)

// Errors
var (
	ErrLotFull       = errors.New("no available spot for this vehicle type")
	ErrNotParked     = errors.New("vehicle is not currently parked")
)

// Occupancy holds a summary of lot usage.
type Occupancy struct {
	SmallUsed, SmallTotal   int
	MediumUsed, MediumTotal int
	LargeUsed, LargeTotal   int
}

func (o Occupancy) String() string {
	return fmt.Sprintf("Small: %d/%d | Medium: %d/%d | Large: %d/%d",
		o.SmallUsed, o.SmallTotal, o.MediumUsed, o.MediumTotal, o.LargeUsed, o.LargeTotal)
}

// ParkingService orchestrates parking operations.
// Dependencies are injected via constructor (DI).
type ParkingService struct {
	repo       repo.ParkingRepo
	allocator  domain.SpotAllocationStrategy
	pricing    domain.PricingStrategy
}

// NewParkingService creates a service with injected dependencies.
func NewParkingService(
	r repo.ParkingRepo,
	alloc domain.SpotAllocationStrategy,
	pricing domain.PricingStrategy,
) *ParkingService {
	return &ParkingService{
		repo:      r,
		allocator: alloc,
		pricing:   pricing,
	}
}

// Park assigns a spot and issues a ticket.
func (s *ParkingService) Park(vehicle domain.Vehicle) (domain.Ticket, error) {
	spots := s.repo.GetSpots()
	spot := s.allocator.Allocate(spots, vehicle.Type)
	if spot == nil {
		return domain.Ticket{}, ErrLotFull
	}

	spot.Park(&vehicle)
	ticketID := s.repo.NextTicketID()
	ticket := domain.NewTicket(ticketID, vehicle, spot.ID)
	s.repo.SaveTicket(ticket)

	return ticket, nil
}

// Unpark frees the spot, calculates fee, and closes the ticket.
func (s *ParkingService) Unpark(licensePlate string) (domain.Ticket, error) {
	ticket, err := s.repo.GetActiveTicketByPlate(licensePlate)
	if err != nil {
		return domain.Ticket{}, ErrNotParked
	}

	// Free the spot
	spot, err := s.repo.GetSpotByID(ticket.SpotID)
	if err != nil {
		return domain.Ticket{}, err
	}
	spot.Free()

	// Calculate fee
	duration := ticket.Duration()
	fee := s.pricing.Calculate(ticket.Vehicle.Type, duration)
	ticket.Close(fee)

	// Save updated ticket
	s.repo.SaveTicket(ticket)

	return ticket, nil
}

// GetOccupancy returns a summary of lot usage by spot size.
func (s *ParkingService) GetOccupancy() Occupancy {
	spots := s.repo.GetSpots()
	var occ Occupancy

	for _, spot := range spots {
		switch spot.Size {
		case domain.Small:
			occ.SmallTotal++
			if spot.Occupied {
				occ.SmallUsed++
			}
		case domain.Medium:
			occ.MediumTotal++
			if spot.Occupied {
				occ.MediumUsed++
			}
		case domain.Large:
			occ.LargeTotal++
			if spot.Occupied {
				occ.LargeUsed++
			}
		}
	}
	return occ
}
