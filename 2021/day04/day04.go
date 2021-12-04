package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type board struct {
	lines   []row
	columns []row
}

type row []int

var spaceRegex = regexp.MustCompile(`\s+`)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay04Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay04Part2(string(input)))
}

//SolveDay04Part1 returns the sum of the unseen numbers on the bingo board multiplied by the last number
func SolveDay04Part1(input string) int {
	numbers, boards := parseInput(input)

	for i := range numbers {
		for b := range boards {
			if boards[b].checkWin(numbers[:i+1]) {
				return boards[b].calcUnmarked(numbers[:i+1]) * numbers[i]
			}
		}
	}

	return 0
}

//SolveDay04Part2 return the score like in SolveDay04Part1 for the last winning bingo board
func SolveDay04Part2(input string) int {
	numbers, boards := parseInput(input)

	for i := range numbers {
		for b := 0; len(boards) > b; b++ {
			if boards[b].checkWin(numbers[:i+1]) {
				if len(boards) == 1 {
					return boards[0].calcUnmarked(numbers[:i+1]) * numbers[i]
				}
				boards = removeFromBoards(boards, b)
				b--
			}
		}
	}

	return 0
}

//checkWin checks if the board wins
func (b board) checkWin(numbers []int) bool {
	for i := range b.columns {
		if b.columns[i].checkRow(numbers) {
			return true
		}
	}
	for i := range b.lines {
		if b.lines[i].checkRow(numbers) {
			return true
		}
	}
	return false
}

//checkWin returns the sum of all unmarked numbers
func (b board) calcUnmarked(numbers []int) int {
	var sum int
	for i := range b.lines {
		sum += utils.SumSlice(b.lines[i].removeFromRow(numbers))
	}
	return sum
}

//checkRow returns true if all numbers of a row are represented in the numbers
func (r row) checkRow(numbers []int) bool {
	return len(r.removeFromRow(numbers)) == 0
}

//checkWin returns a new row without the given numbers
func (r row) removeFromRow(numbers []int) []int {
	var copyRow []int
	copyRow = append(copyRow, r...)

	for n := range numbers {
		for i := range copyRow {
			if numbers[n] == copyRow[i] {
				copyRow = utils.RemoveFromIntSlice(copyRow, i)
				break
			}
		}
		if len(copyRow) == 0 {
			return []int{}
		}
	}
	return copyRow
}

func removeFromBoards(s []board, i int) []board {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func parseInput(input string) ([]int, []board) {
	var numbers []int
	var boards []board

	for _, block := range strings.Split(input, "\n\n") {
		blockLines := strings.Split(block, "\n")
		if len(blockLines) != 5 {
			if len(blockLines) == 1 {
				numbers = parseNumbers(blockLines[0])
			}
			continue
		}
		boards = append(boards, parseBoard(blockLines))
	}
	return numbers, boards
}

func parseBoard(lines []string) board {
	var b board

	b.lines = make([]row, 5)
	b.columns = make([]row, 5)

	for l := range lines {
		lineNumbers := strings.Split(spaceRegex.ReplaceAllString(strings.TrimSpace(lines[l]), " "), " ")
		for n := range lineNumbers {
			i, _ := strconv.Atoi(lineNumbers[n])
			b.lines[l] = append(b.lines[l], i)
			b.columns[n] = append(b.columns[n], i)
		}
	}
	return b
}

func parseNumbers(line string) []int {
	var numbers []int
	parsedNumbers := strings.Split(strings.TrimSpace(line), ",")
	for i := range parsedNumbers {
		numberInt, _ := strconv.Atoi(parsedNumbers[i])
		numbers = append(numbers, numberInt)
	}
	return numbers
}
