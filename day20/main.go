package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dubbe/advent-of-code-2020/helpers"
)

func main() {
	start := time.Now()
	lines, err := helpers.ReadGroups("input")
	helpers.Check(err)
	//fmt.Printf("result A: %v\n", a(lines))
	fmt.Printf("result B: %v\n", b(lines))
	elapsed := time.Since(start)
	fmt.Printf("Solution took %s", elapsed)
}

type PicturePart struct {
	ID          int
	Matrix      map[int]map[int]rune
	MostMatches int
	Matches     []MatchedPicturePart
}

type MatchedPicturePart struct {
	Matrix map[int]map[int]rune
	Left   int
	Right  int
	Top    int
	Bottom int
	ID     int
}

func a(groups []string) int {
	pictureParts := parseParts(groups)

	foundParts := []PicturePart{}
	j := 0
	for _, v := range pictureParts {
		foundParts = append(foundParts, findMatches(v, pictureParts))
		j++
	}

	sum := 1
	for _, part := range foundParts {
		if part.MostMatches == 2 {
			sum *= part.ID
		}
	}
	return sum
}

func b(groups []string) int {
	for {

		pictureParts := parseParts(groups)

		foundParts := map[int]PicturePart{}
		for _, part := range pictureParts {
			p := findMatches(part, pictureParts)
			foundParts[p.ID] = p
		}

		possibleCornerParts := []PicturePart{}
		for _, part := range foundParts {
			if part.MostMatches == 2 {
				possibleCornerParts = append(possibleCornerParts, part)
			}
		}

		m := map[int]map[int]MatchedPicturePart{}
		m[0] = map[int]MatchedPicturePart{}
	Test:
		for _, part := range possibleCornerParts {
			for _, match := range part.Matches {
				if match.Right != 0 && match.Bottom != 0 {
					match.ID = part.ID
					m[0][0] = match
					break Test
				}
			}
		}

		i := 0

		for {
			j := 0
			for {
				p, ok := m[i][j]
				if ok {
					if p.Right == 0 {
						break
					}
					part := foundParts[p.Right]

					for _, match := range part.Matches {
						if match.Left == p.ID {
							found := true
							for x := 0; x < len(p.Matrix); x++ {
								if p.Matrix[x][len(p.Matrix[0])-1] != match.Matrix[x][0] {
									found = false
									break
								}
							}
							if found {
								match.ID = part.ID
								m[i][j+1] = match
								break
							}

						}
					}

				} else {
					break
				}
				j++
			}
			p, ok := m[i][0]
			if ok {
				if p.Bottom == 0 {
					break
				}
				m[i+1] = map[int]MatchedPicturePart{}
				part := foundParts[p.Bottom]
				for _, match := range part.Matches {
					if match.Top == p.ID {
						found := true
						fmt.Printf("-----------\n")
						printMatrix(p.Matrix, true)
						fmt.Printf("\n")
						printMatrix(match.Matrix, true)
						for x := 0; x < len(p.Matrix); x++ {

							fmt.Printf("x: %d, y: %d", x, len(p.Matrix)-1)
							fmt.Printf("%s %s \n", string(p.Matrix[len(p.Matrix)-1][x]), string(match.Matrix[0][x]))

							if p.Matrix[len(p.Matrix)-1][x] != match.Matrix[0][x] {
								found = false
								break
							}
						}
						if found {
							match.ID = part.ID
							m[i+1][0] = match
							break
						}

					}
				}
			} else {
				break
			}
			i++
		}

		finalImage := ""
		finalImageTest := map[int]map[int]rune{}

		// matrix grid
		i = 0
		j := 0

		// image grid
		x := 1
		y := 1
		matrix := m[j][i].Matrix

		w := 0
		z := 0
		finalImageTest[z] = map[int]rune{}
		for {
			finalImage += string(matrix[y][x])
			finalImageTest[z][w] = matrix[y][x]
			w++
			x++
			if x > len(matrix)-2 {
				x = 1
				i++

				if i > len(m)-1 {
					i = 0
					w = 0
					y++

					if y > len(matrix[y])-2 {
						y = 1
						j++

						if j > len(m[j])-1 {
							break
						}

					}
					z++
					finalImageTest[z] = map[int]rune{}
					finalImage += "\n"
				}
				matrix = m[j][i].Matrix
			}
		}

		// fmt.Print(finalImage)
		newImage := ""
	FindImage:
		for x := 0; x < 4; x++ {
			for y := 0; y < 3; y++ {

				image := sprintMatrix(finalImageTest, false)
				found := false
				for l := 0; l < len(finalImageTest[0])-len(".#..#..#..#..#..#..."); l++ {
					re := regexp.MustCompile(fmt.Sprintf(".{%d}..................(#)..*\n.{%d}(#)....(#)(#)....(#)(#)....(#)(#)(#).*\n.{%d}.(#)..(#)..(#)..(#)..(#)..(#)....*", l, l, l))
					for re.MatchString(image) {
						found = true

						index := re.FindStringSubmatchIndex(image)

						for w := 2; w < len(index); w += 2 {
							in := index[w]
							image = image[:in] + "O" + image[in+1:]

						}
					}
				}

				if found {
					newImage = image
					break FindImage
				}

				if x == 0 {
					finalImageTest = flipMatrixHor(finalImageTest)
				} else if x == 1 {
					finalImageTest = flipMatrixHor(finalImageTest)
					finalImageTest = flipMatrixVert(finalImageTest)
				} else {
					finalImageTest = flipMatrixVert(finalImageTest)
				}
			}

			finalImageTest = rotateMatrix(finalImageTest)
		}
		sum := strings.Count(newImage, "#")
		if sum != 0 {
			fmt.Print(newImage)

			return sum
		}
	}
}

