package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, "92658374", a(lines, 10))
	// assert.Equal(t, "67384529", a(lines, 100))
}

func TestASolution(t *testing.T) {
	lines, _ := helpers.ReadLines("input")
	assert.Equal(t, "98645732", a(lines, 10))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 149245887792, b2(lines, 10000000))
}

