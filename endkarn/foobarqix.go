package endkarn

import (
	"strconv"
	"strings"
)

func ComputeFooBarQix(input string) interface{} {
	number, _ := strconv.Atoi(input)
	numberTextSet := [3]string{"3", "5", "7"}
	numberSet := [3]int{3, 5, 7}
	stringSet := [3]string{"Foo", "Bar", "Qix"}
	var output string

	for i := 0; i < len(numberSet); i++ {
		if number%numberSet[i] == 0 {
			output += stringSet[i]
		}
		if strings.Contains(input, numberTextSet[i]) {
			output += stringSet[i]
		}
	}
	if len(output) == 0 {
		return input
	}
	return output
}
