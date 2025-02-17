package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type response struct {
	Name string
	Age int
	Passed  bool
	Skills []string
}

const url = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func main() {
	httpClient := pester.New()

	httpResponse, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards yesno API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of yesno API response"),
		)
	}

	var examResults []response
	err = json.Unmarshal(bodyContent, &examResults)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	f, err := os.Create("name_list.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	defer f.Close()

	for _, result := range examResults {
		if result.Passed == true {
			if contains(result.Skills, "Java") || contains(result.Skills, "Go") {
				f.WriteString(fmt.Sprintf("%s - %s", result.Name, strings.Join(result.Skills, ", ")) + "\n")
			}
		}
	}

	f.Sync()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

