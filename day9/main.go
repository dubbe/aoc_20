package main

import (
	"fmt"
	"sort"

	"github.com/dubbe/advent-of-code-2020/helpers"
)


func main() {
	lines, err := helpers.ReadInts("input")
	helpers.Check(err)
	result := a(lines, 25)
	b(lines, result)
}

func a(lines []int, preamble int) int {

	numbers := lines[0:preamble]
	for i := preamble; i < len(lines); i++ {
		valid := validateNumber(numbers, lines[i])
		if(valid == false) {
			fmt.Printf("Result: %d \n", lines[i]);
			return lines[i]
		}

		numbers = append(numbers, lines[i])[1:]
	}

	return 0
}

func validateNumber(originalNumbers []int, number int) bool {
	numbers := make([]int, len(originalNumbers))
	copy(numbers, originalNumbers)
	sort.Ints(numbers)

	for _, n := range numbers {
		searchFor := number - n
		found := sort.SearchInts(numbers, searchFor)
		if found > 0 {
			return true
		}
	} 
	return false
}

func b(lines []int, findNumber int) int {
	for i := 0; i < len(lines); i++ {
		x := 0
		
		for {
			result := 0
			numbers := make([]int, len(lines[i:x+i]))

			copy(numbers, lines[i:x+i])
			
			for _, v := range numbers {
				result+=v
			}

			if result == findNumber {
				
				sort.Ints(numbers)
				answer :=  numbers[0] + numbers[len(numbers) - 1]
				fmt.Printf("found it: %v, result is: %d \n", numbers, answer)
				return answer
			} else if result > findNumber {
				break;
			}
			x++
		}
		
	}
	return 0
}
