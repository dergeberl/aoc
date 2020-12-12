package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//step contains the action and the number of steps or the number of degrees
type step struct {
	action string
	number int
}

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay12Part1(stringListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay12Part2(stringListToSlice(input)))
}

//SolveDay12Part1 move the ship in the to the given instructions
func SolveDay12Part1(i []string) (s int) {
	steps := sliceToStepSlice(i)
	var east, north int
	currentDirection := 90
	for _, step := range steps {
		switch step.action {
		case "N":
			north += step.number
		case "S":
			north -= step.number
		case "E":
			east += step.number
		case "W":
			east -= step.number
		case "L":
			currentDirection -= step.number
			if currentDirection < 0 {
				currentDirection += 360
			}
		case "R":
			currentDirection += step.number
			if currentDirection > 359 {
				currentDirection -= 360
			}
		case "F":
			if currentDirection % 90 == 0{
				switch currentDirection {
				case 0:
					north += step.number
				case 90:
					east += step.number
				case 180:
					north -= step.number
				case 270:
					east -= step.number
				}
			}
		}
	}
	if east < 0 {
		east = east*(-1)
	}
	if north < 0 {
		north = north*(-1)
	}
	return east+north
}

//SolveDay12Part2 move the ship in the to the given instructions with waypoints
func SolveDay12Part2(i []string) (s int) {
	steps := sliceToStepSlice(i)
	var east, north  int
	wayEast, wayNorth := 10, 1
	for _, step := range steps {
		switch step.action {
		case "N":
			wayNorth += step.number
		case "S":
			wayNorth -= step.number
		case "E":
			wayEast += step.number
		case "W":
			wayEast -= step.number
		case "L":
			count := step.number/90
			for i :=0; i <count; i++{
				wayEastCur := wayNorth*(-1)
				wayNorthCur := wayEast
				wayNorth, wayEast = wayNorthCur, wayEastCur
			}
		case "R":
			count := step.number/90
			for i :=0; i <count; i++{
				wayEastCur := wayNorth
				wayNorthCur := wayEast*(-1)
				wayNorth, wayEast = wayNorthCur, wayEastCur
			}
		case "F":
			north += wayNorth * step.number
			east += wayEast * step.number
		}
	}
	if east < 0 {
		east = east*(-1)
	}
	if north < 0 {
		north = north*(-1)
	}
	return east+north
}

//sliceToStepSlice convert a string slice to a step slice
func sliceToStepSlice(i []string) (steps []step) {
	for _, line := range i {
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil
		}
		curStep := step{
			action: string(line[0]),
			number: num,
		}
		steps = append(steps, curStep)
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
