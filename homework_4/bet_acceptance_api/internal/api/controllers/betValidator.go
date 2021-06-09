package controllers

import "github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers/models"

// BetValidator validates bet requests.
type BetValidator interface {
	BetIsValid(betRequestDto models.BetRequestDto) bool
}
