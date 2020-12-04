package day4

import (
	"fmt"
	"regexp"
	"strconv"
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

func ValidatePassports() {
	raw := parseInput("./day4/input.txt")
	var goodPassports []map[string]string
	var badPassports []map[string]string
	for _, strs := range raw {
		data := strings.Fields(strs)
		pp := parsePassport2(data)
		if hasRequiredValues(pp) {
			goodPassports = append(goodPassports, pp)
		} else {
			badPassports = append(badPassports, pp)
		}
	}
	fmt.Printf("%v\n", len(goodPassports))
}

func parsePassport2(data []string) map[string]string {
	pp := initPassport()
	for _, d := range data {
		kv := strings.Split(d, ":")
		key := kv[0]
		value := kv[1]
		if isValidKey(key, value) {
			pp[key] = value
		}
	}
	return pp
}

func initPassport2() map[string]string {
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

func isValidKey(key string, value string) bool {
	valid := false
	switch key {
	case "byr":
		i, err := strconv.Atoi(value)
		if err != nil || i < 1920 || i > 2002 {
			valid = false
		} else {
			valid = true
		}
	case "iyr":
		i, err := strconv.Atoi(value)
		if err != nil || i < 2010 || i > 2020 {
			valid = false
		} else {
			valid = true
		}
	case "eyr":
		i, err := strconv.Atoi(value)
		if err != nil || i < 2020 || i > 2030 {
			valid = false
		} else {
			valid = true
		}
	case "hgt":
		tail := value[len(value)-2:]
		head := value[0 : len(value)-2]
		if tail == "cm" {
			i, err := strconv.Atoi(head)
			if err != nil || i < 150 || i > 193 {
				valid = false
			} else {
				valid = true
			}
		} else if tail == "in" {
			i, err := strconv.Atoi(head)
			if err != nil || i < 59 || i > 76 {
				valid = false
			} else {
				valid = true
			}
		} else {
			valid = false
		}
	case "hcl":
		matched, err := regexp.Match(`^#[0-9a-f]{6}`, []byte(value))
		if err != nil || matched != true {
			valid = false
		} else {
			valid = true
		}
	case "ecl":
		switch value {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			valid = true
		default:
			valid = false
		}
	case "pid":
		matched, err := regexp.Match(`[0-9]{9}`, []byte(value))
		if err != nil || matched != true {
			valid = false
		} else {
			valid = true
		}
	case "cid":
		valid = true
	default:
		valid = false
	}
	return valid
}
