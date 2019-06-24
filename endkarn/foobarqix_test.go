package endkarn

import "testing"

func Test_FooBarQix_Put_3_Should_Return_FooFoo(t *testing.T) {
	expectedResult := "FooFoo"

	actualResult := ComputeFooBarQix("3")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

func Test_FooBarQix_Put_5_Should_Return_BarBar(t *testing.T) {
	expectedResult := "BarBar"

	actualResult := ComputeFooBarQix("5")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

func Test_FooBarQix_Put_7_Should_Return_QixQix(t *testing.T) {
	expectedResult := "QixQix"

	actualResult := ComputeFooBarQix("7")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

func Test_FooBarQix_Put_1_Should_Return_1(t *testing.T) {
	expectedResult := "1"

	actualResult := ComputeFooBarQix("1")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

func Test_FooBarQix_Put_6_Should_Return_Foo(t *testing.T) {
	expectedResult := "Foo"

	actualResult := ComputeFooBarQix("6")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

func Test_FooBarQix_Put_10_Should_Return_Bar(t *testing.T) {
	expectedResult := "Bar"

	actualResult := ComputeFooBarQix("10")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

func Test_FooBarQix_Put_14_Should_Return_Qix(t *testing.T) {
	expectedResult := "Qix"

	actualResult := ComputeFooBarQix("14")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}
