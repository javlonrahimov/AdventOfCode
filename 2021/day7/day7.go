package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"math"
)

func main() {
	input := utils.ReadFile("input.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	list := utils.ParseStringToList(input)
	q := math.MaxInt64

	for i := findMin(list); i <= findMax(list); i++ {
		temp := absDistanseList(list, i)
		q = int(math.Min(float64(q), float64(temp)))
	}

	return q
}

// todo: optimize
func part2(input string) int {
	list := utils.ParseStringToList(input)
	q := math.MaxInt64

	for i := findMin(list); i <= findMax(list); i++ {
		temp := absDistanseList(list, i)
		temp = temp * (temp + 1) / 2
		q = int(math.Min(float64(q), float64(temp)))
	}

	return q
}

func findMin(list []int) int {
	min := math.MaxInt64
	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min
}

func findMax(list []int) int {
	max := math.MinInt64
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}

func absDistanseList(list []int, pos int) int {
	sum := 0
	for _, v := range list {
		sum += absDistanse(v, pos)
	}
	return sum
}

func absDistanse(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}
