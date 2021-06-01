package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/domain/models"
)

// BetService implements bet related functions.
type BetService interface {
	GetBetById(ctx *gin.Context) (models.Bet, bool, error)
	GetBetsByUser(ctx *gin.Context) ([]models.Bet, error)
	GetBetsByStatus(ctx *gin.Context) ([]models.Bet, error)
}
