package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dubbe/advent-of-code-2020/helpers"
)


func main() {
	lines, err := helpers.ReadLines("input")
	helpers.Check(err)

	
	a(lines)
	b(lines)
}

func a(lines []string) int {
	lineIndex := 0
	accumulator := 0
	ranLines := map[int]bool{}
	
	for {
		ranLines[lineIndex] = true
		executeLine(lines[lineIndex], &accumulator, &lineIndex)
		if(ranLines[lineIndex]) {

			break
		}
	}
	fmt.Printf("accumulator is: %d \n", accumulator)
	return accumulator
}

func executeLine(line string, accumulator *int, lineIndex *int) {
	operation := strings.Split(line, " ")
	value, err := strconv.Atoi(operation[1])

	if(err != nil) {
		fmt.Printf("ERROR, %v", err)
	}

	switch os := operation[0]; os {
	case "acc":
		*lineIndex++
		*accumulator += value
		return
	case "jmp":
		*lineIndex += value
		return
	case "nop":
		*lineIndex++
		return
	default:
		return
	}
}

func b(lines []string) int {

	found := false
	accumulator := 0
	for i := range lines {
		found, accumulator = executeSearch(lines, i)
		if(found) {
			break
		}
	}
	
	fmt.Printf("accumulator is: %d \n", accumulator)
	return accumulator
}

func executeSearch(lines []string, indexToChange int) (bool, int) {

	lineIndex := 0
	accumulator := 0
	ranLines := map[int]bool{}

	for {
		ranLines[lineIndex] = true

		if(lineIndex == indexToChange) {
			executeLine(changeLine(lines[lineIndex]), &accumulator, &lineIndex)
		} else {
			executeLine(lines[lineIndex], &accumulator, &lineIndex)
		}
		
		if(ranLines[lineIndex]) {
			return false, 0
		}
		if(lineIndex == len(lines)) {
			return true, accumulator
		}
	}
}

func changeLine(line string) string {
	lineToChange := strings.Split(line, " ")
	if lineToChange[0]== "jmp" {
		return fmt.Sprintf("%s %s", "nop", lineToChange[1])
	} else if lineToChange[0] == "nop" {
		return fmt.Sprintf("%s %s", "jmp", lineToChange[1])
	} else {
		return line
	}
}
