package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test")
	assert.Equal(t, 4, a(lines))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadLines("input_test_b")
	assert.Equal(t, 126, b(lines))
}

