package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Line struct {
	from Coords
	to   Coords
}

func parseInput(lines []string) (map[Coords]bool, int) {
	blocks := make(map[Coords]bool)
	last := Coords{}
	maxCol := -1
	for _, line := range lines {
		paths := strings.Split(line, " -> ")
		for i, path := range paths {
			split := strings.Split(path, ",")
			row, _ := strconv.Atoi(split[0])
			col, _ := strconv.Atoi(split[1])
			coords := Coords{row, col}
			if i != 0 {
				addToSet(&blocks, Line{last, coords})
			}
			last = coords
			if col > maxCol {
				maxCol = col
			}
		}
	}
	return blocks, maxCol + 1
}

func addToSet(blocks *map[Coords]bool, line Line) {
	if line.from.x == line.to.x {
		from := line.from.y
		to := line.to.y
		if line.to.y < from {
			from = line.to.y
			to = line.from.y
		}
		for i := from; i <= to; i++ {
			(*blocks)[Coords{line.from.x, i}] = true
		}
	} else {
		from := line.from.x
		to := line.to.x
		if line.to.x < from {
			from = line.to.x
			to = line.from.x
		}
		for i := from; i <= to; i++ {
			(*blocks)[Coords{i, line.from.y}] = true
		}
	}
}

func part1(input []string) int {
	blocks, abyss := parseInput(input)
	sandCount := 0
	s := []Coords{{500, 0}}
	for {
		sand := s[len(s)-1]
		s = s[:len(s)-1]
		for {
			if ns := (Coords{sand.x, sand.y + 1}); !blocks[ns] {
				if ns.y >= abyss {
					return sandCount
				}
				s = append(s, sand)
				sand = ns
			} else if ns := (Coords{sand.x - 1, sand.y + 1}); !blocks[ns] {
				s = append(s, sand)
				sand = ns
			} else if ns := (Coords{sand.x + 1, sand.y + 1}); !blocks[ns] {
				s = append(s, sand)
				sand = ns
			} else {
				sandCount++
				blocks[sand] = true
				break
			}
		}
	}
}

func part2(input []string) int {
	blocks, abyss := parseInput(input)
	sum := 0
	s := []Coords{{500, 0}}
	for len(s) > 0 {
		sand := s[len(s)-1]
		s = s[:len(s)-1]
		for {
			if sand.y > abyss {
				blocks[sand] = true
				break
			}
			if ns := (Coords{sand.x, sand.y + 1}); !blocks[ns] {
				s = append(s, sand)
				sand = ns
			} else if ns := (Coords{sand.x - 1, sand.y + 1}); !blocks[ns] {
				s = append(s, sand)
				sand = ns
			} else if ns := (Coords{sand.x + 1, sand.y + 1}); !blocks[ns] {
				s = append(s, sand)
				sand = ns
			} else {
				sum++
				blocks[sand] = true
				break
			}
		}
	}
	return sum
}

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
