package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay1(t *testing.T) {

	lines := utils.ReadFileLineByLine("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(lines)
		var want int64 = 24000

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(lines)
		var want int64 = 45000

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
