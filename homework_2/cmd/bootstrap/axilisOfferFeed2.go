package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/http"
)

func AxilisOfferFeed2() *http.AxilisOfferFeed2 {
	httpClient := &stdhttp.Client{}
	return http.NewAxilisOfferFeed2(httpClient)
}
