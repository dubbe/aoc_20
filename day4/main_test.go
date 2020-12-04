package main

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
  if a == b {
    return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestA(t *testing.T) {
	lines, _ := readLines("input_test")

	result := A(lines)
	assertEqual(t, result, 2, "did not find 2 valid passports")

}

func TestB(t *testing.T) {

	result := validatePassport("pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f")
	assertEqual(t, result, true, "did not validate passport 1")

	result = validatePassport("eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm")
	assertEqual(t, result, true, "did validate passport 2")

	result = validatePassport("hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022")
	assertEqual(t, result, true, "did validate passport 3")

	result = validatePassport("iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719")
	assertEqual(t, result, true, "did validate passport 4")

	fmt.Println("------")

	result = validatePassport("eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926")
	assertEqual(t, result, false, "did validate passport 1")

	result = validatePassport("iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946")
	assertEqual(t, result, false, "did validate passport 2")

	result = validatePassport("hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277")
	assertEqual(t, result, false, "did validate passport 3")

	result = validatePassport("hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007")
	assertEqual(t, result, false, "did validate passport 4")
}