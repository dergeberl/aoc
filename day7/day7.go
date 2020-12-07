package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//SolveDay7Part1 counts the colors that can contain a shiny gold bag
func SolveDay7Part1(i []string) (c int) {
	bags := exportBags(i)
	colorsToCheck := []string{"shiny gold"}
	checkedColors := make(map[string]bool)
	for {
		var newColorsToCheck []string
		for _, checkColor := range colorsToCheck {
			for color, bag := range bags {
				if bag[checkColor] > 0 && !checkedColors[color] {
					newColorsToCheck = append(newColorsToCheck, color)
					checkedColors[color] = true
					c++
				}
			}
		}
		if len(newColorsToCheck) == 0 {
			return
		}
		colorsToCheck = newColorsToCheck
	}
}

//SolveDay7Part2 counts the bags that are contained in one shiny gold
func SolveDay7Part2(i []string) int {
	bags := exportBags(i)
	return countBags("shiny gold", 1, bags) - 1
}

// count the bags that are contained in a bag (incl. the initial bag)
func countBags(color string, i int, bags map[string]map[string]int) int {
	s := i
	for nextColor, num := range bags[color] {
		s += i * countBags(nextColor, num, bags)
	}
	return s
}

//exportBags export all bags that contains other bags in a map that contains a map with all contained bags and the number of contained bags
func exportBags(i []string) (bags map[string]map[string]int) {
	bags = make(map[string]map[string]int)
	for _, line := range i {
		if !strings.HasSuffix(line, "bags contain no other bags.") {
			split := strings.Split(line, " bags contain ")
			bag := make(map[string]int)
			for _, c := range strings.Split(split[1], ", ") {
				t := strings.Split(c, " ")
				containInt, err := strconv.Atoi(t[0])
				if err != nil {
					return nil
				}
				bag[t[1]+" "+t[2]] = containInt
			}
			bags[split[0]] = bag
		}
	}
	return
}

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay7Part1(stringListToSlice(string(i))))
	fmt.Printf("Part 2: %v\n", SolveDay7Part2(stringListToSlice(string(i))))
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
