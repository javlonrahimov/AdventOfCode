package main

import (
	"errors"
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
)

type Coords struct {
	row int
	col int
}

func parseInput(input string, start, end rune) ([]string, Coords, Coords) {
	var lines []string
	line, row, col := "", 0, 0
	startC, endC := Coords{}, Coords{}

	for _, c := range input {
		if c == '\n' {
			lines = append(lines, line)
			line = ""
			col = 0
			row++
			continue
		}
		if c == start {
			startC = Coords{row, col}
		}
		if c == end {
			endC = Coords{row, col}
		}
		if c == 'S' {
			c = 'a'
		} else if c == 'E' {
			c = 'z'
		}
		line += string(c)
		col++
	}
	lines = append(lines, line)
	return lines, startC, endC
}

func bfs(lines []string, start Coords, isEnd func(Coords) bool, canMove func(byte, byte) bool) int {
	visited := make(map[Coords]int)
	var q []Coords
	var curr Coords

	q = append(q, start)
	visited[start] = 0

	for len(q) > 0 {
		curr = q[0]
		q = q[1:]
		v := visited[curr]

		next := getMovableCoords(lines, curr, canMove)

		for _, n := range next {
			if isEnd(n) {
				return v + 1
			}
			_, ok := visited[n]
			if !ok {
				visited[n] = v + 1
				q = append(q, n)
			}
		}
	}
	return -1
}

func getMovableCoords(lines []string, x Coords, canMove func(byte, byte) bool) []Coords {
	var next []Coords
	y, _ := heightOfCoords(lines, x)
	left := Coords{x.row - 1, x.col}
	right := Coords{x.row + 1, x.col}
	up := Coords{x.row, x.col - 1}
	down := Coords{x.row, x.col + 1}

	if v, err := heightOfCoords(lines, left); err == nil && canMove(y, v) {
		next = append(next, left)
	}
	if v, err := heightOfCoords(lines, right); err == nil && canMove(y, v) {
		next = append(next, right)
	}
	if v, err := heightOfCoords(lines, up); err == nil && canMove(y, v) {
		next = append(next, up)
	}
	if v, err := heightOfCoords(lines, down); err == nil && canMove(y, v) {
		next = append(next, down)
	}
	return next
}

func heightOfCoords(lines []string, x Coords) (byte, error) {
	if x.row < 0 || x.row >= len(lines) || x.col < 0 || x.col >= len(lines[x.row]) {
		return '?', errors.New("out of bounds")
	}
	return lines[x.row][x.col], nil
}

func part1(input string) int {
	lines, start, end := parseInput(input, 'S', 'E')
	canMove := func(from, to byte) bool {
		return int(to)-int(from) <= 1
	}
	isEnd := func(b Coords) bool {
		return b.row == end.row && b.col == end.col
	}
	return bfs(lines, start, isEnd, canMove)
}

func part2(input string) int {
	lines, start, _ := parseInput(input, 'E', 'S')
	canMove := func(from, to byte) bool {
		return int(from)-int(to) <= 1
	}
	isEnd := func(c Coords) bool {
		b, _ := heightOfCoords(lines, c)
		return b == 'a'
	}
	return bfs(lines, start, isEnd, canMove)
}

func main() {
	data := utils.ReadFile("input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}
