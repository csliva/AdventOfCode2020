package day4

import (
	"fmt"
	"strings"
)

// split by lines that are only \n\n
// verify required values
// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID) (optional)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func VerifyPassports() {
	raw := parseInput("./day4/input.txt")
	var goodPassports []map[string]string
	var badPassports []map[string]string
	for _, strs := range raw {
		data := strings.Fields(strs)
		pp := parsePassport(data)
		if hasRequiredValues(pp) {
			goodPassports = append(goodPassports, pp)
		} else {
			badPassports = append(badPassports, pp)
		}
	}
	fmt.Printf("%v\n", len(goodPassports))
}

func parsePassport(data []string) map[string]string {
	pp := initPassport()
	for _, d := range data {
		kv := strings.Split(d, ":")
		key := kv[0]
		pp[key] = kv[1]
	}
	return pp
}

func initPassport() map[string]string {
	m := map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		"cid": "",
	}
	return m
}

func hasRequiredValues(pp map[string]string) bool {
	requirements := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	valid := true
	for _, req := range requirements {
		if pp[req] == "" {
			valid = false
		}
	}
	return valid
}
