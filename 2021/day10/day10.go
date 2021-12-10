package main

import (
	"fmt"
	"github.com/dergeberl/aoc/utils"
	"os"
	"sort"
)

var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay10Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay10Part2(string(input)))
}

//SolveDay10Part1 returns the error score of all lines
func SolveDay10Part1(input string) int {
	line, _ := utils.InputToSlice(input)
	var sum int
	for i := range line {
		nextClosing := []rune{}
		errorPoints := 0
		for c := range line[i] {
			switch line[i][c] {
			case '(':
				nextClosing = append(nextClosing, ')')
				continue
			case '{':
				nextClosing = append(nextClosing, '}')
				continue
			case '[':
				nextClosing = append(nextClosing, ']')
				continue
			case '<':
				nextClosing = append(nextClosing, '>')
				continue
			default:
				if nextClosing[len(nextClosing)-1] != rune(line[i][c]) {
					errorPoints = scores[rune(line[i][c])]
				}
				nextClosing = nextClosing[:len(nextClosing)-1]
			}
			if errorPoints != 0 {
				sum += errorPoints
				break
			}
		}
	}
	return sum
}

//SolveDay10Part2 returns middle the completion score
func SolveDay10Part2(input string) int {
	line, _ := utils.InputToSlice(input)
	var sum []int
	for i := range line {
		nextClosing := []rune{}
		syntaxError := false
		completionScore := 0
		for c := range line[i] {
			switch line[i][c] {
			case '(':
				nextClosing = append(nextClosing, ')')
				continue
			case '{':
				nextClosing = append(nextClosing, '}')
				continue
			case '[':
				nextClosing = append(nextClosing, ']')
				continue
			case '<':
				nextClosing = append(nextClosing, '>')
				continue
			default:
				if nextClosing[len(nextClosing)-1] != rune(line[i][c]) {
					syntaxError = true
				}
				nextClosing = nextClosing[:len(nextClosing)-1]
			}
			if syntaxError {
				break
			}
		}
		if !syntaxError {
			for n := range nextClosing {
				completionScore *= 5
				switch nextClosing[len(nextClosing)-1-n] {
				case ')':
					completionScore += 1
				case ']':
					completionScore += 2
				case '}':
					completionScore += 3
				case '>':
					completionScore += 4
				}
			}
			sum = append(sum, completionScore)
		}
	}
	sort.Ints(sum)
	return sum[len(sum)/2]
}
