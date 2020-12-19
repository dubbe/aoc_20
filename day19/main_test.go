package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 2, a(lines))
}

func TestBA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test_b")
	assert.Equal(t, 3, a(lines))
}

func TestBB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test_b")
	assert.Equal(t, 12, b(lines, 0))
}

func TestASolution(t *testing.T) {
	lines, _ := helpers.ReadLines("input")
	assert.Equal(t, 180, a(lines))
}

func TestBSolution(t *testing.T) {
	lines, _ := helpers.ReadLines("input")
	assert.Equal(t, 323, b(lines, 0))
}