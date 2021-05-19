package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http"
)

func AxilisOfferFeed2(httpClient *stdhttp.Client) *http.AxilisOfferFeed2 {
	return http.NewAxilisOfferFeed2(httpClient)
}
