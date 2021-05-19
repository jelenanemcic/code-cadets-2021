package bootstrap

import (
	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/services"
)

func FeedMerger(feed ...services.Feed) *services.FeedMerger {
	return services.NewFeedMerger(feed)
}
