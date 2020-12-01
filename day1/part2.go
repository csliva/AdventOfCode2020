package day1

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// Problem: Find the THREE entries that sum to 2020 and then multiply those numbers together.
// Use an array because its fixed length and faster
// Sort the array first because we need to loop twice
// Binary search for the remaining results

// GetExpenses for the elves
func GetTripleExpenses() (int, error) {
	// load input data
	data, err := ioutil.ReadFile("./day1/input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}
	// split strings by endline
	numStrs := strings.Split(string(data), "\n") // \r\n for windows
	// convert the slice into an array of numbers
	ints := make([]int, len(numStrs))
	for i, s := range numStrs {
		ints[i], _ = strconv.Atoi(s)
	}
	// sort ints
	sort.Ints(ints)

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
