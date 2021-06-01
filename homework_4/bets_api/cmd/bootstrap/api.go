package bootstrap

import (
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/cmd/config"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/api"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/api/controllers"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/domain/mappers"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/domain/services"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite"
)

func newBetService(repository sqlite.BetRepository) *services.BetService {
	return services.NewBetService(repository)
}

func newController(betService controllers.BetService) *controllers.Controller {
	return controllers.NewController(betService)
}

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, betMapper)
}

// Api bootstraps the http server.
func Api(dbExecutor sqlite.DatabaseExecutor) *api.WebServer {
	betMapper := newBetMapper()
	betRepository := newBetRepository(dbExecutor, betMapper)

	betService := newBetService(*betRepository)
	controller := newController(betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
