package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay7(t *testing.T) {

	t.Run("part 1", func(t *testing.T) {
		lines := utils.ReadFileLineByLine("input_test.txt")
		got := part1(lines)
		var want int64 = 95437

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		lines := utils.ReadFileLineByLine("input_test.txt")
		got := part2(lines)
		var want = 24933642

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
