package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"code-cadets-2021/homework_1/zad2/taxlibrary"
	"github.com/pkg/errors"
)

// The program expects N string arguments, where N is an arbitrary number.
// First N-1 arguments represent tax classes and each one has three values: start, end and tax percentage
// (except the last class that does not have the end variable).
// The last argument is a value for which tax will be calculated.
func main() {

	var classes []taxlibrary.TaxClass
	var start, end float64
	var percentage float64

	numOfClasses := len(os.Args) - 2
	value, _ := strconv.ParseFloat(os.Args[len(os.Args)-1], 64)

	for n := 1; n <= numOfClasses; n++ {
		input := strings.Fields(os.Args[n])

		if len(input) != 3 && n != numOfClasses {
			log.Fatal(
				errors.New("invalid input"),
			)
		}

		if len(input) == 2 && n == numOfClasses {
			start, _ = strconv.ParseFloat(input[0],64)
			end  = math.MaxInt64
			percentage, _ = strconv.ParseFloat(input[1],64)
		} else {
			start, _ = strconv.ParseFloat(input[0],64)
			end, _ = strconv.ParseFloat(input[1],64)
			percentage, _ = strconv.ParseFloat(input[2],64)
		}

		class := taxlibrary.TaxClass{
			Start:      start,
			End:        end,
			Percentage: percentage,
		}

		classes = append(classes, class)
	}

	tax, err := taxlibrary.CalculateTax(classes, value)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "error calculating tax"),
		)
	}

	fmt.Printf("%f", tax)
}
