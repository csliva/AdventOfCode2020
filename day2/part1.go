package day2

import (
	"fmt"
	"strings"
)

// CheckPasswords solves for parsing a string into a password policy and a password
// It separates good passwords and bad
// Format: 4-8 g: ggtxgtgbg
// min count of 4 `g`
// max count of 8 `g`
func CheckPasswords() {
	strs := parseInput("./day2/input.txt")
	var goodPasswords []string
	var badPasswords []string
	for _, str := range strs {
		min, max, char, password, err := parsePassword(str)
		if err != nil {
			fmt.Println(err)
		} else {
			charCount := strings.Count(password, char)
			if charCount < min || charCount > max {
				badPasswords = append(badPasswords, password)
			} else {
				goodPasswords = append(goodPasswords, password)
			}
		}
	}
	fmt.Printf("%v\n", len(goodPasswords))
}
