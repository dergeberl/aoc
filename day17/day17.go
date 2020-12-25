package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay17Part1(stringListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay17Part2(stringListToSlice(input)))
}

type coordinates struct {
	x int
	y int
	z int
	w int
}

//SolveDay17Part1 returns the number of active cubes after the six-cycle boot process (with 3 dimensional space)
func SolveDay17Part1(input []string) (s int) {
	runs := 6
	startX := len(input)
	startY := len(input[0])

	curState := make(map[coordinates]bool)
	for y, line := range input {
		for x, state := range line {
			if string(state) == "#" {
				curState[coordinates{x: x, y: y, z: 0, w: 0}] = true
			}
		}
	}
	newState := make(map[coordinates]bool)
	for key, value := range curState {
		newState[key] = value
	}
	for i := 1; i <= runs; i++ {
		for iZ := 0 - i; iZ <= i; iZ++ {
			for iY := 0 - i; iY < startY+i; iY++ {
				for iX := 0 - i; iX < startX+i; iX++ {
					if curState[coordinates{x: iX, y: iY, z: iZ, w: 0}] {
						// enabled
						count := countNeighbors(curState, iZ, iY, iX)
						if count != 2 && count != 3 {
							newState[coordinates{x: iX, y: iY, z: iZ}] = false
						}
					} else {
						// disabled
						count := countNeighbors(curState, iZ, iY, iX)
						if count == 3 {
							newState[coordinates{x: iX, y: iY, z: iZ, w: 0}] = true
						}
					}
				}
			}
		}
		for key, value := range newState {
			curState[key] = value
		}
	}

	//count active
	for _, state := range curState {
		if state {
			s++
		}
	}
	return
}

//SolveDay17Part2 returns the number of active cubes after the six-cycle boot process (with 4 dimensional space)
func SolveDay17Part2(input []string) (s int) {
	runs := 6
	startX := len(input)
	startY := len(input[0])

	curState := make(map[coordinates]bool)
	for y, line := range input {
		for x, state := range line {
			if string(state) == "#" {
				curState[coordinates{x: x, y: y, z: 0, w: 0}] = true
			}
		}
	}
	newState := make(map[coordinates]bool)
	for key, value := range curState {
		newState[key] = value
	}
	for i := 1; i <= runs; i++ {
		for iZ := 0 - i; iZ <= i; iZ++ {
			for iW := 0 - i; iW <= i; iW++ {
				for iY := 0 - i; iY < startY+i; iY++ {
					for iX := 0 - i; iX < startX+i; iX++ {
						if curState[coordinates{x: iX, y: iY, z: iZ, w: iW}] {
							// enabled
							count := countNeighbors(curState, iZ, iY, iX, iW)
							if count != 2 && count != 3 {
								newState[coordinates{x: iX, y: iY, z: iZ, w: iW}] = false
							}
						} else {
							// disabled
							count := countNeighbors(curState, iZ, iY, iX, iW)
							if count == 3 {
								newState[coordinates{x: iX, y: iY, z: iZ, w: iW}] = true
							}
						}
					}
				}
			}
		}

		for key, value := range newState {
			curState[key] = value
		}
	}
	//count active
	for _, state := range curState {
		if state {
			s++
		}
	}
	return
}

//countNeighbors returns the number of active neighbor cubes (maximum 4)
func countNeighbors(curState map[coordinates]bool, nums ...int) (count int) {
	var z, y, x, w int
	check4d := false
	if len(nums) == 3 {
		z = nums[0]
		y = nums[1]
		x = nums[2]
		w = 0
	} else if len(nums) == 4 {
		z = nums[0]
		y = nums[1]
		x = nums[2]
		w = nums[3]
		check4d = true
	} else {
		return 0
	}
	for iz := z - 1; iz <= z+1; iz++ {
		for iw := w - 1; iw <= w+1; iw++ {
			if iw != 0 && !check4d {
				continue
			}
			for iy := y - 1; iy <= y+1; iy++ {
				for ix := x - 1; ix <= x+1; ix++ {
					if iz == z && iy == y && ix == x && iw == w {
						continue
					}
					if curState[coordinates{x: ix, y: iy, z: iz, w: iw}] {
						count++
						if count >= 4 {
							return
						}
					}

				}
			}
		}
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
