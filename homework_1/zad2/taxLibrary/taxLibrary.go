package taxLibrary

import (
	"github.com/pkg/errors"
	"math"
)

type TaxClass struct {
	Start      int
	End        int
	Percentage float64
}

func checkClassesDefinitions(classes []TaxClass) error {

	for n, class := range classes {
		if n != len(classes) - 1 {
			if class.End != classes[n+1].Start {
				return errors.New("Marginal values are not the same.")
			}
		}
		if class.Percentage < 0 {
			return errors.New("Tax percentage must be >= 0.")
		}
	}

	if classes[len(classes) - 1].End != math.MaxInt64 {
		return errors.New("End of the last class must be infinite.")
	}

	return nil
}

func CalculateTax(classes []TaxClass, value int) (float64, error) {
	err := checkClassesDefinitions(classes)
	if err != nil {
		return 0, err
	}

	var tax float64 = 0
	var taxed = 0

	for n := range classes {
		if value > classes[n].End {
			tax = tax + (float64(classes[n].End- taxed) * classes[n].Percentage)
			taxed = classes[n].End
		} else {
			tax = tax + (float64(value - taxed) * classes[n].Percentage)
			break
		}
	}

	return tax, nil
}

