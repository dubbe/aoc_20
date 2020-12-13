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
	fmt.Printf("Solution took %s\n", elapsed)
}

func a(lines []string) int {
	earliestTimestamp, _ := strconv.Atoi(lines[0])
	timetable := map[int]bool{}
	for _, b := range strings.Split(lines[1], ",") {
		if b != "x" {
			nr, _ := strconv.Atoi(b)
			timetable[nr] = true
		}
	}

	i := earliestTimestamp
	bus := 0
out:
	for {
		if i >= earliestTimestamp {
			for k := range timetable {
				remainder := i % k
				if remainder == 0 {
					bus = k
					break out
				}
			}
		}
		i++
	}

	return (i - earliestTimestamp) * bus
}

func b(lines []string) int {
	timetable := strings.Split(lines[1], ",")

	timestamp := 0
	steps, _ := strconv.Atoi(timetable[0])
	for i := 1; i < len(timetable); i++ {
		s := timetable[i]
		bus, err := strconv.Atoi(s)
		if err == nil {
			for {
				timestamp += steps

				if (timestamp+i)%bus == 0 {
					break
				}

			}
			steps *= bus
		}
	}
	return timestamp
}
