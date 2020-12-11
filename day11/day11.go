package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay11Part1(stringListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay11Part2(stringListToSlice(input)))
}

//SolveDay11Part1 count the number of used seats after no seats change state with the rules for part1
func SolveDay11Part1(i []string) (s int) {
	curSeats := sliceToMap(i)
	curSum := 1
	for curSum != s {
		curSum = s
		curSeats = applySeatRulesPart1(curSeats)
		s = countSeats(curSeats)
	}
	return countSeats(curSeats)
}

//SolveDay11Part2 count the number of used seats after no seats change state with the rules for part1
func SolveDay11Part2(i []string) (s int) {
	curSeats := sliceToMap(i)
	curSum := 1
	for curSum != s {
		curSum = s
		curSeats = applySeatRulesPart2(curSeats)
		s = countSeats(curSeats)
	}
	return countSeats(curSeats)
}

//printMap prints the map to show the current state for debugging
func printMap(input map[int]map[int]string) (o string) {
	for rowI := 0; rowI < len(input); rowI++ {
		if rowI != 0{
			o = o+fmt.Sprintf("\n")
		}
		for i := 0; i < len(input[rowI]); i++ {
			o = o+fmt.Sprintf("%v", input[rowI][i])
		}
	}
	return o
}

//countSeats returns the used seats for a given map
func countSeats(input map[int]map[int]string) (seats int) {
	for _, row := range input {
		for _, seat := range row {
			if seat == "#" {
				seats++
			}
		}
	}
	return
}

//sliceToMap convert the string slice to a map
func sliceToMap(input []string) (seats map[int]map[int]string) {
	seats = make(map[int]map[int]string)
	for rowI, row := range input {
		seats[rowI] = make(map[int]string)
		for i, seat := range row {
			seats[rowI][i] = string(seat)
		}
	}
	return
}

//applySeatRulesPart1 the rules to the seats for part1 and return the new seat allocation
func applySeatRulesPart1(input map[int]map[int]string) (newSeats map[int]map[int]string) {
	newSeats = make(map[int]map[int]string)

	for rowI, row := range input {
		newSeats[rowI] = make(map[int]string)
		for i, seat := range row {
			if seat == "." {
				newSeats[rowI][i] = "."
				continue
			}
			if seat == "L" {
				newSeats[rowI][i] = "#"
				//left
				if row[i-1] == "#" {
					newSeats[rowI][i] = "L"
				}
				//right
				if row[i+1] == "#" {
					newSeats[rowI][i] = "L"
				}
				////up
				//left
				if input[rowI-1][i-1] == "#" {
					newSeats[rowI][i] = "L"
				}
				//middle
				if input[rowI-1][i] == "#" {
					newSeats[rowI][i] = "L"
				}
				//right
				if input[rowI-1][i+1] == "#" {
					newSeats[rowI][i] = "L"
				}
				////down
				//left
				if input[rowI+1][i-1] == "#" {
					newSeats[rowI][i] = "L"
				}
				//middle
				if input[rowI+1][i] == "#" {
					newSeats[rowI][i] = "L"
				}
				//right
				if input[rowI+1][i+1] == "#" {
					newSeats[rowI][i] = "L"
				}
				continue
			}
			if seat == "#" {
				seatCount := 0
				//left
				if row[i-1] == "#" {
					seatCount++
				}
				//right
				if row[i+1] == "#" {
					seatCount++
				}
				////up
				//left
				if input[rowI-1][i-1] == "#" {
					seatCount++
				}
				//middle
				if input[rowI-1][i] == "#" {
					seatCount++
				}
				//right
				if input[rowI-1][i+1] == "#" {
					seatCount++
				}
				////down
				//left
				if input[rowI+1][i-1] == "#" {
					seatCount++
				}
				//middle
				if input[rowI+1][i] == "#" {
					seatCount++
				}
				//right
				if input[rowI+1][i+1] == "#" {
					seatCount++
				}
				if seatCount >= 4 {
					newSeats[rowI][i] = "L"
				} else {
					newSeats[rowI][i] = "#"
				}
			}
		}
	}
	return
}

//applySeatRulesPart2 the rules to the seats for part2 and return the new seat allocation
func applySeatRulesPart2(input map[int]map[int]string) (newSeats map[int]map[int]string) {
	newSeats = make(map[int]map[int]string)

	for rowI, row := range input {
		newSeats[rowI] = make(map[int]string)
		for i, seat := range row {
			if seat == "." {
				newSeats[rowI][i] = "."
				continue
			}
			if seat == "L" {
				newSeats[rowI][i] = "#"
				//left
				tempI := 1
				for {
					if row[i-tempI] == "#" {
						newSeats[rowI][i] = "L"
						break
					} else if row[i-tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//right
				tempI = 1
				for {
					if row[i+tempI] == "#" {
						newSeats[rowI][i] = "L"
						break
					} else if row[i+tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				////up
				//left
				tempI = 1
				for {
					if input[rowI-tempI][i-tempI] == "#" {
						newSeats[rowI][i] = "L"
						break
					} else if input[rowI-tempI][i-tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//middle
				tempI = 1
				for {
					if input[rowI-tempI][i] == "#" {
						newSeats[rowI][i] = "L"
						break
					} else if input[rowI-tempI][i] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//right
				tempI = 1
				for {
					if input[rowI-tempI][i+tempI] == "#" {
						newSeats[rowI][i] = "L"
						break
					} else if input[rowI-tempI][i+tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				////down
				//left
				tempI = 1
				for {
					if input[rowI+tempI][i-tempI] == "#" {
						newSeats[rowI][i] = "L"
						break
					} else if input[rowI+tempI][i-tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//middle
				tempI = 1
				for {
					if input[rowI+tempI][i] == "#" {
						newSeats[rowI][i] = "L"
						break
					} else if input[rowI+tempI][i] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//right
				tempI = 1
				for {
					if input[rowI+tempI][i+tempI] == "#" {
						newSeats[rowI][i] = "L"
						break
					} else if input[rowI+tempI][i+tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				continue
			}
			if seat == "#" {
				seatCount := 0
				//left
				tempI := 1
				for {
					if row[i-tempI] == "#" {
						seatCount++
						break
					} else if row[i-tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//right
				tempI = 1
				for {
					if row[i+tempI] == "#" {
						seatCount++
						break
					} else if row[i+tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				////up
				//left
				tempI = 1
				for {
					if input[rowI-tempI][i-tempI] == "#" {
						seatCount++
						break
					} else if input[rowI-tempI][i-tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//middle
				tempI = 1
				for {
					if input[rowI-tempI][i] == "#" {
						seatCount++
						break
					} else if input[rowI-tempI][i] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//right
				tempI = 1
				for {
					if input[rowI-tempI][i+tempI] == "#" {
						seatCount++
						break
					} else if input[rowI-tempI][i+tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				////down
				//left
				tempI = 1
				for {
					if input[rowI+tempI][i-tempI] == "#" {
						seatCount++
						break
					} else if input[rowI+tempI][i-tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//middle
				tempI = 1
				for {
					if input[rowI+tempI][i] == "#" {
						seatCount++
						break
					} else if input[rowI+tempI][i] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				//right
				tempI = 1
				for {
					if input[rowI+tempI][i+tempI] == "#" {
						seatCount++
						break
					} else if input[rowI+tempI][i+tempI] == "." {
						tempI++
						continue
					} else {
						break
					}
				}
				if seatCount >= 5 {
					newSeats[rowI][i] = "L"
				} else {
					newSeats[rowI][i] = "#"
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
