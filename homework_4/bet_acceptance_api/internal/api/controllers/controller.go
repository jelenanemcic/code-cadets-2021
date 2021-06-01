package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers/models"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betValidator BetValidator
	betService   BetService
}

// NewController creates a new instance of Controller
func NewController(betValidator BetValidator, betService BetService) *Controller {
	return &Controller{
		betValidator: betValidator,
		betService:   betService,
	}
}

// ProcessBet handlers bet quest.
func (e *Controller) ProcessBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var betRequestDto models.BetRequestDto
		err := ctx.ShouldBindWith(&betRequestDto, binding.JSON)
		if err != nil {
			ctx.String(http.StatusBadRequest, "bet request is not valid.")
			return
		}

		if !e.betValidator.BetIsValid(betRequestDto) {
			ctx.String(http.StatusBadRequest, "bet is not valid.")
			return
		}

		err = e.betService.ReceiveBet(betRequestDto.CustomerId, betRequestDto.SelectionId, betRequestDto.SelectionCoefficient, betRequestDto.Payment)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.Status(http.StatusOK)
	}
}
