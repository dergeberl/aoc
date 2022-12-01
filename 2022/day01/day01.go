package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

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

// SolveDay01Part1 returns the number of calories from the elf with the most calories
func SolveDay01Part1(input string) int {
	calories := getCaloriesFromInput(input)
	sort.Ints(calories)
	return calories[len(calories)-1]
}

// SolveDay01Part2 returns the sum of calories from the 3 elf with the most calories
func SolveDay01Part2(input string) int {
	calories := getCaloriesFromInput(input)
	sort.Ints(calories)
	return utils.SumSlice(calories[len(calories)-3:])
}

func getCaloriesFromInput(input string) []int {
	line, _ := utils.InputToSlice(input)

	elf := 0
	var calories []int

	for i := range line {
		if line[i] == "" {
			elf++
			continue
		}
		if len(calories)-1 < elf {
			calories = append(calories, 0)
		}
		lintInt, err := strconv.Atoi(line[i])
		if err != nil {
			continue
		}
		calories[elf] += lintInt
	}
	return calories
}
