package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var formulaCheck *regexp.Regexp

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay18Part1(stringListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay18Part2(stringListToSlice(input)))
}

func init() {
	formulaCheck, _ = regexp.Compile("^([0-9]|\\s|\\(|\\)|\\+|\\-|\\*)*$")
}

//SolveDay18Part1 calc each line (from left to right, parentheses first) and return the sum of all lines
func SolveDay18Part1(input []string) (s int) {
	for _, line := range input {
		s += calcPart1(line)
	}
	return
}

//SolveDay18Part2 calc each line (first add and subtract then multiply, parentheses first) and return the sum of all lines
func SolveDay18Part2(input []string) (s int) {
	for _, line := range input {
		s += calcPart2(line)
	}
	return
}

//calcPart1 calculate the given formula from left to right (parentheses first)
func calcPart1(formula string) (sum int) {
	if !formulaCheck.MatchString(formula) {
		return 0
	}
	operator := "+"
	var parenthesesCount, parenthesesPosition int
	for i, char := range formula {
		if parenthesesCount != 0 && char != '(' && char != ')' {
			continue
		}
		switch char {
		case '(':
			parenthesesCount++
			if parenthesesCount == 1 {
				parenthesesPosition = i + 1
			}
		case ')':
			parenthesesCount--
			if parenthesesCount == 0 {
				parenthesesSolve := calcPart1(formula[parenthesesPosition:i])
				if operator == "+" {
					sum += parenthesesSolve
				} else if operator == "-" {
					sum -= parenthesesSolve
				} else if operator == "*" {
					sum *= parenthesesSolve
				}
			}
		case ' ':
			continue
		case '+':
			operator = "+"
		case '-':
			operator = "-"
		case '*':
			operator = "*"
		default:
			if operator == "+" {
				num, _ := strconv.Atoi(string(char))
				sum += num
			} else if operator == "-" {
				num, _ := strconv.Atoi(string(char))
				sum -= num
			} else if operator == "*" {
				num, _ := strconv.Atoi(string(char))
				sum *= num
			}
		}
	}
	if parenthesesCount != 0 {
		return 0
	}
	return
}

//calcPart2 calculate the given formula first add and subtract then multiply (parentheses first)
func calcPart2(formula string) int {
	if !formulaCheck.MatchString(formula) {
		return 0
	}
	var multiply []int
	var parenthesesCount, parenthesesPosition, tempSum int
	operator := "+"

	for i, char := range formula {
		if parenthesesCount != 0 && char != '(' && char != ')' {
			continue
		}
		switch char {
		case '(':
			parenthesesCount++
			if parenthesesCount == 1 {
				parenthesesPosition = i + 1
			}
		case ')':
			parenthesesCount--
			if parenthesesCount == 0 {
				parenthesesSolve := calcPart2(formula[parenthesesPosition:i])
				if operator == "+" {
					tempSum += parenthesesSolve
				} else if operator == "-" {
					tempSum -= parenthesesSolve
				}
			}
		case ' ':
			continue
		case '+':
			operator = "+"
		case '-':
			operator = "-"
		case '*':
			operator = "+"
			multiply = append(multiply, tempSum)
			tempSum = 0
		default:
			if operator == "+" {
				num, _ := strconv.Atoi(string(char))
				tempSum += num
			} else if operator == "-" {
				num, _ := strconv.Atoi(string(char))
				tempSum -= num
			}
		}
	}
	if parenthesesCount != 0 {
		return 0
	}
	multiply = append(multiply, tempSum)
	sum := 1
	for _, num := range multiply {
		sum *= num
	}
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
