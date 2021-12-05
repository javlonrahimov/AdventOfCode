package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay4(t *testing.T) {

	lines := utils.ReadFile("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(lines)
		want := 4512

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(lines)
		want := 1924

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
