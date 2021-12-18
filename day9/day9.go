package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
)

func main() {
	lines := utils.ReadFileLineByLine("input_test.txt")
	fmt.Print(part1(lines))
}

func part1(lines []string) int {
	result := 0
	for index, value := range lines {
		fmt.Println(index, value)
	}
	return result
}
