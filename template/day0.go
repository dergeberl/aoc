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
	fmt.Printf("Part 1: %v\n", SolveDay0Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay0Part2(string(input)))
}

//SolveDay0Part1
func SolveDay0Part1(input string) int {
	_, _ = utils.InputToIntSlice(input)
	var solution int

	return solution
}

//SolveDay0Part2
func SolveDay0Part2(input string) int {
	_, _ = utils.InputToIntSlice(input)
	var solution int

	return solution
}
