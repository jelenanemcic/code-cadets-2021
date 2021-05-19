package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http"
)

func AxilisOfferFeed(httpClient *stdhttp.Client) *http.AxilisOfferFeed {
	return http.NewAxilisOfferFeed(httpClient)
}
