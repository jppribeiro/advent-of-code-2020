package day2

import (
	"aoc-2020/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

type dbEntry struct {
	min int
	max int
	char string
	password string
}

func Solution() {
	lines := utils.ReadInput("internal/day2/input2")

	numValidA := 0
	numValidB := 0

	for _, line := range lines {
		entry := parseLine(line)

		if isValidA(entry) {
			numValidA++
		}

		if isValidB(entry) {
			numValidB++
		}
	}

	fmt.Println(":: Part 1 ::")
	fmt.Printf("There are %d invalid passwords according to the sled rental criteria\n", numValidA)

	fmt.Println("\n:: Part 2 ::")
	fmt.Printf("There are %d invalid passwords according to the Toboggan Corporate criteria\n", numValidB)
}

func parseLine(line string) dbEntry {
	contents := strings.Split(line, ": ")

	min, _ := strconv.Atoi(strings.Split(strings.Split(contents[0], " ")[0], "-")[0])
	max, _ := strconv.Atoi(strings.Split(strings.Split(contents[0], " ")[0], "-")[1])
	char := strings.Split(contents[0], " ")[1]

	return dbEntry{min, max, char, contents[1] }
}

func isValidA(e dbEntry) bool {
	instances := 0

	for _, c := range strings.Split(e.password, "") {
		if c == e.char {
			instances++
		}
	}

	if instances >= e.min && instances <= e.max {
		return true
	}

	return false
}

func isValidB(e dbEntry) bool {
	return (e.char == strings.Split(e.password, "")[e.min - 1]) != (e.char == strings.Split(e.password, "")[e.max - 1])
}