func joinMaps(left, right map[int]rune) map[int]rune {
	for key, rightVal := range right {
		left[key] = rightVal
	}
	return left
}

func parseParts(groups []string) map[int]PicturePart {
	pictureParts := map[int]PicturePart{}
	for _, group := range groups {
		part := parsePart(group)
		pictureParts[part.ID] = part

	}
	return pictureParts
}

func parsePart(part string) PicturePart {
	picturePart := PicturePart{}
	title := strings.Split(strings.Split(part, "\n")[0], " ")
	id, _ := strconv.Atoi(strings.TrimRight(title[1], ":"))
	picturePart.ID = id

	picturePart.Matrix = parseMatrix(part)
	return picturePart
}

func findMatches(part PicturePart, parts map[int]PicturePart) PicturePart {
	finalMatch := 0
	finalMatches := []MatchedPicturePart{}
	matrix := part.Matrix
	for i := 0; i < 4; i++ {
		for x := 0; x < 3; x++ {
			t := getTopFromMatrix(matrix)
			b := getBottomFromMatrix(matrix)
			l := getLeftFromMatrix(matrix)
			r := getRightFromMatrix(matrix)
			left := []int{}
			right := []int{}
			top := []int{}
			bottom := []int{}
			for _, p := range parts {
				if p.ID == part.ID {
					continue
				}
				pMatrix := p.Matrix
				for y := 0; y < 6; y++ {
					for w := 0; w < 3; w++ {

						if t == getBottomFromMatrix(pMatrix) {
							top = append(top, p.ID)
						}
						if b == getTopFromMatrix(pMatrix) {
							bottom = append(bottom, p.ID)
						}
						if l == getRightFromMatrix(pMatrix) {
							left = append(left, p.ID)
						}
						if r == getLeftFromMatrix(pMatrix) {
							right = append(right, p.ID)
						}
						if w == 0 {
							pMatrix = flipMatrixHor(pMatrix)
						} else if w == 1 {
							pMatrix = flipMatrixHor(pMatrix)
							pMatrix = flipMatrixVert(pMatrix)
						} else {
							pMatrix = flipMatrixVert(pMatrix)
						}
					}
					pMatrix = rotateMatrix(pMatrix)
				}

			}

			top = removeDuplicates(top)
			right = removeDuplicates(right)
			bottom = removeDuplicates(bottom)
			left = removeDuplicates(left)

			match := len(top) + len(right) + len(bottom) + len(left)
			if match >= finalMatch {
				m := MatchedPicturePart{}

				if len(top) == 1 {
					m.Top = top[0]
				}
				if len(right) == 1 {
					m.Right = right[0]
				}
				if len(bottom) == 1 {
					m.Bottom = bottom[0]
				}
				if len(left) == 1 {
					m.Left = left[0]
				}
				m.Matrix = matrix

				if match == finalMatch {
					finalMatches = append(finalMatches, m)
				} else {
					finalMatch = match
					finalMatches = []MatchedPicturePart{m}
				}

			}

			if x == 0 {
				matrix = flipMatrixHor(matrix)
			} else if x == 1 {
				matrix = flipMatrixHor(matrix)
				matrix = flipMatrixVert(matrix)
			} else {
				matrix = flipMatrixVert(matrix)
			}
		}

		matrix = rotateMatrix(matrix)
	}

	part.MostMatches = finalMatch
	part.Matches = finalMatches
	return part
}

