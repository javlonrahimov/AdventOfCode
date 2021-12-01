package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readFileLineByLine("input.txt")
	part1(lines)
	part2(lines)
	sum(5, 6, 7)
}

func part1(lines []int64) {
	increaseCount := 0
	prevVal := lines[0]
	for _, line := range lines {
		if prevVal < line {
			increaseCount++
		}
		prevVal = line
	}
	fmt.Println(increaseCount)
}

func part2(lines []int64) {
	increaseCount := 0
	prevVal := sum(lines[0], lines[1], lines[2])
	for index := range lines {
		if index > len(lines)-3 {break}
		sum := sum(lines[index], lines[index+1], lines[index+2])
		if prevVal < sum {
			increaseCount++
		}
		prevVal = sum
	}
	fmt.Println(increaseCount)
}

func sum(numbers ...int64) int64 {
	var sum int64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func readFileLineByLine(filePath string) []int64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, int64(data))
	}

	return lines
}
