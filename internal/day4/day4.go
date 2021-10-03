package day4

import (
	"aoc-2020/internal/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
 */

type height struct {
	cm func(val int) bool
	in func(val int) bool
}

type rules struct {
	byr func(val int) bool
	iyr func(val int) bool
	eyr func(val int) bool
	hgt height
	hcl string
	ecl string
	pid string
}

var validator = rules{
	func(val int) bool {
		return val >= 1920 && val <= 2002
	},
	func(val int) bool {
		return val >= 2010 && val <= 2020
	},
	func(val int) bool {
		return val >= 2020 && val <= 2030
	},
	height{
		func(val int) bool {
			return val >= 150 && val <= 193
		},
		func(val int) bool {
			return val >= 59 && val <= 76
		},
	},
	"^#[a-f0-9]{6,6}$",
	"^(amb|blu|brn|gry|grn|hzl|oth)$",
	"^[0-9]{9}$",
}

func Solution() {
	lines := utils.ReadInput("internal/day4/input4")

	passports := extractPassports(lines)

	fmt.Println(":: Part1 ::")
	fmt.Printf("There are %d valid passports.\n\n", bulkCheckValid(passports))

	fmt.Println(":: Part2 ::")
	fmt.Printf("There are %d valid passports.\n\n", checkValidB(passports))
}

func extractPassports(rawData []string) []string {
	var passports []string
	running := true
	start := 0

	for running {
		passport := ""

		for i := start; i < len(rawData); i++ {
			if rawData[i] == "" {
				start = i + 1
				break
			}

			if i == len(rawData) - 1 {
				running = false
			}

			passport += fmt.Sprintf(" %s", rawData[i])
		}

		passports = append(passports, strings.TrimSpace(passport))
	}

	return passports
}

func bulkCheckValid(passports []string) int {
	num := 0

	for _, p := range passports {
		if checkValid(p) {
			num++
		}
	}

	return num
}

func checkValid(p string) bool {
	if len(strings.Split(p, " ")) < 7 {
		return false
	}

	if len(strings.Split(p, " ")) == 8 {
		return true
	}

	if match, _ := regexp.MatchString("cid", p); !match {
		// If there is 1 field missing and it is NOT cdi then consider valid
		return true
	}

	return false
}

func checkValidB(passports []string) int {
	numValid := 0

	for _, p := range passports {
		isValid := true

		if !checkValid(p) {
			continue
		}

		for _, field := range strings.Split(p, " ") {
			key := strings.Split(field, ":")[0]

			switch key {
			case "byr":
				isValid = isValid && executeValidator(validator.byr, strings.Split(field, ":")[1])
			case "iyr":
				isValid = isValid && executeValidator(validator.iyr, strings.Split(field, ":")[1])
			case "eyr":
				isValid = isValid && executeValidator(validator.eyr, strings.Split(field, ":")[1])
			case "hgt":
				unit := strings.Split(field, ":")[1][len(strings.Split(field, ":")[1])-2:]

				value := strings.Split(field, ":")[1][0:len(strings.Split(field, ":")[1])-2]

				if unit == "cm" {
					isValid = isValid && executeValidator(validator.hgt.cm, value)
					break
				} else if unit == "in" {
					isValid = isValid && executeValidator(validator.hgt.in, value)
					break
				}

				isValid = false
			case "hcl":
				match, _ := regexp.MatchString(validator.hcl, strings.Split(field, ":")[1])
				isValid = isValid && match
			case "ecl":
				match, _ := regexp.MatchString(validator.ecl, strings.Split(field, ":")[1])
				isValid = isValid && match
			case "pid":
				match, _ := regexp.MatchString(validator.pid, strings.Split(field, ":")[1])
				isValid = isValid && match
			}

			// No need to check more fields when one is not valid
			if !isValid {
				break
			}
		}

		if isValid { numValid++ }
	}

	return numValid
}

func executeValidator(cb func(val int) bool, val string) bool {
	value, err := strconv.Atoi(val)
	if err != nil { return false }

	return cb(value)
}