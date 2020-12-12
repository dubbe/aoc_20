package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 25, a(lines))
}

func TestTurn(t *testing.T) {
	assert.Equal(t, string('S'), string(turnShip('E', 90, 'R')))
	assert.Equal(t, string('N'), string(turnShip('E', 90, 'L')))
	assert.Equal(t, string('N'), string(turnShip('S', 180, 'L')))
	assert.Equal(t, string('N'), string(turnShip('W', 90, 'R')))
	assert.Equal(t, string('E'), string(turnShip('W', 180, 'R')))
	assert.Equal(t, string('S'), string(turnShip('W', 90, 'L')))
	assert.Equal(t, string('E'), string(turnShip('W', 180, 'L')))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 286, b(lines))
}

func TestB2(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test_b")
	assert.Equal(t, 374, b(lines))
}

func TestBSolution(t *testing.T) {
	lines, _ := helpers.ReadLines("input")
	assert.Equal(t, 47806, b(lines))
}

func TestRotateShip(t *testing.T) {
	d, x, y := rotateShip('E', 90, 'R', 10, 4)
	assert.Equal(t, string('S'), string(d))
	assert.Equal(t, 4, x)
	assert.Equal(t, -10, y)

	// d, x, y = rotateShip('E', 90, 'L', 10, 4)
	// assert.Equal(t, string('N'), string(d))
	// assert.Equal(t, -4, x)
	// assert.Equal(t, 10, y)

	// d, x, y = rotateShip('E', 180, 'R', 10, 4)
	// assert.Equal(t, string('W'), string(d))
	// assert.Equal(t, -10, x)
	// assert.Equal(t, -4, y)

	// d, x, y = rotateShip('E', 180, 'L', 10, 4)
	// assert.Equal(t, string('W'), string(d))
	// assert.Equal(t, -10, x)
	// assert.Equal(t, -4, y)

	// d, x, y = rotateShip('W', 90, 'R', -10, 4)
	// assert.Equal(t, string('N'), string(d))
	// assert.Equal(t, 4, x)
	// assert.Equal(t, 10, y)

	// d, x, y = rotateShip('W', 90, 'R', 1, -8)
	// assert.Equal(t, string('N'), string(d))
	// assert.Equal(t, -8, x)
	// assert.Equal(t, -1, y)

	// d, x, y = rotateShip('W', 270, 'R', -10, 4)
	// assert.Equal(t, string('S'), string(d))
	// assert.Equal(t, -4, x)
	// assert.Equal(t, -10, y)
}

func TestRotateSlice(t *testing.T) {
	slice := []rune{'N', 'E', 'S', 'W'}
	expected := []rune{'E', 'S', 'W', 'N'}
	assert.Equal(t, expected, rotateRuneSlice(slice, 1))

	slice = []rune{'N', 'E', 'S', 'W'}
	expected = []rune{'S', 'W', 'N', 'E'}
	assert.Equal(t, expected, rotateRuneSlice(slice, 2))

	slice = []rune{'N', 'E', 'S', 'W'}
	expected = []rune{'W', 'N', 'E', 'S'}
	assert.Equal(t, expected, rotateRuneSlice(slice, 3))

	slice = []rune{'N', 'E', 'S', 'W'}
	expected = []rune{'W', 'N', 'E', 'S'}
	assert.Equal(t, expected, rotateRuneSlice(slice, -1))

	slice = []rune{'N', 'E', 'S', 'W'}
	expected = []rune{'S', 'W', 'N', 'E'}
	assert.Equal(t, expected, rotateRuneSlice(slice, -2))
}
