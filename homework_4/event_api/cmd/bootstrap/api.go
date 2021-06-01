package bootstrap

import (
	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/cmd/config"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/internal/api"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/internal/api/controllers"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/internal/api/controllers/validators"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/internal/domain/services"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/internal/infrastructure/rabbitmq"
	"github.com/streadway/amqp"
)

func newEventUpdateValidator() *validators.EventUpdateValidator {
	return validators.NewEventUpdateValidator()
}

func newEventUpdatePublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.EventUpdatePublisher {
	return rabbitmq.NewEventUpdatePublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherEventUpdateQueueQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newEventService(publisher services.EventUpdatePublisher) *services.EventService {
	return services.NewEventService(publisher)
}

func newController(eventUpdateValidator controllers.EventUpdateValidator, eventService controllers.EventService) *controllers.Controller {
	return controllers.NewController(eventUpdateValidator, eventService)
}

// Api bootstraps the http server.
func Api(rabbitMqChannel *amqp.Channel) *api.WebServer {
	eventUpdateValidator := newEventUpdateValidator()
	eventUpdatePublisher := newEventUpdatePublisher(rabbitMqChannel)
	eventService := newEventService(eventUpdatePublisher)
	controller := newController(eventUpdateValidator, eventService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
