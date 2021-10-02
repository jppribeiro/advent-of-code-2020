package day3

import (
	"strings"
	"testing"
)

func Test_getPath(t *testing.T) {
	testInput := "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"
	testArray := [][]string{}

	for _, line := range strings.Split(testInput, "\n") {
		testArray = append(testArray, strings.Split(line, ""))
	}

	type args struct {
		travelMap [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Expect correct number of trees", args{testArray}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPath(tt.args.travelMap); got != tt.want {
				t.Errorf("getPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
