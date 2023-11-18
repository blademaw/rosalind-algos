package utils

import "strconv"

// Converts a list of strings to list of integers.
func StringArrToInts(arr []string) []int {
	res := make([]int, len(arr))
	for i, e := range arr {
		e_int, _ := strconv.Atoi(e)
		res[i] = e_int
	}

	return res
}
