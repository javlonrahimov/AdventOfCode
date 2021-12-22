package main

import (
	"javlonrahimov/AdventOfCode/utils"
	"testing"
)

func TestDay9(t *testing.T) {
	lines := utils.ReadFileLineByLine("input_test.txt")

	got := part1(lines)
	want := 15

	if got != want {
		t.Errorf("want '%d' got '%d", want, got)
	}
}
