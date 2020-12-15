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
	fmt.Printf("result A: %v \n", a(lines))
	fmt.Printf("result B: %v \n", b(lines))
	elapsed := time.Since(start)
	fmt.Printf("Solution took %s", elapsed)
}

func a(lines []string) int {
	startingNumbers := strings.Split(lines[0], ",")

	saidNumbers := map[int][]int{}

	turn := 1
	lastNumberSpoken := 0

	for _, number := range startingNumbers {
		n, _ := strconv.Atoi(number)
		if saidNumbers[n] == nil {
			saidNumbers[n] = []int{ turn }
			lastNumberSpoken = n
		} else {
			saidNumbers[n] = append(saidNumbers[n], turn)
			lastNumberSpoken = n
		}
		turn++
	} 

	for {
		n := 0
		if len(saidNumbers[lastNumberSpoken]) > 1 {
			
			lastTwo := saidNumbers[lastNumberSpoken][len(saidNumbers[lastNumberSpoken])-2:]
			n = lastTwo[1] - lastTwo[0]
			saidNumbers[n] = append(saidNumbers[n], turn)
			lastNumberSpoken = n
		} else {
			saidNumbers[0] = append(saidNumbers[0], turn)
			lastNumberSpoken = 0
		}

		if(turn == 2020) {
			break
		}
		turn++
	}

	return lastNumberSpoken
}

func b(lines []string) int {

	startingNumbers := strings.Split(lines[0], ",")

	saidNumbers := map[int][]int{}

	turn := 1
	lastNumberSpoken := 0

	for _, number := range startingNumbers {
		n, _ := strconv.Atoi(number)
		if saidNumbers[n] == nil {
			saidNumbers[n] = []int{ turn }
			lastNumberSpoken = n
		} else {
			saidNumbers[n] = append(saidNumbers[n], turn)
			lastNumberSpoken = n
		}
		turn++
	} 

	for {
		n := 0
		if len(saidNumbers[lastNumberSpoken]) > 1 {
			
			lastTwo := saidNumbers[lastNumberSpoken][len(saidNumbers[lastNumberSpoken])-2:]
			n = lastTwo[1] - lastTwo[0]

			saidNumbers[n] = append(saidNumbers[n], turn)
			lastNumberSpoken = n
		} else {
			saidNumbers[0] = append(saidNumbers[0], turn)
			lastNumberSpoken = 0
		}

		if(turn == 30000000) {
			break
		}
		turn++
	}

	return lastNumberSpoken
}
