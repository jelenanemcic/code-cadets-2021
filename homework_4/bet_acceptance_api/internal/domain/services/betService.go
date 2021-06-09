package services

import (
	uuid "github.com/nu7hatch/gouuid"
	"log"
)

// BetService implements bet related functions.
type BetService struct {
	betPublisher BetPublisher
}

// NewBetService creates a new instance of BetService.
func NewBetService(betPublisher BetPublisher) *BetService {
	return &BetService{
		betPublisher: betPublisher,
	}
}

// ReceiveBet gives the bet an id and sends bet message to the queue.
func (e BetService) ReceiveBet(customerId string, selectionId string, selectionCoefficient float64, payment float64) error {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("%s: %s", "failed to generate a uuid", err)
	}

	return e.betPublisher.Publish(id.String(), customerId, selectionId, selectionCoefficient, payment)
}