func getTopFromMatrix(matrix map[int]map[int]rune) string {
	ret := ""
	for i := 0; i < len(matrix); i++ {
		ret += fmt.Sprintf("%s", string(matrix[0][i]))
	}
	return ret
}

func getBottomFromMatrix(matrix map[int]map[int]rune) string {
	ret := ""
	j := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		ret += fmt.Sprintf("%s", string(matrix[j][i]))
	}
	return ret
}

func getLeftFromMatrix(matrix map[int]map[int]rune) string {
	ret := ""
	for i := 0; i < len(matrix); i++ {
		ret += fmt.Sprintf("%s", string(matrix[i][0]))
	}
	return ret
}

func getRightFromMatrix(matrix map[int]map[int]rune) string {
	ret := ""
	j := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		ret += fmt.Sprintf("%s", string(matrix[i][j]))
	}
	return ret
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func removeDuplicates(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func parseMatrix(part string) map[int]map[int]rune {
	lines := strings.Split(part, "\n")

	ret := map[int]map[int]rune{}
	for i, line := range lines {
		if i == 0 {
			continue
		}
		ret[i-1] = map[int]rune{}
		for j, r := range line {
			ret[i-1][j] = r
		}
	}
	return ret
}

func rotateMatrix(matrix map[int]map[int]rune) map[int]map[int]rune {
	ret := map[int]map[int]rune{}
	n := len(matrix[0])
	for i := 0; i < n; i++ {
		ret[i] = map[int]rune{}
		for j := 0; j < n; j++ {
			ret[i][j] = matrix[n-j-1][i]
		}
	}

	return ret
}

func flipMatrixVert(matrix map[int]map[int]rune) map[int]map[int]rune {
	ret := map[int]map[int]rune{}
	n := len(matrix[0])
	for i := 0; i < n; i++ {
		ret[i] = matrix[n-i-1]
	}

	return ret
}

func flipMatrixHor(matrix map[int]map[int]rune) map[int]map[int]rune {
	ret := map[int]map[int]rune{}
	n := len(matrix[0])
	for i := 0; i < n; i++ {
		ret[i] = map[int]rune{}
		for j := 0; j < n; j++ {
			ret[i][j] = matrix[i][n-j-1]
		}
	}

	return ret
}

func printPicturePart(pp PicturePart) {
	fmt.Printf("\nid: %d\n", pp.ID)

	printMatches(pp.Matches)
	printMatrix(pp.Matrix, true)
}

func printMatches(matches []MatchedPicturePart) {
	for _, match := range matches {

		fmt.Printf("top: %d, right: %d, bottom: %d, left: %d \n ", match.Top, match.Right, match.Bottom, match.Left)
		printMatrix(match.Matrix, true)
		fmt.Printf("\n")
	}
}

func printMatrix(matrix map[int]map[int]rune, printSpace bool) {
	for i := 0; i < len(matrix); i++ {
		value := matrix[i]
		for j := 0; j < len(value); j++ {
			v := value[j]
			if printSpace {
				fmt.Printf("%s ", string(v))
			} else {
				fmt.Printf("%s", string(v))
			}
		}
		fmt.Printf("\n")
	}
}

func sprintMatrix(matrix map[int]map[int]rune, printSpace bool) string {
	ret := ""
	for i := 0; i < len(matrix); i++ {
		value := matrix[i]
		for j := 0; j < len(value); j++ {
			v := value[j]
			if printSpace {
				ret += fmt.Sprintf("%s ", string(v))
			} else {
				ret += fmt.Sprintf("%s", string(v))
			}
		}
		ret += fmt.Sprintf("\n")
	}
	return ret
}
