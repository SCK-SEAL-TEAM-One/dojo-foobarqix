package mo_foobarqix

import "testing"

func Test_DivisibleFooBarQix_Input_3_Should_Foo(t *testing.T){
	expectedResult := "FooFoo"

	actualResult := DivisibleFooBarQix("3")

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v",expectedResult,actualResult)
	}
}

func Test_DivisibleFooBarQix_Input_33_Should_FooFooFoo(t *testing.T){
	expectedResult := "FooFooFoo"

	actualResult := DivisibleFooBarQix("33")

	if expectedResult != actualResult {
		t.Errorf("Expect %v but get %v",expectedResult,actualResult)
	}
}