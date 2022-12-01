package main

import (
	"javlonrahimov/AdventOfCode/2021/utils"
	"testing"
)

func TestDay1(t *testing.T) {

	lines := convertStringListToIntList(utils.ReadFileLineByLine("input_test.txt"))

	t.Run("part 1", func(t *testing.T) {
		got := part1(lines)
		var want int64 = 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(lines)
		var want int64 = 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
