package bootstrap

import (
	"code-cadets-2021/homework_2/internal/domain/services"
)

func FeedMerger(feeds ...services.Feed) *services.FeedMerger {
	return services.NewFeedMerger(feeds)
}
