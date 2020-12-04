package day2

import (
	"fmt"
	"strings"
)

// ReCheckPasswords solves for parsing a string into a password policy and a password
// It separates good passwords and bad
// Format: 4-8 g: ggtxgtgbg
// `g` at index of 4 or 8 (3 and 7 in 0 based index)
func ReCheckPasswords() {
	strs := parseInput("./day2/input.txt")
	var goodPasswords []string
	var badPasswords []string
	for _, str := range strs {
		min, max, char, password, err := parsePassword(str)
		passwordArr := strings.Split(password, "")
		if err != nil {
			fmt.Println(err)
		} else {
			idxA := min - 1
			idxB := max - 1
			if idxA >= 0 && idxB < len(password) {
				if passwordArr[idxA] == char && passwordArr[idxB] != char {
					goodPasswords = append(goodPasswords, password)
				} else if passwordArr[idxA] != char && passwordArr[idxB] == char {
					goodPasswords = append(goodPasswords, password)
				} else {
					badPasswords = append(badPasswords, password)
				}
			}
		}
	}
	fmt.Printf("%v\n", len(goodPasswords))
}
