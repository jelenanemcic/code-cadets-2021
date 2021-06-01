package engine

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_4/calculator/internal/infrastructure/rabbitmq/models"
)

type Publisher interface {
	PublishBetsCalculated(ctx context.Context, betsCalculated <-chan rabbitmqmodels.BetCalculated)
}
