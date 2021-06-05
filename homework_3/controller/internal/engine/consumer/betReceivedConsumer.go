package consumer

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_3/controller/internal/infrastructure/rabbitmq/models"
)

type BetReceivedConsumer interface {
	Consume(ctx context.Context) (<-chan rabbitmqmodels.BetReceived, error)
}
