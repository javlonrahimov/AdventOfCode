package main

import (
	"testing"
)

func TestDay6(t *testing.T) {

	t.Run("part 1", func(t *testing.T) {
		tests := []struct {
			input string
			want  int
		}{
			{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 7},
			{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 5},
			{input: "nppdvjthqldpwncqszvftbrmjlhg", want: 6},
			{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 10},
			{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 11},
		}
		for _, tc := range tests {
			got := part1(tc.input)
			var want = tc.want

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})

	t.Run("part 2", func(t *testing.T) {
		tests := []struct {
			input string
			want  int
		}{
			{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 19},
			{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 23},
			{input: "nppdvjthqldpwncqszvftbrmjlhg", want: 23},
			{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 29},
			{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 26},
		}
		for _, tc := range tests {
			got := part2(tc.input)
			var want = tc.want

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		}
	})

}
