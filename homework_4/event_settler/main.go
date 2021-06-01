package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"

	"github.com/jelenanemcic/code-cadets-2021/homework_4/bet_acceptance_api/c"
)

func main() {
	response, err := http.Get("http://127.0.0.1:8080//bets?status=active")
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error in HTTP GET request"))
	}

	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(errors.WithMessage(err, "error reading body of HTTP response"))
	}

	var bets
	err := json.Unmarshal(bodyContent, &pokemonInfo)


}