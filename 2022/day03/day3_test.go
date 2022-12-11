package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay2(t *testing.T) {

	bags := utils.ReadFileLineByLine("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(bags)
		var want = 157

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(bags)
		var want = 70

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
