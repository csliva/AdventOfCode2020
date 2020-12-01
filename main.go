package main

import (
	"AoC2020/day1"
	"fmt"
)

func main() {
	twoExpenses, err := day1.GetExpenses()
	if err != nil {
		fmt.Println(err)
		return
	}
	threeExpenses, err := day1.GetTripleExpenses()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", twoExpenses)
	fmt.Printf("%v\n", threeExpenses)
}
