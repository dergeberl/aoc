package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/dergeberl/aoc/utils"
)

type flowMap [][]int

type point struct {
	x, y int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay09Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay09Part2(string(input)))
}

//SolveDay09Part1 returns the risk level by sum all low points and add a 1 per low point
func SolveDay09Part1(input string) int {
	flow := parseInput(input)
	var sum int
	for x := range flow {
		for y := range flow[x] {
			if flow.checkLowPoint(point{x: x, y: y}) {
				sum += flow[x][y] + 1
			}
		}
	}

	return sum
}

//SolveDay09Part2 multiply the 3 highest basin counts for all low points
func SolveDay09Part2(input string) int {
	flow := parseInput(input)
	var sum []int
	for x := range flow {
		for y := range flow[x] {
			if flow.checkLowPoint(point{x: x, y: y}) {
				sum = append(sum, flow.calcBaseIn(x, y, nil))
			}
		}
	}
	sort.Ints(sum)
	return sum[len(sum)-1] * sum[len(sum)-2] * sum[len(sum)-3]
}

//checkLowPoint returns true if the given point is a low point
func (f flowMap) checkLowPoint(p point) bool {
	if p.x != 0 {
		if f[p.x-1][p.y] <= f[p.x][p.y] {
			return false
		}
	}
	if p.y != 0 {
		if f[p.x][p.y-1] <= f[p.x][p.y] {
			return false
		}
	}
	if p.x+1 != len(f) {
		if f[p.x+1][p.y] <= f[p.x][p.y] {
			return false
		}
	}
	if p.y+1 != len(f[p.x]) {
		if f[p.x][p.y+1] <= f[p.x][p.y] {
			return false
		}
	}
	return true
}

//calcBaseIn returns the number base in of a given point
func (f flowMap) calcBaseIn(x, y int, cache map[point]bool) int {
	if cache == nil {
		cache = make(map[point]bool)
	}

	if f[x][y] == 9 || cache[point{x: x, y: y}] {
		return 0
	}
	cache[point{x: x, y: y}] = true
	if x > 0 {
		if (f[x-1][y] - 1) >= f[x][y] {
			f.calcBaseIn(x-1, y, cache)
		}
	}
	if y > 0 {
		if (f[x][y-1] - 1) >= f[x][y] {
			f.calcBaseIn(x, y-1, cache)
		}
	}
	if x < len(f)-1 {
		if (f[x+1][y] - 1) >= f[x][y] {
			f.calcBaseIn(x+1, y, cache)
		}
	}
	if y < len(f[x])-1 {
		if (f[x][y+1] - 1) >= f[x][y] {
			f.calcBaseIn(x, y+1, cache)
		}
	}

	return len(cache)
}

func parseInput(input string) flowMap {
	var fm flowMap
	inputLines, _ := utils.InputToSlice(input)
	fm = make([][]int, len(inputLines))
	for i := range inputLines {
		for _, r := range inputLines[i] {
			fm[i] = append(fm[i], int(r-48))
		}
	}

	return fm
}
