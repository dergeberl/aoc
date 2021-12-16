package utils

import "sort"

//SumSlice returns the sum of all numbers of a slice
func SumSlice(slice []int) int {
	var sum int
	for i := range slice {
		sum += slice[i]
	}
	return sum
}

//ProductSlice returns the product of all numbers of a slice
//returns 1 if slice is empty
//if the len of the slice is 1 it returns the number of this one item
func ProductSlice(slice []int) int {
	product := 1
	for i := range slice {
		product *= slice[i]
	}
	return product
}

//MinimumSlice returns the lowest number from the slice (0 if slice is empty)
func MinimumSlice(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	sort.Ints(slice)
	return slice[0]
}

//MaximumSlice returns the highest number from the slice (0 if slice is empty)
func MaximumSlice(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	sort.Ints(slice)
	return slice[len(slice)-1]
}

//GetDiff returns the difference between a and b
func GetDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

//GetGaussscheSummenformel returns sum with the Gaußsche Summenformel
func GetGaussscheSummenformel(n int) int {
	return (n * (n + 1)) / 2
}
