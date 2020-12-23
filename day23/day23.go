package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay23Part1(input))
	fmt.Printf("Part 2: %v\n", SolveDay23Part2(input))
}

//SolveDay23Part1 plays 100 runs with 9 given cups and returns the cup order after cup 1
func SolveDay23Part1(input string) (s int) {
	return playCupGame(input, len(input), 100, false)
}

//SolveDay23Part2 plays 10000000 runs with 1000000 cups and returns the product of the to cup numbers after cup 1
func SolveDay23Part2(input string) (s int) {
	return playCupGame(input, 1000000, 10000000, true)
}

//playCupGame runs the cup game with the defined startNumbers, a totalNumbers and runs and returns the output in the part2 oder part2 format
func playCupGame(startNumbers string, totalNumbers int, runs int, part2 bool) (solution int) {
	r := ring.New(totalNumbers)
	destinationFinder := make(map[int]*ring.Ring)

	for _, num := range startNumbers {
		numInt, _ := strconv.Atoi(string(num))
		r.Value = numInt
		destinationFinder[numInt] = r
		r = r.Next()
	}

	for i := len(startNumbers) + 1; i <= totalNumbers; i++ {
		r.Value = i
		destinationFinder[i] = r
		r = r.Next()
	}

	for i := 0; i < runs; i++ {
		destination := r.Value.(int) - 1
		pick := r.Unlink(3)

		for {
			if destination < 1 {
				destination = totalNumbers
			}
			if destination != pick.Value && destination != pick.Next().Value && destination != pick.Next().Next().Value {
				break
			}
			destination--
		}
		destinationFinder[destination].Link(pick)
		r = r.Next()
	}

	if part2 {
		return destinationFinder[1].Next().Value.(int) * destinationFinder[1].Next().Next().Value.(int)
	} else {
		var result int
		r = destinationFinder[1].Next()
		for i := 0; i < r.Len()-1; i++ {
			result *= 10
			result += r.Value.(int)
			r = r.Next()
		}
		return result
	}
}
