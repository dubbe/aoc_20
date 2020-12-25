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
