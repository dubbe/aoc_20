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

func a(lines []string) int {
	facingDirection := 'E'
	x := 0
	y := 0

	for _, line := range lines {
		newDirection := facingDirection
		for i, r := range line {
			if i == 0 {
				newDirection = r
				break
			}
		}
		distance, _ := strconv.Atoi(line[1:])

		if(newDirection == 'F') {
			newDirection = facingDirection
		} else if (newDirection == 'L' || newDirection == 'R' ) {
			facingDirection = turnShip(facingDirection, distance, newDirection)
			continue;
		} 

		switch newDirection {
		case 'N':
			y += distance
		case 'E':
			x += distance
		case 'S':
			y += (distance*-1)
		case 'W':
			x += (distance*-1)
		default:
			fmt.Printf("ERROR, direction: %v, distance: %d \n", string(newDirection), distance)
		}
	}

	if(x < 0) {
		x = x*-1
	}

	if(y < 0) {
		y = y*-1
	}
	return x+y
}

func b(lines []string) int {
	facingDirection := 'E'
	x := 0
	y := 0
	waypointX := 10
	waypointY := 1

	for _, line := range lines {
		newDirection := facingDirection
		for i, r := range line {
			if i == 0 {
				newDirection = r
				break
			}
		}
		distance, _ := strconv.Atoi(line[1:])

		if(newDirection == 'F') {
			x += waypointX * distance
			y += waypointY * distance
			continue;
		} else if (newDirection == 'L' || newDirection == 'R' ) {
			facingDirection, waypointX, waypointY = rotateShip(facingDirection, distance, newDirection, waypointX, waypointY)
			continue;
		} 

		switch newDirection {
		case 'N':
			waypointY += distance
		case 'E':
			waypointX += distance
		case 'S':
			waypointY += (distance*-1)
		case 'W':
			waypointX += (distance*-1)
		default:
			fmt.Printf("ERROR, direction: %v, distance: %d \n", string(newDirection), distance)
		}
	}

	if(x < 0) {
		x = x*-1
	}

	if(y < 0) {
		y = y*-1
	}
	return x+y

}

func turnShip(currentDirection rune, degrees int, direction rune) rune {
	possibleDirections := []rune{'N', 'E', 'S', 'W'}
	currentIndex := 0
	for i, d := range possibleDirections {
		if(d == currentDirection) {
			currentIndex = i
		}
	}

	change := degrees / 90

	if direction == 'R' {
		newIndex := change + currentIndex
		for {
			if newIndex >= 4  {
				newIndex = newIndex - len(possibleDirections)
			} else {
				break
			}
		}
		return possibleDirections[newIndex]
	} else {
		newIndex := currentIndex - change 
		for {
			if newIndex < 0 {
				newIndex = newIndex + len(possibleDirections)
			} else {
				break
			}
		}
		return possibleDirections[newIndex]
	}

}

func rotateShip(currentDirection rune, degrees int, direction rune, x int, y int) (rune, int, int) {
	newDirection := turnShip(currentDirection, degrees, direction)

	n := 0
	e := 0
	s := 0
	w := 0

	if(x > 0) {
		e = x
	} else {
		w = x * -1
	}

	if(y > 0) {
		n = y
	} else {
		s = y * -1
	}

	newN := 0
	newE := 0
	newS := 0
	newW := 0
	if(degrees == 180) {
		newN = s
		newS = n
		newE = w
		newW = e
	} else if(degrees == 90) {
		if(direction == 'R') {
			newE = n
			newS = e
			newW = s
			newN = w
		} else if(direction == 'L') {
			newE = s
			newS = w
			newW = n
			newN = e
		}
	} else if(degrees == 270) {
		if(direction == 'R') {
			newE = s
			newS = w
			newW = n
			newN = e
		} else if(direction == 'L') {
			newE = n
			newS = e
			newW = s
			newN = w	
		}
	} else {
		fmt.Printf("ERROR, not a degree I like... %d \n", degrees)
	}

	if(newN != 0) {
		y = newN
	} else {
		y = newS * -1
	}

	if(newE != 0) {
		x = newE
	} else {
		x = newW * -1
	}

	return newDirection, x, y
}
