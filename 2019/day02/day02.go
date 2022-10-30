package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay02Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay02Part2(string(input)))
}

// SolveDay02Part1
func SolveDay02Part1(input string) int {
	return runWithReplace(input, 12, 2)
}

// SolveDay02Part2
func SolveDay02Part2(input string) int {
	_, _ = utils.InputToIntSlice(input)
	var solution int
	for a := 0; a <= 99; a++ {
		for b := 0; b <= 99; b++ {
			if runWithReplace(input, a, b) == 19690720 {
				solution = a*100 + b
			}
		}
	}
	return solution
}

func runWithReplace(input string, a, b int) int {
	inputInt, _ := utils.InputToIntSlice(strings.ReplaceAll(strings.ReplaceAll(input, "\n", ""), ",", "\n"))
	position := 0
	inputInt[1] = a
	inputInt[2] = b
	for inputInt[position] != 99 {
		switch inputInt[position] {
		case 1:
			// add
			inputInt[inputInt[position+3]] = inputInt[inputInt[position+1]] + inputInt[inputInt[position+2]]
		case 2:
			// multiply
			inputInt[inputInt[position+3]] = inputInt[inputInt[position+1]] * inputInt[inputInt[position+2]]
		}
		position += 4
	}
	return inputInt[0]
}
