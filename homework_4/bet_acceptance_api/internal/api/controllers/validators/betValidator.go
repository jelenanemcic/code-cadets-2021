package validators

import (
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers/models"
)

// BetValidator validates bet requests.
type BetValidator struct {
	MaxCoefficient float64
	MinPayment     float64
	MaxPayment     float64
}

// NewBetValidator creates a new instance of BetValidator.
func NewBetValidator(maxCoefficient, minPayment, maxPayment float64) *BetValidator {
	return &BetValidator{
		MaxCoefficient: maxCoefficient,
		MinPayment:     minPayment,
		MaxPayment:     maxPayment,
	}
}

// BetIsValid checks if bet is valid.
// CustomerId, SelectionId are not empty
// SelectionCoefficient is <= 10.0
// Payment is >=2.0 && <= 100.0
func (e *BetValidator) BetIsValid(betRequestDto models.BetRequestDto) bool {
	if betRequestDto.CustomerId != "" && betRequestDto.SelectionId != "" &&
		betRequestDto.SelectionCoefficient <= e.MaxCoefficient &&
		betRequestDto.Payment >= e.MinPayment &&
		betRequestDto.Payment <= e.MaxPayment {
		return true
	}

	return false
}
