package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay9(t *testing.T) {

	t.Run("part 1", func(t *testing.T) {
		lines := utils.ReadFileLineByLine("input_test.txt")
		got := part1(lines)
		want := 13

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		lines := utils.ReadFileLineByLine("input_2_test.txt")
		got := part2(lines)
		want := 36

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
