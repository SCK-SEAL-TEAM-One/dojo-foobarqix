package mo_foobarqix

import (
	"fmt"
	"strconv"
)

func DivisibleFooBarQix(number string) string{
	intNumber,_ := strconv.Atoi(number)
	var fooBarQix string
	for _, c := range number {
		if string(c) == "3" {
			fooBarQix += fmt.Sprint("Foo")
		}
	}
	if intNumber%3==0 {
		fooBarQix += fmt.Sprint("Foo")
	}
	return fooBarQix
}