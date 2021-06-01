package models

// BetReduced represents a domain model representation of a bet without CustomerId.
type BetReduced struct {
	Id                   string
	Status               string
	SelectionId          string
	SelectionCoefficient float64
	Payment              float64
	Payout               float64
}
