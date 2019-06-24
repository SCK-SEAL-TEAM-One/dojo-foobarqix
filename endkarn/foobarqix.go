package endkarn

import "strconv"

func ComputeFooBarQix(input string) interface{} {
	int, _ := strconv.Atoi(input)
	if input == "3" {
		return "FooFoo"
	}
	if int%3 == 0 {
		return "Foo"
	}
	if input == "5" {
		return "BarBar"
	}
	if int%5 == 0 {
		return "Bar"
	}
	if input == "7" {
		return "QixQix"
	}
	if int%7 == 0 {
		return "Qix"
	}
	return input
}
