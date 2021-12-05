package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type line struct {
	start position
	end   position
}
type position struct {
	x, y int
}
type lines []line

type grid [][]int

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay05Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay05Part2(string(input)))
}

//SolveDay05Part1 returns the number of overlapping positions with straight lines
func SolveDay05Part1(input string) int {
	l := getLinesFromInput(input)
	g := l.getGrid()
	for i := range l {
		g.drawStraightLine(l[i])
	}
	return g.getOverlapping()
}

//SolveDay05Part2 returns the number of overlapping positions with straight and diagonal lines
func SolveDay05Part2(input string) int {
	l := getLinesFromInput(input)
	g := l.getGrid()
	for i := range l {
		g.drawAllLine(l[i])
	}
	return g.getOverlapping()
}

//getLinesFromInput returns lines for the input
func getLinesFromInput(input string) lines {
	inputLines, _ := utils.InputToSlice(input)
	var output []line
	for l := range inputLines {
		positions := strings.Split(inputLines[l], " -> ")
		if len(positions) != 2 {
			continue
		}
		start := strings.Split(positions[0], ",")
		end := strings.Split(positions[1], ",")
		if len(start) != 2 || len(end) != 2 {
			continue
		}
		startx, _ := strconv.Atoi(start[0])
		starty, _ := strconv.Atoi(start[1])
		endx, _ := strconv.Atoi(end[0])
		endy, _ := strconv.Atoi(end[1])
		output = append(output, line{
			start: position{
				x: startx,
				y: starty,
			},
			end: position{
				x: endx,
				y: endy,
			},
		})
	}
	return output
}

//getGrid returns an initialized gird with the needed size for the lines
func (l lines) getGrid() grid {
	var hightX, hightY int

	for i := range l {
		if l[i].start.x > hightX {
			hightX = l[i].start.x
		}
		if l[i].end.x > hightX {
			hightX = l[i].end.x
		}
		if l[i].start.y > hightY {
			hightY = l[i].start.y
		}
		if l[i].end.x > hightY {
			hightY = l[i].end.y
		}
	}
	g := make([][]int, hightX+1)

	for i := range g {
		g[i] = make([]int, hightY+1)
	}
	return g
}

//drawStraightLine draws a vertical or horizontal line, diagonal lines are ignored
func (g *grid) drawStraightLine(l line) {
	if l.start.x == l.end.x {
		yNumbers := getNumbersBetween(l.start.y, l.end.y)
		for i := range yNumbers {
			(*g)[l.start.x][yNumbers[i]]++
		}
		return
	}
	if l.start.y == l.end.y {
		xNumbers := getNumbersBetween(l.start.x, l.end.x)
		for i := range xNumbers {
			(*g)[xNumbers[i]][l.start.y]++
		}
	}
}

//drawAllLine draws all lines incl. diagonal lines
func (g *grid) drawAllLine(l line) {
	if l.start.x == l.end.x || l.start.y == l.end.y {
		g.drawStraightLine(l)
		return
	}
	xNumbers := getNumbersBetween(l.start.x, l.end.x)
	yNumbers := getNumbersBetween(l.start.y, l.end.y)
	if len(xNumbers) != len(yNumbers) {
		return
	}
	for i := range xNumbers {
		(*g)[xNumbers[i]][yNumbers[i]]++
	}
}

//getOverlapping returns the positions that have more than one line on it
func (g grid) getOverlapping() int {
	var output int
	for x := range g {
		for y := range g[x] {
			if g[x][y] > 1 {
				output++
			}
		}
	}
	return output
}

//toString returns a string for troubleshooting. DO NOT USE WITH BIG INPUT!
func (g grid) toString() string { //nolint: unused
	var output string
	for y := range g[0] {
		for x := range g {
			output += fmt.Sprint(g[x][y])
		}
		output += "\n"
	}
	return output
}

func getNumbersBetween(a, b int) []int {
	var s []int
	if a > b {
		for i := a; i >= b; i-- {
			s = append(s, i)
		}
		return s
	}
	for i := a; i <= b; i++ {
		s = append(s, i)
	}
	return s
}
