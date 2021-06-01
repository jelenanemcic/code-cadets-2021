package publisher

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_4/controller/internal/infrastructure/rabbitmq/models"
)

type BetPublisher interface {
	Publish(ctx context.Context, bets <-chan rabbitmqmodels.Bet)
}
