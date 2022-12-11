package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strconv"
	"strings"
)

type Move struct {
	amount int
	from   int
	to     int
}

func loadCrates(lines []string) [][]rune {
	numStacks := (len(lines[0]) + 1) / 4
	stacks := make([][]rune, numStacks)

	for _, line := range lines {
		if strings.Contains(line, "1") {
			break
		}
		for i := 0; i < numStacks; i++ {
			index := i*4 + 1
			if line[index] != ' ' {
				stacks[i] = append([]rune{[]rune(line)[index]}, stacks[i]...)
			}
		}
	}
	return stacks
}

func parseMoves(lines []string) (moves []Move) {
	for _, line := range lines {
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])
		moves = append(moves, Move{amount, from - 1, to - 1})
	}
	return moves
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	data := utils.ReadFile("input.txt")
	parts := strings.Split(data, "\n\n")
	linesStacks := strings.Split(parts[0], "\n")
	linesMoves := strings.Split(parts[1], "\n")
	stacks := loadCrates(linesStacks)
	moves := parseMoves(linesMoves)

	fmt.Println(part1(stacks, moves))

	stacks = loadCrates(linesStacks)
	moves = parseMoves(linesMoves)
	fmt.Println(part2(stacks, moves))
}

func part1(stacks [][]rune, moves []Move) string {
	res := ""
	for _, move := range moves {
		moving := stacks[move.from][len(stacks[move.from])-move.amount:]
		moving = reverse(moving)

		for _, crate := range moving {
			stacks[move.from] = stacks[move.from][:len(stacks[move.from])-1]
			stacks[move.to] = append(stacks[move.to], crate)
		}
	}

	for _, stack := range stacks {
		res += string(stack[len(stack)-1])
	}

	return res
}

func part2(stacks [][]rune, moves []Move) string {
	res := ""
	for _, move := range moves {
		moving := stacks[move.from][len(stacks[move.from])-move.amount:]

		for _, crate := range moving {
			stacks[move.from] = stacks[move.from][:len(stacks[move.from])-1]
			stacks[move.to] = append(stacks[move.to], crate)
		}
	}

	for _, stack := range stacks {
		res += string(stack[len(stack)-1])
	}

	return res
}
