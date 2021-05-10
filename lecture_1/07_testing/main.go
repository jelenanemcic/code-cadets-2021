package main

import (
	"flag"
	"log"

	"code-cadets-2021/lecture_1/07_testing/filter"
)

func main() {
	var countFrom, countTo, divisor int

	// flag -> za parsiranje iz command linea
	// (ime varijable, ime koje se stavlja u cmd, defaultna vrijednost, opis)
	flag.IntVar(&countFrom, "count-from", 1, "Value (inclusive) from which to start counting")
	flag.IntVar(&countTo, "count-to", 10, "Value (inclusive) to count to")
	flag.IntVar(&divisor, "divisor", 1, "The divisor used for filtering")

	// za parsiranje command line argumenata

	numbers, err := filter.GetDivisibleFromRange(countFrom, countTo, divisor)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", numbers)
}
