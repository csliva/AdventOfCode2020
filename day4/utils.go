package day4

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseInput(filename string) []string {
	// load input data
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}
	// split strings by endline and newline
	strs := strings.Split(string(data), "\n\n") // \r\n for windows
	return strs
}
