package main

import (
	"log"

	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/cmd/bootstrap"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/cmd/config"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	signalHandler := bootstrap.SignalHandler()
	api := bootstrap.Api(rabbitMqChannel)

	log.Println("Bootstrap finished. Event API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Event API finished gracefully")
}
