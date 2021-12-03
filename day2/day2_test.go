package main

import (
	utils "javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay2(t *testing.T) {

	lines := utils.ReadFileLineByLine("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(lines)
		want := 150

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(lines)
		want := 900

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
