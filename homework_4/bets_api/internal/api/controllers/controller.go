package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betService BetService
}

// NewController creates a new instance of Controller
func NewController(betService BetService) *Controller {
	return &Controller{
		betService: betService,
	}
}

// BetByIdHandler handles betById request.
func (e *Controller) BetByIdHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bet, exists, err := e.betService.GetBetById(ctx)

		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		if !exists {
			ctx.String(http.StatusNotFound, "bet with this ID does not exist.")
			return
		}

		ctx.JSON(http.StatusOK, bet)
	}
}

// BetsByUserHandler handles betsByUser request.
func (e *Controller) BetsByUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bets, err := e.betService.GetBetsByUser(ctx)

		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		if len(bets) == 0 {
			ctx.String(http.StatusNotFound, "there are no bets with this userID.")
			return
		}

		ctx.JSON(http.StatusOK, bets)
	}
}

// BetsByStatusHandler handles betsByStatus request.
func (e *Controller) BetsByStatusHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bets, err := e.betService.GetBetsByStatus(ctx)

		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		if len(bets) == 0 {
			ctx.String(http.StatusNotFound, "there are no bets with status.")
			return
		}

		ctx.JSON(http.StatusOK, bets)
	}
}
