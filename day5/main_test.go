package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestA(t *testing.T) {
}

func TestParseTicket(t *testing.T) {
	result := parseTicket("FBFBBFFRLR");
	expected := [2]int{44, 5}
	assert.Equal(t, expected, result)
	assert.Equal(t, 357, calculateSeatID(result))

	result = parseTicket("BFFFBBFRRR");
	expected = [2]int{70, 7}
	assert.Equal(t, expected, result)
	assert.Equal(t, 567, calculateSeatID(result))

	result = parseTicket("FFFBBBFRRR");
	expected = [2]int{14, 7}
	assert.Equal(t, expected, result)
	assert.Equal(t, 119, calculateSeatID(result))

	result = parseTicket("BBFFBBFRLL");
	expected = [2]int{102, 4}
	assert.Equal(t, expected, result)
	assert.Equal(t, 820, calculateSeatID(result))
}

func TestB(t *testing.T) {

}