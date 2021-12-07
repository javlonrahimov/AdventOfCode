package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay5(t *testing.T) {

	lines := utils.ReadFile("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(lines)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(lines)
		want := 12

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
