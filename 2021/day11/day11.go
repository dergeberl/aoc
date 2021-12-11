package main

import (
	"fmt"
	"os"

	"github.com/dergeberl/aoc/utils"
)

type coordinate struct {
	x, y int
}

type octopusEnergy struct {
	value int
	set   bool
}

type octopusField map[coordinate]octopusEnergy

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay11Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay11Part2(string(input)))
}

//SolveDay11Part1 returns the number of flashes after 100 steps
func SolveDay11Part1(input string) int {
	o := parseInput(input)
	var solution int
	for i := 0; i < 100; i++ {
		o.increaseAll()
		solution += o.flashAll()
	}

	return solution
}

//SolveDay11Part2 returns the number of steps after which all octopus are in sync
func SolveDay11Part2(input string) int {
	o := parseInput(input)
	var solution int
	for {
		solution++
		o.increaseAll()
		_ = o.flashAll()
		if o.checkInSync() {
			break
		}
	}
	return solution
}

//increaseAll increases all octopuses energy by 1
func (o *octopusField) increaseAll() {
	for k := range *o {
		o.increaseOctopus(k)
	}
}

//increaseOctopus increases given octopus energy by 1
func (o *octopusField) increaseOctopus(c coordinate) {
	oct := (*o)[c]
	if !oct.set {
		return
	}
	oct.value++
	(*o)[c] = oct
}

//flashAll returns number of flashes from all octopuses if they reached there energy level until no more flashes appear this step
func (o *octopusField) flashAll() int {
	var sum int
	var flashedOctopus []coordinate
	for {
		var flashesAppear bool
		for c := range *o {
			if o.flashOctopus(c) {
				flashesAppear = true
				sum++
				flashedOctopus = append(flashedOctopus, c)
			}
		}
		if !flashesAppear {
			break
		}
	}
	for i := range flashedOctopus {
		(*o)[flashedOctopus[i]] = octopusEnergy{value: 0, set: true}
	}
	return sum
}

//flashOctopus flash an octopus if the reached energy level and increase the adjacent octopuses
func (o *octopusField) flashOctopus(c coordinate) bool {
	if !(*o)[c].set || (*o)[c].value < 10 {
		return false
	}
	o.increaseOctopus(coordinate{x: c.x - 1, y: c.y})
	o.increaseOctopus(coordinate{x: c.x + 1, y: c.y})
	o.increaseOctopus(coordinate{x: c.x, y: c.y - 1})
	o.increaseOctopus(coordinate{x: c.x, y: c.y + 1})
	o.increaseOctopus(coordinate{x: c.x - 1, y: c.y - 1})
	o.increaseOctopus(coordinate{x: c.x + 1, y: c.y + 1})
	o.increaseOctopus(coordinate{x: c.x - 1, y: c.y + 1})
	o.increaseOctopus(coordinate{x: c.x + 1, y: c.y - 1})
	(*o)[c] = octopusEnergy{
		value: 0,
		set:   true,
	}
	return true
}

//checkInSync returns true if all octopuses are in sync
func (o *octopusField) checkInSync() bool {
	ref := (*o)[coordinate{x: 0, y: 0}]
	for c := range *o {
		if ref.value != (*o)[c].value {
			return false
		}
	}
	return true
}

//parseInput parses the input into an octopusField
func parseInput(input string) octopusField {
	inputLines, _ := utils.InputToSlice(input)
	o := make(octopusField)
	for x := range inputLines {
		for y, r := range inputLines[x] {
			o[coordinate{
				x: x,
				y: y,
			}] = octopusEnergy{
				value: int(r - 48),
				set:   true,
			}
		}
	}
	return o
}
