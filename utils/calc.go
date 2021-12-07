package utils

//SumSlice returns the sum of all numbers of a slice
func SumSlice(slice []int) int {
	var sum int
	for i := range slice {
		sum += slice[i]
	}
	return sum
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
