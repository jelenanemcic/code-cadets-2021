package main

import (
	"log"

	"github.com/jelenanemcic/code-cadets-2021/homework_3/controller/cmd/bootstrap"
	"github.com/jelenanemcic/code-cadets-2021/homework_3/controller/cmd/config"
	"github.com/jelenanemcic/code-cadets-2021/homework_3/controller/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	db := bootstrap.Sqlite()

	signalHandler := bootstrap.SignalHandler()
	engine := bootstrap.Engine(rabbitMqChannel, db)

	log.Println("Bootstrap finished. Engine is starting")

	tasks.RunTasks(signalHandler, engine)

	log.Println("Service finished gracefully")
}
