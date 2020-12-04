package day2

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(filename string) []string {
	// load input data
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}
	// split strings by endline
	strs := strings.Split(string(data), "\n") // \r\n for windows
	return strs
}

func parsePassword(str string) (min int, max int, char string, password string, err error) {

	//split in into heead and password
	head, password, err := split(str, ": ")
	if err != nil {
		fmt.Println(err)
		return 0, 0, "", "", errors.New("Could not parse password")
	}
	nums, char, err := split(head, " ")
	if err != nil {
		fmt.Println(err)
		return 0, 0, "", "", errors.New("Could not parse character")
	}
	minStr, maxStr, err := split(nums, "-")
	if err != nil {
		fmt.Println(err)
		return 0, 0, "", "", errors.New("Could not parse min max")
	}
	min, _ = strconv.Atoi(minStr)
	max, _ = strconv.Atoi(maxStr)
	return min, max, char, password, nil
}

//New string split function to split string in 2
func split(str string, sep string) (string, string, error) {
	s := strings.Split(string(str), sep)
	if len(s) < 2 {
		return "", "", errors.New("Minimum match not found")
	}
	return s[0], s[1], nil
}
