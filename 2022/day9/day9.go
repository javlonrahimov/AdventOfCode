package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"math"
	"strconv"
	"strings"
)

type Loc struct {
	x int
	y int
}

func (c Loc) Up(amount int) Loc {
	return Loc{c.x, c.y + amount}
}
func (c Loc) Down(amount int) Loc {
	return Loc{c.x, c.y - amount}
}
func (c Loc) Right(amount int) Loc {
	return Loc{c.x + amount, c.y}
}
func (c Loc) Left(amount int) Loc {
	return Loc{c.x - amount, c.y}
}

func (c Loc) Previous(direction string) Loc {
	switch direction {
	case "U":
		return c.Down(1)
	case "D":
		return c.Up(1)
	case "R":
		return c.Left(1)
	case "L":
		return c.Right(1)
	}
	return c
}

func (c Loc) isTouching(other Loc) bool {
	xDistance := math.Abs(float64(c.x) - float64(other.x))
	yDistance := math.Abs(float64(c.y) - float64(other.y))

	return xDistance <= 1 && yDistance <= 1
}

func part1(lines []string) int {
	visited := make(map[Loc]bool)
	head := Loc{0, 0}
	tail := Loc{0, 0}
	count := 1
	visited[tail] = true
	for _, line := range lines {
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])

		for a := 0; a < amount; a++ {
			switch split[0] {
			case "U":
				head = head.Up(1)
			case "D":
				head = head.Down(1)
			case "R":
				head = head.Right(1)
			case "L":
				head = head.Left(1)
			}
			if !tail.isTouching(head) {
				tail = head.Previous(split[0])
				if !visited[tail] {
					count++
				}
				visited[tail] = true
			}
		}
	}
	return count
}

func dir(from, to int) int {
	if to == from {
		return 0
	} else if to > from {
		return 1
	}
	return -1
}

func part2(lines []string) int {
	visited := make(map[Loc]bool)
	var heads []Loc
	for i := 0; i < 10; i++ {
		heads = append(heads, Loc{0, 0})
	}
	visited[Loc{0, 0}] = true
	count := 1
	for _, line := range lines {
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])

		for a := 0; a < amount; a++ {
			switch split[0] {
			case "U":
				heads[0] = heads[0].Up(1)
			case "D":
				heads[0] = heads[0].Down(1)
			case "R":
				heads[0] = heads[0].Right(1)
			case "L":
				heads[0] = heads[0].Left(1)
			}

			for knot := 1; knot < len(heads); knot++ {

				if !heads[knot].isTouching(heads[knot-1]) {
					heads[knot] = Loc{
						heads[knot].x + dir(heads[knot].x, heads[knot-1].x),
						heads[knot].y + dir(heads[knot].y, heads[knot-1].y),
					}
				}
			}
			if !visited[heads[len(heads)-1]] {
				count++
			}
			visited[heads[len(heads)-1]] = true
		}
	}
	return count
}

func main() {

	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
