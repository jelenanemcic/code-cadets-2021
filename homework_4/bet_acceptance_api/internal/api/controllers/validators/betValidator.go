package validators

import "github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers/models"

const maxCoefficient = 10.0
const minPayment = 2.0
const maxPayment = 100.0

// BetValidator validates bet requests.
type BetValidator struct{}

// NewBetValidator creates a new instance of BetValidator.
func NewBetValidator() *BetValidator {
	return &BetValidator{}
}

// BetIsValid checks if bet is valid.
// CustomerId, SelectionId are not empty
// SelectionCoefficient is <= 10.0
// Payment is >=2.0 && <= 100.0
func (e *BetValidator) BetIsValid(betRequestDto models.BetRequestDto) bool {
	if betRequestDto.CustomerId != "" && betRequestDto.SelectionId != "" &&
		betRequestDto.SelectionCoefficient <= maxCoefficient &&
		betRequestDto.Payment >= minPayment && betRequestDto.Payment <= maxPayment {
		return true
	}

	return false
}
