package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type area struct {
	x1, x2 int
	y1, y2 int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay17Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay17Part2(string(input)))
}

//SolveDay17Part1 returns the highest reachable y value for a target area
func SolveDay17Part1(input string) int {
	a := parseInput(input)
	return utils.GetGaussscheSummenformel((a.y1 * -1) - 1)
}

//SolveDay17Part2  returns the number of possible shots for a target area
func SolveDay17Part2(input string) int {
	a := parseInput(input)
	var hits int
	for x := 0; x <= a.x2; x++ {
		for y := a.y1; y <= (a.y1 * -1); y++ {
			if a.checkShot(x, y) {
				hits++
			}
		}
	}
	return hits
}

//checkShot returns true shot is possible and returns the highest reached y
func (a area) checkShot(x, y int) bool {
	curX, curY := 0, 0
	for a.isReachableByPoint(curX, curY) {
		curX += x
		curY += y
		if a.checkPoint(curX, curY) {
			return true
		}
		if x > 0 {
			x--
		}
		y--
	}
	return false
}

//checkPoint returns true if point is in area
func (a area) checkPoint(x, y int) bool {
	return y >= a.y1 && y <= a.y2 && x >= a.x1 && x <= a.x2
}

//isReachableByPoint returns if the area is reachable by a point
func (a area) isReachableByPoint(x, y int) bool {
	return !(y < a.y1 || x > a.x2)
}

//parseInput returns the given target area
func parseInput(input string) area {
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimPrefix(input, "target area: ")

	var a area

	for _, v := range strings.Split(input, ",") {
		v = strings.TrimSpace(v)
		if strings.HasPrefix(v, "x") {
			xVals := strings.Split(strings.TrimPrefix(v, "x="), "..")
			if len(xVals) != 2 {
				panic("wrong input")
			}
			a.x1, _ = strconv.Atoi(xVals[0])
			a.x2, _ = strconv.Atoi(xVals[1])
		}
		if strings.HasPrefix(v, "y") {
			yVals := strings.Split(strings.TrimPrefix(v, "y="), "..")
			if len(yVals) != 2 {
				panic("wrong input")
			}
			a.y1, _ = strconv.Atoi(yVals[0])
			a.y2, _ = strconv.Atoi(yVals[1])
		}
	}

	if a.y1 > a.y2 {
		a.y1, a.y2 = a.y2, a.y1
	}
	if a.x1 > a.x2 {
		a.x1, a.x2 = a.x2, a.x1
	}
	return a
}
