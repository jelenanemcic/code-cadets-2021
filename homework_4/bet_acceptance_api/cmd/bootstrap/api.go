package bootstrap

import (
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/cmd/config"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/api"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/api/controllers/validators"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/domain/services"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/internal/infrastructure/rabbitmq"
	"github.com/streadway/amqp"
)

func newBetValidator() *validators.BetValidator {
	return validators.NewBetValidator(config.Cfg.Validator.MaxCoefficient, config.Cfg.Validator.MinPayment, config.Cfg.Validator.MaxPayment)
}

func newBetPublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.BetPublisher {
	return rabbitmq.NewBetPublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherBetQueueQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newEventService(publisher services.BetPublisher) *services.BetService {
	return services.NewBetService(publisher)
}

func newController(betValidator controllers.BetValidator, betService controllers.BetService) *controllers.Controller {
	return controllers.NewController(betValidator, betService)
}

// Api bootstraps the http server.
func Api(rabbitMqChannel *amqp.Channel) *api.WebServer {
	betValidator := newBetValidator()
	betPublisher := newBetPublisher(rabbitMqChannel)
	betService := newEventService(betPublisher)
	controller := newController(betValidator, betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
