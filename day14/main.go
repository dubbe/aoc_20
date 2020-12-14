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

func a(lines []string) int64 {
	memory := map[int]int64{}
	mask := map[int]rune{}
	for _, line := range lines {
		l := strings.Split(line, " = ")
		op := l[0]
		value := l[1]
		if op == "mask" {
			mask = map[int]rune{}
			for i, r := range value {
				if r != 'X' {
					mask[i] = r
				}
			}
		} else {
			op = strings.Replace(op, "mem[", "", -1)
			op = strings.Replace(op, "]", "", -1)
			o, _ := strconv.Atoi(op)
			n, _ := strconv.Atoi(value)
			binary := fmt.Sprintf("%036s", strconv.FormatInt(int64(n), 2))
			for k, v := range mask {
				binary = replaceAtIndex(binary, v, k)

			}
			i, err := strconv.ParseInt(binary, 2, 64)
			if err == nil {
				memory[o] = i
			}
		}
	}
	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	return sum
}

func b(lines []string) int64 {
	memory := map[int64]int64{}
	mask := map[int]rune{}
	for _, line := range lines {
		l := strings.Split(line, " = ")
		op := l[0]
		value := l[1]
		if op == "mask" {
			mask = map[int]rune{}
			for i, r := range value {
				if r != '0' {
					mask[i] = r
				}
			}
		} else {
			op = strings.Replace(op, "mem[", "", -1)
			op = strings.Replace(op, "]", "", -1)
			o, _ := strconv.Atoi(op)
			n, _ := strconv.Atoi(value)
			binary := fmt.Sprintf("%036s", strconv.FormatInt(int64(o), 2))
			binaries := []string{binary}

			for k, v := range mask {
				newBinearies := []string{}
				if v == 'X' {
					for _, b := range binaries {
						newBinearies = append(newBinearies, replaceAtIndex(b, '1', k))
						newBinearies = append(newBinearies, replaceAtIndex(b, '0', k))
					}
				} else {
					for _, b := range binaries {
						newBinearies = append(newBinearies, replaceAtIndex(b, v, k))
					}
				}
				binaries = newBinearies
			}

			for _, b := range binaries {
				i, err := strconv.ParseInt(b, 2, 64)
				if err == nil {
					memory[i] = int64(n)
				}
			}
		}
	}
	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	return sum

}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
