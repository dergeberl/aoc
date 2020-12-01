package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

//SolveDay1Part1 search for two numbers in a slice that sum is 2020 and multiply them
func SolveDay1Part1(nums []int) int {
	for _, num := range nums {
		for _, num2 := range nums {
			if (num + num2) == 2020 {
				return num * num2
			}
		}
	}
	return 0
}

//SolveDay1Part2 search for three numbers in a slice that sum is 2020 and multiply them
func SolveDay1Part2(nums []int) int {
	sort.Ints(nums)
	for _, num := range nums {
		for _, num2 := range nums {
			if (num + num2) > 2020 {
				break
			}
			for _, num3 := range nums {
				if (num + num2 + num3) == 2020 {
					return num * num2 * num3
				}
				if (num + num2 + num3) > 2020 {
					break
				}
			}
		}
	}
	return 0
}

//listToSlice converts the list of numbers (each number one row) to a slice
func listToSlice(list string) (i []int) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		lineint, err := strconv.Atoi(line)
		if err != nil {
			return nil
		}
		i = append(i, lineint)
	}
	return
}

func main() {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay1Part1(listToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay1Part2(listToSlice(input)))

}
