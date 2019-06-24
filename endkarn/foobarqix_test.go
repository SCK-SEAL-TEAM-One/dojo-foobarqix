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

func Test_FooBarQix_Put_15_Should_Return_FooBarBar(t *testing.T) {
	expectedResult := "FooBarBar"

	actualResult := ComputeFooBarQix("15")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

func Test_FooBarQix_Put_21_Should_Return_FooQix(t *testing.T) {
	expectedResult := "FooQix"

	actualResult := ComputeFooBarQix("21")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

func Test_FooBarQix_Put_33_Should_Return_FooFooFoo(t *testing.T) {
	expectedResult := "FooFooFoo"

	actualResult := ComputeFooBarQix("33")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

func Test_FooBarQix_Put_51_Should_Return_FooBar(t *testing.T) {
	expectedResult := "FooBar"

	actualResult := ComputeFooBarQix("51")

	if expectedResult != actualResult {
		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
	}
}

//func Test_FooBarQix_Put_53_Should_Return_BarFoo(t *testing.T) {
//	expectedResult := "BarFoo"
//
//	actualResult := ComputeFooBarQix("53")
//
//	if expectedResult != actualResult {
//		t.Errorf(" Expected value %v but got %v ", expectedResult, actualResult)
//	}
//}
