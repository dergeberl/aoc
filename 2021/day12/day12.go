package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type connection map[string][]string

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay12Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay12Part2(string(input)))
}

//SolveDay12Part1 returns the number of path if one small cave can be used only once
func SolveDay12Part1(input string) int {
	connections := parseInput(input)
	return connections.findPath("start", false, nil)
}

//SolveDay12Part2 returns the number of path if one small cave can be used twice
func SolveDay12Part2(input string) int {
	connections := parseInput(input)
	return connections.findPath("start", true, nil)
}

func parseInput(input string) connection {
	line, _ := utils.InputToSlice(input)

	connections := make(connection)

	for i := range line {
		tmp := strings.Split(line[i], "-")
		if len(tmp) != 2 {
			panic("input false")
		}
		connections[tmp[0]] = append(connections[tmp[0]], tmp[1])
		connections[tmp[1]] = append(connections[tmp[1]], tmp[0])
	}
	return connections
}

// returns the number of paths for a given start cave
func (c connection) findPath(cave string, visitOneSmallCaveTwice bool, smallCaves []string) int {
	if cave == "end" {
		return 1
	}
	if cave == "start" && len(smallCaves) != 0 {
		return 0
	}

	if strings.ToLower(cave) == cave {
		if smallCaves == nil {
			smallCaves = make([]string, 0)
		}
		if !checkSmallCave(cave, smallCaves, visitOneSmallCaveTwice) {
			return 0
		}
		smallCaves = append(smallCaves, cave)
	}

	var sum int
	for _, p := range c[cave] {
		sum += c.findPath(p, visitOneSmallCaveTwice, smallCaves)
	}

	return sum
}

//checkSmallCave checks if this small cave can be used in path.
//returns true if this cave is used the first time
//returns true if this cave is used the second time but no other small cave is used (with visitOneSmallCaveTwice flag).
func checkSmallCave(cave string, smallCaves []string, visitOneSmallCaveTwice bool) bool {
	numberOfSmallCavesVisits := make(map[string]int)
	for i := range smallCaves {
		numberOfSmallCavesVisits[smallCaves[i]]++
	}
	if numberOfSmallCavesVisits[cave] >= 1 {
		if !visitOneSmallCaveTwice {
			return false
		}
		for _, i := range numberOfSmallCavesVisits {
			if i >= 2 {
				return false
			}
		}
	}
	return true
}
