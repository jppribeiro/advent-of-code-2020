package day1

import "testing"

var testInput = []int{1721,979,366, 299, 675, 1456}

func Test_calc(t *testing.T) {
	type args struct {
		lines *[]int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Expect 514579", args{&testInput}, 514579, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcTwoNum(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcThreeNum(t *testing.T) {
	type args struct {
		lines *[]int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Expect 241861950", args{&testInput}, 241861950, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcThreeNum(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcThreeNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calcThreeNum() got = %v, want %v", got, tt.want)
			}
		})
	}
}