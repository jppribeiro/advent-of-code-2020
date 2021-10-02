package day4

import (
	"reflect"
	"testing"
)

func Test_extractPassports(t *testing.T) {
	testData := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}

	expectedData := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
		"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
		"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
	}

	type args struct {
		rawData []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Return expected data", args{testData}, expectedData},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractPassports(tt.args.rawData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractPassports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkValid(t *testing.T) {
	testData := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
		"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
		"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
	}

	type args struct {
		passports []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Return expected number of valid passports", args{testData}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValid(tt.args.passports); got != tt.want {
				t.Errorf("checkValid() = %v, want %v", got, tt.want)
			}
		})
	}
}