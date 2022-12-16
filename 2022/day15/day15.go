package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"regexp"
)

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

type Coors struct {
	x, y int
}

func (p Coors) Equal(o Coors) bool {
	return p.x == o.x && p.y == o.y
}

func manhattan(a, b Coors) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

type Sensor struct {
	pos, beacon Coors
	d           int
}

func (s Sensor) Contains(p Coors) bool {
	return s.d >= manhattan(s.pos, p)
}

func parseInput(lines []string) []Sensor {
	var out []Sensor
	for _, line := range lines {
		re := regexp.MustCompile("-?[0-9]+")
		numbers := re.FindAllString(line, -1)
		xSensor, ySensor, xBeacon, yBeacon := utils.ToInt(numbers[0]), utils.ToInt(numbers[1]), utils.ToInt(numbers[2]), utils.ToInt(numbers[3])
		sp := Coors{xSensor, ySensor}
		bp := Coors{xBeacon, yBeacon}
		out = append(out, Sensor{
			sp,
			bp,
			manhattan(sp, bp),
		})
	}
	return out
}

func part1(lines []string, y int) int {
	sensors := parseInput(lines)
	xMin := 99_999_999
	xMax := 0
	for _, s := range sensors {
		if s.pos.x-s.d < xMin {
			xMin = s.pos.x - s.d
		}
		if s.pos.x+s.d > xMax {
			xMax = s.pos.x + s.d
		}
	}

	covered := 0
	testPos := Coors{xMin, y}
	for x := xMin; x <= xMax; x++ {
		testPos.x = x
		for _, s := range sensors {
			if s.Contains(testPos) && !s.pos.Equal(testPos) && !s.beacon.Equal(testPos) {
				covered++
				break
			}
		}
	}

	return covered
}

func part2(lines []string, maxCoordinates int) int64 {
	sensors := parseInput(lines)
	notDetectedBeaconCoors := Coors{
		x: 0,
		y: 0,
	}
	for y := 0; y < maxCoordinates; y++ {
	loop:
		for x := 0; x < maxCoordinates; x++ {
			notDetectedBeaconCoors.x = x
			notDetectedBeaconCoors.y = y
			for _, sensor := range sensors {
				if sensor.Contains(notDetectedBeaconCoors) {
					x += sensor.d - abs(sensor.pos.y-y) + abs(sensor.pos.x-x)
					continue loop
				}
			}
			return int64(x*4_000_000) + int64(y)
		}
	}
	return 0
}

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines, 2_000_000))
	fmt.Println(part2(lines, 4_000_000))
}
