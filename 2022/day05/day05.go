package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []string
type compartment map[int]stack

type step struct {
	from, to, count int
}
type procedure []step

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay05Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay05Part2(string(input)))
}

// SolveDay05Part1 returns the crate which are in top of each stack
func SolveDay05Part1(input string) string {
	c, p := parseInput(input)

	for i := range p {
		for x := 0; x < p[i].count; x++ {
			var item string
			c[p[i].from], item = popLatestFromStack(c[p[i].from])
			c[p[i].to] = append(c[p[i].to], item)
		}
	}

	var solution string
	for i := 1; i <= len(c); i++ {
		_, r := popLatestFromStack(c[i])
		solution += r
	}
	return solution
}

// SolveDay05Part2 returns the crate which are in top of each stack
func SolveDay05Part2(input string) string {
	c, p := parseInput(input)

	for i := range p {
		var items stack
		c[p[i].from], items = popMultipleFromStack(c[p[i].from], p[i].count)
		c[p[i].to] = append(c[p[i].to], items...)
	}

	var solution string
	for i := 1; i <= len(c); i++ {
		_, r := popLatestFromStack(c[i])
		solution += r
	}
	return solution
}

// parseInput return a compartment and procedure for the input
func parseInput(input string) (compartment, procedure) {
	inputSplit := strings.Split(input, "\n\n")
	cargoLines := strings.Split(inputSplit[0], "\n")
	procedureLines := strings.Split(inputSplit[1], "\n")

	comp := make(compartment)
	pro := make(procedure, 0)

	for i := range cargoLines[len(cargoLines)-1] {
		if cargoLines[len(cargoLines)-1][i] == ' ' {
			continue
		}
		compartmentInt, _ := strconv.Atoi(string(cargoLines[len(cargoLines)-1][i]))

		comp[compartmentInt] = make([]string, 0)
		for x := len(cargoLines) - 2; x >= 0; x-- {
			if cargoLines[x][i] == ' ' {
				continue
			}
			comp[compartmentInt] = append(comp[compartmentInt], string(cargoLines[x][i]))
		}
	}

	for i := range procedureLines {
		s := step{}
		lineSplit := strings.Split(procedureLines[i], " ")
		if len(lineSplit) != 6 {
			continue
		}
		s.count, _ = strconv.Atoi(lineSplit[1])
		s.from, _ = strconv.Atoi(lineSplit[3])
		s.to, _ = strconv.Atoi(lineSplit[5])
		pro = append(pro, s)
	}

	return comp, pro
}

// popLatestFromStack returns the new stack without the last item and returns last item
func popLatestFromStack(s stack) (stack, string) {
	newStack := s[:len(s)-1]

	return newStack, s[len(s)-1]
}

// popMultipleFromStack returns the new stack without the count last item and returns last count items
func popMultipleFromStack(s stack, count int) (stack, stack) {
	newStack := s[:len(s)-count]
	popStack := s[len(s)-count:]
	return newStack, popStack
}
