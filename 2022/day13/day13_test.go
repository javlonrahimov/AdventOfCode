package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"strings"
	"testing"
)

func TestDay13(t *testing.T) {

	packetPairs := strings.Split(utils.ReadFile("input_test.txt"), "\n\n")
	packets := strings.Split(utils.ReadFile("input_test.txt"), "\n")

	t.Run("part 1", func(t *testing.T) {
		got := part1(packetPairs)
		want := 13

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		got := part2(packets)
		want := 140

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
