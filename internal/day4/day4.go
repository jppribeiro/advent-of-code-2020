package day4

import (
	"aoc-2020/internal/utils"
	"fmt"
	"regexp"
	"strings"
)

var mandatoryFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	// "cid",
}

func Solution() {
	lines := utils.ReadInput("internal/day4/input4")

	passports := extractPassports(lines)

	fmt.Println(":: Part1 ::")
	fmt.Printf("There are %d valid passports.\n\n", checkValid(passports))
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

func checkValid(passports []string) int {
	num := 0

	for _, p := range passports {
		if len(strings.Split(p, " ")) < 7 {
			continue
		}

		if len(strings.Split(p, " ")) == 8 {
			num++
			continue
		}

		if match, _ := regexp.MatchString("cid", p); !match {
			// If there is 1 field missing and it is NOT cdi then consider valid
			num++
		}
	}

	return num
}