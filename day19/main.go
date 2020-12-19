package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dubbe/advent-of-code-2020/helpers"
)

func main() {
	start := time.Now()
	lines, err := helpers.ReadLines("input")
	helpers.Check(err)
	fmt.Printf("result A: %v\n", a(lines))
	fmt.Printf("result B: %v\n", b(lines, 0))
	elapsed := time.Since(start)
	fmt.Printf("Solution took %s", elapsed)
}

func a(lines []string) int {
	rules := map[int]string{}
	sum := 0
	messageStartIndex := 0
	for i, line := range lines {
		if line == "" {
			messageStartIndex = i+1
			break
		}
		test := findFirstIndex(line, ':')

		number, err := strconv.Atoi(line[0:test]) 
		if(err != nil) {
			panic("error")
		}
		rules[number] = line[test+2:]	
	}

	rule := regexp.MustCompile(`\A` + parseRules(rules, 0) + `\z`)
	for i := messageStartIndex; i < len(lines); i++ {
		if rule.MatchString(lines[i]) {
			sum++
		}
	}
	return sum
}


func parseRules(rules map[int]string, index int) string {

	rule := rules[index]
	find := regexp.MustCompile("(\\d+)")
	recursion := 0
	for find.MatchString(rule) {
		if(recursion > 50) {
			break;
		}
		rule = find.ReplaceAllStringFunc(rule, func(s string) string {
			number, err := strconv.Atoi(s)
			if err != nil {
				panic(fmt.Sprintf("could not parse %v \n %v", s, err))
			}

			str := rules[number]
			str = strings.Trim(str, "\"")
			if len(str) == 1 && (str[0] == 'b' || str[0] == 'a') {			
				return str
			}

			if strings.Contains(rules[number], fmt.Sprintf(" %d ", number))  {
				fmt.Printf("recursion %s, %s \n", rules[number], s)
				recursion++

			}

			return fmt.Sprintf("( %s )", rules[number])
		})
	}


	rule = strings.ReplaceAll(rule, " ", "")

	return rule
	
}

func findFirstIndex(str string, find rune) int {
	for i, s := range str {
		if s == find {
			return i
		}
	}
	return 0
}

func b(lines []string, startRule int) int {

	rules := map[int]string{}
	sum := 0
	messageStartIndex := 0
	for i, line := range lines {
		if line == "" {
			messageStartIndex = i+1
			break
		}
		test := findFirstIndex(line, ':')

		number, err := strconv.Atoi(line[0:test]) 
		if(err != nil) {
			panic("error")
		}
		rules[number] = line[test+2:]	
	}
	
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"

	rule := regexp.MustCompile(`\A` + parseRules(rules, startRule) + `\z`)
	fmt.Println(parseRules(rules, startRule))
	for i := messageStartIndex; i < len(lines); i++ {
		if rule.MatchString(lines[i]) {
			sum++
		}
	}
	return sum
}
