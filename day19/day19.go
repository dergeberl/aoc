package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var loopDetect map[int]int
var cache map[int]string

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)

	fmt.Printf("Part 1: %v\n", SolveDay19Part1(input))
	fmt.Printf("Part 2: %v\n", SolveDay19Part2(input))
}

func init() {
	loopDetect = make(map[int]int)
}

//SolveDay19Part1 return the number of strings that matched the rules
func SolveDay19Part1(input string) (s int) {
	blocks := strings.Split(input, "\n\n")
	if len(blocks) != 2 {
		return 0
	}
	rules := make(map[int]string)
	for _, line := range strings.Split(blocks[0], "\n") {
		splitLine := strings.Split(line, ": ")
		ruleId, _ := strconv.Atoi(splitLine[0])
		rules[ruleId] = splitLine[1]

	}
	cache = make(map[int]string)
	regex := "^" + resolveRulesRegex(0, rules) + "$"
	r, _ := regexp.Compile(regex)
	for _, line := range strings.Split(blocks[1], "\n") {
		if r.MatchString(line) {
			s++
		}
	}
	return
}

//SolveDay19Part2 return the number of strings that matched the rules (with updated rule 8 und 11)
func SolveDay19Part2(input string) (s int) {
	blocks := strings.Split(input, "\n\n")
	if len(blocks) != 2 {
		return 0
	}
	rules := make(map[int]string)
	for _, line := range strings.Split(blocks[0], "\n") {
		splitLine := strings.Split(line, ": ")
		ruleId, _ := strconv.Atoi(splitLine[0])
		rules[ruleId] = splitLine[1]

	}
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"
	cache = make(map[int]string)
	regex := "^" + resolveRulesRegex(0, rules) + "$"
	r, _ := regexp.Compile(regex)
	for _, line := range strings.Split(blocks[1], "\n") {
		if r.MatchString(line) {
			s++
		}
	}
	return
}

//resolveRulesRegex resolve a rule to a regex string (follow a loop max 20 times)
func resolveRulesRegex(ruleID int, rules map[int]string) (resolvedRule string) {
	if cache[ruleID] != "" {
		return cache[ruleID]
	}
	if strings.HasPrefix(rules[ruleID], "\"") {
		return strings.Trim(rules[ruleID], "\"")
	}
	regex := "(:?"
	orBlocks := strings.Split(rules[ruleID], " | ")
	for i, ruleBlock := range orBlocks {
		if i != 0 {
			regex += "|"
		}
		for _, otherRules := range strings.Split(ruleBlock, " ") {
			otherRulesNum, _ := strconv.Atoi(otherRules)
			if otherRulesNum == ruleID {
				loopDetect[ruleID]++
				if loopDetect[ruleID] > 20 {
					continue
				}
			}
			regex += resolveRulesRegex(otherRulesNum, rules)
		}
	}
	regex += ")"
	cache[ruleID] = regex
	return regex
}
