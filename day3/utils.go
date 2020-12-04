package day3

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
	// split strings by endline
	strs := strings.Split(string(data), "\n") // \r\n for windows
	return strs
}
