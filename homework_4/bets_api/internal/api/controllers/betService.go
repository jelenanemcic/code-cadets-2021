package controllers

import (
	"context"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/domain/models"
)

// BetService implements bet related functions.
type BetService interface {
	GetBetById(ctx context.Context, id string) (models.BetReduced, bool, error)
	GetBetsByUser(ctx context.Context, userId string) ([]models.BetReduced, error)
	GetBetsByStatus(ctx context.Context, status string) ([]models.BetReduced, error)
}
