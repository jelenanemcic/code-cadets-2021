package sqlite

import (
	domainmodels "github.com/jelenanemcic/code-cadets-2021/homework_4/calculator/internal/domain/models"
	storagemodels "github.com/jelenanemcic/code-cadets-2021/homework_4/calculator/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapDomainBetToStorageBet(domainBet domainmodels.Bet) storagemodels.Bet
	MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.Bet
}
