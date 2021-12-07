package main

import (
	"fmt"
	"github.com/dergeberl/aoc/utils"
	"os"
	"strconv"
	"strings"
)

type crabs []int

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay07Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay07Part2(string(input)))
}

//SolveDay07Part1 returns the sum of fuel for all submarines on the optimal position
func SolveDay07Part1(input string) int {
	c := getNumbersFromList(input)
	return c.getSumOfNeededFuel(false)
}

//SolveDay07Part2 returns the sum of fuel for all submarines on the optimal position
func SolveDay07Part2(input string) int {
	c := getNumbersFromList(input)
	return c.getSumOfNeededFuel(true)
}

// getSumOfNeededFuel returns the sum of fuel for all submarines on the optimal position
func (c crabs) getSumOfNeededFuel(crabVersion bool) int {
	l, h := c.getRange()
	fuel := len(c) * getFuelConsumptionCrab(l, h)
	for i := l; i <= h; i++ {
		tmpFuel := 0
		for cr := range c {
			if crabVersion {
				tmpFuel += getFuelConsumptionCrab(i, c[cr])
				continue
			}
			tmpFuel += utils.GetDiff(i, c[cr])

		}
		if tmpFuel < fuel {
			fuel = tmpFuel
		}
		if tmpFuel > fuel {
			break
		}
	}

	return fuel
}

// getRange returns the lowest and highest position of the crab submarines
func (c crabs) getRange() (int, int) {
	var height, low int
	for i := range c {
		if height < i {
			height = i
		}
	}
	low = height
	for i := range c {
		if low > i {
			low = i
		}
	}

	return low, height
}

// getFuelConsumptionCrab returns the fuel consumption as crab engineering
func getFuelConsumptionCrab(a, b int) int {
	d := utils.GetDiff(a, b)
	var s int
	for i := 1; i <= d; i++ {
		s += i
	}
	return s
}

// getNumbersFromList returns the crab submarines
func getNumbersFromList(input string) crabs {
	numbers := strings.Split(input, ",")
	var output []int
	for i := range numbers {
		tmpInt, _ := strconv.Atoi(numbers[i])
		output = append(output, tmpInt)
	}
	return output
}
