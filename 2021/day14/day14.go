package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type polymerTemplate map[string]int
type polymerRules map[string][]string

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay14Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay14Part2(string(input)))
}

//SolveDay14Part1 applies the polymerRules 10 times and returns the highest polymer count - the lowest polymer count
func SolveDay14Part1(input string) int {
	template, rules := parseInput(input)

	for i := 0; i < 10; i++ {
		template = rules.applyRules(template)
	}

	count := template.getPolymerCount()
	return count[len(count)-1] - count[0]
}

//SolveDay14Part2 applies the polymerRules 40 times and returns the highest polymer count - the lowest polymer count
func SolveDay14Part2(input string) int {
	template, rules := parseInput(input)

	for i := 0; i < 40; i++ {
		template = rules.applyRules(template)
	}

	count := template.getPolymerCount()
	return count[len(count)-1] - count[0]
}

//applyRules returns a new polymerTemplate for a given polymerTemplate after applying the polymerRules
func (p polymerRules) applyRules(template polymerTemplate) polymerTemplate {
	newTemplate := make(polymerTemplate)
	for t := range template {
		for _, r := range p[t] {
			newTemplate[r] += template[t]
		}
		// keep start and end
		if len(t) == 1 {
			newTemplate[t]++
		}
	}
	return newTemplate
}

//getPolymerCount returns a sorted int slice with the count of each polymer in a polymerTemplate
func (p polymerTemplate) getPolymerCount() []int {
	tempCount := make(map[string]int)
	for t, i := range p {
		for _, r := range t {
			tempCount[string(r)] += i
		}
	}
	var count []int
	for _, s := range tempCount {
		count = append(count, s/2)
	}
	sort.Ints(count)
	return count
}

//parseInput returns the initial polymer template and polymer rules
func parseInput(input string) (polymerTemplate, polymerRules) {
	startTemplatePolymerRules := strings.Split(input, "\n\n")
	if len(startTemplatePolymerRules) != 2 {
		panic("wrong input")
	}

	// create initial polymerTemplate
	tmpStartTemplate := startTemplatePolymerRules[0]
	startTemplate := make(polymerTemplate)
	for i := 0; i < len(tmpStartTemplate)-1; i++ {
		startTemplate[string(tmpStartTemplate[i])+string(tmpStartTemplate[i+1])]++
	}
	// add start and end as single for tha count at the end
	startTemplate[string(tmpStartTemplate[0])]++
	startTemplate[string(tmpStartTemplate[len(tmpStartTemplate)-1])]++

	// create rules
	// for example:
	// CH -> B creates CB and BH polymer
	tmpRules, _ := utils.InputToSlice(startTemplatePolymerRules[1])
	rules := make(polymerRules)
	for i := range tmpRules {
		tmpPolymerRules := strings.Split(tmpRules[i], " -> ")
		if len(tmpPolymerRules) != 2 {
			panic("wrong input")
		}
		rules[tmpPolymerRules[0]] = make([]string, 2)
		rules[tmpPolymerRules[0]][0] = string(tmpPolymerRules[0][0]) + tmpPolymerRules[1]
		rules[tmpPolymerRules[0]][1] = tmpPolymerRules[1] + string(tmpPolymerRules[0][1])

	}

	return startTemplate, rules
}
