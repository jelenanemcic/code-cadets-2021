package main

import (
	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/services"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/tasks"
	"fmt"
)

func main() {

	queue := bootstrap.NewOrderedQueue()
	offerFeed := bootstrap.NewAxilisOfferFeed()
	service := services.NewFeedProcessorService(offerFeed, queue)
	signalHandler := tasks.NewSignalHandler()

	tasks.RunTasks(offerFeed, service, queue, signalHandler)

	fmt.Println("program finished gracefully")
}
