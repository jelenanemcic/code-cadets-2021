package consumer

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq/models"
)

type BetConsumer interface {
	Consume(ctx context.Context) (<-chan rabbitmqmodels.Bet, error)
}
