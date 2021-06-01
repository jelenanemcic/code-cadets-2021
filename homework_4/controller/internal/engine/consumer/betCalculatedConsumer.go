package consumer

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_4/controller/internal/infrastructure/rabbitmq/models"
)

type BetCalculatedConsumer interface {
	Consume(ctx context.Context) (<-chan rabbitmqmodels.BetCalculated, error)
}
