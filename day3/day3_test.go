package day3

import (
	"strings"
	"testing"
)

func Test_numTreesFound(t *testing.T) {
	testInput := "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"
	var testArray [][]string

	for _, line := range strings.Split(testInput, "\n") {
		testArray = append(testArray, strings.Split(line, ""))
	}

	type args struct {
		travelMap [][]string
		xSlope    int
		ySlope    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Expect correct number of trees", args{testArray, 1, 1}, 2},
		{"Expect correct number of trees", args{testArray, 3, 1}, 7},
		{"Expect correct number of trees", args{testArray, 5, 1}, 3},
		{"Expect correct number of trees", args{testArray, 7, 1}, 4},
		{"Expect correct number of trees", args{testArray, 1, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTreesFound(tt.args.travelMap, tt.args.xSlope, tt.args.ySlope); got != tt.want {
				t.Errorf("numTreesFound() = %v, want %v", got, tt.want)
			}
		})
	}
}