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

// func TestDay01(t *testing.T) {
// 	test := []string{"..##.......",
// 		"#...#...#..",
// 		".#....#..#.",
// 		"..#.#...#.#",
// 		".#...##..#.",
// 		"..#.##.....",
// 		".#.#.#....#",
// 		".#........#",
// 		"#.##...#...",
// 		"#...##....#",
// 		".#..#...#.#"}

// 	fmt.Println(test[1][3] == '.')
// 	result := A(test)
// 	assertEqual(t, result, 7, "did not find 7 trees")

// }

func TestDay01B(t *testing.T) {
	test := []string{"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#"}

	result := checkSlope(test, 1, 1)
	assertEqual(t, result, 2, "did not find 2 trees")

	result = checkSlope(test, 3, 1)
	assertEqual(t, result, 7, "did not find 7 trees")

	result = checkSlope(test, 5, 1)
	assertEqual(t, result, 3, "did not find 3 trees")

	result = checkSlope(test, 7, 1)
	assertEqual(t, result, 4, "did not find 4 trees")

	result = checkSlope(test, 1, 2)
	assertEqual(t, result, 2, "did not find 2 trees")

	result = B(test)
	assertEqual(t, result, 336, "sum is not 336")

}