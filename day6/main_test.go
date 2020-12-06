package main

import (
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadGroups("input_test")
	assert.Equal(t, 11, a(lines))
}

func TestAResullt(t *testing.T) {
	lines, _ := helpers.ReadGroups("input")
	assert.Equal(t, 6799, a(lines))
}

func TestBResult(t *testing.T) {
	lines, _ := helpers.ReadGroups("input")
	assert.Equal(t, 3354, b(lines))
}

func TestB(t *testing.T) {
	lines, _ := helpers.ReadGroups("input_test")
	assert.Equal(t, 6, b(lines))
}

