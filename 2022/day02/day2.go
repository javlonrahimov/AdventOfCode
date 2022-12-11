package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
)

func main() {
	rounds := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(rounds))
	fmt.Println(part2(rounds))

}

func part1(rounds []string) int {
	var score = 0
	for _, round := range rounds {
		switch round[2] {
		case 'X':
			score += 1
		case 'Y':
			score += 2
		case 'Z':
			score += 3
		}
		score += roundOutcome(round[0], round[2])
	}
	return score
}

func part2(rounds []string) int {
	var score = 0
	for _, round := range rounds {
		switch nextMove(round[0], round[2]) {
		case 'X':
			score += 1
		case 'Y':
			score += 2
		case 'Z':
			score += 3
		}
		score += roundOutcome(round[0], nextMove(round[0], round[2]))
	}
	return score
}

func roundOutcome(opponentMove, selfMove uint8) int {
	if opponentMove == 'A' {
		switch selfMove {
		case 'X':
			return 3
		case 'Y':
			return 6
		case 'Z':
			return 0
		}
	} else if opponentMove == 'B' {
		switch selfMove {
		case 'X':
			return 0
		case 'Y':
			return 3
		case 'Z':
			return 6
		}
	} else if opponentMove == 'C' {
		switch selfMove {
		case 'X':
			return 6
		case 'Y':
			return 0
		case 'Z':
			return 3
		}
	}
	return 0
}

func nextMove(opponentMove, hint uint8) uint8 {
	if hint == 'X' {
		switch opponentMove {
		case 'A':
			return 'Z'
		case 'B':
			return 'X'
		case 'C':
			return 'Y'
		}
	} else if hint == 'Y' {
		switch opponentMove {
		case 'A':
			return 'X'
		case 'B':
			return 'Y'
		case 'C':
			return 'Z'
		}
	} else if hint == 'Z' {
		switch opponentMove {
		case 'A':
			return 'Y'
		case 'B':
			return 'Z'
		case 'C':
			return 'X'
		}
	}
	return 'A'
}
