package main

import (
	"fmt"
	"log"
	"strings"

	"code-cadets-2021/homework_1/zad3/pokemonAPI"
	"github.com/pkg/errors"
)

func main() {

	fmt.Println("Pokemon name or number:")

	var pokemonName string
	_, err := fmt.Scanln(&pokemonName)
	if err != nil {
		log.Fatal(
			errors.New("error reading user input"),
		)
	}

	httpResponse, err := pokemonAPI.MakeHTTPRequest(pokemonName, true)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error in HTTP get towards pokemonAPI API"),
		)
	}

	bodyContent, err := pokemonAPI.ReadContent(httpResponse)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error reading body of pokemonAPI API response"),
		)
	}

	pokemonInfo, err := pokemonAPI.GetPokemonInfo(bodyContent)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error unmarshalling the JSON body content"),
		)
	}
	log.Printf("Pokemon name: %s", pokemonInfo.Name)

	httpResponse, err = pokemonAPI.MakeHTTPRequest(pokemonInfo.LocationAreaEncounters, false)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error in HTTP get towards pokemonAPI encounters API"),
		)
	}

	bodyContent, err = pokemonAPI.ReadContent(httpResponse)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error reading body of pokemonAPI encounters API response"),
		)
	}

	pokemonLocations, err := pokemonAPI.GetPokemonLocations(bodyContent)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error unmarshalling the JSON body content"),
		)
	}

	var pokemonEncounters []string
	for _, location := range pokemonLocations {
		pokemonEncounters = append(pokemonEncounters, location.LocationArea.Name)
	}
	log.Printf("Pokemon encounters: %v", strings.Join(pokemonEncounters, ", "))

}
