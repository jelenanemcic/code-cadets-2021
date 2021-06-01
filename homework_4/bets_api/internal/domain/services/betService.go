package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/domain/models"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite"
	"github.com/pkg/errors"
)

// BetService implements bet related functions.
type BetService struct {
	betRepository sqlite.BetRepository
}

// NewBetService creates a new instance of BetService.
func NewBetService(repository sqlite.BetRepository) *BetService {
	return &BetService{
		betRepository: repository,
	}
}

//GetBetById returns a bet with the specific id.
func (e BetService) GetBetById(ctx *gin.Context) (models.BetReduced, bool, error) {
	id := ctx.Param("id")
	return e.betRepository.GetBetByID(ctx, id)
}

//GetBetsByUser returns all bets with the specific user id.
func (e BetService) GetBetsByUser(ctx *gin.Context) ([]models.BetReduced, error) {
	userId := ctx.Param("id")
	return e.betRepository.GetBetsByUserID(ctx, userId)
}

//GetBetsByStatus returns all bets with the specific status.
func (e BetService) GetBetsByStatus(ctx *gin.Context) ([]models.BetReduced, error) {
	status := ctx.Query("status")
	if status != "won" && status != "lost" && status != "active" {
		return []models.BetReduced{}, errors.New("invalid status")
	}
	return e.betRepository.GetBetsByStatus(ctx, status)
}
