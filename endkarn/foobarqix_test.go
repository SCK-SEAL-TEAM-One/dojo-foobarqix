package endkarn

import "testing"

func Test_FooBarQix_Put_3_Should_Return_FooFoo(t *testing.T) {
	expectedResult := "FooFoo"

	actualResult := ComputeFooBarQix("3")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}
