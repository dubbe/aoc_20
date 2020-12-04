package main

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
  if a == b {
    return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestDay01(t *testing.T) {
	valid := checkValid("abcde", 1, 3, 'a')
	assertEqual(t, valid, true, "not true")

	valid = checkValid("cdefg", 1, 3, 'b')
	assertEqual(t, valid, false, "not true")

	valid = checkValid("ccccccccc", 2, 9, 'c')
	assertEqual(t, valid, false, "not true")
}