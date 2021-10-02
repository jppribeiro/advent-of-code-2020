package day1

import (
	utils "aoc-2020/utils"

	"fmt"
)

func Solution() {
	lines := utils.ReadIntInput("day1/input1")

	res1, err := calcTwoNum(&lines)
	if err != nil {
		fmt.Println("error on solution 1")
	}

	res2, err := calcThreeNum(&lines)
	if err != nil {
		fmt.Println("error on solution 1")
	}

	fmt.Printf("Day 1a solution: %d\n", res1)
	fmt.Printf("Day 1b solution: %d\n", res2)

}

func calcTwoNum(lines *[]int) (int, error) {
	for i := 0; i < len(*lines) - 1; i++ {
		for j := i + 1; j < len(*lines); j++ {
			if (*lines)[i] + (*lines)[j] == 2020 {
				return (*lines)[i] * (*lines)[j], nil
			}
		}
	}

	return 0, fmt.Errorf("error on solution 1")
}

func calcThreeNum(lines *[]int) (int, error) {
	for i := 0; i < len(*lines) - 2; i++ {
		for j := i + 1; j < len(*lines) - 1; j++ {
			for k := j + 1; k < len(*lines); k++ {
				if (*lines)[i] + (*lines)[j] + (*lines)[k] == 2020 {
					return (*lines)[i] * (*lines)[j] * (*lines)[k], nil
				}
			}
		}
	}

	return 0, fmt.Errorf("error on solution 1")
}