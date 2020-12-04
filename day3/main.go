package main

import (
	"bufio"
	"os"
	"fmt"
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
	lines, err := readLines("input")
	check(err)
	A(lines)
	B(lines)
}

func A(lines []string) (int) {
	// count the number of columns
	columns := len(lines[0])
	rows := len(lines)
	
	x := 0
	i := 0

	for y := 1; y < rows; y++ {
		x = x+3

		if(x >= columns) {
			x = x - columns
		}

		if(lines[y][x] == '#') {
			i++
		} 
	}

	fmt.Println(i)
	return i

}

func B(lines []string) int {
	a := checkSlope(lines, 1, 1)
	b := checkSlope(lines, 3, 1)
	c := checkSlope(lines, 5, 1)
	d := checkSlope(lines, 7, 1)
	e := checkSlope(lines, 1, 2)

	fmt.Println(a*b*c*d*e)
	return a*b*c*d*e
	
}

func checkSlope(lines []string, right int, down int) int {
	columns := len(lines[0])
	rows := len(lines)
	
	x := 0
	i := 0
	fmt.Println("START")
	for y := 0; y < rows; y=y+down {
		fmt.Print(y)

		if(lines[y][x] == '#') {
			fmt.Print("#")
			i++
		} else {
			fmt.Print(".")
		}

		x = x+right

		if(x >= columns) {
			x = x - columns
		}
		
	}
	fmt.Println("-----")
	fmt.Println(i)
	return i
}
