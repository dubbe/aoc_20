package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 26335, a(lines))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 0, b(lines))
}
func TestCalculate(t *testing.T) {
	assert.Equal(t, 71, calculate("1 + 2 * 3 + 4 * 5 + 6"))
	assert.Equal(t, 51, calculate("1 + (2 * 3) + (4 * (5 + 6))"))
	assert.Equal(t, 26, calculate("2 * 3 + (4 * 5)"))
	assert.Equal(t, 437, calculate("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	assert.Equal(t, 12240, calculate("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	assert.Equal(t, 13632, calculate("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
	assert.Equal(t, 8, calculate("(2 + 2) + (2 + 2)"))
	assert.Equal(t, 18, calculate("((1 + 1) + (1 + 1)) + 2 * 3"))
	assert.Equal(t, 312, calculate("9 * (6 + 7 + (7 * 3)) + 6"))
}

func TestCalculateB(t *testing.T) {
	assert.Equal(t, 51, calculateB("1 + (2 * 3) + (4 * (5 + 6))"))
	assert.Equal(t, 46, calculateB("2 * 3 + (4 * 5)"))
	assert.Equal(t, 1445, calculateB("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	assert.Equal(t, 669060, calculateB("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	assert.Equal(t, 23340, calculateB("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
}