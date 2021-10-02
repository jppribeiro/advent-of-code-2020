package day3

import (
	"aoc-2020/internal/utils"
	"fmt"
	"strings"
)

func Solution() {
	lines := utils.ReadInput("internal/day3/input3")

	var travelMap [][]string

	for _, line := range lines {
		travelMap = append(travelMap, strings.Split(line, ""))
	}

	fmt.Println(":: Part1 ::")
	fmt.Printf("The sled will encounter %d trees\n", numTreesFound(travelMap, 3, 1))

	fmt.Println("\n::Part 2 ::")

	slopes := [][]int{ {1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2} }

	multiplied := 1

	for _, slope := range slopes {
		multiplied *= numTreesFound(travelMap, slope[0], slope[1])
	}

	fmt.Printf("Result: %d\n", multiplied)
}

func numTreesFound(travelMap [][]string, xSlope int, ySlope int) int {
	numTrees := 0
	x := 0
	y := 0

	for {
		if travelMap[y][x % len(travelMap[y])] == "#" {
			numTrees++
		}

		x += xSlope
		y += ySlope

		if y >= len(travelMap) {
			break
		}
	}

	return numTrees
}
