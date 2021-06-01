package main

import (
	"log"

	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/cmd/bootstrap"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/cmd/config"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/bets_api/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	signalHandler := bootstrap.SignalHandler()
	db := bootstrap.Sqlite()
	api := bootstrap.Api(db)

	log.Println("Bootstrap finished. Bets API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Bets API finished gracefully")
}
