package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strconv"
	"strings"
)

var coordinates = initCoordinates()

func main() {
	lines := utils.ReadFile("input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines string) int {
	return solve(lines, false)
}

func part2(lines string) int {
	return solve(lines, true)
}

func initCoordinates() [1000][1000]int {
	return [1000][1000]int{}
}

func parse(input string) []string {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = strings.TrimSpace(line)
	}
	return result
}

func solve(input string, diagonal bool) int {
	counter := make(map[Point]int)
	lines := parse(input)
	for _, line := range lines {
		a, b := parseLine(line)
		deltaX, deltaY := 0, 0
		if b.y < a.y {
			deltaY = -1
		} else if b.y > a.y {
			deltaY = 1
		}
		if b.x < a.x {
			deltaX = -1
		} else if b.x > a.x {
			deltaX = 1
		}

		if !diagonal && (a.x != b.x && a.y != b.y) {
			continue
		}
		start := a
		for {
			counter[start]++
			if start.x == b.x && start.y == b.y {
				break
			}
			start.x += deltaX
			start.y += deltaY
		}
	}

	var total int
	for _, v := range counter {
		if v > 1 {
			total += 1
		}
	}

	return total
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

type Point struct {
	x, y int
}

func parseLine(line string) (Point, Point) {
	parts := strings.Split(line, " -> ")
	return parsePoint(parts[0]), parsePoint(parts[1])
}

func parsePoint(point string) Point {
	parts := strings.Split(point, ",")
	return Point{
		x: toInt(parts[0]),
		y: toInt(parts[1]),
	}
}
