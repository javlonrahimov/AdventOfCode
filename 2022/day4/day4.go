package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strings"
)

func main() {
	data := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(data []string) int {
	counter := 0
	for _, line := range data {
		firstPair := strings.Split(line, ",")[0]
		secondPair := strings.Split(line, ",")[1]
		fpStart := utils.ToInt(strings.Split(firstPair, "-")[0])
		fpEnd := utils.ToInt(strings.Split(firstPair, "-")[1])
		spStart := utils.ToInt(strings.Split(secondPair, "-")[0])
		spEnd := utils.ToInt(strings.Split(secondPair, "-")[1])

		if (fpStart >= spStart && fpStart <= spEnd && fpEnd <= spEnd) || (spStart >= fpStart && spStart <= fpEnd && spEnd <= fpEnd) {
			counter++
		}
	}
	return counter
}

func part2(data []string) int {
	counter := 0
	for _, line := range data {
		firstPair := strings.Split(line, ",")[0]
		secondPair := strings.Split(line, ",")[1]
		fpStart := utils.ToInt(strings.Split(firstPair, "-")[0])
		fpEnd := utils.ToInt(strings.Split(firstPair, "-")[1])
		spStart := utils.ToInt(strings.Split(secondPair, "-")[0])
		spEnd := utils.ToInt(strings.Split(secondPair, "-")[1])

		if fpStart <= spEnd && fpEnd >= spStart {
			counter++
		}
	}
	return counter
}
