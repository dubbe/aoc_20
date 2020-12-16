package main

import (
	"errors"
	"fmt"
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
	fmt.Printf("result B: %v\n", b(lines))
	elapsed := time.Since(start)
	fmt.Printf("Solution took %s", elapsed)
}

func a(lines []string) int {

	rules, _, _, nearbyTickets := parseTickets(lines)
	sum := 0
	for _, ticket := range nearbyTickets {
		for _, t := range ticket {
			found := false
			for _, rule := range rules {
				if t >= rule[0] && t <= rule[1] {
					found = true
				}
			}
			if !found {
				sum += t
			}
		}
	}

	return sum
}

func b(lines []string) int {
	rules, ruleNames, myTicket, nearbyTickets := parseTickets(lines)
	validTickets := [][]int{}
	for _, ticket := range nearbyTickets {
		valid := true
		for _, t := range ticket {
			found := false
			for _, rule := range rules {
				if t >= rule[0] && t <= rule[1] {
					found = true
				}
			}
			if !found {
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	r := 0
	validColumns := map[int][]int{}
	for i := range ruleNames {

		validColumn := []int{}

		rule1 := rules[r]
		rule2 := rules[r+1]
		r += 2

		for pos := range validTickets[0] {
			valid := true
			for _, ticket := range validTickets {
				t := ticket[pos]
				if (t < rule1[0] || t > rule1[1]) && (t < rule2[0] || t > rule2[1]) {
					valid = false
				}
			}
			if valid {
				validColumn = append(validColumn, pos)
			}
		}
		validColumns[i] = validColumn
	}

	matchedColumns := map[int]int{}
	findMatchedColumns(validColumns, matchedColumns)

	sum := 0
	for key, value := range matchedColumns {

		rule := ruleNames[key]
		if strings.HasPrefix(rule, "departure") {
			if sum == 0 {
				sum = myTicket[value]
			} else {
				sum *= myTicket[value]
			}
		}
	}

	return sum
}

func parseTickets(lines []string) ([][]int, []string, []int, [][]int) {
	rules := [][]int{}
	ruleNames := []string{}
	myTicket := []int{}
	nearbyTickets := [][]int{}
	parser := 0
	for _, line := range lines {
		if line == "" {
			continue
		} else if line == "your ticket:" {
			parser = 1
			continue
		} else if line == "nearby tickets:" {
			parser = 2
			continue
		}
		switch parser {
		case 0:
			splitLine := strings.Split(line, ": ")
			ruleNames = append(ruleNames, splitLine[0])
			values := strings.Split(splitLine[1], " or ")
			for _, value := range values {
				v := strings.Split(value, "-")
				min, _ := strconv.Atoi(v[0])
				max, _ := strconv.Atoi(v[1])
				r := []int{min, max}
				rules = append(rules, r)
			}
		case 1:
			for _, s := range strings.Split(line, ",") {
				i, _ := strconv.Atoi(s)
				myTicket = append(myTicket, i)
			}

		case 2:
			ticket := []int{}
			for _, s := range strings.Split(line, ",") {
				i, _ := strconv.Atoi(s)
				ticket = append(ticket, i)
			}
			nearbyTickets = append(nearbyTickets, ticket)
		}
	}
	return rules, ruleNames, myTicket, nearbyTickets
}

func removeElementFromSlice(slice []int, elem int) ([]int, error) {
	i, err := findIndexInSlice(slice, elem)
	if err == nil {
		copy(slice[i:], slice[i+1:])
		slice[len(slice)-1] = 0
		slice = slice[:len(slice)-1]
		return slice, nil
	}
	return nil, errors.New("error")
}

func findIndexInSlice(slice []int, search int) (int, error) {
	for i, s := range slice {
		if s == search {
			return i, nil
		}
	}
	return 0, errors.New("error")
}

func findMatchedColumns(columns map[int][]int, matchedColumns map[int]int) {
	found := 0
	for i, col := range columns {
		if len(col) == 1 {
			found = col[0]
			matchedColumns[i] = found
			break
		}
	}
	if found != 0 {
		findMatchedColumns(removeFromColumns(columns, found), matchedColumns)
	}
}

func removeFromColumns(columns map[int][]int, value int) map[int][]int {
	ret := map[int][]int{}
	for i, col := range columns {

		ret[i], _ = removeElementFromSlice(col, value)

	}
	return ret
}
