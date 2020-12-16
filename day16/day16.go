package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay16Part1(input))
	fmt.Printf("Part 2: %v\n", SolveDay16Part2(input))
}

//SolveDay16Part1 sum all invalid ticket numbers in the nearby tickets
func SolveDay16Part1(i string) (s int) {
	_, _, tickets, valid := getTicketData(i)
	for _, ticket := range tickets {
		for _, num := range strings.Split(ticket, ",") {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				return 0
			}
			if !valid[numInt] {
				s += numInt
			}
		}
	}
	return
}

//SolveDay16Part2 search for the right rule for each number in the ticket an multiply all numbers that start with departure
func SolveDay16Part2(i string) (s int) {
	rules, myTicket, nearbyTickets, valid := getTicketData(i)

	//clean nearby tickets
	var cleanNearbyTickets []string
	for _, ticket := range nearbyTickets {
		val := true
		for _, num := range strings.Split(ticket, ",") {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				return 0
			}
			if !valid[numInt] {
				val = false
				break
			}
		}
		if val {
			cleanNearbyTickets = append(cleanNearbyTickets, ticket)
		}
	}

	//init maps
	ruleMatchesPerColumn := make(map[string]int)

	ColumnRuleMatches := make(map[string]map[int]int)
	for ruleName, _ := range rules {
		ColumnRuleMatches[ruleName] = make(map[int]int)
	}

	columnMatchRule := make(map[int]map[string]bool)
	for i, _ := range myTicket {
		columnMatchRule[i] = make(map[string]bool)
	}

	//count rule match per column
	for _, nearbyTicket := range cleanNearbyTickets {
		for i, number := range strings.Split(nearbyTicket, ",") {
			num, _ := strconv.Atoi(number)
			for ruleName, rule := range rules {
				if rule[num] {
					ColumnRuleMatches[ruleName][i]++
				}
			}

		}
	}

	//search for full column matches per rule
	for ruleName, rules := range ColumnRuleMatches {
		for column, sum := range rules {
			if sum == len(cleanNearbyTickets) {
				columnMatchRule[column][ruleName] = true
			}
		}
	}

	//search for columns with only one rule and delete this rule from all other columns until all columns are set
	for {
		var found []string
		for column, rule := range columnMatchRule {
			if len(rule) == 1 {
				for ruleName, _ := range rule {
					found = append(found, ruleName)
					ruleMatchesPerColumn[ruleName] = column
				}
			}
		}
		for _, foundName := range found {
			for column, rule := range columnMatchRule {
				for ruleName, _ := range rule {
					if ruleName == foundName {
						delete(columnMatchRule[column], ruleName)
					}
				}

			}
		}
		if len(ruleMatchesPerColumn) == len(columnMatchRule) {
			break
		}
	}

	//multiply departure* values
	s = 1
	for ruleName, column := range ruleMatchesPerColumn {
		if strings.HasPrefix(ruleName, "departure") {
			s *= myTicket[column]
		}
	}
	return
}

//getTicketData extract the ticket data from the input
func getTicketData(input string) (rules map[string]map[int]bool, myTicket []int, nearbyTickets []string, validNumbers map[int]bool) {
	rules, validNumbers = make(map[string]map[int]bool), make(map[int]bool)

	inputParts := strings.Split(input, "\n\n")
	if len(inputParts) != 3 {
		return nil, nil, nil, nil
	}

	for _, rule := range strings.Split(inputParts[0], "\n") {
		ruleParts := strings.Split(rule, ": ")
		if len(ruleParts) != 2 {
			return nil, nil, nil, nil
		}
		numbers := strings.Split(ruleParts[1], " or ")
		rules[ruleParts[0]] = make(map[int]bool)
		for _, num := range numbers {
			minmax := strings.Split(num, "-")
			if len(minmax) != 2 {
				return nil, nil, nil, nil
			}
			min, err := strconv.Atoi(minmax[0])
			if err != nil {
				return nil, nil, nil, nil
			}
			max, err := strconv.Atoi(minmax[1])
			if err != nil {
				return nil, nil, nil, nil
			}
			for i := min; i <= max; i++ {
				validNumbers[i] = true
				rules[ruleParts[0]][i] = true
			}
		}
	}
	for _, num := range strings.Split(strings.Split(inputParts[1], "\n")[1], ",") {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			return nil, nil, nil, nil
		}
		myTicket = append(myTicket, numInt)
	}

	for _, ticket := range strings.Split(inputParts[2], "\n")[1:] {
		nearbyTickets = append(nearbyTickets, ticket)
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
