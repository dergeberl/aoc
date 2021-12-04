package utils

//SumSlice returns the sum of all numbers of a slice
func SumSlice(slice []int) int {
	var sum int
	for i := range slice {
		sum += slice[i]
	}
	return sum
}
