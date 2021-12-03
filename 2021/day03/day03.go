package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dergeberl/aoc/utils"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay03Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay03Part2(string(input)))
}

//SolveDay03Part1
func SolveDay03Part1(input string) int {
	line, _ := utils.InputToSlice(input)
	var gamma, epsilon int
	calculated := make([]int, len(line[0]))
	for i := range line {
		chars := []byte(line[i])
		for b := range chars {
			if chars[b] == '1' {
				calculated[b]++
			}
		}
	}

	for i := range calculated {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if calculated[i] >= len(line)/2 {
			gamma++
		} else {
			epsilon++
		}
	}
	return gamma * epsilon
}

//SolveDay03Part2
func SolveDay03Part2(input string) int {
	lineInput0, _ := utils.InputToSlice(input)
	lineInput1 := lineInput0
	for char := range lineInput0[0] {
		var line0, line1 []string
		for i := range lineInput0 {
			if lineInput0[i][char] == '1' {
				line1 = append(line1, lineInput0[i])
			} else {
				line0 = append(line0, lineInput0[i])
			}
		}
		if len(line1) >= len(line0) {
			lineInput0 = line1
		} else {
			lineInput0 = line0
		}

		if len(lineInput1) == 1 {
			break
		}
	}
	oxi, _ := strconv.ParseInt(lineInput0[0], 2, 32)

	for char := range lineInput1[0] {
		var line0, line1 []string
		for i := range lineInput1 {
			if lineInput1[i][char] == '1' {
				line1 = append(line1, lineInput1[i])
			} else {
				line0 = append(line0, lineInput1[i])
			}
		}
		if len(line1) >= len(line0) {
			lineInput1 = line0
		} else {
			lineInput1 = line1
		}
		if len(lineInput1) == 1 {
			break
		}
	}
	co2, _ := strconv.ParseInt(lineInput1[0], 2, 32)

	return int(oxi * co2)
}
