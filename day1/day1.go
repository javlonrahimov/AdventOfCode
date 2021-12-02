package main

import (
	"fmt"
	utils "javlonrahimov/AdventOfCode/utils"
)

func main() {
	lines := convertStringListToIntList(utils.ReadFileLineByLine("input.txt"))
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []int64) int64 {
	increaseCount := 0
	prevVal := lines[0]
	for _, line := range lines {
		if prevVal < line {
			increaseCount++
		}
		prevVal = line
	}
	return int64(increaseCount)
}

func part2(lines []int64) int64 {
	increaseCount := 0
	prevVal := sum(lines[0], lines[1], lines[2])
	for index := range lines {
		if index > len(lines)-3 {
			break
		}
		sum := sum(lines[index], lines[index+1], lines[index+2])
		if prevVal < sum {
			increaseCount++
		}
		prevVal = sum
	}
	return int64(increaseCount)
}

func sum(numbers ...int64) int64 {
	var sum int64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func convertStringListToIntList(lines []string) []int64 {
	var intList []int64
	for _, line := range lines {
		intList = append(intList, StringToInt64(line))
	}
	return intList
}

func StringToInt64(s string) int64 {
	var x int64
	fmt.Sscanf(s, "%d", &x)
	return x
}
