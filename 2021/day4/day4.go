package main

import (
	"fmt"
	utils2 "javlonrahimov/AdventOfCode/2021/utils"
	"strings"
)

const (
	Row = 5
)

func main() {
	lines := utils2.ReadFile("input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func stringToStringList(input string) []string {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = strings.TrimSpace(line)
	}
	return result
}

func part1(input string) int {
	plays, boards := initBoard(input)
	var scores []int
	for i := range plays {
		for board := range boards {
			if !win(boards[board], plays[:i]) {
				continue
			}
			set := make(map[int]bool)
			for _, play := range plays[:i] {
				set[play] = true
			}

			var total int
			for _, row := range boards[board] {
				for _, col := range row {
					if set[col] {
						continue
					}
					total += col
				}
			}
			scores = append(scores, total*plays[i-1])
			delete(boards, board)
		}
	}
	return scores[0]
}

func part2(input string) int {
	plays, boards := initBoard(input)
	var scores []int
	for i := range plays {
		for board := range boards {
			if !win(boards[board], plays[:i]) {
				continue
			}
			set := make(map[int]bool)
			for _, play := range plays[:i] {
				set[play] = true
			}

			var total int
			for _, row := range boards[board] {
				for _, col := range row {
					if set[col] {
						continue
					}
					total += col
				}
			}
			scores = append(scores, total*plays[i-1])
			delete(boards, board)
		}
	}
	return scores[len(scores)-1]
}

func initBoard(input string) (plays []int, boards map[int][][]int) {
	lines := stringToStringList(input)
	boards = make(map[int][][]int)

	board := 0
	rows := 0
	for i, line := range lines {
		if i == 0 {
			play := strings.Split(strings.TrimSpace(line), ",")
			plays = make([]int, len(play))
			for j, s := range play {
				plays[j] = utils2.ToInt(s)
			}
			continue
		}
		if strings.TrimSpace(line) == "" {
			board++
			rows = 0
			boards[board] = make([][]int, Row)
			continue
		}
		columns := strings.Fields(line)
		boards[board][rows] = make([]int, len(columns))
		for k, col := range columns {
			boards[board][rows][k] = utils2.ToInt(col)
		}
		rows++
	}
	return
}

func win(board [][]int, numbers []int) bool {
	played := make(map[int]bool)
	for _, n := range numbers {
		played[n] = true
	}
	for i := 0; i < len(board); i++ {
		var rowMatch, colMatch int
		for j := 0; j < len(board[i]); j++ {
			if played[board[i][j]] {
				rowMatch++
			}
			if played[board[j][i]] {
				colMatch++
			}
		}
		if colMatch == len(board) || rowMatch == len(board) {
			return true
		}
	}
	return false
}
