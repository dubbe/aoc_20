package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadInts("input_test")
	assert.Equal(t, 127, a(lines, 5))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadInts("input_test")
	assert.Equal(t, 62, b(lines, 127))
}

