package main

import (
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
	sum := 0
	for _, line := range lines {
		sum += calculateLine(line)
	}
	return sum
}

func b(lines []string) int {

	return 0
}

func calculateLine(line string) int {
	splitted := strings.Split(line, " ")
	sum, _ := calculate(splitted)
	return sum
}

func calculate(operations []string) (int, []string)  {
	sum := 0
	nextOperand := ""


	for {
		op := operations[0]

		if op == "+" || op == "*"  {
			nextOperand = op
			operations = operations[1:]
			continue
		}

		number, err := strconv.Atoi(op)
		if err != nil {
			if strings.HasPrefix(op, "(") {
				operations[0] = op[1:]
				number, operations = calculate(operations)
			} else if strings.HasSuffix(op, ")") {
				number, err = strconv.Atoi(strings.TrimRight(op, ")"))
				if(err != nil) {
					fmt.Println("ERROR")
					operations = operations[1:]
					continue
				}
				sum = calc(nextOperand, sum, number)
				return sum, operations
			}
		}
		sum = calc(nextOperand, sum, number)
		
		if len(operations) == 0 {
			break;
		}

		operations = operations[1:]

		if len(operations) == 0 {
			break;
		}

	}

	return sum, []string{}
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