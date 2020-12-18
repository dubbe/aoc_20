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

	return 0
}

func calculate(line string) int {
	find := regexp.MustCompile("\\(([0-9 +*])+\\)")
		for find.MatchString(line) {
			line = find.ReplaceAllStringFunc(line, func(s string) string {
				sum := calculateLine(s)
				return strconv.Itoa(sum)
			})
		}
		
	return calculateLine(line)
}

func calculateLine(line string) int {
	splitted := strings.Split(line, " ")
	sum := calculateOperations(splitted)
	return sum
}

func calculateOperations(operations []string) int  {
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
			
			fmt.Printf("ERROR for %s \n", op)				
			break
	
		}
		sum = calc(nextOperand, sum, number)
		
		operations = operations[1:]

		if len(operations) == 0 {
			break
		}

	}

	return sum
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

