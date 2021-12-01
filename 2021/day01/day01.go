package main

import (
	"fmt"
	"os"

	"github.com/dergeberl/aoc/utils"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay01Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay01Part2(string(input)))
}

//SolveDay01Part1 returns the sum of numbers that are higher as the numbers before
func SolveDay01Part1(input string) int {
	inputInt, _ := utils.InputToIntSlice(input)
	return calculateIncrements(inputInt)
}

//SolveDay01Part2 returns the sum if numbers that are higher as the numbers before after always sum 3 numbers together
func SolveDay01Part2(input string) int {
	inputInt, _ := utils.InputToIntSlice(input)
	avg := make([]int, len(inputInt)-2)
	for i := range avg {
		avg[i] += inputInt[i]
		avg[i] += inputInt[i+1]
		avg[i] += inputInt[i+2]
	}
	return calculateIncrements(avg)
}

func calculateIncrements(input []int) int {
	var increments int
	for i := range input {
		if i == 0 {
			continue
		}
		if input[i] > input[i-1] {
			increments++
		}
	}
	return increments
}
