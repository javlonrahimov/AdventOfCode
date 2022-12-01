package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/2021/utils"
	"strings"
)

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))
}

func part1(lines []string) int {
	t := 0
	for _, k := range lines {
		b := strings.Split(k, " | ")[1]
		for _, j := range strings.Split(b, " ") {
			if len(j) == 2 || len(j) == 3 || len(j) == 4 || len(j) == 7 {
				t++
			}
		}
	}
	return t
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
