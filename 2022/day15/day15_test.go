package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay15(t *testing.T) {

	lines := utils.ReadFileLineByLine("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(lines, 10)
		want := 26

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(lines, 20)
		var want int64 = 56000011

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
