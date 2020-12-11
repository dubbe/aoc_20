package main

import (
	"fmt"

	"github.com/dubbe/advent-of-code-2020/helpers"
)


func main() {
	lines, err := helpers.ReadLines("input")
	helpers.Check(err)
	a(lines)
	b(lines)
}

func a(lines []string) int {
	seats := getSeats(lines)
	occupiedSeats := seatingRound(seats)
	fmt.Printf("Result A: %d \n", occupiedSeats)
	return occupiedSeats
}

func b(lines []string) int {

	seats := getSeats(lines)
	occupiedSeats := seatingRound2(seats)
	fmt.Printf("Result B: %d \n", occupiedSeats)
	return occupiedSeats
}

func printSeats(seats [][]rune) {
	for _, row := range seats {
		for _, seat := range row {
			fmt.Printf("%v ", string(seat))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func seatingRound(seats [][]rune) int {
	tmpSeats := [][]rune{}

	changes := 0

	for y, row := range seats {
		rowOfSeats := []rune{}
		for x, seat := range row {
			if(checkSeatChange(x, y, seat, seats)) {
				rowOfSeats = append(rowOfSeats, switchSeat(seat))
				changes++
			}  else {
				rowOfSeats = append(rowOfSeats, seat)
			}
		}
		tmpSeats = append(tmpSeats, rowOfSeats);

	}

	if(changes == 0) {
		return countOccupiedSeats(tmpSeats)
		
	} 
		
	return seatingRound(tmpSeats)
	
}

func seatingRound2(seats [][]rune) int {
	tmpSeats := [][]rune{}

	changes := 0

	for y, row := range seats {
		rowOfSeats := []rune{}
		for x, seat := range row {
			if(checkSeatChange2(x, y, seat, seats)) {
				rowOfSeats = append(rowOfSeats, switchSeat(seat))
				changes++
			}  else {
				rowOfSeats = append(rowOfSeats, seat)
			}
		}
		tmpSeats = append(tmpSeats, rowOfSeats);

	}

	if(changes == 0) {
		return countOccupiedSeats(tmpSeats)		
	} 
		
	return seatingRound2(tmpSeats)
	
}

func getSeats(lines [] string) [][]rune {
	seats := [][]rune{} 

	for _, row := range lines {
		rowOfSeats := []rune{}
		for _, seat := range row {
			rowOfSeats = append(rowOfSeats, seat)
		}
		seats = append(seats, rowOfSeats);
	}
	return seats
}

func checkSeatChange(x int, y int, seat rune, seats [][]rune) bool {
	if(seat == '.') {
		return false
	}
	adjacentSeats := getAdjacentSeats(x, y, seats)

	if(seat == 'L' && countRune(adjacentSeats, '#') == 0) {
		return true
	} else if (seat == '#' && countRune(adjacentSeats, '#') >= 4) {
		return true
	}
	return false
}

func checkSeatChange2(x int, y int, seat rune, seats [][]rune) bool {
	if(seat == '.') {
		return false
	}
	adjacentSeats := getSeatsInSight(x, y, seats)

	if(seat == 'L' && countRune(adjacentSeats, '#') == 0) {
		return true
	} else if (seat == '#' && countRune(adjacentSeats, '#') >= 5) {
		return true
	}
	return false
}

func getAdjacentSeats(x int, y int, seats [][]rune) []rune {
	
	adjacentSeats := []rune{}
	for yi := y-1; yi <= y+1; yi++ {
		for xi := x-1; xi <= x+1; xi++ {
				if (xi >= 0 && yi >= 0) && (xi < len(seats[0]) && yi < len(seats)) {
					if !(xi == x && yi == y) {
						
						adjacentSeats = append(adjacentSeats, seats[yi][xi])
					}
				}
		}
	}	
	return adjacentSeats
}

func getSeatsInSight(x int, y int, seats [][]rune) []rune {
	adjacentSeats := []rune{}


	for yi := -1; yi <= 1; yi++ {
		for xi := -1; xi <= 1; xi++ {
			if !(xi == 0 && yi == 0) {
				adjacentSeats = append(adjacentSeats, traverseLine(x, y, seats, yi, xi))
			}
		}
	}

	return adjacentSeats
}


func traverseLine(startX int, startY int, seats [][]rune, changeX int, changeY int) rune {
	// fmt.Printf("traverse line \n")
	i := 1
	for {

		newX := startX+(i*changeX)
		newY := startY+(i*changeY)

		if(newX < 0 || newY < 0 || newX >= len(seats[0]) || newY >= len(seats)) {
			break;
		}

		if(seats[newY][newX] != '.') {
			return seats[newY][newX]
		}
		i++
	}
	return '.'
}

func countRune(runes []rune, find rune) int {
	sum := 0
	for _, r := range runes {
		if r == find {
			sum++
		}
	}
	return sum
}

func switchSeat(seat rune) rune {
	if seat == '#' {
		return 'L'
	}
	return '#'
}

func countOccupiedSeats(seats [][]rune) int {
	sum := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				sum++
			}
		}
	}
	return sum
}