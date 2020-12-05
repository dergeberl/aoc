package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//SolveDay5Part1 returns the number of the highest seatId
func SolveDay5Part1(boardingPass []string) (highestNumber int) {
	for _, p := range boardingPass {
		s := getSeatID(p)
		if s > highestNumber {
			highestNumber = s
		}
	}
	return
}

//SolveDay5Part2 returns my seat (the unused seat with the id-1 and id+2 is used)
func SolveDay5Part2(boardingPass []string) (mySeat int) {
	seatIds := make(map[int]bool)
	for _, p := range boardingPass {
		seatIds[getSeatID(p)] = true
	}
	for i, _ := range seatIds {
		if !seatIds[i+1] && seatIds[i+2] {
			return i + 1
		}
	}
	return
}

//listToSlice converts the list of strings (each string one row) to a slice
func listToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}

// getSeatValue returns row, column, seatId of a given seat
func getSeatValue(s string) (row int, column int, seatId int) {
	if len(s) != 10 {
		return 0, 0, 0
	}
	if (s[0] != 'F' && s[0] != 'B') ||
		(s[1] != 'F' && s[1] != 'B') ||
		(s[2] != 'F' && s[2] != 'B') ||
		(s[3] != 'F' && s[3] != 'B') ||
		(s[4] != 'F' && s[4] != 'B') ||
		(s[5] != 'F' && s[5] != 'B') ||
		(s[6] != 'F' && s[6] != 'B') ||
		(s[7] != 'L' && s[7] != 'R') ||
		(s[8] != 'L' && s[8] != 'R') ||
		(s[9] != 'L' && s[9] != 'R') {
		return 0, 0, 0
	}
	if s[0] == 'B' {
		row += 64
	}
	if s[1] == 'B' {
		row += 32
	}
	if s[2] == 'B' {
		row += 16
	}
	if s[3] == 'B' {
		row += 8
	}
	if s[4] == 'B' {
		row += 4
	}
	if s[5] == 'B' {
		row += 2
	}
	if s[6] == 'B' {
		row++
	}
	if s[7] == 'R' {
		column += 4
	}
	if s[8] == 'R' {
		column += 2
	}
	if s[9] == 'R' {
		column++
	}
	seatId = (row * 8) + column
	return
}

// getSeatID returns seatId of a given seat
func getSeatID(s string) (seatId int) {
	if len(s) != 10 {
		return 0
	}
	cur := 512
	for _, c := range s {
		if c == 82 || c == 66 {
			seatId += cur
		} else if c != 70 && c != 76 {
			return 0
		}
		cur = cur / 2
	}
	return
}

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay5Part1(listToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay5Part2(listToSlice(input)))
}
