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
	sort.Sort(sort.Reverse(sort.IntSlice(lines)))

	configurations := map[int]int{}
	for i, adapter := range lines {

		from := i+1
		to := i+4
		if(to > len(lines)) {
			to = len(lines)
		}
		possibleNextAdapters := lines[from:to]
		found := 0 
		for _, v := range possibleNextAdapters {
			if(adapter == 0) {
				found = 1
			} else if v != 0 && adapter - v <= 3 {
				found++
			}
		}
		configurations[adapter] = found
	}

	sort.Ints(lines)
	sum := 0
	for i, adapter := range lines {
		lowest := i-3
		if(lowest < 0) {
			lowest = 0
		}
		possibleNextAdapters := lines[lowest:i]
		sum = 0
		for _, v := range possibleNextAdapters {
			if adapter - v <= 3 {
				sum += configurations[v]
			}
		}
		if sum == 0 {
			sum = 1
		}
		configurations[adapter] = sum
	}

	fmt.Printf("Result B is: %d", sum)
	return sum
}

