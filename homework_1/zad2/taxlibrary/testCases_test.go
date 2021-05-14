package taxlibrary_test

import (
	"math"

	"code-cadets-2021/homework_1/zad2/taxlibrary"
)

type testCase struct {
	classes []taxlibrary.TaxClass
	value float64

	expectedOutput float64
	expectingError bool
}

func getTestCases() []testCase {
	return []testCase {
		{
			classes: []taxlibrary.TaxClass {
				{
					Start:      0,
					End:        500,
					Percentage: 0,
				},
				{
					Start:      500,
					End:        2000,
					Percentage: 0.15,
				},
				{
					Start:      2000,
					End:        6000,
					Percentage: 0.25,
				},
				{
					Start:      6000,
					End:        math.MaxInt64,
					Percentage: 0.4,
				},
			},
			value: 8000,

			expectedOutput: 2025,
			expectingError: false,
		},
		{
			classes: []taxlibrary.TaxClass {
				{
					Start:      0,
					End:        1000,
					Percentage: 0,
				},
				{
					Start:      1000,
					End:        5000,
					Percentage: 0.1,
				},
				{
					Start:      5000,
					End:        10000,
					Percentage: 0.2,
				},
				{
					Start:      10000,
					End:        math.MaxInt64,
					Percentage: 0.3,
				},
			},
			value: 7000,

			expectedOutput: 800,
			expectingError: false,
		},
		{
			classes: []taxlibrary.TaxClass {
				{
					Start:      0,
					End:        1000,
					Percentage: 0.1,
				},
				{
					Start:      500,
					End:        1000,
					Percentage: 0.2,
				},
				{
					Start:      1000,
					End:        math.MaxInt64,
					Percentage: 0.4,
				},
			},

			expectingError: true,
		},
		{
			classes: []taxlibrary.TaxClass {
				{
					Start:      0,
					End:        1000,
					Percentage: 0,
				},
				{
					Start:      1000,
					End:        5000,
					Percentage: -0.1,
				},
				{
					Start:      5000,
					End:        10000,
					Percentage: -0.2,
				},
				{
					Start:      10000,
					End:        math.MaxInt64,
					Percentage: 0.3,
				},
			},

			expectingError: true,
		},
		{
			classes: []taxlibrary.TaxClass {
				{
					Start:      0,
					End:        1000,
					Percentage: 0,
				},
				{
					Start:      1000,
					End:        5000,
					Percentage: 0.1,
				},
				{
					Start:      5000,
					End:        10000,
					Percentage: 0.2,
				},
				{
					Start:      10000,
					End:        11000,
					Percentage: 0.3,
				},
			},

			expectingError: true,
		},
		{
			classes: []taxlibrary.TaxClass {
				{
					Start:      0,
					End:        1500.5,
					Percentage: 0.1,
				},
				{
					Start:      1500.5,
					End:        4000,
					Percentage: 0.2,
				},
				{
					Start:      4000,
					End:        math.MaxInt64,
					Percentage: 0.5,
				},
			},
			value: 5250.5,

			expectedOutput: 1275.2,
			expectingError: false,
		},
	}
}
