package main

import (
	"javlonrahimov/AdventOfCode/2021/utils"
	"testing"
)

func TestDay6(t *testing.T) {
	lines := utils.ReadFile("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(lines, 18)
		want := 26

		if got != want {
			t.Errorf("got %d want %d", got, want)

		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(lines)
		want := 26984457539

		if got != want {
			t.Errorf("got %d want %d", got, want)

		}
	})
}
