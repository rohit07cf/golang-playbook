package domain

import "time"

// Ticket is issued when a vehicle enters the lot.
type Ticket struct {
	ID        int
	Vehicle   Vehicle
	SpotID    int
	EntryTime time.Time
	ExitTime  time.Time
	Fee       float64
	Active    bool
}

// NewTicket creates a new active ticket.
func NewTicket(id int, vehicle Vehicle, spotID int) Ticket {
	return Ticket{
		ID:        id,
		Vehicle:   vehicle,
		SpotID:    spotID,
		EntryTime: time.Now(),
		Active:    true,
	}
}

// Close marks the ticket as inactive and sets exit time and fee.
func (t *Ticket) Close(fee float64) {
	t.ExitTime = time.Now()
	t.Fee = fee
	t.Active = false
}

// Duration returns how long the vehicle has been parked.
func (t *Ticket) Duration() time.Duration {
	if t.Active {
		return time.Since(t.EntryTime)
	}
	return t.ExitTime.Sub(t.EntryTime)
}
