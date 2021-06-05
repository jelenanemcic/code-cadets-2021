package engine

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq/models"
)

type Publisher interface {
	PublishBetsCalculated(ctx context.Context, bets <-chan rabbitmqmodels.BetCalculated)
}
