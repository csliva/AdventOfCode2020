package day3

import "fmt"

// CheckTrees that wil get bumped into
func CheckTrees() {
	// Unicode
	// 46 = path
	// 35 = tree

	// parse input
	strs := parseInput("./day3/input.txt")
	slopes := []struct {
		right int
		down  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	// multiplier
	totalTrees := 1
	//loop through slopes multiplying
	for _, slope := range slopes {
		totalTrees = totalTrees * checkSlope(slope.right, slope.down, strs)
		fmt.Printf("Tree multiplier %v\n", totalTrees)
	}
}

func checkSlope(right int, down int, strs []string) int {
	treeCount := 0
	// loop through rows
	for idx, str := range strs {
		// start counting down rows
		if idx > 0 && idx%down == 0 {
			// calculate offset for slope
			offset := idx * right
			// use mod to get true index
			trueIdx := offset % len(str)
			// if character at the offset location is a tree, increment
			if str[trueIdx] == 35 {
				treeCount = treeCount + 1
			}
		}
	}
	fmt.Printf("Right %v and down %v runs into %v trees\n", right, down, treeCount)
	return treeCount
}
