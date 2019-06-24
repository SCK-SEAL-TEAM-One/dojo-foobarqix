package praw

import (
	"strconv"
)

func FooBarQixCompute(input string) string {
	number, _ := strconv.Atoi(input)
	if number % 3 == 0{
		return "Foo"
	}
	return "Bar"
}