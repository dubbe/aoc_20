package main

import (
	"fmt"
	"testing"

	"github.com/dubbe/advent-of-code-2020/helpers"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	lines, _ := helpers.ReadGroups("input_test")
	assert.Equal(t, 20899048083289, a(lines))
}

func TestASolution(t *testing.T) {
	lines, _ := helpers.ReadGroups("input")
	assert.Equal(t, 5966506063747, a(lines))
}


func TestB(t *testing.T) {
	lines, _ := helpers.ReadGroups("input_test")
	assert.Equal(t, 0, b(lines))
}

func TestRotateMatrix(t *testing.T) {
	lines, _ := helpers.ReadGroups("input_test_2")

	fmt.Println("------ Original")
	matrix := parseMatrix(lines[0])
	printMatrix(matrix)

	fmt.Println("------ Rotate")
	matrix = rotateMatrix(matrix)
	printMatrix(matrix)

	fmt.Println("------ Flip")
	matrix = flipMatrixVert(matrix)
	printMatrix(matrix)

	fmt.Println("------ Flip Horizontal")
	matrix = flipMatrixHor(matrix)
	printMatrix(matrix)

	fmt.Println("------")

	assert.Equal(t, 1, 1)
}

func TestGetSidesFromMatrix(t *testing.T) {
	lines, _ := helpers.ReadGroups("input_test_2")
	matrix := parseMatrix(lines[0])

	printMatrix(matrix)

	assert.Equal(t, "#.#.##.##.", getTopFromMatrix(matrix))
	assert.Equal(t, "#..###.#.#", getBottomFromMatrix(matrix))
	assert.Equal(t, "##...#.#.#", getLeftFromMatrix(matrix))
	assert.Equal(t, ".##..#...#", getRightFromMatrix(matrix))
}