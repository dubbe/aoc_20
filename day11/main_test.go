package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 37, a(lines))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 26, b(lines))
}

func TestAdjacentSeat(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	seats := getSeats(lines)

	expected := []rune {'L', '.', 'L'}
	assert.Equal(t, expected, getAdjacentSeats(0,0, seats))
	assert.Equal(t, 2, countRune(expected, 'L'))

	expected = []rune {'L', 'L', 'L'}
	assert.Equal(t, expected, getAdjacentSeats(9,0, seats))

	expected = []rune {'L', '.', '.'}
	assert.Equal(t, expected, getAdjacentSeats(0,9, seats))

	expected = []rune {'.', 'L', 'L'}
	assert.Equal(t, expected, getAdjacentSeats(9,9, seats))

	expected = []rune {'L', 'L', 'L', '.', '.', 'L', 'L', 'L'}
	assert.Equal(t, expected, getAdjacentSeats(1,1, seats))
}

func TestAdjacentSeatB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test_b")
	seats := getSeats(lines)

	expected := []rune {'#', '#', '.', '#', '.'}
	assert.Equal(t, expected, getAdjacentSeats(0,1, seats))
	assert.Equal(t, 3, countRune(expected, '#'))
	assert.Equal(t, false, checkSeatChange(0,1, seats[0][1], seats))
}

func TestCount(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	seats := getSeats(lines)
	assert.Equal(t, 0, countOccupiedSeats(seats))

	lines, _ = helpers.ReadLines("input_test_b")
	seats = getSeats(lines)
	assert.Equal(t, 71, countOccupiedSeats(seats))
}

func TestSeatInSightB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test_b")
	seats := getSeats(lines)

	expected := []rune {'.', '.', '.', '#', '#', '.', '#', '#'}

	assert.Equal(t, expected, getSeatsInSight(0,1, seats))
	assert.Equal(t, 4, countRune(expected, '#'))
	assert.Equal(t, false, checkSeatChange2(0,1, seats[0][1], seats))

	expected = []rune {'#', '#', '#', '.', '#', '#', '#', '#'}

	assert.Equal(t, expected, getSeatsInSight(1,1, seats))
	assert.Equal(t, 7, countRune(expected, '#'))
	assert.Equal(t, false, checkSeatChange2(0,1, seats[0][1], seats))
}
