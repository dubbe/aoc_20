package main

import (
	"fmt"
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

	return 0
}

func b(lines []string) int {

	return 0
}
