package main

import (
	"strings"
	"unicode"

	"github.com/dubbe/advent-of-code-2020/helpers"
)


func main() {
	lines, err := helpers.ReadGroups("input")
	helpers.Check(err)
	a(lines)
	b(lines)
}

func a(groups []string) int {
	count := 0
	for _, lines := range groups {
		count += countUnique(strings.Split(lines, "\n"))
	}

	return count
}

func b(groups []string) int {
	count := 0

	for _, lines := range groups {
		count += countSum(lines)
	}

	return count
}

func countUnique(lines []string) int {
	uniqueChars := map[rune]bool{}
	for _, line := range lines {
		for _, char := range line {
			uniqueChars[char] = true
		}
	}
	return len(uniqueChars)
}

func countSum(lines string) int {
	passenger := strings.Count(lines, "\n") + 1

	uniqueChars := map[rune]int{}
	count := 0
	for _, char := range lines {
		if(unicode.IsLetter(char)) {
			uniqueChars[char]++
		}
	}
	for _, test := range uniqueChars {
		if(test == passenger) {
			count++
		}
	}

	return count
}

