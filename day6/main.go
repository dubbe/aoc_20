package main

import (
	"bufio"
	"fmt"
	"os"
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
	a(lines)
	b(lines)
}

func a(lines []string) int {
	set := make(map[rune]bool)
	count := 0
	for _, line := range lines {
		for _, char := range line {
			set[char] = true
		}
		if(line == "") {
			count = len(set) + count 
			set = make(map[rune]bool)
		}
	}

	count = len(set) + count 

	fmt.Printf("Result for a is: %v \n", count)
	return count
}

func b(lines []string) int {
	set := make(map[rune]int)
	count := 0
	i := 0
	for _, line := range lines {
		for _, char := range line {
			set[char] = set[char] + 1
		}
		
		if(line == "") {
			for _, v := range set {
				if(v == i) {
					count++
				} 
			}
			i = 0
			set = make(map[rune]int)
		} else {
			i++
		}
	}

	for _, v := range set {
		if(v == i) {
			count++
		} 
	}
	fmt.Printf("Result for b is: %v \n", count)
	return count
}

