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
	fmt.Printf("result B: %v\n", b(lines))
	elapsed := time.Since(start)
	fmt.Printf("Solution took %s", elapsed)
}

func a(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += calculate(line)
	}
	return sum
}

func b(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += calculateB(line)
	}
	return sum
}

func calculate(line string) int {
	find := regexp.MustCompile("\\(([0-9 +*])+\\)")
		for find.MatchString(line) {
			line = find.ReplaceAllStringFunc(line, func(s string) string {
				sum, _ := calculateLine(s)
				return strconv.Itoa(sum)
			})
		}
	sum, _ := calculateLine(line)
	return sum
}

func calculateB(line string) int {
	find := regexp.MustCompile("\\(([0-9 +*])+\\)")
		for find.MatchString(line) {
			line = find.ReplaceAllStringFunc(line, func(s string) string {
				sum := calculateLineB(s)
				return strconv.Itoa(sum)
			})
		}
		
	return calculateLineB(line)
}

func calculateLineB(line string) int {
	find := regexp.MustCompile("\\d*\\s\\+\\s\\d*")
		for find.MatchString(line) {
			line = find.ReplaceAllStringFunc(line, func(s string) string {
				sum, err := calculateLine(s)
				if err != nil {
					return s
				}
				return strconv.Itoa(sum)
			})
		}
	sum, err := calculateLine(line)
	if err != nil {
		panic("argh")
	}
	return sum
}

func calculateLine(line string) (int, error) {
	splitted := strings.Split(line, " ")
	sum, err := calculateOperations(splitted)
	return sum, err
}

func calculateOperations(operations []string) (int, error) {
	sum := 0
	nextOperand := ""
	for {
		op := operations[0]

		if op == "+" || op == "*"  {
			nextOperand = op
			operations = operations[1:]
			continue
		}

		op = strings.Trim(op, "(")
		op = strings.Trim(op, ")")
		number, err := strconv.Atoi(op)
		if err != nil {			
			return 0, err
	
		}
		sum = calc(nextOperand, sum, number)
		
		operations = operations[1:]

		if len(operations) == 0 {
			break
		}

	}

	return sum, nil
}

func calc(operand string, a int, b int) int {
	switch operand {
	case "+":
		return a + b
	case "*":
		return a * b
	default:
		return b
	}
}

