package bootstrap

import (
	"github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/cmd/config"
	"github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/internal/domain/mappers"
	"github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/internal/engine"
	"github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/internal/engine/consumer"
	"github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/internal/engine/handler"
	"github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/internal/engine/publisher"
	"github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/internal/infrastructure/rabbitmq"
	"github.com/jelenanemcic/code-cadets-2021/homework_3/calculator/internal/infrastructure/sqlite"
)

func newBetConsumer(channel rabbitmq.Channel) *rabbitmq.BetConsumer {
	betReceivedConsumer, err := rabbitmq.NewBetConsumer(
		channel,
		rabbitmq.ConsumerConfig{
			Queue:             config.Cfg.Rabbit.ConsumerBetQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			ConsumerName:      config.Cfg.Rabbit.ConsumerBetName,
			AutoAck:           config.Cfg.Rabbit.ConsumerAutoAck,
			Exclusive:         config.Cfg.Rabbit.ConsumerExclusive,
			NoLocal:           config.Cfg.Rabbit.ConsumerNoLocal,
			NoWait:            config.Cfg.Rabbit.ConsumerNoWait,
			Args:              nil,
		},
	)
	if err != nil {
		panic(err)
	}
	return betReceivedConsumer
}

func newEventUpdateConsumer(channel rabbitmq.Channel) *rabbitmq.EventUpdateConsumer {
	betCalculatedConsumer, err := rabbitmq.NewEventUpdateConsumer(
		channel,
		rabbitmq.ConsumerConfig{
			Queue:             config.Cfg.Rabbit.ConsumerEventUpdateQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			ConsumerName:      config.Cfg.Rabbit.ConsumerEventUpdateName,
			AutoAck:           config.Cfg.Rabbit.ConsumerAutoAck,
			Exclusive:         config.Cfg.Rabbit.ConsumerExclusive,
			NoLocal:           config.Cfg.Rabbit.ConsumerNoLocal,
			NoWait:            config.Cfg.Rabbit.ConsumerNoWait,
			Args:              nil,
		},
	)
	if err != nil {
		panic(err)
	}
	return betCalculatedConsumer
}

func newConsumer(betConsumer consumer.BetConsumer, betCalculatedConsumer consumer.EventUpdateConsumer) *consumer.Consumer {
	return consumer.New(betConsumer, betCalculatedConsumer)
}

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, betMapper)
}

func newHandler(betRepository handler.BetRepository) *handler.Handler {
	return handler.New(betRepository)
}

func newBetPublisher(channel rabbitmq.Channel) *rabbitmq.BetPublisher {
	betPublisher, err := rabbitmq.NewBetPublisher(
		channel,
		rabbitmq.PublisherConfig{
			Queue:             config.Cfg.Rabbit.PublisherBetCalculatedQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			Exchange:          config.Cfg.Rabbit.PublisherExchange,
			Mandatory:         config.Cfg.Rabbit.PublisherMandatory,
			Immediate:         config.Cfg.Rabbit.PublisherImmediate,
		},
	)
	if err != nil {
		panic(err)
	}
	return betPublisher
}

func newPublisher(betPublisher publisher.BetCalculatedPublisher) *publisher.Publisher {
	return publisher.New(betPublisher)
}

func Engine(rabbitMqChannel rabbitmq.Channel, dbExecutor sqlite.DatabaseExecutor) *engine.Engine {
	betReceivedConsumer := newBetConsumer(rabbitMqChannel)
	betCalculatedConsumer := newEventUpdateConsumer(rabbitMqChannel)
	consumer := newConsumer(betReceivedConsumer, betCalculatedConsumer)

	betMapper := newBetMapper()
	betRepository := newBetRepository(dbExecutor, betMapper)
	handler := newHandler(betRepository)

	betPublisher := newBetPublisher(rabbitMqChannel)
	publisher := newPublisher(betPublisher)

	return engine.New(consumer, handler, publisher)
}
