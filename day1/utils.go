package day1

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

//returns -1 for false, and an index position for true
func binarySearch(needle int, haystack []int) int {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(haystack) || haystack[low] != needle {
		return -1
	}
	return low
}

func parseAndSort(filename string) []int {
	// load input data
	data, err := ioutil.ReadFile(filename)
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
	return ints
}
