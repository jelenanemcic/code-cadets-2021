package handler

import (
	"context"
	"log"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Handler handles bets received and bets calculated.
type Handler struct {
	betRepository BetRepository
}

// New creates and returns a new Handler.
func New(betRepository BetRepository) *Handler {
	return &Handler{
		betRepository: betRepository,
	}
}

// HandleBets handles bets received.
func (h *Handler) HandleBets(
	ctx context.Context,
	bets <-chan rabbitmqmodels.Bet,
) <-chan rabbitmqmodels.BetCalculated {
	resultingBets := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(resultingBets)

		for bet := range bets {
			log.Println("Processing bet, betId:", bet.Id)

			// Calculate the calculated bet based on the incoming bet.
			domainBet := domainmodels.BetCalculated{
				Id:                   bet.Id,
				SelectionId:          bet.SelectionId,
				SelectionCoefficient: bet.SelectionCoefficient,
				Payment:              bet.Payment * bet.SelectionCoefficient,
			}

			// Insert the calculated bet into the repository.
			err := h.betRepository.InsertBetCalculated(ctx, domainBet)
			if err != nil {
				log.Println("Failed to insert bet, error: ", err)
				continue
			}
		}
	}()

	return resultingBets
}

// HandleEventUpdates handles event updates.
func (h *Handler) HandleEventUpdates(
	ctx context.Context,
	eventUpdates <-chan rabbitmqmodels.EventUpdate,
) <-chan rabbitmqmodels.BetCalculated {
	resultingBets := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(resultingBets)

		for eventUpdate := range eventUpdates {
			log.Println("Processing event update, selectionId:", eventUpdate.Id)

			// Fetch all domain bets with this selectionID
			domainBets, exists, err := h.betRepository.GetBetCalculatedBySelectionID(ctx, eventUpdate.Id)
			if err != nil {
				log.Println("Failed to fetch calculated bets, error: ", err)
				continue
			}
			if !exists {
				log.Println("There are no calculated bets with selectionId: ", eventUpdate.Id)
				continue
			}

			for _, domainBet := range domainBets {
				log.Printf("%v\n", domainBet)
				// Calculate the resulting calculated bet, which should be published.
				resultingBet := rabbitmqmodels.BetCalculated{
					Id:     domainBet.Id,
					Status: eventUpdate.Outcome,
					Payout: domainBet.Payment,
				}

				select {
				case resultingBets <- resultingBet:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return resultingBets
}
