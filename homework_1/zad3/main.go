package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"strings"
)

type pokemonInformation struct {
	Name string
	LocationAreaEncounters string `json:"location_area_encounters"`
}

type pokemonLocation struct {
	LocationArea locationArea `json:"location_area"`
}

type locationArea struct {
	Name string
}

const pokemonURL= "https://pokeapi.co/api/v2/pokemon/"

func main() {

	fmt.Println("Pokemon name or number:")

	var pokemon string
	_, err := fmt.Scanln(&pokemon)
	if err != nil {
		log.Fatal(
			errors.New("Error reading user input."),
		)
	}

	httpClient := pester.New()
	httpResponse, err := httpClient.Get(pokemonURL + pokemon)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error in HTTP get towards pokemon API."),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error reading body of pokemon API response."),
		)
	}


	var pokemonInfo pokemonInformation
	err = json.Unmarshal(bodyContent, &pokemonInfo)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error unmarshalling the JSON body content."),
		)
	}

	log.Printf("Pokemon name: %s", pokemonInfo.Name)

	httpResponse, err = httpClient.Get(pokemonInfo.LocationAreaEncounters)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error in HTTP get towards pokemon encounters API."),
		)
	}

	bodyContent, err = ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error reading body of pokemon encounters API response."),
		)
	}

	var pokemonEncountersJSON []pokemonLocation

	err = json.Unmarshal(bodyContent, &pokemonEncountersJSON)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error unmarshalling the JSON body content."),
		)
	}

	var pokemonEncounters []string

	for _, location := range pokemonEncountersJSON {
		pokemonEncounters = append(pokemonEncounters, location.LocationArea.Name)
	}

	log.Printf("Pokemon encounters: %v", strings.Join(pokemonEncounters, ", "))

}