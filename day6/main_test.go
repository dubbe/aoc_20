package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestA(t *testing.T) {
	lines, _ := readLines("input_test")
	assert.Equal(t, 11, a(lines))
}

func TestAResullt(t *testing.T) {
	lines, _ := readLines("input")
	assert.Equal(t, 6799, a(lines))
}

func TestBResullt(t *testing.T) {
	lines, _ := readLines("input")
	assert.Equal(t, 3354, b(lines))
}

func TestB(t *testing.T) {
	lines, _ := readLines("input_test")
	assert.Equal(t, 6, b(lines))
}