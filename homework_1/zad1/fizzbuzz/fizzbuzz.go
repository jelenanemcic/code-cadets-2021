package fizzbuzz

import (
	"fmt"

	"github.com/pkg/errors"
)

// CalculateFizzBuzz expects two integers (the starting number and the ending number), and calculates the FizzBuzz output.
// The return values are FizzBuzz string array and the encountered error (if there is one).
func CalculateFizzBuzz(start, end int) ([]string, error) {

	if start > end {
		return nil, errors.New("Value start is greater than the value end.")
	}

	var gameResult []string

	for n := start; n <= end; n++ {
		if n % 3 == 0 && n % 5 == 0 {
			gameResult = append(gameResult, "FizzBuzz")
		} else if n % 3 == 0 {
			gameResult = append(gameResult, "Fizz")
		} else if n % 5 == 0 {
			gameResult = append(gameResult, "Buzz")
		} else {
			gameResult = append(gameResult, fmt.Sprintf("%d", n))
		}
	}

	return gameResult, nil
}
