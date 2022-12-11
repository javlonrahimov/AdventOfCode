package main

import (
	"fmt"
	"javlonrahimov/AdventOfCode/utils"
	"strconv"
	"strings"
)

type CPU struct {
	cycle                int
	value                int
	lastCalculationCycle int
}

func (c *CPU) noop(onCycle func(cycle int)) {
	c.cycle++
	onCycle(c.cycle)
}

func (c *CPU) addx(value int, onCycle func(cycle int)) {
	c.noop(func(cycle int) {
		onCycle(c.cycle)
	})
	c.noop(func(cycle int) {
		onCycle(c.cycle)
	})
	c.value += value
}

func (c *CPU) calculateSignalStrength() int {
	c.lastCalculationCycle = c.cycle
	return c.value * c.cycle
}

func part1(lines []string) int {
	totalSignalStrength := 0
	var cpu = &CPU{
		cycle: 0,
		value: 1,
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "addx") {
			value := utils.ToInt(strings.Split(line, " ")[1])
			cpu.addx(value, func(cycle int) {
				if cycle == 20 || cycle-cpu.lastCalculationCycle == 40 {
					totalSignalStrength += cpu.calculateSignalStrength()
				}
			})
		} else {
			cpu.noop(func(cycle int) {
				if cycle == 20 || cycle-cpu.lastCalculationCycle == 40 {
					totalSignalStrength += cpu.calculateSignalStrength()
				}
			})
		}
	}

	return totalSignalStrength
}

func incrementAndControl(cpu *CPU) {
	if (cpu.cycle)%40 == 0 && cpu.cycle <= 220 {
		fmt.Println()
	}
	if cpu.value-1 == cpu.cycle%40 || cpu.value == cpu.cycle%40 || cpu.value+1 == cpu.cycle%40 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	cpu.cycle++
}

func part2(lines []string) {
	cpu := &CPU{
		cycle: 0,
		value: 1,
	}
	for _, line := range lines {
		operations := strings.Split(line, " ")
		incrementAndControl(cpu)
		if operations[0] == "addx" {
			value, _ := strconv.Atoi(operations[1])
			incrementAndControl(cpu)
			cpu.value += value
		}
	}
}

func main() {
	lines := utils.ReadFileLineByLine("input.txt")
	fmt.Println(part1(lines))
	part2(lines)
}
