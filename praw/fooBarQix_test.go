package praw

import (
	"testing"
)

func Test_FooBarQixCompute_Input_Is_6_Should_Get_Foo(t *testing.T){
	expectedResult := "Foo"

	actualResult := FooBarQixCompute("6")

	if expectedResult != actualResult {
		t.Errorf("Expect %v but got %v", expectedResult, actualResult)
	}
}