package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dergeberl/aoc/utils"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay01Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay01Part2(string(input)))
}

//SolveDay01Part1
func SolveDay01Part1(input string) int {
	inputInt, _ := utils.InputToIntSlice(input)
	return calculateIncrements(inputInt)
}

//SolveDay01Part2
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
