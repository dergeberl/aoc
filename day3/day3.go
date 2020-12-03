package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//SolveDay3Part1 count the trees that hit by a path
func SolveDay3Part1(trees []string, right int, down int) (sum int) {
	i := 0
	for in, p := range trees {
		if (in % down) == 0 {
			if len(p)-1 < i {
				i = i - len(p)
			}
			if p[i] == '#' {
				sum++
			}
			i = i + right
		}
	}
	return
}

//SolveDay3Part2 count the trees that hit for 5 path and multiply them
func SolveDay3Part2(trees []string) (sum int) {
	return SolveDay3Part1(trees, 1, 1) * SolveDay3Part1(trees, 3, 1) * SolveDay3Part1(trees, 5, 1) * SolveDay3Part1(trees, 7, 1) * SolveDay3Part1(trees, 1, 2)
}

//listToSlice converts the list of numbers (each number one row) to a slice
func listToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}

func main() {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay3Part1(listToSlice(input), 3, 1))
	fmt.Printf("Part 2: %v\n", SolveDay3Part2(listToSlice(input)))

}
