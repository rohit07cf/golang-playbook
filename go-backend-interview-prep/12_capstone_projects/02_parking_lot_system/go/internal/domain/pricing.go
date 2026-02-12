package domain

import (
	"math"
	"time"
)

// PricingStrategy calculates the fee for a parking session.
type PricingStrategy interface {
	Calculate(vehicleType VehicleType, duration time.Duration) float64
}

// --- Hourly Slab Pricing ---

// HourlyPricing charges by the hour with per-vehicle-type rates.
type HourlyPricing struct {
	Rates map[VehicleType]float64 // rate per hour
}

// NewHourlyPricing creates a pricing strategy with default rates.
func NewHourlyPricing() *HourlyPricing {
	return &HourlyPricing{
		Rates: map[VehicleType]float64{
			Motorcycle: 1.0, // $1/hour
			Car:        2.0, // $2/hour
			Truck:      3.0, // $3/hour
		},
	}
}

// Calculate returns the fee based on vehicle type and duration.
// Rounds up to the nearest hour (minimum 1 hour).
func (p *HourlyPricing) Calculate(vehicleType VehicleType, duration time.Duration) float64 {
	hours := math.Ceil(duration.Hours())
	if hours < 1 {
		hours = 1
	}
	rate, ok := p.Rates[vehicleType]
	if !ok {
		rate = 2.0 // default
	}
	return hours * rate
}

// --- Flat Rate Pricing (alternative strategy) ---

// FlatPricing charges a fixed amount regardless of duration.
type FlatPricing struct {
	Rate float64
}

// Calculate returns the flat rate.
func (p *FlatPricing) Calculate(_ VehicleType, _ time.Duration) float64 {
	return p.Rate
}
