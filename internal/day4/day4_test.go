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
			if got := bulkCheckValid(tt.args.passports); got != tt.want {
				t.Errorf("bulkCheckValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeValidator(t *testing.T) {
	testCb := func(val int) bool {
		return val >= 1920 && val <= 2001
	}

	okVal := "1921"
	notOkVal := "1919"

	type args struct {
		cb  func(val int) bool
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Expect true for value in range", args{testCb, okVal}, true},
		{"Expect true for value in range", args{testCb, notOkVal}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeValidator(tt.args.cb, tt.args.val); got != tt.want {
				t.Errorf("executeValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkValidB(t *testing.T) {
	invalidInputs := []string{
		"eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946",
		"hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007",
	}

	validInputs := []string{
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f",
		"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	}

	type args struct {
		passports []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Return no valid passports", args{invalidInputs}, 0},
		{"Return 4 valid passports", args{validInputs}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidB(tt.args.passports); got != tt.want {
				t.Errorf("checkValidB() = %v, want %v", got, tt.want)
			}
		})
	}
}