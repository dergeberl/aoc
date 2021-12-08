package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type inputData struct {
	uniqueSignalPatterns []digit
	outputDigit          []digit
}

type digit string

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay08Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay08Part2(string(input)))
}

//SolveDay08Part1 returns the number of the digits that has an uniq segment count
func SolveDay08Part1(input string) int {
	data := parseInput(input)
	var sum int

	for d := range data {
		for o := range data[d].outputDigit {
			if data[d].outputDigit[o].checkNumber(1, nil) ||
				data[d].outputDigit[o].checkNumber(4, nil) ||
				data[d].outputDigit[o].checkNumber(7, nil) ||
				data[d].outputDigit[o].checkNumber(8, nil) {
				sum++
			}
		}
	}
	return sum
}

//SolveDay08Part2 returns the sum of all output digits
func SolveDay08Part2(input string) int {
	data := parseInput(input)

	var sum []int

	for d := range data {

		cache := make(map[int]digit)

		// get easy one
		cache[1] = data[d].getNumber(1, cache)
		cache[4] = data[d].getNumber(4, cache)
		cache[7] = data[d].getNumber(7, cache)
		cache[8] = data[d].getNumber(8, cache)

		// get 2 -> len(5) and 2 same as number4
		cache[2] = data[d].getNumber(2, cache)

		// get 9 -> len(6) and 4 same as number4
		cache[9] = data[d].getNumber(9, cache)

		// get 3 -> len(5) and 3 same as number7
		cache[3] = data[d].getNumber(3, cache)

		// get 5 -> len(5) and 2 same as number7 and 5 same as number9
		cache[5] = data[d].getNumber(5, cache)

		// get 6 -> len(6) and 3 same as number4 and 1 same as number1
		cache[6] = data[d].getNumber(6, cache)

		// get 0 -> len(6) and 3 same as number4 and 2 same as number1
		cache[0] = data[d].getNumber(0, cache)

		sol := make(map[digit]int)
		for i, s := range cache {
			sol[s] = i
		}
		var tempSum int
		for o := range data[d].outputDigit {
			tempSum *= 10
			tempSum += sol[data[d].outputDigit[o]]

		}
		sum = append(sum, tempSum)
	}
	return utils.SumSlice(sum)
}

//getNumber returns the digit to a number for some digits a cache of known digits in needed
func (i inputData) getNumber(number int, cache map[int]digit) digit {
	for d := range i.uniqueSignalPatterns {
		if i.uniqueSignalPatterns[d].checkNumber(number, cache) {
			return i.uniqueSignalPatterns[d]
		}
	}
	return ""
}

//checkNumber returns true if the digit is the given number for some digits a cache of known digits in needed
func (i digit) checkNumber(number int, cache map[int]digit) bool {
	switch number {
	case 0:
		// get 0 -> len(6) and 3 same as number4 and 2 same as number1
		if len(i) == 6 &&
			countSameSegments(i, cache[4]) == 3 &&
			countSameSegments(i, cache[1]) == 2 {
			return true
		}
	case 1:
		if len(i) == 2 {
			return true
		}
	case 2:
		// get 2 -> len(5) and 2 same as number4
		if len(i) == 5 &&
			countSameSegments(i, cache[4]) == 2 {
			return true
		}
	case 3:
		// get 3 -> len(5) and 3 same as number7
		if len(i) == 5 &&
			countSameSegments(i, cache[7]) == 3 {
			return true
		}
	case 4:
		if len(i) == 4 {
			return true
		}
	case 5:
		// get 5 -> len(5) and 2 same as number7 and 5 same as number9
		if len(i) == 5 &&
			countSameSegments(i, cache[7]) == 2 &&
			countSameSegments(i, cache[9]) == 5 {
			return true
		}
	case 6:
		// get 6 -> len(6) and 3 same as number4 and 1 same as number1
		if len(i) == 6 &&
			countSameSegments(i, cache[4]) == 3 &&
			countSameSegments(i, cache[1]) == 1 {
			return true
		}
	case 7:
		if len(i) == 3 {
			return true
		}
	case 8:
		if len(i) == 7 {
			return true
		}
	case 9:
		// get 9 -> len(6) and 4 same as number4
		if len(i) == 6 &&
			countSameSegments(i, cache[4]) == 4 {
			return true
		}
	}
	return false
}

//countSameSegments returns the number if same segments in 2 digits
func countSameSegments(a, b digit) int {
	var sum int
	for i := range a {
		for n := range b {
			if a[i] == b[n] {
				sum++
				break
			}
		}
	}
	return sum
}

//parseInput returns a slice of the numbers that needs to be solved
func parseInput(input string) []inputData {
	lines, _ := utils.InputToSlice(input)
	var data []inputData
	for l := range lines {
		outputSection := strings.Split(lines[l], " | ")
		if len(outputSection) != 2 {
			panic("")
		}
		var uniqueSignalPatterns []digit
		var outputDigit []digit
		for _, d := range strings.Split(outputSection[0], " ") {
			uniqueSignalPatterns = append(uniqueSignalPatterns, digit(utils.SortString(d)))
		}
		for _, d := range strings.Split(outputSection[1], " ") {
			outputDigit = append(outputDigit, digit(utils.SortString(d)))
		}
		data = append(data, inputData{
			uniqueSignalPatterns: uniqueSignalPatterns,
			outputDigit:          outputDigit,
		})
	}

	return data
}
