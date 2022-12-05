package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"strings"
	"testing"
)

func TestDay5(t *testing.T) {
	data := utils.ReadFile("input_test.txt")
	parts := strings.Split(data, "\n\n")
	linesStacks := strings.Split(parts[0], "\n")
	linesMoves := strings.Split(parts[1], "\n")

	t.Run("part 1", func(t *testing.T) {
		stacks := loadCrates(linesStacks)
		moves := parseMoves(linesMoves)
		got := part1(stacks, moves)
		var want = "CMZ"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		stacks := loadCrates(linesStacks)
		moves := parseMoves(linesMoves)
		got := part2(stacks, moves)
		var want = "MCD"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
