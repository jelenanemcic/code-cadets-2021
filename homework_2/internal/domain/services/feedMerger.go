package services

import (
	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
	"context"
	"log"
	"sync"
)

type FeedMerger struct {
	updates chan models.Odd
	feeds   []Feed
}

func NewFeedMerger(
	feeds []Feed,
) *FeedMerger {
	return &FeedMerger{
		updates: make(chan models.Odd),
		feeds:   feeds,
	}
}

func (f *FeedMerger) Start(ctx context.Context) error {
	out := f.GetUpdates()

	defer close(out)
	defer log.Printf("shutting down %s", f)

	var wg sync.WaitGroup
	wg.Add(len(f.feeds))

	for _, feed := range f.feeds {
		go func(c chan models.Odd, out chan models.Odd) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(feed.GetUpdates(), out)
	}

	wg.Wait()
	return nil
}

func (f *FeedMerger) GetUpdates() chan models.Odd {
	return f.updates
}

func (f *FeedMerger) String() string {
	return "feed merger"
}
