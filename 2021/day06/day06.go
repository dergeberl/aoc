package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type fishes []int

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay06Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay06Part2(string(input)))
}

//SolveDay06Part1 returns the number of lantern fishes after 80 days
func SolveDay06Part1(input string) int {
	fish := getNumbersFromList(input)
	return fish.calcLanternFish(80)
}

//SolveDay06Part2 returns the number of lantern fishes after 256 days
func SolveDay06Part2(input string) int {
	fish := getNumbersFromList(input)
	return fish.calcLanternFish(256)
}

// calcLanternFish2 returns the number of fishes after x days
func (f fishes) calcLanternFish(days int) int {
	states := make([]int, 9)
	for i := range f {
		states[f[i]]++
	}
	for i := 1; i <= days; i++ {
		birthFish := states[0]
		states[0] = states[1]
		states[1] = states[2]
		states[2] = states[3]
		states[3] = states[4]
		states[4] = states[5]
		states[5] = states[6]
		states[6] = states[7] + birthFish
		states[7] = states[8]
		states[8] = birthFish
	}
	return utils.SumSlice(states)
}

// getNumbersFromList returns the fishes from an input
func getNumbersFromList(input string) fishes {
	numbers := strings.Split(input, ",")
	var output []int
	for i := range numbers {
		tmpInt, _ := strconv.Atoi(numbers[i])
		output = append(output, tmpInt)
	}
	return output
}
