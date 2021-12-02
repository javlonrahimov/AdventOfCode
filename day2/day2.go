package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "javlonrahimov/AdventOfCode/utils"
)

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))

}

func part1(lines []string) int {
	depth, length := 0, 0
	for _, line := range lines {
		lineArr := strings.Split(line, " ")
		switch lineArr[0] {
		case "forward":
			result, _ := strconv.Atoi(lineArr[1])
			length += result
		case "up":
			result, _ := strconv.Atoi(lineArr[1])
			depth -= result
		case "down":
			result, _ := strconv.Atoi(lineArr[1])
			depth += result
		}
	}
	return depth * length
}

func part2(lines []string) int {
	return 0
}
