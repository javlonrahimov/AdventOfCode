package main

import (
	"encoding/json"
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"math"
	"strings"
)

func parseLine(input string) []interface{} {
	var distressSignal []interface{}
	err := json.Unmarshal([]byte(input), &distressSignal)
	if err != nil {
		panic("cannot parse")
	}
	return distressSignal
}

func isSlice(input interface{}) bool {
	switch input.(type) {
	case []interface{}:
		return true
	}
	return false
}

func compare(left, right interface{}) int {
	isLeftSlice, is1Slice := isSlice(left), isSlice(right)

	if !isLeftSlice && is1Slice {
		return compare([]interface{}{left}, right)
	} else if isLeftSlice && !is1Slice {
		return compare(left, []interface{}{right})
	}

	leftSlice, rightSlice := left.([]interface{}), right.([]interface{})
	minLen := int(math.Min(float64(len(leftSlice)), float64(len(rightSlice))))

	for i := 0; i < minLen; i++ {
		leftItem, rightItem := leftSlice[i], rightSlice[i]

		if isSlice(leftItem) || isSlice(rightItem) {
			if result := compare(leftItem, rightItem); result != 0 {
				return result
			}
			continue
		}

		if leftItem.(float64) > rightItem.(float64) {
			return -1
		}

		if leftItem.(float64) < rightItem.(float64) {
			return 1
		}
	}

	if len(rightSlice) == len(leftSlice) {
		return 0
	}

	if len(rightSlice) == minLen && len(leftSlice) != minLen {
		return -1
	}

	return 1
}

func part1(signalPacketPairs []string) int {
	sum := 0
	i := 1

	for _, pairs := range signalPacketPairs {
		split := strings.Split(pairs, "\n")
		left, right := parseLine(split[0]), parseLine(split[1])

		if compare(left, right) == 1 {
			sum += i
		}
		i++
	}
	return sum
}

func part2(signalPackets []string) int {
	two := parseLine("[[2]]")
	six := parseLine("[[6]]")

	two1Order, six1Order := 1, 1

	for _, packet := range signalPackets {
		if packet == "" {
			continue
		}

		transmission := parseLine(packet)

		if compare(transmission, two) == 1 {
			two1Order++
		}
		if compare(transmission, six) == 1 {
			six1Order++
		}
	}

	return two1Order * (six1Order + 1)
}

func main() {
	packetPairs := strings.Split(utils.ReadFile("input.txt"), "\n\n")
	fmt.Println(part1(packetPairs))

	packets := strings.Split(utils.ReadFile("input.txt"), "\n")
	fmt.Println(part2(packets))
}
