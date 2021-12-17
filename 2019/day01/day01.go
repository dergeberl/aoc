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
	var solution int
	for _, i := range inputInt {
		solution += i/3 - 2
	}
	return solution
}

//SolveDay01Part2
func SolveDay01Part2(input string) int {
	inputInt, _ := utils.InputToIntSlice(input)
	var solution int
	for _, i := range inputInt {
		for {
			i = i/3 - 2
			if i <= 0 {
				break
			}
			solution += i
		}
	}
	return solution
}
