package services

// BetPublisher handles bet queue publishing.
type BetPublisher interface {
	Publish(id string, customerId string, selectionId string, selectionCoefficient float64, payment float64) error
}
