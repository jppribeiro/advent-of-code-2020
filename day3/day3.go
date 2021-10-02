package day3

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

func Solution() {
	lines := utils.ReadInput("day3/input3")

	var travelMap [][]string

	for _, line := range lines {
		travelMap = append(travelMap, strings.Split(line, ""))
	}

	fmt.Printf("The sled will encounter %d trees", getPath(travelMap))
}

func getPath(travelMap [][]string) int {
	numTrees := 0
	x := 0
	y := 0

	for {
		if travelMap[y][x % len(travelMap[y])] == "#" {
			numTrees++
		}

		x += 3
		y++

		if y == len(travelMap) {
			break
		}
	}

	return numTrees
}
