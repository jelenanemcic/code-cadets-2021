package bootstrap

import (
	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/services"
)

func FeedMerger(feeds ...services.Feed) *services.FeedMerger {
	return services.NewFeedMerger(feeds)
}
