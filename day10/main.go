package main

import (
	"fmt"
	"sort"

	"github.com/dubbe/advent-of-code-2020/helpers"
)


func main() {
	lines, err := helpers.ReadInts("input")
	helpers.Check(err)
	a(lines)
	b(lines)
}

func a(lines []int) int {
	sort.Ints(lines)

	i1 := 0
	i3 := 1

	previousVoltage := 0
	for _, voltage := range lines {
		if voltage - previousVoltage == 1 {
			i1++
		} else if voltage - previousVoltage == 3 {
			i3++
		}
		previousVoltage = voltage
	}
	fmt.Printf("Result A is: %d\n", i1*i3)
	return i1*i3
}

func b(lines []int) int {
	lines = append(lines, 0)
	sort.Ints(lines)

	configurations := map[int]int{}

	for i, adapter := range lines {
		lowest := i-3
		if(lowest < 0) {
			lowest = 0
		}
		for _, v := range lines[lowest:i] {
			if adapter - v <= 3 {
				fmt.Printf("a: %v, sum: %v \n", adapter, configurations[adapter])
				configurations[adapter] += configurations[v]
			}
			
		}
		if(configurations[adapter] == 0) {
			configurations[adapter] = 1
		}

	}
	fmt.Printf("Result B is: %d", configurations[lines[len(lines)-1]])
	return configurations[lines[len(lines)-1]]
}

