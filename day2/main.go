package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
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
	day02A(lines)
	day02B(lines)
}

func day02A(lines []string) {
	i := 0
	for _, line := range lines {
		pass := strings.Split(line, " ")
		lengths := strings.Split(pass[0], "-")
		char := pass[1][0:1]
		count := strings.Count(pass[2], char)
		min, _ := strconv.Atoi(lengths[0])
		max, _ := strconv.Atoi(lengths[1])
		if(count >= min  && count <= max) {
			i++
		}
	}
	fmt.Println(i)
}

func day02B(lines []string) {
	i := 0
	for _, line := range lines {
		pass := strings.Split(line, " ")
		lengths := strings.Split(pass[0], "-")
		char := pass[1][0]

		first, _ := strconv.Atoi(lengths[0]) 
		second, _ := strconv.Atoi(lengths[1])
		if(checkValid(pass[2], first, second, char)) {
			i++
		}
	}
	fmt.Println(i)
	
}

func checkValid(password string, first int, second int, c byte) (bool) {
	valid := false
	first = first - 1
	second = second - 1
	if(password[first] == c) {
		if(password[second] != c) {
			valid = true
		}
	}

	if(password[second] == c) {
		if(password[first] != c) {
			valid = true
		}
	}
	return valid
}