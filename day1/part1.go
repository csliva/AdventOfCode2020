package day1

import (
	"errors"
)

// Problem: Find the two entries that sum to 2020 and then multiply those two numbers together.
// Use an array because its fixed length and faster
// Sort the array first because we need to loop twice
// Binary search for the remaining results

// GetExpenses for the elves
func GetExpenses() (int, error) {
	// load input data
	ints := parseAndSort("./day1/input.txt")

	var result int = -1

	for idx, num := range ints {
		needle := 2020 - num
		foundAt := binarySearch(needle, ints)
		if foundAt >= 0 {
			result = ints[idx] * ints[foundAt]
		}
	}
	if result < 0 {
		return result, errors.New("Failed to find number")
	}
	return result, nil
}
