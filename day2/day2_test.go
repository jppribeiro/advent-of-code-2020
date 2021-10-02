package day2

import (
	"reflect"
	"testing"
)

func Test_parseLine(t *testing.T) {
	line := "1-3 a: abcde"

	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want dbEntry
	}{
		{"Expect valid entry", args{line}, dbEntry{1, 3, "a", "abcde"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	validEntry1 := dbEntry{
		min:      1,
		max:      3,
		char:     "a",
		password: "abcde",
	}

	validEntry2 := dbEntry{
		min:      2,
		max:      9,
		char:     "c",
		password: "ccccccccc",
	}

	invalidEntry1 := dbEntry{
		min:      1,
		max:      3,
		char:     "b",
		password: "cdefg",
	}

	invalidEntry2 := dbEntry{
		min:      4,
		max:      5,
		char:     "s",
		password: "snssj",
	}
	type args struct {
		e dbEntry
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Expect true", args{validEntry1}, true},
		{"Expect true", args{validEntry2}, true},
		{"Expect false", args{invalidEntry1}, false},
		{"Expect false", args{invalidEntry2}, false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.e); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}