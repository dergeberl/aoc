package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type point struct {
	x, y, z int
}

type scanner map[point]bool

type absoluteMap map[point]bool

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay19Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay19Part2(string(input)))
}

//SolveDay19Part1 returns the number of points after align all scanner
func SolveDay19Part1(input string) int {
	s := parseInput(input)
	abs := absoluteMap(s[0])
	_, abs = abs.solve(s)

	return len(abs)
}

//SolveDay19Part2 returns the largest Manhattan distance between any 2 points
func SolveDay19Part2(input string) int {
	s := parseInput(input)
	abs := absoluteMap(s[0])
	scannerPoints, _ := abs.solve(s)
	var max int
	for i1 := range scannerPoints {
		for i2 := range scannerPoints {
			sum := math.Abs(float64(scannerPoints[i1].x-scannerPoints[i2].x)) +
				math.Abs(float64(scannerPoints[i1].y-scannerPoints[i2].y)) +
				math.Abs(float64(scannerPoints[i1].z-scannerPoints[i2].z))
			if max < int(sum) {
				max = int(sum)
			}
		}
	}
	return max
}

//solve search the position for each scanner and add it to the absoluteMap
// returns a list of points (where the scanner are located) and a new absoluteMap
func (a absoluteMap) solve(s []scanner) ([]point, absoluteMap) {
	var scannerPoints []point
	abs := absoluteMap(s[0])
	for {
		var newScanner []scanner
		for scan := range s {
			x, y, z, sc := abs.findOffset(s[scan])
			for p := range sc {
				abs[point{
					x: p.x + x,
					y: p.y + y,
					z: p.z + z,
				}] = true
			}
			if len(sc) == 0 {
				newScanner = append(newScanner, s[scan])
				continue
			}
			scannerPoints = append(scannerPoints, point{x: x, y: y, z: z})
		}
		s = newScanner
		if len(s) == 0 {
			break
		}
	}
	return scannerPoints, a
}

//getPossibility returns all possible orientations for a scanner
func (s scanner) getPossibility() []scanner {
	possibilities := make([]scanner, 24)
	possibilities[0] = make(scanner)
	possibilities[1] = make(scanner)
	possibilities[2] = make(scanner)
	possibilities[3] = make(scanner)
	possibilities[4] = make(scanner)
	possibilities[5] = make(scanner)
	possibilities[6] = make(scanner)
	possibilities[7] = make(scanner)
	possibilities[8] = make(scanner)
	possibilities[9] = make(scanner)
	possibilities[10] = make(scanner)
	possibilities[11] = make(scanner)
	possibilities[12] = make(scanner)
	possibilities[13] = make(scanner)
	possibilities[14] = make(scanner)
	possibilities[15] = make(scanner)
	possibilities[16] = make(scanner)
	possibilities[17] = make(scanner)
	possibilities[18] = make(scanner)
	possibilities[19] = make(scanner)
	possibilities[20] = make(scanner)
	possibilities[21] = make(scanner)
	possibilities[22] = make(scanner)
	possibilities[23] = make(scanner)
	for p, _ := range s {
		x, y, z := p.x, p.y, p.z
		possibilities[0][point{x, y, z}] = true
		possibilities[1][point{x, -z, y}] = true
		possibilities[2][point{x, -y, -z}] = true
		possibilities[3][point{x, z, -y}] = true
		possibilities[4][point{-x, -y, z}] = true
		possibilities[5][point{-x, -z, -y}] = true
		possibilities[6][point{-x, y, -z}] = true
		possibilities[7][point{-x, z, y}] = true
		possibilities[8][point{-z, x, -y}] = true
		possibilities[9][point{y, x, -z}] = true
		possibilities[10][point{z, x, y}] = true
		possibilities[11][point{-y, x, z}] = true
		possibilities[12][point{z, -x, -y}] = true
		possibilities[13][point{y, -x, z}] = true
		possibilities[14][point{-z, -x, y}] = true
		possibilities[15][point{-y, -x, -z}] = true
		possibilities[16][point{-y, -z, x}] = true
		possibilities[17][point{z, -y, x}] = true
		possibilities[18][point{y, z, x}] = true
		possibilities[19][point{-z, y, x}] = true
		possibilities[20][point{z, y, -x}] = true
		possibilities[21][point{-y, z, -x}] = true
		possibilities[22][point{-z, -y, -x}] = true
		possibilities[23][point{y, -z, -x}] = true
	}
	return possibilities
}

//findOffset calc the offset between points and returns x,y,z and the scanner if it found minimum 2 matching points
func (a absoluteMap) findOffset(s scanner) (int, int, int, scanner) {
	possibility := s.getPossibility()
	points := make(map[point]int)
	for p := range possibility {
		for possiblePoint := range possibility[p] {
			for mapPoint := range a {
				points[point{
					x: mapPoint.x - possiblePoint.x,
					y: mapPoint.y - possiblePoint.y,
					z: mapPoint.z - possiblePoint.z,
				}]++
				if points[point{
					x: mapPoint.x - possiblePoint.x,
					y: mapPoint.y - possiblePoint.y,
					z: mapPoint.z - possiblePoint.z,
				}] > 11 {
					return mapPoint.x - possiblePoint.x, mapPoint.y - possiblePoint.y, mapPoint.z - possiblePoint.z, possibility[p]
				}
			}
		}
	}
	return 0, 0, 0, nil
}

//parseInput returns a list of scanner from the string input
func parseInput(input string) []scanner {
	s := strings.Split(input, "\n\n")
	scannerList := make([]scanner, 0)
	for i := range s {
		scan := make(scanner)
		lines, _ := utils.InputToSlice(s[i])
		for l := range lines {
			if strings.HasPrefix(lines[l], "---") {
				continue
			}
			nums := strings.Split(lines[l], ",")
			if len(nums) != 3 {
				parseInput("wrong input")
			}
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			z, _ := strconv.Atoi(nums[2])
			scan[point{x: x, y: y, z: z}] = true
		}
		scannerList = append(scannerList, scan)
	}
	return scannerList
}
