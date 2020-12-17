package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 112, a(lines))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 848, b(lines))
}

func TestCountNeighbours(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	grid := getGrid(lines)
	//assert.Equal(t, 0, getActiveNeighbours(grid, 0,0,0))

	//assert.Equal(t, 3, getActiveNeighbours(grid, 0,1,0))
	//assert.Equal(t, 5, getActiveNeighbours(grid, 0,1,1))
	assert.Equal(t, 3, getActiveNeighbours(grid, 0, 1, 2))
}
