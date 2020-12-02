package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var invalidMinMaxError = errors.New("invalid min max values")

//SolveDay1Part1 count valid password that contains the right number of chars
func SolveDay2Part1(passwordlist []string) (sum int) {
	for _, p := range passwordlist {
		vmin, vmax, char, password, err := getPasswordValues(p)
		if err == nil {
			count := strings.Count(password, char)
			if count <= vmax && count >= vmin {
				sum++
			}
		}
	}
	return
}

//SolveDay1Part2 count valid password that contains the right chars on the right place
func SolveDay2Part2(passwordlist []string) (sum int) {
	for _, p := range passwordlist {
		vmin, vmax, char, password, err := getPasswordValues(p)
		if err == nil {
			passwordslice := strings.Split(password, "")
			if len(passwordslice) >= vmin-1 && len(passwordslice) >= vmax-1 {
				if passwordslice[vmin-1] == char && passwordslice[vmax-1] != char || passwordslice[vmin-1] != char && passwordslice[vmax-1] == char {
					sum++
				}
			}
		}
	}
	return
}

//getPasswordValues extract values from password line
func getPasswordValues(s string) (int, int, string, string, error) {
	t := strings.Split(s, " ")
	v := strings.Split(t[0], "-")
	vmin, err := strconv.Atoi(v[0])
	if err != nil {
		return 0, 0, "", "", err
	}
	vmax, _ := strconv.Atoi(v[1])
	if err != nil {
		return 0, 0, "", "", err
	}
	char := strings.Split(t[1], ":")[0]
	if vmin == 0 || vmax == 0 || vmin > vmax {
		return 0, 0, "", "", invalidMinMaxError
	}
	return vmin, vmax, char, t[2], nil
}

//listToSlice converts the list of numbers (each number one row) to a slice
func listToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}

func main() {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay2Part1(listToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay2Part2(listToSlice(input)))

}
