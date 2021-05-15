package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"github.com/pkg/errors"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeed(httpClient http.Client) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	// repeatedly:
	// - get odds from HTTP server
	// - write them to updates channel
	// - if context is finished, exit and close updates channel
	// (test your program from cmd/main.go)

	defer close(a.updates)
	defer fmt.Println("Gasi se feed.")

	for {
		select {

		case <-ctx.Done():
			fmt.Println("finished")
			return nil

		case <-time.After(time.Second*3):
			httpResponse, err := a.httpClient.Get(axilisFeedURL)
			if err != nil {
				return errors.WithMessage(err, "error in HTTP request")
			}

			bodyContent, err := ioutil.ReadAll(httpResponse.Body)
			if err != nil {
				return errors.WithMessage(err, "error reading HTTP response")
			}

			var offerOdds []axilisOfferOdd

			err = json.Unmarshal(bodyContent, &offerOdds)
			if err != nil {
				return errors.WithMessage(err, "error unmarshalling JSON")
			}

			for i := range offerOdds {
				odd := models.Odd{
					Id:          offerOdds[i].Id,
					Name:        offerOdds[i].Name,
					Match:       offerOdds[i].Match,
					Coefficient: offerOdds[i].Details.Price,
					Timestamp:   time.Now(),
				}
				a.updates <- odd
			}
		}
	}
}

func (a *AxilisOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}

type axilisOfferOdd struct {
	Id      string
	Name    string
	Match   string
	Details axilisOfferOddDetails
}

type axilisOfferOddDetails struct {
	Price float64
}
