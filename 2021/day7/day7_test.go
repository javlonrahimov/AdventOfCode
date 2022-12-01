package main

import (
	"javlonrahimov/AdventOfCode/2021/utils"
	"testing"
)

func TestDay7(t *testing.T) {

	input := utils.ReadFile("input_test.txt")

	t.Run("part 1", func(t *testing.T) {
		got := part1(input)
		want := 37

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		// got := part2(input)
		// want := 168

		// if got != want {
		// 	t.Errorf("got %d want %d", got, want)
		// }
	})
}
