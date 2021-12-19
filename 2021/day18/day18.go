package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type snailFishNumber struct {
	x                *int
	y                *int
	xSnailFishNumber *snailFishNumber
	ySnailFishNumber *snailFishNumber
	parent           *snailFishNumber
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay18Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay18Part2(string(input)))
}

//SolveDay18Part1 returns the magnitude after add all snailFishNumbers together
func SolveDay18Part1(input string) int {
	s := parseInput(input)
	var sum *snailFishNumber
	for i := range s {
		s[i].reduce()
		sum = sum.add(s[i])
		sum.reduce()
	}
	return sum.magnitude()
}

//SolveDay18Part2 returns the maxMagnitude which is possible by adding 2 snailFishNumbers
func SolveDay18Part2(input string) int {
	s := parseInput(input)
	var maxMagnitude int
	for i1 := range s {
		for i2 := range s {
			if i1 == i2 {
				continue
			}
			n1 := s[i1].copy(nil)
			n2 := s[i2].copy(nil)
			sum := n1.add(n2)
			sum.reduce()
			if sum.magnitude() > maxMagnitude {
				maxMagnitude = sum.magnitude()
			}

		}
	}
	return maxMagnitude
}

//reduce the snailFishNumber to the minimal possible
func (s *snailFishNumber) reduce() {
	for {
		stop := true

		for s.reduceExplode() {
			stop = false
		}
		if s.reduceSplit() {
			stop = false
		}

		if stop {
			break
		}
	}
}

//reduceExplode reduce the snailFishNumber by explode a number which is 4 times nested (only the leftest one)
func (s *snailFishNumber) reduceExplode() bool {
	e := s.getFoursNumber(0)

	if e == nil {
		return false
	}
	if e == e.parent.xSnailFishNumber {
		e.parent.xSnailFishNumber = nil
		n := 0
		e.parent.x = &n

		if e.parent.y != nil {
			tmp := *e.parent.y + *e.y
			e.parent.y = &tmp
		} else {
			e.parent.ySnailFishNumber.addRight2(*e.y)
		}
		e.parent.addLeft(*e.x)
		return true
	}
	if e == e.parent.ySnailFishNumber {
		e.parent.ySnailFishNumber = nil
		n := 0
		e.parent.y = &n

		if e.parent.x != nil {
			tmp := *e.parent.x + *e.x
			e.parent.x = &tmp
		} else {
			e.parent.xSnailFishNumber.addLeft2(*e.y)
		}
		e.parent.addRight(*e.y)
		return true
	}
	return false
}

//getFoursNumber gets the snailFishNumber which is 4 times nested (only the leftest one)
func (s *snailFishNumber) getFoursNumber(i int) *snailFishNumber {
	if i == 4 {
		if s.xSnailFishNumber != nil {
			return s.xSnailFishNumber.getFoursNumber(i)
		}
		if s.ySnailFishNumber != nil {
			return s.ySnailFishNumber.getFoursNumber(i)
		}
		return s
	}
	i++
	var found *snailFishNumber
	if s.xSnailFishNumber != nil {
		found = s.xSnailFishNumber.getFoursNumber(i)
	}
	if found != nil {
		return found
	}
	if s.ySnailFishNumber != nil {
		found = s.ySnailFishNumber.getFoursNumber(i)
	}
	return found
}

func (s *snailFishNumber) addLeft(i int) {
	if s == nil || s.parent == nil {
		return
	}

	if s == s.parent.xSnailFishNumber {
		s.parent.addLeft(i)
		return
	}
	if s.parent.x != nil {
		i += *s.parent.x
		s.parent.x = &i
		return
	}
	s.parent.xSnailFishNumber.addLeft2(i)
}

func (s *snailFishNumber) addLeft2(i int) {
	if s == nil {
		return
	}
	if s.y != nil {
		i += *s.y
		s.y = &i
		return
	}
	if s.ySnailFishNumber != nil {
		s.ySnailFishNumber.addLeft2(i)
	}
}

func (s *snailFishNumber) addRight(i int) {
	if s == nil || s.parent == nil {
		return
	}

	if s == s.parent.ySnailFishNumber {
		s.parent.addRight(i)
		return
	}
	if s.parent.y != nil {
		i += *s.parent.y
		s.parent.y = &i
		return
	}
	s.parent.ySnailFishNumber.addRight2(i)
}

