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
	fmt.Printf("result A: %v \n", a(lines))
	fmt.Printf("result B: %v \n", b(lines))
	elapsed := time.Since(start)
	fmt.Printf("Solution took %s \n", elapsed)
}

const (
	forward = 'F'
	left    = 'L'
	right   = 'R'
	north   = 'N'
	east    = 'E'
	south   = 'S'
	west    = 'W'
)

func a(lines []string) int {
	facingDirection := east
	x, y := 0, 0

	for _, line := range lines {
		newDirection, distance := getTravel(facingDirection, line)

		if newDirection == forward {
			newDirection = facingDirection
		} else if newDirection == left || newDirection == right {
			facingDirection = turnShip(facingDirection, distance, newDirection)
			continue
		}

		x, y = travel(x, y, newDirection, distance)
	}

	return abs(x) + abs(y)
}

func b(lines []string) int {
	facingDirection := east
	x, y := 0, 0

	waypointX := 10
	waypointY := 1

	for _, line := range lines {
		newDirection, distance := getTravel(facingDirection, line)

		if newDirection == forward {
			x += waypointX * distance
			y += waypointY * distance
			continue
		} else if newDirection == left || newDirection == right {
			facingDirection, waypointX, waypointY = rotateShip(facingDirection, distance, newDirection, waypointX, waypointY)
			continue
		}

		waypointX, waypointY = travel(waypointX, waypointY, newDirection, distance)
	}

	return abs(x) + abs(y)

}

func travel(x int, y int, direction rune, distance int) (int, int) {
	switch direction {
	case north:
		y += distance
	case east:
		x += distance
	case south:
		y += -distance
	case west:
		x += -distance
	default:
		fmt.Printf("ERROR, direction: %v, distance: %d \n", string(distance), distance)
	}

	return x, y
}

func getTravel(direction rune, line string) (rune, int) {
	newDirection := direction
	for i, r := range line {
		if i == 0 {
			newDirection = r
			break
		}
	}
	distance, _ := strconv.Atoi(line[1:])
	return newDirection, distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func turnShip(currentDirection rune, degrees int, direction rune) rune {
	possibleDirections := []rune{north, east, south, west}
	currentIndex := findIndex(possibleDirections, currentDirection)
	
	change := degrees / 90

	if direction == left {
		change = -change
	} 
	rotatedSlice := rotateRuneSlice(possibleDirections, change)
	return rotatedSlice[currentIndex]
}

func findIndex(slice []rune, find rune) int {
	for i, d := range slice {
		if d == find {
			return i
		}
	}
	return 0
}

func rotateRuneSlice(slice []rune, d int) []rune {
	if d < 0 {
		d = len(slice)+d
	}
	slice = append(slice[d:], slice[0:d]...)
	return slice
}

func rotateIntSlice(slice []int, d int) []int {
	if d < 0 {
		d = len(slice)+d
	}
	slice = append(slice[d:], slice[0:d]...)
	return slice
}

func rotateShip(currentDirection rune, degrees int, direction rune, x int, y int) (rune, int, int) {
	newDirection := turnShip(currentDirection, degrees, direction)

	coordinates := []int{0,0,0,0}

	if y > 0 {
		coordinates[0] = y
	} else {
		coordinates[2] = y * -1
	}

	if x > 0 {
		coordinates[1] = x
	} else {
		coordinates[3] = x * -1
	}

	change := degrees / 90

	if direction == right {
		change = -change
	} 

	coordinates = rotateIntSlice(coordinates, change)


	if coordinates[0] != 0 {
		y = coordinates[0]
	} else {
		y = coordinates[2] * -1
	}

	if coordinates[1] != 0 {
		x = coordinates[1]
	} else {
		x = coordinates[3] * -1
	}

	return newDirection, x, y
}
