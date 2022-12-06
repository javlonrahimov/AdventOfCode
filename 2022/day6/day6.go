package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
)

func main() {
	stream := utils.ReadFile("input.txt")
	fmt.Println(part1(stream))
	fmt.Println(part2(stream))
}

func getWindow(input string, length int) []rune {
	runes := []rune(input)
	window := make([]rune, length)
	for i := 0; i < length; i++ {
		window[i] = runes[i]
	}
	return window
}

func allDifferent(window []rune) bool {
	hashSet := make(map[rune]struct{})
	for _, value := range window {
		hashSet[value] = struct{}{}
	}
	return len(window) == len(hashSet)
}

func part1(stream string) int {
	return solve(stream, 4)
}

func part2(stream string) int {
	return solve(stream, 14)
}

func solve(stream string, length int) int {
	window := getWindow(stream, length)
	for i, v := range stream[length:] {
		if allDifferent(window) {
			return i + length
		}
		window[i%length] = v
	}
	return len(stream)
}
