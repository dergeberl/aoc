package main

import (
	"errors"
	"fmt"
	"github.com/dergeberl/aoc/utils"
	"os"
	"strconv"
	"strings"
)

type section struct {
	start, end int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay04Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay04Part2(string(input)))
}

// SolveDay04Part1 returns the number pairs that fully contain the other
func SolveDay04Part1(input string) int {
	lines, _ := utils.InputToSlice(input)
	var solution int
	for i := range lines {
		e1, e2, err := lineToSections(lines[i])
		if err != nil {
			continue
		}
		if e1.checkIfCompleteFit(e2) || e2.checkIfCompleteFit(e1) {
			solution++
		}
	}
	return solution
}

// SolveDay04Part2 returns the number pairs that partial contains the other
func SolveDay04Part2(input string) int {
	lines, _ := utils.InputToSlice(input)
	var solution int
	for i := range lines {
		e1, e2, err := lineToSections(lines[i])
		if err != nil {
			continue
		}
		if e1.checkIfPartialOverlap(e2) || e2.checkIfPartialOverlap(e1) {
			solution++
		}
	}
	return solution
}

// lineToSections returns 2 sections for each line
func lineToSections(line string) (section, section, error) {
	split := strings.Split(line, ",")
	if len(split) != 2 {
		return section{}, section{}, errors.New("conversion failed")
	}
	e1 := strings.Split(split[0], "-")
	e2 := strings.Split(split[1], "-")
	if len(e1) != 2 || len(e2) != 2 {
		return section{}, section{}, errors.New("conversion failed")
	}
	e1start, err := strconv.Atoi(e1[0])
	if err != nil {
		return section{}, section{}, err
	}
	e1end, err := strconv.Atoi(e1[1])
	if err != nil {
		return section{}, section{}, err
	}

	e2start, err := strconv.Atoi(e2[0])
	if err != nil {
		return section{}, section{}, err
	}
	e2end, err := strconv.Atoi(e2[1])
	if err != nil {
		return section{}, section{}, err
	}
	return section{start: e1start, end: e1end}, section{e2start, e2end}, nil
}

// checkIfCompleteFit return true if b fits completely in a
func (a section) checkIfCompleteFit(b section) bool {
	return b.start >= a.start && b.end <= a.end
}

// checkIfPartialOverlap return true if b fits partially in a
func (a section) checkIfPartialOverlap(b section) bool {
	return (b.start >= a.start && b.start <= a.end) || (b.end >= a.start && b.end <= a.end)
}
