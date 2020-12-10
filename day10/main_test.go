package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA1(t *testing.T) {
	lines, _ := helpers.ReadInts("input_test_a_1")
	assert.Equal(t, 35, a(lines))
}

func TestA2(t *testing.T) {
	lines, _ := helpers.ReadInts("input_test_a_2")
	assert.Equal(t, 220, a(lines))
}


func TestB1(t *testing.T) {
	lines, _ := helpers.ReadInts("input_test_a_1")
	assert.Equal(t, 8, b(lines))
}

func TestB2(t *testing.T) {
	lines, _ := helpers.ReadInts("input_test_a_2")
	assert.Equal(t, 19208, b(lines))
}

func TestB3(t *testing.T) {
	lines, _ := helpers.ReadInts("input")
	assert.Equal(t, 453551299002368, b(lines))
}

