package pokemonAPI

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sethgrid/pester"
)

type PokemonInformation struct {
	Name                   string
	LocationAreaEncounters string `json:"location_area_encounters"`
}

type PokemonLocation struct {
	LocationArea LocationArea `json:"location_area"`
}

type LocationArea struct {
	Name string `json:"name"`
}

const pokemonURL = "https://pokeapi.co/api/v2/pokemon/"

func MakeHTTPRequest(name string, isName bool) (resp *http.Response, err error) {
	httpClient := pester.New()
	var httpResponse *http.Response
	if isName {
		httpResponse, err = httpClient.Get(pokemonURL + name)
	} else {
		httpResponse, err = httpClient.Get(name)
	}
	return httpResponse, err
}

func ReadContent(httpResponse *http.Response) ([]byte, error) {
	return ioutil.ReadAll(httpResponse.Body)
}

func GetPokemonInfo(bodyContent []byte) (PokemonInformation, error) {
	var pokemonInfo PokemonInformation
	err := json.Unmarshal(bodyContent, &pokemonInfo)
	return pokemonInfo, err
}

func GetPokemonLocations(bodyContent []byte) ([]PokemonLocation, error) {
	var pokemonEncountersJSON []PokemonLocation
	err := json.Unmarshal(bodyContent, &pokemonEncountersJSON)
	return pokemonEncountersJSON, err
}
