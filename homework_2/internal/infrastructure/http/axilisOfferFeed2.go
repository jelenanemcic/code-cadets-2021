package http

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code-cadets-2021/homework_2/internal/domain/models"
)

const axilisFeedURL2 = "http://18.193.121.232/axilis-feed-2"

type AxilisOfferFeed2 struct {
	updates    chan models.Odd
	httpClient *http.Client
}

func NewAxilisOfferFeed2(
	httpClient *http.Client,
) *AxilisOfferFeed2 {
	return &AxilisOfferFeed2{
		updates:    make(chan models.Odd),
		httpClient: httpClient,
	}
}

func (a *AxilisOfferFeed2) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-time.After(time.Second * 3):
			response, err := a.httpClient.Get(axilisFeedURL2)
			if err != nil {
				log.Println("axilis offer feed 2, http get", err)
				continue
			}
			a.processResponse(ctx, response)
		}
	}
}

func (a *AxilisOfferFeed2) GetUpdates() chan models.Odd {
	return a.updates
}

func (a *AxilisOfferFeed2) String() string {
	return "axilis offer feed 2"
}

func (a *AxilisOfferFeed2) processResponse(ctx context.Context, response *http.Response) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("axilis offer feed 2, read all", err)
		return
	}

	bodyText := string(body)

	lines := strings.Split(bodyText, "\n")

	for _, line := range lines {
		fields := strings.Split(line, ",")
		coeff, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			log.Println("axilis offer feed 2, parsing coefficient", err)
			continue
		}

		odd := models.Odd{
			Id:          fields[0],
			Name:        fields[1],
			Match:       fields[2],
			Coefficient: coeff,
			Timestamp:   time.Now(),
		}

		select {
		case <-ctx.Done():
			return
		case a.updates <- odd:
		}
	}
}
