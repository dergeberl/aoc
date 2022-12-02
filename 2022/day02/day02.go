package main

import (
	"fmt"
	"github.com/dergeberl/aoc/utils"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay02Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay02Part2(string(input)))
}

// SolveDay02Part1 returns the total score of the rock, paper, scissors game
func SolveDay02Part1(input string) int {
	inputLines, _ := utils.InputToSlice(input)
	var solution int
	for i := range inputLines {
		if len(inputLines[i]) != 3 {
			continue
		}
		switch inputLines[i][0] {
		case 'A': // other choose rock
			switch inputLines[i][2] {
			case 'X': // I choose rock (1), draw (3)
				solution += 1 + 3
			case 'Y': // I choose paper (2), win (6)
				solution += 2 + 6
			case 'Z': // I choose scissors (3), loss (0)
				solution += 3 + 0
			}
		case 'B': // other choose paper
			switch inputLines[i][2] {
			case 'X': // I choose rock (1), loss (0)
				solution += 1 + 0
			case 'Y': // I choose paper (2), draw (3)
				solution += 2 + 3
			case 'Z': // I choose scissors (3), win (6)
				solution += 3 + 6
			}
		case 'C': // other choose scissors
			switch inputLines[i][2] {
			case 'X': // I choose rock (1), win (6)
				solution += 1 + 6
			case 'Y': // I choose paper (2), loss (0)
				solution += 2 + 0
			case 'Z': // I choose scissors (3), draw (3)
				solution += 3 + 3
			}
		}
	}
	return solution
}

// SolveDay02Part2 returns the total score of the rock, paper, scissors game
func SolveDay02Part2(input string) int {
	inputLines, _ := utils.InputToSlice(input)
	var solution int
	for i := range inputLines {
		if len(inputLines[i]) != 3 {
			continue
		}
		switch inputLines[i][0] {
		case 'A': // other choose rock
			switch inputLines[i][2] {
			case 'X': // I should lose (0), chose scissors (3)
				solution += 0 + 3
			case 'Y': // I should draw (3), chose rock (1)
				solution += 3 + 1
			case 'Z': // I should win (6), chose paper (2)
				solution += 6 + 2
			}
		case 'B': // other choose paper
			switch inputLines[i][2] {
			case 'X': // I should lose (0), chose rock (1)
				solution += 0 + 1
			case 'Y': // I should draw (3), chose paper (2)
				solution += 3 + 2
			case 'Z': // I should win (6), chose scissors (3)
				solution += 6 + 3
			}
		case 'C': // other choose scissors
			switch inputLines[i][2] {
			case 'X': // I should lose (0), chose paper (2)
				solution += 0 + 2
			case 'Y': // I should draw (3), chose scissors (3)
				solution += 3 + 3
			case 'Z': // I should win (6), chose rock (1)
				solution += 6 + 1
			}
		}
	}
	return solution
}
