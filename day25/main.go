package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dubbe/advent-of-code-2020/helpers"
)


func main() {
	start := time.Now()
	lines, err := helpers.ReadLines("input")
	helpers.Check(err)
	fmt.Printf("result A: %v\n", a(lines))
	elapsed := time.Since(start)
	fmt.Printf("result A took %s\n", elapsed)

	startB := time.Now()
	fmt.Printf("result B: %v\n", b(lines))
	elapsed = time.Since(startB)
	fmt.Printf("result B took %s\n", elapsed)
	
	elapsed = time.Since(start)
	fmt.Printf("Solution took %s±n", elapsed)
}

func a(lines []string) int {
	card, _ := strconv.Atoi(lines[0])
	door, _ := strconv.Atoi(lines[1])


	i := 0
	for k := 1; k != card; i++ {
		k = k * 7 % 20201227
	}

	fmt.Printf("i: %d \n", i)

	key := 1
	for l := 0; l < i; l++ {
		key = key * door % 20201227
	}

	return key
}

func b(lines []string) int {

	return 0
}
