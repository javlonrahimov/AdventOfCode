package main

import (
	"javlonrahimov/AdventOfCode/2021/utils"
	"testing"
)

func TestDay3(t *testing.T) {

	lines := utils.ReadFileLineByLine("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(lines)
		want := 198

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(lines)
		want := 230
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
