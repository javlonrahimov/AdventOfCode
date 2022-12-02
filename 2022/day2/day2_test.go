package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay2(t *testing.T) {

	rounds := utils.ReadFileLineByLine("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(rounds)
		var want = 15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(rounds)
		var want = 12

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
