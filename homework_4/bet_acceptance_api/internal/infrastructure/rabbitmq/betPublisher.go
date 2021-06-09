package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/infrastructure/rabbitmq/models"
	"github.com/streadway/amqp"
)

const contentTypeTextPlain = "application/json"

// BetPublisher handles bet queue publishing.
type BetPublisher struct {
	exchange  string
	queueName string
	mandatory bool
	immediate bool
	publisher QueuePublisher
}

// NewBetPublisher create a new instance of BetPublisher.
func NewBetPublisher(
	exchange string,
	queueName string,
	mandatory bool,
	immediate bool,
	publisher QueuePublisher,
) *BetPublisher {
	return &BetPublisher{
		exchange:  exchange,
		queueName: queueName,
		mandatory: mandatory,
		immediate: immediate,
		publisher: publisher,
	}
}

// Publish publishes a bet to the queue.
func (p *BetPublisher) Publish(id string, customerId string, selectionId string, selectionCoefficient float64, payment float64) error {
	bet := &models.BetDto{
		Id:                   id,
		CustomerId:           customerId,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}

	betJson, err := json.Marshal(bet)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		p.exchange,
		p.queueName,
		p.mandatory,
		p.immediate,
		amqp.Publishing{
			ContentType: contentTypeTextPlain,
			Body:        betJson,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent %s", betJson)
	return nil
}
