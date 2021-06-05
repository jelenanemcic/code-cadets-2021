package engine

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_3/controller/internal/infrastructure/rabbitmq/models"
)

type Publisher interface {
	PublishBets(ctx context.Context, bets <-chan rabbitmqmodels.Bet)
}
