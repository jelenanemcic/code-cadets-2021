package taxLibrary_test

import (
	"code-cadets-2021/homework_1/zad2/taxLibrary"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	for _, tc := range getTestCases() {
		actualOutput, actualErr := taxLibrary.CalculateTax(tc.classes, tc.value)

		if tc.expectingError {
			if actualErr == nil {
				t.Errorf("Expected an error but got `nil` error.")
			}
		} else {
			if actualErr != nil {
				t.Errorf("Expected no error but got non-nil error: %v", actualErr)
			}

			if actualOutput != tc.expectedOutput {
				t.Errorf(
					"Actual and expected output is not equal - actual: %f, expected: %f",
					actualOutput,
					tc.expectedOutput)
			}
		}
	}
}