package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"regexp"
)

func check(e error) {
	if e != nil {
			panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
			return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("input")
	check(err)
	A(lines)
	B(lines)
}

func A(lines []string) (int) {
	passports := getPassports(lines)
	requiredFields := [7]string{
		"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
	}
	validPassports := 0
	for _, passport := range passports {
		validFields := 0
		for _, field := range requiredFields {
			if(strings.Count(passport, field) == 1) {
				validFields++
			}
		}
		if(validFields == len(requiredFields)) {
			validPassports++
		}
	}
	fmt.Println(validPassports)
	return validPassports
}

func B(lines []string) int {
	passports := getPassports(lines)

	validPassports := 0
	for _, passport := range passports {
		if(validatePassport(passport)) {
			validPassports++
		}
	}
	fmt.Println(validPassports)
	return validPassports
	
}



func getPassports(lines []string) []string {
	i := 0
	passports := make([]string, 0)
	passport := ""
	for _, line := range lines  {
		if(line == "") {
			passports = append(passports, passport)
			i++
			passport = ""
		} else {
			passport = passport + " " + line
		}
	}

	passports = append(passports, passport)
	return passports
}

type validator struct {
	name string
	regex string
}

func validatePassport(passport string) bool {
	requiredFields := make(map[string]string)
	requiredFields["byr"] = "^(19[2-8][0-9]|199[0-9]|200[0-2])$"
	requiredFields["iyr"] = "^(201[0-9]|2020)$"
	requiredFields["eyr"] = "^(202[0-9]|2030)$"
	requiredFields["hgt"] = "^(((1[5-8][0-9]|19[0-3])cm)|((59|6[0-9]|7[0-6])in))$"
	requiredFields["hcl"] = "^#((\\d|[a-f]){6})$"
	requiredFields["ecl"] = "^amb|blu|brn|gry|grn|hzl|oth$"
	requiredFields["pid"] = "^\\d{9}$"
		//"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",


	fields := strings.Split(passport, " ")

	validFields := 0
	for _, field := range fields {
		f := strings.Split(field, ":")
		name := f[0]

		if(len(requiredFields[name]) == 0) {
			continue
		}

		re := regexp.MustCompile(requiredFields[name])
		match := re.MatchString(f[1])
		if(match) {
			validFields++
		}
		


	}
	

	return validFields == len(requiredFields)
} 

