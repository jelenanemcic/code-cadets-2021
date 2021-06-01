package consumer

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_4/calculator/internal/infrastructure/rabbitmq/models"
)

type EventUpdateConsumer interface {
	Consume(ctx context.Context) (<-chan rabbitmqmodels.EventUpdate, error)
}
