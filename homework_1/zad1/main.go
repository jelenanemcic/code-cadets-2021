package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strings"

	"code-cadets-2021/homework_1/zad1/fizzbuzz"
)

func main() {
	var start, end int

	flag.IntVar(&start, "start", 1, "Starting value for game Fizzbuzz")
	flag.IntVar(&end, "end", 10, "Ending value for game Fizzbuzz")

	flag.Parse()

	result, err := fizzbuzz.PlayFizzBuzz(start, end)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error playing FizzBuzz."),
		)
	}

	fmt.Printf("%v", strings.Join(result, " "))
}