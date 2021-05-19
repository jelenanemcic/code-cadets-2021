package main

import (
	"fmt"

	"code-cadets-2021/lecture_2/06_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/06_offerfeed/internal/tasks"
)

func main() {
	signalHandler := bootstrap.SignalHandler()

	feed1 := bootstrap.AxilisOfferFeed()
	feed2 := bootstrap.AxilisOfferFeed2()
	feed := bootstrap.FeedMerger(feed1, feed2)
	queue := bootstrap.OrderedQueue()
	processingService := bootstrap.FeedProcessingService(feed, queue)

	// blocking call, start "the application"
	tasks.RunTasks(signalHandler, feed1, feed2, feed, queue, processingService)

	fmt.Println("program finished gracefully")
}
