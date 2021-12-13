package main

import (
	"fmt"
	"github.com/dergeberl/aoc/utils"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type paper [][]bool

type fold struct {
	axis  string
	value int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay13Part1(string(input)))
	fmt.Printf("Part 2: \n%v\n", SolveDay13Part2(string(input)))
}

//SolveDay13Part1 returns the number of dots after the first fold
func SolveDay13Part1(input string) int {
	p, ins := parseInput(input)

	if ins[0].axis == "x" {
		p = foldX(p, ins[0].value)
	}
	if ins[0].axis == "y" {
		p = foldY(p, ins[0].value)
	}

	return p.countDots()
}

//SolveDay13Part2 returns the dots as readable string
func SolveDay13Part2(input string) string {
	p, ins := parseInput(input)

	for i := range ins {
		if ins[i].axis == "x" {
			p = foldX(p, ins[i].value)
		}

		if ins[i].axis == "y" {
			p = foldY(p, ins[i].value)
		}
	}

	return p.toSting()
}

//foldX fold on the x-axis on a number and returns new paper
func foldX(p paper, v int) paper {
	for x := 0; x < v; x++ {
		fromX := (v * 2) - x
		for y := 0; y < len(p[x]); y++ {
			if fromX > len(p)-1 {
				continue
			}
			if p[fromX][y] {
				p[x][y] = true
			}
		}
	}
	return p[:v]
}

//foldY fold on the y-axis on a number and returns new paper
func foldY(p paper, v int) paper {
	for y := 0; y < v; y++ {
		fromY := (v * 2) - y
		for x := 0; x < len(p); x++ {
			if fromY > len(p[x])-1 {
				continue
			}
			if p[x][fromY] {
				p[x][y] = true
			}
		}
	}
	for x := range p {
		p[x] = p[x][:v]
	}
	return p
}

//parseInput returns an initial paper and a list of folds for an input
func parseInput(input string) (paper, []fold) {
	pointsInstruction := strings.Split(input, "\n\n")
	if len(pointsInstruction) != 2 {
		panic("wrong input")
	}
	points, _ := utils.InputToSlice(pointsInstruction[0])
	poi := make([]point, len(points))
	for i := range points {
		tmpPoints := strings.Split(points[i], ",")
		if len(tmpPoints) != 2 {
			panic("wrong input")
		}
		x, _ := strconv.Atoi(tmpPoints[0])
		y, _ := strconv.Atoi(tmpPoints[1])
		poi[i] = point{x: x, y: y}
	}
	var sizeX, sizeY int
	for i := range poi {
		if poi[i].x > sizeX {
			sizeX = poi[i].x
		}
		if poi[i].y > sizeY {
			sizeY = poi[i].y
		}
	}
	p := make([][]bool, sizeX+1)
	for i := 0; i <= sizeX; i++ {
		p[i] = make([]bool, sizeY+1)
	}
	for i := range poi {
		p[poi[i].x][poi[i].y] = true
	}

	instructions, _ := utils.InputToSlice(pointsInstruction[1])
	ins := make([]fold, len(instructions))
	for i := range instructions {
		tmpInstructions := strings.Split(instructions[i], "=")
		if len(tmpInstructions) != 2 {
			panic("wrong input")
		}
		value, _ := strconv.Atoi(tmpInstructions[1])
		axis := "y"
		if strings.HasSuffix(tmpInstructions[0], "x") {
			axis = "x"
		}
		ins[i].value = value
		ins[i].axis = axis
	}

	return p, ins
}

//countDots returns the number of dots in a paper
func (p paper) countDots() int {
	var sum int
	for x := range p {
		for y := range p[x] {
			if p[x][y] {
				sum++
			}
		}
	}
	return sum
}

//toSting returns a printable string from a paper
func (p paper) toSting() string {
	var out string
	for y := range p[0] {
		for x := range p {
			if p[x][y] {
				out += "#"
				continue
			}
			out += "."
		}
		out += fmt.Sprintln("")
	}
	return out
}
