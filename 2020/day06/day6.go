package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay6Part1(input))
	fmt.Printf("Part 2: %v\n", SolveDay6Part2(input))
}

//SolveDay6Part1 returns the sum of yes answers of groups
func SolveDay6Part1(i string) (s int) {
	for _, answerGroup := range strings.Split(i, "\n\n") {
		s += len(deleteDuplicates(strings.Replace(answerGroup, "\n", "", -1)))
	}
	return
}

//SolveDay6Part2 returns the sum of answers that in one croup all answered with yes
func SolveDay6Part2(i string) (s int) {
	for _, answerGroup := range strings.Split(i, "\n\n") {
		cur := make(map[int32]int)
		answerGroupAnswers := stringListToSlice(answerGroup)
		for _, answers := range answerGroupAnswers {
			answers = deleteDuplicates(answers)
			for _, char := range answers {
				cur[char]++
			}
		}
		for _, answerCount := range cur {
			if answerCount == len(answerGroupAnswers) {
				s++
			}
		}
	}
	return
}

func deleteDuplicates(i string) (output string) {
	cur := make(map[int32]bool)
	for _, chars := range i {
		if cur[chars] {
			continue
		}
		cur[chars] = true
		output = output + string(chars)
	}
	return
}

//Helper functions
//stringListToSlice converts the list of strings (each string one row) to a slice
func stringListToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}
