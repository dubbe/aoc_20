package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 306, a(lines))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 291, b(lines))
}

func TestB2(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test_2")
	assert.Equal(t, 78, b(lines))
}
