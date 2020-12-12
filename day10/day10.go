package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

//wayCache is used as a cache for a memoized recursion
var wayCache map[int]int64

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay10Part1(intListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay10Part2(intListToSlice(input)))
	fmt.Printf("Part 2 Recursive: %v\n", SolveDay10Part2Recursive(intListToSlice(input)))
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
func SolveDay10Part2(input []int) int64 {
	sort.Ints(input)
	counter := make(map[int]int64)
	counter[0] = 1
	var lastVolt int
	for _, volt := range input {
		counter[volt] = counter[volt-1] + counter[volt-2] + counter[volt-3]
		if counter[volt] != 0 {
			lastVolt = volt
		}
	}
	return counter[lastVolt]
}

//SolveDay10Part2Recursive returns the total number of distinct ways but in a recursive method with caching
func SolveDay10Part2Recursive(input []int) int64 {
	input = append(input, 0)
	sort.Ints(input)
	wayCache = make(map[int]int64)
	return countWays(input)
}

//countWays counts the way of distinct ways recursive
func countWays(input []int) (sum int64) {
	if wayCache[input[0]] != 0 {
		return wayCache[input[0]]
	}
	for i := 1; i <= 3 && i < len(input); i++ {
		if input[i]-3 <= input[0] {
			sum += countWays(input[i:])
		}
	}
	if sum == 0 {
		sum = 1
	}
	wayCache[input[0]] = sum
	return sum
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