func (s *snailFishNumber) addRight2(i int) {
	if s == nil {
		return
	}
	if s.x != nil {
		i += *s.x
		s.x = &i
		return
	}
	if s.xSnailFishNumber != nil {
		s.xSnailFishNumber.addRight2(i)
	}
}

//reduceSplit splits numbers which are higher than 9 into a new snailFishNumber
func (s *snailFishNumber) reduceSplit() bool {
	if s.xSnailFishNumber != nil {
		if s.xSnailFishNumber.reduceSplit() {
			return true
		}
	}
	if s.x != nil && *s.x > 9 {
		num := *s.x
		numX := num / 2
		numY := num - numX
		s.x = nil
		s.xSnailFishNumber = &snailFishNumber{
			x:      &numX,
			y:      &numY,
			parent: s,
		}
		return true
	}

	if s.ySnailFishNumber != nil {
		if s.ySnailFishNumber.reduceSplit() {
			return true
		}
	}

	if s.y != nil && *s.y > 9 {
		num := *s.y
		numX := num / 2
		numY := num - numX
		s.y = nil
		s.ySnailFishNumber = &snailFishNumber{
			x:      &numX,
			y:      &numY,
			parent: s,
		}
		return true
	}
	return false
}

//add 2 snailFishNumber together and returns the new one
func (s *snailFishNumber) add(add *snailFishNumber) *snailFishNumber {
	if s == nil {
		return add
	}
	if add == nil {
		return s
	}
	n := &snailFishNumber{
		xSnailFishNumber: s,
		ySnailFishNumber: add,
		parent:           nil,
	}
	s.parent = n
	add.parent = n
	return n
}

//toString returns a printable string (much needed while development)
func (s snailFishNumber) toString() string {
	var str string
	str += "["
	if s.x != nil {
		str += fmt.Sprint(*s.x)
		str += ","
	}
	if s.xSnailFishNumber != nil {
		str += s.xSnailFishNumber.toString()
		str += ","
	}
	if s.y != nil {
		str += fmt.Sprint(*s.y)
		str += ""
	}
	if s.ySnailFishNumber != nil {
		str += s.ySnailFishNumber.toString()
	}
	str += "]"

	return str
}

// returns the magnitude of a snailFishNumber (x*3 + y*4). Follows the nested objects if needed.
func (s snailFishNumber) magnitude() int {
	var tmpX, tmpY int
	if s.x != nil {
		tmpX = *s.x
	} else {
		tmpX = s.xSnailFishNumber.magnitude()
	}
	if s.y != nil {
		tmpY = *s.y
	} else {
		tmpY = s.ySnailFishNumber.magnitude()
	}
	return tmpX*3 + tmpY*2
}

//copy returns a copy of a snailFishNumber
func (s *snailFishNumber) copy(parent *snailFishNumber) *snailFishNumber {
	n := snailFishNumber{}
	n.parent = parent
	if s.x != nil {
		x := *s.x
		n.x = &x
	}
	if s.y != nil {
		y := *s.y
		n.y = &y
	}
	if s.xSnailFishNumber != nil {
		n.xSnailFishNumber = s.xSnailFishNumber.copy(&n)
	}
	if s.ySnailFishNumber != nil {
		n.ySnailFishNumber = s.ySnailFishNumber.copy(&n)
	}
	return &n
}

//parseInput returns a list of snailFishNumbers from a multiline string input
func parseInput(input string) []*snailFishNumber {
	lines, _ := utils.InputToSlice(input)
	var numbers []*snailFishNumber
	for i := range lines {
		numbers = append(numbers, getSnailFishNumberFromString(lines[i], nil))
	}
	return numbers
}

//parseInput returns a snailFishNumbers from a string input
func getSnailFishNumberFromString(s string, parent *snailFishNumber) *snailFishNumber {
	var num snailFishNumber
	num.parent = parent

	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	var open, start int
	for i, r := range s {
		if r == '[' {
			open++
		}
		if r == ']' {
			open--
		}
		if open == 0 {
			if r == ',' {
				start = i + 1
				continue
			}
			intNum, err := strconv.Atoi(s[start : i+1])
			if err == nil {
				if num.x != nil || num.xSnailFishNumber != nil {
					num.y = &intNum
				} else {
					num.x = &intNum
				}
			}
			if err != nil {
				tepNum := getSnailFishNumberFromString(s[start:i+1], &num)
				if num.x != nil || num.xSnailFishNumber != nil {
					num.ySnailFishNumber = tepNum
				} else {
					num.xSnailFishNumber = tepNum
				}
			}
			start = i + 1
		}
	}
	return &num
}
