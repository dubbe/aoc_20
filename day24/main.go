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
	elapsed := time.Since(start)
	fmt.Printf("result A took %s\n", elapsed)

	startB := time.Now()
	fmt.Printf("result B: %v\n", b(lines))
	elapsed = time.Since(startB)
	fmt.Printf("result B took %s\n", elapsed)

	elapsed = time.Since(start)
	fmt.Printf("Solution took %s±n", elapsed)
}

func a(lines []string) int {
	tiles := startTiles(lines)

	sum := 0
	for _, values := range tiles {
		for _, val := range values {
			for _, v := range val {
				if v {
					sum++
				}
			}
		}
	}

	return sum
}

func b(lines []string) int {
	tiles := startTiles(lines)

	for i:=0;i<100;i++ {

		for z, values := range tiles {
			for y, value := range values {
				for x := range value {
					for _, d := range getCubeDirections() {
						_, ok := tiles[d.dz+z][d.dy+y][d.dx+x]
						if !ok {
							_, okz := tiles[d.dz+z]
							if !okz {
								tiles[d.dz+z] = map[int]map[int]bool{}
							}

							_, oky := tiles[d.dz+z][d.dy+y]
							if !oky {
								tiles[d.dz+z][d.dy+y] = map[int]bool{}
							}
							tiles[d.dz+z][d.dy+y][d.dx+x] = false
						}
					}
				}
			}
		}

		flippedTiles := map[int]map[int]map[int]bool{}


		for z, values := range tiles {
			for y, value := range values {
				for x, tile := range value {
					newTile := tile
					blackTiles := calculateAdjacentBlackTiles(tiles, z, y, x)

					if tile && (blackTiles == 0 || blackTiles > 2) {
						newTile = false
					} else if !tile &&  blackTiles == 2 {
						newTile = !tile
					}
					
					_, ok := flippedTiles[z]
					if !ok {
						flippedTiles[z] = map[int]map[int]bool{}
					}

					_, ok = flippedTiles[z][y]
					if !ok {
						flippedTiles[z][y] = map[int]bool{}
					}

					flippedTiles[z][y][x] = newTile

				}
			}
		}
		tiles = flippedTiles
	}

	return calculateBlackTiles(tiles)
}

func calculateBlackTiles(tiles map[int]map[int]map[int]bool) int {
	sum := 0
	for _, values := range tiles {
		for _, val := range values {
			for _, v := range val {
				if v {
					sum++
				}
			}
		}
	}
	return sum
}

func printTiles(tiles map[int]map[int]map[int]bool) { 
	for z, values := range tiles {
		for y, val := range values {
			for x, v := range val {
				if v {
					fmt.Printf("x: %d, y: %d, z: %d: black \n", x, y, z)
				} else {
					fmt.Printf("x: %d, y: %d, z: %d: white \n", x, y, z)
				}
			}
		}
	}
}

type cubeDirection struct {
	dz int
	dy int 
	dx int
}

func getCubeDirections() []cubeDirection {
	return []cubeDirection{
		{+1, -1, 0},
		{+1, 0, -1},
		{0, +1, -1},
		{-1, +1, 0},
		{-1, 0, +1},
		{0, -1, +1},
	}
}


func calculateAdjacentBlackTiles(tiles map[int]map[int]map[int]bool, z int, y int, x int) int {
	ret := 0
	for _, a := range getCubeDirections() {			
		tile, ok := tiles[a.dz+z][a.dy+y][a.dx+x]
		if ok && tile {
				ret++
			} 
	}

	return ret
}

func startTiles(lines []string) map[int]map[int]map[int]bool {
	directions := [][]string{}
	for _, line := range lines {
		dir := []string{}
		x := 0
		for {
			if line[x] == 'e' || line[x] == 'w' {
				dir = append(dir, fmt.Sprintf("%s", string(line[x])))
				x++
			} else {
				dir = append(dir, fmt.Sprintf("%s%s", string(line[x]), string(line[x+1])))
				x = x + 2
			}
			if x == len(line) {
				break
			}
		}
		directions = append(directions, dir)
	}

	tiles := map[int]map[int]map[int]bool{}

	for _, dir := range directions {

		x := 0
		y := 0
		z := 0

		for _, direction := range dir {
			switch direction {
			case "e":
				x++
				y--
			case "se":
				y--
				z++
			case "sw":
				x--
				z++
			case "w":
				x--
				y++
			case "nw":
				z--
				y++
			case "ne":
				z--
				x++
			}

		}
		_, ok := tiles[z]
		if !ok {
			tiles[z] = map[int]map[int]bool{}
		}

		_, ok = tiles[z][y]
		if !ok {
			tiles[z][y] = map[int]bool{}
		}

		value, ok := tiles[z][y][x]
		if !ok || !value {
			tiles[z][y][x] = true
		} else {
			tiles[z][y][x] = false
		}

	}

	return tiles
}