package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	a(lines)
	b(lines)
}

func a(lines []string) int {
	lowest, heighest := findHeighestAndLowestSeat(lines)
	fmt.Printf("Lowest: %v Heighest: %v \n", lowest, heighest)
	return heighest
}

func b(lines []string) int {
	seats := [902]bool{}
	for _, line := range lines {
		ticket := parseTicket(line)
		
		seatID := calculateSeatID(ticket)
		seats[seatID] = true
	}
	
	for i, seat := range seats {
		if(i > 84 && seat == false) {
			fmt.Printf("free seat: %v \n", i)
		}
	}

	return 0
}

func calculateSeatID(seat [2]int) int {
	return (seat[0] * 8) + seat[1]
}

func findHeighestAndLowestSeat(lines []string) (int, int) {
	heighest := 0
	lowest := 10000
	for _, line := range lines {
		ticket := parseTicket(line)
		
		seatID := calculateSeatID(ticket)
		if(seatID > heighest) {
			heighest = seatID
		}
		if(seatID < lowest) {
			lowest = seatID
		}
	}
	return lowest, heighest
}


func parseTicket(line string) [2]int {
	rows := [2]int{0,127}
	columns := [2]int{0,7}
	col := 0
	row := 0
	for i := range line {
		letter := line[i]
		if(letter == 'F' || letter == 'B') {
			r, err := parseLetter(letter, &rows, 'F', 'B')
			if(err == nil) {
				row = r
			}
		} else if (letter == 'L' ||  letter == 'R') {
			c, err := parseLetter(letter, &columns, 'L', 'R')
			if(err == nil) {
				col = c
			}
		}
	}
	return [2]int{row, col}
}

func parseLetter(letter byte, rows *[2]int, first byte, last byte) (int, error) {

	remaining := rows[1] - rows[0]
	half := remaining - (remaining / 2)
	if(letter == last) {
		rows[0] = rows[0] + half
	} 
	if(letter == first) {
		rows[1] = rows[1] - half
	}
	if(rows[0] - rows[1] == 0) {
		return rows[0], nil
	}
	return 0, errors.New("Not a seat")
}

