package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"code-cadets-2021/homework_1/zad2/taxLibrary"
	"github.com/pkg/errors"
)

// program očekuje kao argumente N stringova, gdje je N proizvoljan
// N-1 argumenata predstavlja porezne razrede i svaki ima 3 vrijednosti: početak, kraj i postotak (osim zadnjeg razreda koji nema kraj)
// zadnji argument je vrijednost za koju računamo porez

func main() {

	var classes []taxLibrary.TaxClass
	var start, end int64
	var percentage float64

	numOfClasses := len(os.Args) - 2
	value, _ := strconv.ParseInt(os.Args[len(os.Args)-1], 10, 64)

	for n := 1; n <= numOfClasses; n++ {
		input := strings.Fields(os.Args[n])

		if len(input) != 3 && n != numOfClasses {
			log.Fatal(
				errors.New("Invalid input"),
			)
		}

		if len(input) == 2 && n == numOfClasses {
			start, _ = strconv.ParseInt(input[0], 10, 64)
			end  = math.MaxInt64
			percentage, _ = strconv.ParseFloat(input[1], 64)
		} else {
			start, _ = strconv.ParseInt(input[0], 10, 64)
			end, _ = strconv.ParseInt(input[1], 10, 64)
			percentage, _ = strconv.ParseFloat(input[2], 64)
		}

		class := taxLibrary.TaxClass{
			Start:      int(start),
			End:        int(end),
			Percentage: percentage,
		}

		classes = append(classes, class)
	}

	tax, err := taxLibrary.CalculateTax(classes, int(value))
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Error calculating tax"),
		)
	}

	fmt.Printf("%f", tax)
}

