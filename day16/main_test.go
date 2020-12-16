package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 71, a(lines))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test_b")
	assert.Equal(t, 0, b(lines))
}

func TestBSolution(t *testing.T) {
	lines, _ := helpers.ReadLines("input")
	assert.Equal(t, 3709435214239, b(lines))
}

func TestFindIndexInSlice(t *testing.T) {
	slice := []int{4, 5, 6, 8}
	index, err := findIndexInSlice(slice, 6)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, index)
}
