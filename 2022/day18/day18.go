package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"math"
	"regexp"
)

type Point struct {
	x, y, z int
}

var re = regexp.MustCompile("[0-9]+")

var delta = []Point{
	{-1, 0, 0}, {0, -1, 0}, {0, 0, -1},
	{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
}

func (p Point) Add(q Point) Point {
	p.x = p.x + q.x
	p.y = p.y + q.y
	p.z = p.z + q.z
	return p
}

func parsePoint(line string) Point {
	numbers := re.FindAllString(line, -1)
	return Point{
		x: utils.ToInt(numbers[0]),
		y: utils.ToInt(numbers[1]),
		z: utils.ToInt(numbers[2]),
	}
}

func createObsidian(lines []string) (map[Point]bool, Point, Point) {
	obsidian := map[Point]bool{}
	min := Point{math.MaxInt, math.MaxInt, math.MaxInt}
	max := Point{math.MinInt, math.MinInt, math.MinInt}
	for _, line := range lines {
		p := parsePoint(line)
		obsidian[p] = true
		min = Point{utils.Min(min.x, p.x), utils.Min(min.y, p.y), utils.Min(min.z, p.z)}
		max = Point{utils.Max(max.x, p.x), utils.Max(max.y, p.y), utils.Max(max.z, p.z)}
	}
	return obsidian, min, max
}

func part1(lines []string) int {

	surface := 0
	obsidian, _, _ := createObsidian(lines)

	for p := range obsidian {
		for _, d := range delta {
			if !obsidian[p.Add(d)] {
				surface++
			}
		}
	}
	return surface
}

func part2(lines []string) int {

	surface := 0
	obsidian, min, max := createObsidian(lines)

	for x := min.x - 1; x <= max.x+1; x++ {
		for y := min.y - 1; y <= max.y+1; y++ {
			for z := min.z - 1; z <= max.z+1; z++ {
				obsidian[Point{x, y, z}] = obsidian[Point{x, y, z}]
			}
		}
	}

	queue := []Point{min}
	visited := map[Point]struct{}{min: {}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range delta {
			next := cur.Add(d)
			if cube, valid := obsidian[next]; cube {
				surface++
			} else if _, seen := visited[next]; valid && !seen {
				visited[next] = struct{}{}
				queue = append(queue, next)
			}
		}
	}

	return surface
}

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
