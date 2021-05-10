package filter

import "github.com/pkg/errors"

// bolje pisati za svaki parametar kojeg je tipa, čitljivije

func GetDivisibleFromRange(start, end, divisor int) ([]int, error) {
	if start > end {
		return nil, errors.New("range start is greater than range end")
	}

	if divisor <= 0 {
		return nil, errors.New("divisor is 0 or negative")
	}

	// slice
	var filtered []int

	for n := start; n <= end; n++ {
		if n % divisor == 0 {
			filtered = append(filtered, n)
		}
	}

	return filtered, nil
}
