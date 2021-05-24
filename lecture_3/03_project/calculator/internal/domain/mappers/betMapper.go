package mappers

import (
	"math"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

// BetMapper maps storage bets to domain bets and vice versa.
type BetMapper struct {
}

// NewBetMapper creates and returns a new BetMapper.
func NewBetMapper() *BetMapper {
	return &BetMapper{}
}

// MapDomainBetToStorageBet maps the given domain bet into storage bet. Floating point values will
// be converted to corresponding integer values of the storage bet by multiplying them with 100.
func (m *BetMapper) MapDomainBetToStorageBet(domainBet domainmodels.BetCalculated) storagemodels.BetCalculated {
	return storagemodels.BetCalculated{
		Id:                   domainBet.Id,
		SelectionId:          domainBet.SelectionId,
		SelectionCoefficient: int(math.Round(domainBet.SelectionCoefficient * 100)),
		Payment:              int(math.Round(domainBet.Payment * 100)),
	}
}

// MapStorageBetToDomainBet maps the given storage bet into domain bet. Floating point values will
// be converted from corresponding integer values of the storage bet by dividing them with 100.
func (m *BetMapper) MapStorageBetToDomainBet(storageBet storagemodels.BetCalculated) domainmodels.BetCalculated {
	return domainmodels.BetCalculated{
		Id:                   storageBet.Id,
		SelectionId:          storageBet.SelectionId,
		SelectionCoefficient: float64(storageBet.SelectionCoefficient) / 100,
		Payment:              float64(storageBet.Payment) / 100,
	}
}
