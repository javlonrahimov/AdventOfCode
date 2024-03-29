package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay11(t *testing.T) {

	data := utils.ReadFile("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(data)
		var want = 31

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(data)
		var want = 29

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
