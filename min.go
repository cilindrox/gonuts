package main

// Min recursively searches for the smallest number in a list of integers
func Min(col []int, min int) int {

	if col[0] < min {
		min = col[0]
	}

	if len(col) == 1 {
		return min
	}

	return Min(col[1:], min)
}
