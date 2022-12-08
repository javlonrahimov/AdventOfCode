package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strings"
)

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	trees := createGrid(lines)

	fmt.Println(part1(trees))
	fmt.Println(part2(trees))
}

func createGrid(lines []string) [][]int {
	grid := make([][]int, 0)
	for _, line := range lines {
		grid = append(grid, utils.Map(strings.Split(line, ""), func(t string) int {
			return utils.ToInt(t)
		}))
	}
	return grid
}

func createEmptyGrid(trees [][]int) [][]bool {
	visibleTrees := make([][]bool, len(trees))
	for i := range visibleTrees {
		visibleTrees[i] = make([]bool, len(trees[0]))
	}
	return visibleTrees
}

func numberOfVisibleTrees(visibleTrees [][]bool) int {
	count := 0
	for _, line := range visibleTrees {
		for _, isVisible := range line {
			if isVisible {
				count++
			}
		}
	}
	return count
}

func part1(trees [][]int) int {
	visibleTrees := createEmptyGrid(trees)
	for i, line := range trees {
		leftMax := -1
		rightMax := -1

		for j, tree := range line {
			if tree > leftMax {
				visibleTrees[i][j] = true
				leftMax = tree
			}
		}
		for j := len(line) - 1; j >= 0; j-- {
			if line[j] > rightMax {
				visibleTrees[i][j] = true
				rightMax = line[j]
			}
		}

	}

	for column := 0; column < len(visibleTrees[0]); column++ {

		leftMax := -1
		rightMax := -1

		for row := 0; row < len(visibleTrees); row++ {
			if trees[row][column] > leftMax {
				visibleTrees[row][column] = true
				leftMax = trees[row][column]
			}
		}

		for row := len(visibleTrees) - 1; row >= 0; row-- {
			if trees[row][column] > rightMax {
				visibleTrees[row][column] = true
				rightMax = trees[row][column]
			}
		}
	}

	return numberOfVisibleTrees(visibleTrees)
}

func rowScore(trees [][]int, x, y int) int {
	left, right := 0, 0
	t := trees[y][x]

	for i := x - 1; i >= 0; i-- {
		tree := trees[y][i]
		left++
		if tree >= t {
			break
		}
	}

	for i := x + 1; i < len(trees[0]); i++ {
		tree := trees[y][i]
		right++
		if tree >= t {
			break
		}
	}

	return left * right
}

func colScore(trees [][]int, x, y int) int {
	up, down := 0, 0
	t := trees[y][x]

	for i := y - 1; i >= 0; i-- {
		tree := trees[i][x]
		up++
		if tree >= t {
			break
		}
	}

	for i := y + 1; i < len(trees); i++ {
		tree := trees[i][x]
		down++
		if tree >= t {
			break
		}
	}

	return down * up
}

func calcScore(trees [][]int, x, y int) int {
	if y == 0 || y == len(trees)-1 || x == 0 || x == len(trees[0])-1 {
		return 0
	}
	return rowScore(trees, x, y) * colScore(trees, x, y)
}

func part2(trees [][]int) int {
	max := 0
	for i := range trees {
		for j := range trees[i] {
			if score := calcScore(trees, i, j); score > max {
				max = score
			}
		}
	}
	return max
}
