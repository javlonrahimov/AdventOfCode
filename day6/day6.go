package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
)

func main() {
	lines := utils.ReadFile("input.txt")
	fmt.Println(part1(lines, 80))
	fmt.Println(part2(lines))
}

func part1(input string, days int) int {
	list := utils.ParseStringToList(input)

	for i := 0; i < days; i++ {
		addCount := 0
		for j := 0; j < len(list); j++ {
			if list[j] > 0 {
				list[j]--
			} else {
				list[j] = 6
				addCount++
			}
		}
		for j := 0; j < addCount; j++ {
			list = append(list, 8)
		}
	}
	return len(list)
}

func part2(input string) int {

	school := utils.ParseStringToList(input)

	counts := map[int]int{}
	for _, timer := range school {
		counts[timer] += 1
	}

	for day := 1; day <= 256; day++ {
		// school = util.FlatMap(school, AdvanceOne)
		nextCounts := map[int]int{}
		for timer, count := range counts {
			if timer == 0 {
				nextCounts[6] += count
				nextCounts[8] += count
			} else {
				nextCounts[timer-1] += count
			}
		}
		counts = nextCounts

		schoolSize := 0
		for _, count := range counts {
			schoolSize += count
		}

		if day == 256 {
			return schoolSize
		}
	}
	return 0
}
