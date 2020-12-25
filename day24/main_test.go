package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 10, a(lines))
}

func TestSmall(t *testing.T) {
	lines := []string{"nwwswee"}
	assert.Equal(t, 1, a(lines))
}

func TestSolvedA(t *testing.T) {
	lines, _ := helpers.ReadLines("input")
	assert.Equal(t, 438, a(lines))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 2208, b(lines))
}

func TestCalculateAdjacentBlackTiles(t *testing.T) {

	tiles := map[int]map[int]map[int]bool{}
	tiles[0] = map[int]map[int]bool{}
	tiles[0][0] = map[int]bool{}
	tiles[0][0][0] = true
	assert.Equal(t, 0, calculateAdjacentBlackTiles(tiles, 0,0,0))

	tiles[0][-1] = map[int]bool{}
	tiles[0][-1][1] = true
	assert.Equal(t, 1, calculateAdjacentBlackTiles(tiles, 0,0,0))

	tiles[0][1] = map[int]bool{}
	tiles[0][1][-1] = true
	assert.Equal(t, 2, calculateAdjacentBlackTiles(tiles, 0,0,0))

	tiles[1] = map[int]map[int]bool{}
	tiles[1][-1] = map[int]bool{}
	tiles[1][-1][0] = true
	assert.Equal(t, 3, calculateAdjacentBlackTiles(tiles, 0,0,0))

	tiles[1][0] = map[int]bool{}
	tiles[1][0][-1] = true
	assert.Equal(t, 4, calculateAdjacentBlackTiles(tiles, 0,0,0))

	tiles[-1] = map[int]map[int]bool{}
	tiles[-1][1] = map[int]bool{}
	tiles[-1][1][0] = true
	assert.Equal(t, 5, calculateAdjacentBlackTiles(tiles, 0,0,0))

	tiles[-1][0] = map[int]bool{}
	tiles[-1][0][1] = true
	assert.Equal(t, 6, calculateAdjacentBlackTiles(tiles, 0,0,0))
}