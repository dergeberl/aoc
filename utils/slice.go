package utils

import (
	"strconv"
	"strings"
)

//InputToIntSlice returns a list of numbers from the input
func InputToIntSlice(input string) ([]int, error) {
	var i []int
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		i = append(i, lineInt)
	}
	return i, nil
}

//InputToSlice returns a list of strings from the input
func InputToSlice(input string) ([]string, error) {
	return strings.Split(strings.TrimSuffix(input, "\n"), "\n"), nil
}

//RemoveFromIntSlice removes a given position from an int slice NOT ORDER SAVE
func RemoveFromIntSlice(s []int, i int) []int {
	l := len(s)
	if l <= i {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
