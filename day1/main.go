package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
			panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
			return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	numbers, err := readLines("input")
	check(err)
	day01A(numbers)
	day01B(numbers)
}

func day01A(numbers []string) {
	numbers2 := numbers
	for _, a := range numbers {
		for _, b := range numbers2 {
			x, _ := strconv.ParseInt(a, 10, 64)
			y, _ := strconv.ParseInt(b, 10, 64)
			if(x + y == 2020) {
				fmt.Println(x*y)
				return
			}
		}
	}
}

func day01B(numbers []string) {
	numbers2 := numbers
	numbers3 := numbers
	for _, a := range numbers {
		for _, b := range numbers2 {
			for _, c := range numbers3 {
			x, _ := strconv.ParseInt(a, 10, 64)
			y, _ := strconv.ParseInt(b, 10, 64)
			z, _ := strconv.ParseInt(c, 10, 64)
			if(x + y + z == 2020) {
				fmt.Println(x*y*z)
				return
			}
		}
		}
	}
}