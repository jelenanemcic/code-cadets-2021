package engine

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_3/controller/internal/infrastructure/rabbitmq/models"
)

type Handler interface {
	HandleBetsReceived(ctx context.Context, betsReceived <-chan rabbitmqmodels.BetReceived) <-chan rabbitmqmodels.Bet
	HandleBetsCalculated(ctx context.Context, betsCalculated <-chan rabbitmqmodels.BetCalculated) <-chan rabbitmqmodels.Bet
}
