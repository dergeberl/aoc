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
	fmt.Printf("Part 1: %v\n", SolveDay13Part1(stringListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay13Part2(stringListToSlice(input)))
}

//SolveDay13Part1 calc the time to wait for the next bus, for a given timestamp (input line1) and a bus plan (input line2)
func SolveDay13Part1(input []string) int {
	if len(input) < 2 {
		return 0
	}
	timestamp, err := strconv.Atoi(input[0])
	if err != nil {
		return 0
	}
	curTimestamp := timestamp
	var busNumbers []int
	for _, number := range strings.Split(input[1], ",") {
		num, _ := strconv.Atoi(number)
		if num != 0 {
			busNumbers = append(busNumbers, num)
		}
	}
	if len(busNumbers) == 0 {
		return 0
	}
	for {
		for _, busNumber := range busNumbers {
			if curTimestamp%busNumber == 0 {
				return (curTimestamp - timestamp) * busNumber
			}
		}
		curTimestamp++
	}
}

//SolveDay13Part2 returns the timestamp of the first bus when all bus departure in the right order
func SolveDay13Part2(input []string) int {
	if len(input) < 2 {
		return 0
	}
	var busNumbers []int
	for _, number := range strings.Split(input[1], ",") {
		numberInt, err := strconv.Atoi(number)
		if err != nil || numberInt == 0 {
			busNumbers = append(busNumbers, 1)
			continue
		}
		busNumbers = append(busNumbers, numberInt)
	}
	startTime, jump := busNumbers[0], 1
	for i, busNumber := range busNumbers {
		for (startTime+i)%busNumber != 0 {
			startTime += jump
		}
		jump *= busNumber
	}
	return startTime
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
