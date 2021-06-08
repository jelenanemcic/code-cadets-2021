package models

// BetReduced represents a domain model representation of a bet without CustomerId.
type BetReduced struct {
	Id                   string  `json:"Id"`
	Status               string  `json:"Status"`
	SelectionId          string  `json:"SelectionId"`
	SelectionCoefficient float64 `json:"SelectionCoefficient"`
	Payment              float64 `json:"Payment"`
	Payout               float64 `json:"Payout"`
}
