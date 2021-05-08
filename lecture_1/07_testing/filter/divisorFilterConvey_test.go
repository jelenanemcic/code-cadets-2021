package filter_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"code-cadets-2021/lecture_1/07_testing/filter"
)

// točka . znači da prisupamo ovom paketu s točkom, sva imena i metode iz ovog paketa mi stavi u globalni workspace -> potencijalno opasno (npr. ako imamo dva paketa s istoimenom metodom)
// možemo staviti i ime prije importa i onda s tim imenom referenciramo taj paket

// NOTE - Convey infix in the function name is here just to prevent a name
// clash with the method in `divisorFilter_test.go`
func TestConveyGetDivisibleFromRange(t *testing.T) {

	for idx, tc := range getTestCases() {
		Convey(fmt.Sprintf("Given test case #%v: %+v", idx, tc), t, func() {

			actualOutput, actualErr := filter.GetDivisibleFromRange(tc.inputStart, tc.inputEnd, tc.inputDivisor)

			if tc.expectingError {
				So(actualErr, ShouldNotBeNil)
			} else {
				So(actualErr, ShouldBeNil)
				So(actualOutput, ShouldResemble, tc.expectedOutput)
			}

		})
	}
}
