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
	for _, line := range lines {
		// fmt.Println(index, line)
		for _, letter := range line {
			fmt.Println(letter, "hello")
		}
	}
	return result
}
