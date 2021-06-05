package handler

import (
	"context"

	domainmodels "github.com/jelenanemcic/code-cadets-2021/homework_3/controller/internal/domain/models"
)

type BetRepository interface {
	InsertBet(ctx context.Context, bet domainmodels.Bet) error
	UpdateBet(ctx context.Context, bet domainmodels.Bet) error
	GetBetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error)
}
