package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strconv"
)

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(input []string) int {

	gamma, epsilon := gamma_epsilon(input)

	return gamma * epsilon
}

func part2(input []string) int {

	o2, err := strconv.ParseInt(filter(input, true), 2, 64)
	if err != nil {
		fmt.Printf("error1 = %s \n", err)
	}
	co2, err := strconv.ParseInt(filter(input, false), 2, 64)
	if err != nil {
		fmt.Printf("error2 = %s \n", err)
	}

	return int(o2) * int(co2)
}

func gamma_epsilon(inputs []string) (int, int) {
	zeroes, ones := hist(inputs)

	gamma, epsilon := 0, 0
	for i := 0; i < len(zeroes); i++ {
		gamma = gamma << 1
		epsilon = epsilon << 1
		zeroes := zeroes[i]
		ones := ones[i]
		if zeroes > ones {
			epsilon++
		} else if zeroes <= ones {
			gamma++
		}
	}

	return gamma, epsilon
}

func filter(inputs []string, use_mcb bool) string {
	size := len(inputs[0])
	pos := 0
	candidates := make([]string, len(inputs))
	copy(candidates, inputs)
	for len(candidates) > 1 || pos > size {
		pass := make([]string, 0)
		gamma, epsilon := gamma_epsilon(candidates)
		if !use_mcb {
			gamma = epsilon
		}
		mask := (1 << (size - pos - 1))
		bit := strconv.Itoa((gamma & mask) >> (size - pos - 1))
		for _, candidate := range candidates {
			switch string(candidate[pos]) {
			case bit:
				pass = append(pass, candidate)
			default:
			}
		}
		candidates = pass

		pos++
	}
	return candidates[0]
}

func hist(strs []string) ([]int, []int) {
	size := len(strs[0])
	zeroes, ones := make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		for _, str := range strs {
			d, _ := strconv.Atoi(string(str[i]))
			zeroes[i] += 1 - d
			ones[i] += d
		}
	}
	return zeroes, ones
}
