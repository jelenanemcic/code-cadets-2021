package services

import (
	"context"
	"fmt"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
)

type FeedProcessorService struct {
	FeedService Feed
	QueueService Queue
}

func NewFeedProcessorService(feed Feed, queue Queue) *FeedProcessorService {
	// it should receive "Feed" & "Queue" interfaces through constructor
	return &FeedProcessorService{feed, queue}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	// initially:
	// - get updates channel from feed interface
	// - get source channel from queue interface
	//
	// repeatedly:
	// - range over updates channel
	// - multiply each odd with 2
	// - send it to source channel
	//
	// finally:
	// - when updates channel is closed, exit
	// - when exiting, close source channel

	updates := f.FeedService.GetUpdates()
	source := f.QueueService.GetSource()

	defer fmt.Println("Gasi se service.")
	defer close(source)

	for odd := range updates {
		odd.Coefficient = odd.Coefficient * 2
		source <- odd
		fmt.Printf("%v\n", odd)
	}

	return nil
}

// define feed interface here
type Feed interface {
	Start(ctx context.Context) error
	GetUpdates() chan models.Odd
}

// define queue interface here
type Queue interface {
	Start(ctx context.Context) error
	GetSource() chan models.Odd
}
