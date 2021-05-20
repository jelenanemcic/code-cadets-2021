package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/homework_2/internal/infrastructure/http"
)

func AxilisOfferFeed(httpClient *stdhttp.Client) *http.AxilisOfferFeed {
	return http.NewAxilisOfferFeed(httpClient)
}
