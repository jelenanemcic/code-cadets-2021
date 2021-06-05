package engine

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_3/controller/internal/infrastructure/rabbitmq/models"
)

type Consumer interface {
	ConsumeBetsReceived(ctx context.Context) (<-chan rabbitmqmodels.BetReceived, error)
	ConsumeBetsCalculated(ctx context.Context) (<-chan rabbitmqmodels.BetCalculated, error)
}
