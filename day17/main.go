package main

import (
	"fmt"
	"time"

	"github.com/dubbe/advent-of-code-2020/helpers"
)

func main() {
	start := time.Now()
	lines, err := helpers.ReadLines("input")
	helpers.Check(err)
	fmt.Printf("result A: %v\n", a(lines))
	fmt.Printf("result B: %v\n", b(lines))
	elapsed := time.Since(start)
	fmt.Printf("Solution took %s", elapsed)
}

func a(lines []string) int {

	grid := getGrid(lines)
	for i := 1; i <= 6; i++ {
		grid = cycle(grid, -i)
	}
	sum := countAllActive(grid)
	return sum
}

func b(lines []string) int {
	grid := getGridB(lines)
	for i := 1; i <= 6; i++ {
		grid = cycleB(grid, -i)
	}
	sum := countAllActiveB(grid)
	return sum
}

func getGrid(lines []string) map[int]map[int]map[int]rune {
	grid := map[int]map[int]map[int]rune{}

	for z := 0; z < 1; z++ {
		rows := map[int]map[int]rune{}
		for y, line := range lines {
			cubes := map[int]rune{}
			for x, r := range line {
				cubes[x] = r
			}
			rows[y] = cubes
		}
		grid[z] = rows
	}
	return grid
}

func getGridB(lines []string) map[int]map[int]map[int]map[int]rune {
	grid := map[int]map[int]map[int]map[int]rune{}

	for w := 0; w < 1; w++ {
		panes := map[int]map[int]map[int]rune{}
		for z := 0; z < 1; z++ {
			rows := map[int]map[int]rune{}
			for y, line := range lines {
				cubes := map[int]rune{}
				for x, r := range line {
					cubes[x] = r
				}
				rows[y] = cubes
			}
			panes[z] = rows
		}
		grid[w] = panes
	}
	return grid
}

func cycle(grid map[int]map[int]map[int]rune, start int) map[int]map[int]map[int]rune {
	length := len(grid[0][0]) + start
	paneLength := len(grid) + start

	newPanes := map[int]map[int]map[int]rune{}
	for z := start; z <= paneLength+1; z++ {
		newPanes[z] = map[int]map[int]rune{}
		for y := start; y <= length+1; y++ {
			newPanes[z][y] = map[int]rune{}
			for x := start; x <= length+1; x++ {
				value, found := grid[z][y][x]
				if !found {
					value = '.'
				}
				activeNeighbours := getActiveNeighbours(grid, z, y, x)
				if value == '#' {
					if activeNeighbours == 2 || activeNeighbours == 3 {
						newPanes[z][y][x] = '#'
					} else {
						newPanes[z][y][x] = '.'
					}
				} else if value == '.' {
					if activeNeighbours == 3 {
						newPanes[z][y][x] = '#'
					} else {
						newPanes[z][y][x] = '.'
					}
				} else {
					newPanes[z][y][x] = value
				}

			}
		}
	}
	return newPanes
}

func cycleB(grid map[int]map[int]map[int]map[int]rune, start int) map[int]map[int]map[int]map[int]rune {
	length := len(grid[0][0][0]) + start
	paneLength := len(grid) + start

	newPanes := map[int]map[int]map[int]map[int]rune{}
	for w := start; w <= paneLength+1; w++ {
		newPanes[w] = map[int]map[int]map[int]rune{}
		for z := start; z <= paneLength+1; z++ {
			newPanes[w][z] = map[int]map[int]rune{}
			for y := start; y <= length+1; y++ {
				newPanes[w][z][y] = map[int]rune{}
				for x := start; x <= length+1; x++ {
					value, found := grid[w][z][y][x]
					if !found {
						value = '.'
					}
					activeNeighbours := getActiveNeighboursB(grid, w, z, y, x)
					if value == '#' {
						if activeNeighbours == 2 || activeNeighbours == 3 {
							newPanes[w][z][y][x] = '#'
						} else {
							newPanes[w][z][y][x] = '.'
						}
					} else if value == '.' {
						if activeNeighbours == 3 {
							newPanes[w][z][y][x] = '#'
						} else {
							newPanes[w][z][y][x] = '.'
						}
					} else {
						newPanes[w][z][y][x] = value
					}

				}
			}
		}
	}
	return newPanes
}

func printGrid(grid map[int]map[int]map[int]rune, start int) {
	length := len(grid[0][0]) + start
	paneLength := len(grid) + start

	for z := start; z < paneLength; z++ {
		fmt.Printf("pane: %d \n", z)
		for y := start; y < length; y++ {
			for x := start; x < length; x++ {
				fmt.Printf("%s ", string(grid[z][y][x]))
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func getActiveNeighbours(grid map[int]map[int]map[int]rune, sz int, sy int, sx int) int {
	countActive := 0
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if !(z == 0 && y == 0 && x == 0) {

					value, found := grid[sz+z][sy+y][sx+x]
					if found && value == '#' {
						countActive++
					}
				}
			}
		}
	}
	return countActive
}

func getActiveNeighboursB(grid map[int]map[int]map[int]map[int]rune, sw int, sz int, sy int, sx int) int {
	countActive := 0
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if !(w == 0 && z == 0 && y == 0 && x == 0) {
						value, found := grid[sw+w][sz+z][sy+y][sx+x]
						if found && value == '#' {
							countActive++
						}
					}
				}
			}
		}
	}
	return countActive
}

func countAllActive(grid map[int]map[int]map[int]rune) int {
	sum := 0
	for _, panes := range grid {
		for _, rows := range panes {
			for _, cube := range rows {
				if cube == '#' {
					sum++
				}
			}
		}
	}
	return sum
}

func countAllActiveB(grid map[int]map[int]map[int]map[int]rune) int {
	sum := 0
	for _, panes := range grid {
		for _, grids := range panes {
			for _, rows := range grids {
				for _, cube := range rows {
					if cube == '#' {
						sum++
					}
				}
			}
		}
	}
	return sum
}
