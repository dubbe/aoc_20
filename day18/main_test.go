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
func TestCalculateLine(t *testing.T) {
	assert.Equal(t, 71, calculateLine("1 + 2 * 3 + 4 * 5 + 6"))
	assert.Equal(t, 51, calculateLine("1 + (2 * 3) + (4 * (5 + 6))"))
	assert.Equal(t, 26, calculateLine("2 * 3 + (4 * 5)"))
	assert.Equal(t, 437, calculateLine("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	assert.Equal(t, 12240, calculateLine("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	assert.Equal(t, 13632, calculateLine("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))
	assert.Equal(t, 8, calculateLine("(2 + 2) + (2 + 2)"))
}

