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
	fmt.Printf("Part 1: %v\n", SolveDay25Part1(intListToSlice(input)))
}

//SolveDay25Part1 returns the encryption key for the door and card
func SolveDay25Part1(input []int) (s int) {
	if len(input) != 2 {
		return 0
	}

	value, subjectNumber := 1, 7
	var loopSize, keySubjectNumber int

	for {
		loopSize++
		value *= subjectNumber
		value %= 20201227
		if value == input[0] {
			keySubjectNumber = input[1]
			break
		}
		if value == input[1] {
			keySubjectNumber = input[0]
			break
		}
	}

	value = 1
	for i := 0; i < loopSize; i++ {
		value *= keySubjectNumber
		value %= 20201227
	}

	return value
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
