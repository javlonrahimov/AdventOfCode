package utils

import (
	"strconv"
	"strings"
)

func ToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func ParseStringToList(input string) []int {
	stringList := strings.Split(input, ",")
	intList := make([]int, len(stringList))
	for i, s := range stringList {
		intList[i] = ToInt(s)
	}

	return intList
}

func ToASCIICode(arg interface{}) int {
	var asciiVal int
	switch arg.(type) {
	case string:
		str := arg.(string)
		if len(str) != 1 {
			panic("can only convert ascii Code for string of length 1")
		}
		asciiVal = int(str[0])
	case byte:
		asciiVal = int(arg.(byte))
	case rune:
		asciiVal = int(arg.(rune))
	}

	return asciiVal
}
