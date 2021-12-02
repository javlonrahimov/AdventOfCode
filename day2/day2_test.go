package main

import (
	"testing"
	utils "javlonrahimov/AdventOfCode/utils"
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

}
