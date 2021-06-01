package publisher

import (
	"context"

	rabbitmqmodels "github.com/jelenanemcic/code-cadets-2021/homework_4/controller/internal/infrastructure/rabbitmq/models"
)

// Publisher offers methods for publishing into output queues.
type Publisher struct {
	betPublisher BetPublisher
}

// New creates and returns a new Publisher.
func New(betPublisher BetPublisher) *Publisher {
	return &Publisher{
		betPublisher: betPublisher,
	}
}

// PublishBets publishes into bets queue.
func (p *Publisher) PublishBets(ctx context.Context, bets <-chan rabbitmqmodels.Bet) {
	p.betPublisher.Publish(ctx, bets)
}
