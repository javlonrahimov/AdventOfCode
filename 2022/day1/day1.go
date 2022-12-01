package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
)

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int64 {
	var max int64 = 0
	var temp int64 = 0
	for i, line := range lines {
		if line == "" {
			if max < temp {
				max = temp
			}
			temp = 0
		} else {
			temp = temp + StringToInt64(line)
		}
		if i == len(lines) {
			if max < temp {
				max = temp
			}
			temp = 0
		}
	}
	return max
}

func part2(lines []string) int64 {
	var max1 int64 = 0
	var max2 int64 = 0
	var max3 int64 = 0
	var temp int64 = 0
	for i, line := range lines {
		if line == "" {
			if temp > max1 {
				max3 = max2
				max2 = max1
				max1 = temp
			} else if temp > max2 {
				max3 = max2
				max2 = temp
			} else if temp > max3 {
				max3 = temp
			}
			temp = 0
		} else {
			temp = temp + StringToInt64(line)
		}
		if i == len(lines)-1 {
			if temp > max1 {
				max3 = max2
				max2 = max1
				max1 = temp
			} else if temp > max2 {
				max3 = max2
				max2 = temp
			} else if temp > max3 {
				max3 = temp
			}
			temp = 0
		}
	}
	return max1 + max2 + max3
}

func StringToInt64(s string) int64 {
	var x int64
	fmt.Sscanf(s, "%d", &x)
	return x
}
