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
	numbers := strings.Split(lines[0], ",")
	return calculateLatestNumberOnTurn(parseSliceInt(numbers), 2020)

}

func b(lines []string) int {
	numbers := strings.Split(lines[0], ",")
	return calculateLatestNumberOnTurn(parseSliceInt(numbers), 30000000)
}

func parseSliceInt(sSlice []string) []int {
	iSlice := []int{}
	for _, s := range sSlice {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		iSlice = append(iSlice, n)
	}
	return iSlice
}

func calculateLatestNumberOnTurn(startingNumbers []int, lastTurn int) int {

	saidNumbers := make(map[int][]int, lastTurn/4)

	turn := 1
	lastNumberSpoken := 0

	for _, n := range startingNumbers {

		if saidNumbers[n] == nil {
			saidNumbers[n] = []int{turn}
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
		}

		saidNumbers[n] = append(saidNumbers[n], turn)
		lastNumberSpoken = n

		if turn == lastTurn {
			break
		}
		turn++
	}

	return lastNumberSpoken
}
