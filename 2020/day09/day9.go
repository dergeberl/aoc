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
	fmt.Printf("Part 1: %v\n", SolveDay9Part1(intListToSlice(input), 25))
	fmt.Printf("Part 2: %v\n", SolveDay9Part2(intListToSlice(input), 25))
}

//SolveDay9Part1 returns the first value that can not reached with a sum of 2 values of the last $start values
func SolveDay9Part1(i []int, start int) (s int) {
	cur := start
	if start > len(i) || start < 2 {
		return 0
	}
	for {
		found := false
		if len(i) == cur {
			break
		}
		for _, num := range i[cur-start : cur] {
			if found {
				break
			}
			for _, num2 := range i[cur-start : cur] {
				if num+num2 == i[cur] && num != num2 {
					found = true
					break
				}
			}
		}
		if found {
			cur++
			continue
		}
		return i[cur]
	}
	return 0
}

//SolveDay9Part2 search for 5 numbers that sum is the value from SolveDay9Part1 and sum the lowest and highest
func SolveDay9Part2(i []int, start int) (s int) {
	numToSolve := SolveDay9Part1(i, start)
	var index, currentIndex, currentSum int
	var currentNumbers []int
	for {
		currentSum += i[currentIndex]
		currentNumbers = append(currentNumbers, i[currentIndex])
		if currentSum == numToSolve {
			break
		}
		if currentSum > numToSolve {
			index++
			currentIndex = index
			currentSum = 0
			currentNumbers = []int{}
			continue
		}
		currentIndex++
	}
	sort.Ints(currentNumbers)
	return currentNumbers[0] + currentNumbers[len(currentNumbers)-1]

}

//Helper functions
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
