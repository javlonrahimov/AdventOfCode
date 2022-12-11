package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strings"
)

func main() {
	bags := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(bags))
	fmt.Println(part2(bags))
}

func part1(input []string) int {
	score := 0
	for _, r := range input {
		first := strings.Split(r[:len(r)/2], "")
		second := strings.Split(r[len(r)/2:], "")
		charMap := make(map[string]bool)
		var commonChar string
		for _, char := range first {
			charMap[char] = true
		}
		for _, char := range second {
			if _, found := charMap[char]; found {
				commonChar = char
				break
			}
		}
		score += getPriorityValue(commonChar)
	}
	return score
}

func part2(input []string) int {
	score := 0
	for i := 0; i+2 < len(input); i += 3 {
		first := strings.Split(input[i], "")
		second := strings.Split(input[i+1], "")
		third := strings.Split(input[i+2], "")
		charMap := make(map[string]int)
		var commonChar string
		for _, char := range first {
			charMap[char] = 1
		}
		for _, char := range second {
			if _, found := charMap[char]; found {
				charMap[char] = 2
			}
		}
		for _, char := range third {
			if num := charMap[char]; num == 2 {
				commonChar = char
				break
			}
		}

		score += getPriorityValue(commonChar)
	}
	return score
}

func getPriorityValue(s string) int {
	ascii := utils.ToASCIICode(s)
	if ascii >= utils.ToASCIICode("a") {
		ascii = ascii - utils.ToASCIICode("a") + 1
	} else {
		ascii = ascii - utils.ToASCIICode("A") + 27
	}
	return ascii
}
