package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay10Part1(intListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay10Part2(intListToSlice(input)))
}

//SolveDay10Part1 returns the product of the possible 1 and 3 volts steps
func SolveDay10Part1(input []int) int {
	sort.Ints(input)
	result := make(map[int]int)
	for i, volt := range input {
		if i == 0 {
			result[volt]++
			continue
		}
		result[volt-input[i-1]]++
	}
	result[3]++
	return result[1] * result[3]
}

//SolveDay10Part2 returns the total number of distinct ways
func SolveDay10Part2(index []int) int {
	sort.Ints(index)
	counter := make(map[int]int)
	counter[0] = 1
	var lastVolt int
	for _, volt := range index {
		counter[volt] = counter[volt-1] + counter[volt-2] + counter[volt-3]
		if counter[volt] != 0 {
			lastVolt = volt
		}
	}
	return counter[lastVolt]
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
