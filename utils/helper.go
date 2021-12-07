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