package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type command struct {
	direction string
	step      int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay02Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay02Part2(string(input)))
}

// SolveDay02Part1 returns the depth multiplied by the horizontal position of the submarine, after following the steps
func SolveDay02Part1(input string) int {
	c := convertInputInCommands(input)
	var horizontal, depth int
	for i := range c {
		switch c[i].direction {
		case "forward":
			horizontal += c[i].step

		case "down":
			depth += c[i].step

		case "up":
			depth -= c[i].step
		}
	}
	return horizontal * depth
}

// SolveDay02Part2 returns the depth multiplied by the horizontal position of the submarine,
// after following the steps with no direct up and down use an aim instead
func SolveDay02Part2(input string) int {
	c := convertInputInCommands(input)
	var aim, horizontal, depth int
	for i := range c {
		switch c[i].direction {
		case "forward":
			horizontal += c[i].step
			depth += c[i].step * aim

		case "down":
			aim += c[i].step

		case "up":
			aim -= c[i].step
		}
	}
	return horizontal * depth
}

func convertInputInCommands(input string) []command {
	lines, _ := utils.InputToSlice(input)
	var c []command
	for i := range lines {
		line := strings.Split(lines[i], " ")
		if len(line) != 2 {
			panic("do not end here")
		}
		stepInt, _ := strconv.Atoi(line[1])
		c = append(c, command{
			direction: line[0],
			step:      stepInt,
		})
	}
	return c
}
