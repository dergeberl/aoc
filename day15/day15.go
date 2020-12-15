package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay15Part1(input))
	fmt.Printf("Part 2: %v\n", SolveDay15Part2(input))
}

//SolveDay15Part1 plays the Elves memory game and returns the 2020th spoken number
func SolveDay15Part1(i string) (s int) {
	return playElvesGame(i, 2020)
}

//SolveDay15Part2 plays the Elves memory game and returns the 30000000th spoken number
func SolveDay15Part2(i string) (s int) {
	return playElvesGame(i, 30000000)
}

//playElvesGame plays the Elves game until the given iteration
func playElvesGame(input string, iteration int) int {
	said := []int{0}
	lastTime, beforeLastTime := make(map[int]int), make(map[int]int)
	for i, numbers := range strings.Split(input, ",") {
		say, err := strconv.Atoi(numbers)
		if err != nil {
			return 0
		}
		said = append(said, say)
		beforeLastTime[say] = lastTime[say]
		lastTime[say] = i + 1
	}

	for i := len(said); i <= iteration; i++ {
		last := said[i-1]
		say := 0
		if beforeLastTime[last] != 0 {
			say = lastTime[last] - beforeLastTime[last]
		}
		said = append(said, say)
		beforeLastTime[say] = lastTime[say]
		lastTime[say] = i
	}
	return said[iteration]
}

//Helper functions
//stringListToSlice converts the list of strings (each string one row) to a slice
func stringListToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}

//intListToSlice converts the list of numbers (each number one row) to a slice
func intListToSlice(list string) (i []int) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			return nil
		}
		i = append(i, lineInt)
	}
	return
}
