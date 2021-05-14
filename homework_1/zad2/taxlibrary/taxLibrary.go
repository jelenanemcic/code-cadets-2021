package taxlibrary

import (
	"math"

	"github.com/pkg/errors"
)

type TaxClass struct {
	Start      float64
	End        float64
	Percentage float64
}

func checkClassesDefinitions(classes []TaxClass) error {
	for n, class := range classes {

		if class.Percentage < 0 {
			return errors.New("tax percentage must be >= 0")
		}

		if class.Start >= class.End {
			return errors.New("end value must be larger than the start value")
		}

		if n != len(classes) - 1 {
			if class.End != classes[n+1].Start {
				return errors.New("marginal values are not the same")
			}
		}
	}

	if classes[len(classes) - 1].End != math.MaxInt64 {
		return errors.New("end of the last class must be infinite")
	}

	return nil
}

// CalculateTax receives tax classes and an integer value.
// It calculates the tax for the received value and returns the tax and an encountered error (if there is one).
func CalculateTax(classes []TaxClass, value float64) (float64, error) {
	err := checkClassesDefinitions(classes)
	if err != nil {
		return 0, err
	}

	var tax float64 = 0
	var taxed = 0.

	for n := range classes {
		if value > classes[n].End {
			taxedValue := classes[n].End- taxed
			tax += taxedValue * classes[n].Percentage
			taxed = classes[n].End
		} else {
			taxedValue := value - taxed
			tax += taxedValue * classes[n].Percentage
			break
		}
	}

	return tax, nil
}
