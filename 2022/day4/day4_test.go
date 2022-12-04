package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay4(t *testing.T) {
	data := utils.ReadFileLineByLine("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(data)
		var want = 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(data)
		var want = 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
