package repo

import (
	"errors"
	"sync"

	"go-backend-interview-prep/12_capstone_projects/02_parking_lot_system/go/internal/domain"
)

// Errors
var (
	ErrTicketNotFound = errors.New("ticket not found")
	ErrSpotNotFound   = errors.New("spot not found")
)

// ParkingRepo abstracts persistence for the parking lot.
type ParkingRepo interface {
	GetSpots() []*domain.Spot
	GetSpotByID(id int) (*domain.Spot, error)
	SaveTicket(ticket domain.Ticket)
	GetTicket(id int) (domain.Ticket, error)
	GetActiveTicketByPlate(plate string) (domain.Ticket, error)
	NextTicketID() int
}

// InMemoryRepo stores spots and tickets in memory with mutex protection.
type InMemoryRepo struct {
	mu       sync.Mutex
	spots    []*domain.Spot
	tickets  map[int]domain.Ticket
	ticketID int
}

// NewInMemoryRepo creates a repo with the given spots.
func NewInMemoryRepo(spots []*domain.Spot) *InMemoryRepo {
	return &InMemoryRepo{
		spots:   spots,
		tickets: make(map[int]domain.Ticket),
	}
}

func (r *InMemoryRepo) GetSpots() []*domain.Spot {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.spots
}

func (r *InMemoryRepo) GetSpotByID(id int) (*domain.Spot, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, s := range r.spots {
		if s.ID == id {
			return s, nil
		}
	}
	return nil, ErrSpotNotFound
}

func (r *InMemoryRepo) SaveTicket(ticket domain.Ticket) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tickets[ticket.ID] = ticket
}

func (r *InMemoryRepo) GetTicket(id int) (domain.Ticket, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.tickets[id]
	if !ok {
		return domain.Ticket{}, ErrTicketNotFound
	}
	return t, nil
}

func (r *InMemoryRepo) GetActiveTicketByPlate(plate string) (domain.Ticket, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, t := range r.tickets {
		if t.Active && t.Vehicle.LicensePlate == plate {
			return t, nil
		}
	}
	return domain.Ticket{}, ErrTicketNotFound
}

func (r *InMemoryRepo) NextTicketID() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ticketID++
	return r.ticketID
}
