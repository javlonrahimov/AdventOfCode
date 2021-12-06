package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strconv"
	"strings"
)

var coordinates = initCoordinates()

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	for _, line := range lines {
		x1, y1, x2, y2 := parseLine(line)
		drawLine(x1, y1, x2, y2, false)
	}
	return countCrossings()
}

func part2(lines []string) int {
	for _, line := range lines {
		x1, y1, x2, y2 := parseLine(line)
		drawLine(x1, y1, x2, y2, true)
	}
	return countCrossings()
}

func initCoordinates() [1000][1000]int {
	return [1000][1000]int{}
}

func drawLine(x1, y1, x2, y2 int, canDrawDioganal bool) {
	if x1 == x2 {
		if y1 < y2 {
			for y := y1; y <= y2; y++ {
				coordinates[x1][y] += 1
			}
		} else {
			for y := y2; y <= y1; y++ {
				coordinates[x1][y] += 1
			}
		}
	} else if y1 == y2 {
		if x1 < x2 {
			for x := x1; x <= x2; x++ {
				coordinates[x][y1] += 1
			}
		} else {
			for x := x2; x <= x1; x++ {
				coordinates[x][y1] += 1
			}
		}

	} else if canDrawDioganal {
		
	}
}

func countCrossings() int {
	var count int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if coordinates[x][y] > 1 {
				count++
			}
		}
	}
	coordinates = initCoordinates()
	return count
}

func parseLine(line string) (int, int, int, int) {
	x1, y1, x2, y2 := 0, 0, 0, 0
	x1, y1 = stringPairToIntPair(strings.Split(line, " -> ")[0])
	x2, y2 = stringPairToIntPair(strings.Split(line, " -> ")[1])
	return x1, y1, x2, y2
}

func stringPairToIntPair(pair string) (int, int) {
	x, _ := strconv.Atoi(strings.Split(pair, ",")[0])
	y, _ := strconv.Atoi(strings.Split(pair, ",")[1])
	return x, y
}
