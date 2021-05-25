package handler

import (
	"context"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
)

type BetRepository interface {
	InsertBet(ctx context.Context, bet domainmodels.Bet) error
	UpdateBet(ctx context.Context, bet domainmodels.Bet) error
	GetBetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error)
	GetBetBySelectionID(ctx context.Context, selectionId string) ([]domainmodels.Bet, error)
}
