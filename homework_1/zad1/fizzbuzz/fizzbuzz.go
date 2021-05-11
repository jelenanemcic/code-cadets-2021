package fizzbuzz

import (
	"fmt"
	"github.com/pkg/errors"
)

func PlayFizzBuzz(start, end int) error {

	if start > end {
		return errors.New("Value start is greater than the value end.")
	}

	for n := start; n <= end; n++ {
		if n % 3 == 0 && n % 5 == 0 {
			fmt.Printf("%s ", "FizzBuzz")
		} else if n % 3 == 0 {
			fmt.Printf("%s ", "Fizz")
		} else if n % 5 == 0 {
			fmt.Printf("%s ", "Buzz")
		} else {
			fmt.Printf("%d ", n)
		}
	}

	return nil
}
