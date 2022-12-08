package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay8(t *testing.T) {

	lines := utils.ReadFileLineByLine("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		trees := createGrid(lines)
		got := part1(trees)
		var want = 21

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		trees := createGrid(lines)
		got := part2(trees)
		var want = 8

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
