package main

import (
	"bytes"
	"encoding/json"
	"github.com/jelenanemcic/code-cadets-2021/homework_4/event_settler/models"
	"github.com/pkg/errors"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	var settled []string

	response, err := http.Get("http://127.0.0.1:8081/bets?status=active")
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error in HTTP GET request"))
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error reading body of HTTP response"))
	}

	var bets []models.BetReduced
	err = json.Unmarshal(body, &bets)
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error unmarshalling response body"))
	}

	for _, bet := range bets {
		if !contains(settled, bet.SelectionId) {
			var outcome string
			if rand.Float64() > 0.5 {
				outcome = "lost"
			} else {
				outcome = "won"
			}

			eventUpdate := models.EventUpdate{
				Id:      bet.SelectionId,
				Outcome: outcome,
			}

			eventUpdateBuf := new(bytes.Buffer)
			err := json.NewEncoder(eventUpdateBuf).Encode(eventUpdate)
			if err != nil {
				log.Fatal(errors.WithMessage(err, "error encoding event update"))
			}

			response, err = http.Post("http://127.0.0.1:8080/event/update", "application/json", eventUpdateBuf)
			if err != nil {
				log.Fatal(errors.WithMessage(err, "error in HTTP POST request"))
			}

			log.Printf("Done: %v", bet.SelectionId)
			settled = append(settled, bet.SelectionId)
		}
	}
}

func contains(slice []string, element string) bool {
	for _, i := range slice {
		if i == element {
			return true
		}
	}
	return false
}
