package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

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

// SolveDay03Part1 will return the sum of all priorities of items which are in both compartment.
func SolveDay03Part1(input string) int {
	line, _ := utils.InputToSlice(input)
	var solution int
	for i := range line {
		compartment1 := line[i][0 : len(line[i])/2]
		compartment2 := line[i][len(line[i])/2:]
		for c1 := range compartment1 {
			if strings.Contains(compartment2, string(compartment1[c1])) {
				solution += getNumberForRune(rune(compartment1[c1]))
				break
			}
		}
	}
	return solution
}

// SolveDay03Part2 will return the sum of all batch items of each elf group
func SolveDay03Part2(input string) int {
	lines, _ := utils.InputToSlice(input)
	var solution int
	for i := range lines {
		if i%3 != 0 {
			continue
		}
		for e := range lines[i] {
			if strings.Contains(lines[i+1], string(lines[i][e])) &&
				strings.Contains(lines[i+2], string(lines[i][e])) {
				solution += getNumberForRune(rune(lines[i][e]))
				break
			}
		}
	}
	return solution
}

// getNumberForRune returns a number for a rune.
// a through z will return 1 through 26
// A through Z will return 27 through 52
func getNumberForRune(r rune) int {
	if unicode.IsUpper(r) {
		return int(r - 38)
	}
	return int(r - 96)
}
