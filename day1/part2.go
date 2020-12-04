package day1

import (
	"errors"
)

// Problem: Find the THREE entries that sum to 2020 and then multiply those numbers together.
// Use an array because its fixed length and faster
// Sort the array first because we need to loop twice
// Binary search for the remaining results

// GetTripleExpenses for the elves
func GetTripleExpenses() (int, error) {
	// load input data
	ints := parseAndSort("./day1/input.txt")

	var result int = -1

	for idxA, numA := range ints {
		for idxB, numB := range ints {
			needle := 2020 - numA - numB
			// needle shouldn't be less than the smallest int
			if needle >= ints[0] {
				foundAt := binarySearch(needle, ints)
				if foundAt >= 0 {
					result = ints[idxA] * ints[idxB] * ints[foundAt]
				}
			}
		}
	}
	if result < 0 {
		return result, errors.New("Failed to find number")
	}
	return result, nil
}
