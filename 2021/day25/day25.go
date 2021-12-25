package main

import (
	"fmt"
	"os"

	"github.com/dergeberl/aoc/utils"
)

type location struct {
	x, y int
}

type seafloor map[location]*string

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay25Part1(string(input)))
}

//SolveDay25Part1 returns the number of steps until all sea cucumbers stops
func SolveDay25Part1(input string) int {
	floor := parseInput(input)
	var count int
	moved := true
	for moved {
		count++
		moved, floor = floor.move()
	}
	return count
}

//move returns a new seafloor after every sea cucumber is moved
func (s seafloor) move() (bool, seafloor) {
	var newSeafloor seafloor
	var movedWest, movedSouth bool
	movedWest, newSeafloor = s.moveWest()
	movedSouth, newSeafloor = newSeafloor.moveSouth()
	return movedWest || movedSouth, newSeafloor
}

//moveWest returns a new seafloor after every sea cucumber facing to the west is moved
func (s seafloor) moveWest() (bool, seafloor) {
	var moved bool
	empty := "."
	newSeafloor := make(seafloor)
	for l, c := range s {
		if *c == ">" {
			newLocation := location{x: l.x + 1, y: l.y}
			if s[newLocation] == nil {
				newLocation.x = 0
			}
			if *s[newLocation] == "." {
				newSeafloor[l] = &empty
				newSeafloor[newLocation] = c
				moved = true
				continue
			}
			newSeafloor[l] = c
			continue
		}
		if newSeafloor[l] == nil {
			newSeafloor[l] = c
		}
	}
	return moved, newSeafloor
}

//moveWest returns a new seafloor after every sea cucumber facing to the south is moved
func (s seafloor) moveSouth() (bool, seafloor) {
	var moved bool
	empty := "."
	newSeafloor := make(seafloor)
	for l, c := range s {
		if *c == "v" {
			newLocation := location{x: l.x, y: l.y + 1}
			if s[newLocation] == nil {
				newLocation.y = 0
			}
			if *s[newLocation] == "." {
				newSeafloor[newLocation] = c
				newSeafloor[l] = &empty
				moved = true
				continue
			}
			newSeafloor[l] = c
			continue
		}
		if newSeafloor[l] == nil {
			newSeafloor[l] = c
		}
	}
	return moved, newSeafloor
}

//parseInput returns a seafloor for a string input
func parseInput(input string) seafloor {
	lines, _ := utils.InputToSlice(input)
	floor := make(seafloor)
	for y := range lines {
		for x := range lines[y] {
			s := string(lines[y][x])
			floor[location{x: x, y: y}] = &s
		}
	}
	return floor
}

//print prints a seafloor for debugging
func (s seafloor) print() {
	maxX := 0
	maxY := 0
	for i := range s {
		if i.x > maxX {
			maxX = i.x
		}
		if i.y > maxY {
			maxY = i.y
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			fmt.Print(*s[location{x: x, y: y}])
		}
		fmt.Println()
	}
}
