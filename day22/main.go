package main

import (
	"fmt"
	"strconv"
	"strings"
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
	hands := parseHands(lines)
	

	for {
		hand1 := 0
		hand2 := 0

		hand1, hands[1] = hands[1][0], hands[1][1:]
		hand2, hands[2] = hands[2][0], hands[2][1:]

		if hand1 > hand2 {
			hands[1] = append(hands[1],hand1)
			hands[1] = append(hands[1], hand2)
		}

		if hand1 < hand2 {
			hands[2] = append(hands[2], hand2)
			hands[2] = append(hands[2], hand1)
		}

		if len(hands[1]) == 0 || len(hands[2]) == 0 {
			break
		}
	}

	return calculateWinningHand(hands)
}

func b(lines []string) int {
	hands := parseHands(lines)

	
	playGame(hands, 1, 1)

	return calculateWinningHand(hands)
}

func playGame(hands map[int][]int, round int, game int) int {
	previousHands := map[int][][]int{}

	for {
		hand := map[int]int{}

		if contains(previousHands[1], hands[1]) && contains(previousHands[2], hands[2]) {
			return 1
		}

		previousHands[1] = append(previousHands[1], hands[1])
		previousHands[2] = append(previousHands[2], hands[2])
		
		hand[1], hands[1] = hands[1][0], hands[1][1:]
		hand[2], hands[2] = hands[2][0], hands[2][1:]

		roundWinner := 0
		
		if hand[1] <= len(hands[1]) && hand[2] <= len(hands[2]) {
			newHands := copyMap(hands)
			newHands[1] = newHands[1][:hand[1]]
			newHands[2] = newHands[2][:hand[2]]
			roundWinner = playGame(newHands, 1, game + 1)
		} else {

			if hand[1] > hand[2] {
				roundWinner = 1			
			}

			if hand[1] < hand[2] {
				roundWinner = 2
			}
		}

		hands[roundWinner] = append(hands[roundWinner], hand[roundWinner])
		roundLoser := 2
		if roundWinner == 2 {
			roundLoser = 1
		}		
		hands[roundWinner] = append(hands[roundWinner], hand[roundLoser])

		round++

		if len(hands[1]) == 0 {
			return 2
		} else if len(hands[2]) == 0 {
			return 1
		}
	}
}

func reverse(a  []int) []int {
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

func parseHands(lines []string) map[int][]int {
	hands := map[int][]int{}
	currentParsingPlayer := 0
	for _, line := range lines {
		if(strings.HasPrefix(line, "Player")) {
			player := strings.Split(line, " ")
			currentParsingPlayer, _ =strconv.Atoi(strings.TrimRight(player[1], ":"))
			hands[currentParsingPlayer] = []int{}
		} else {
			card, err := strconv.Atoi(line)
			if err == nil {
				hands[currentParsingPlayer] = append(hands[currentParsingPlayer], card)
			}
		}
	}
	return hands
}

func calculateWinningHand(hands map[int][]int) int {
	winningHand := reverse(hands[1])
	if len(hands[2]) > 0 {
		winningHand = reverse(hands[2])
	}

	sum := 0
	for i, hand := range winningHand {
		sum += (i+1) * hand
	}
	return sum
}

func copyMap(a map[int][]int) map[int][]int {
	ret := map[int][]int{}
	for k, v := range a {
		l := []int{}
		for _, i := range v {
			l = append(l, i)
		}
		ret[k] = l
	}
	return ret
}



func contains(s [][]int, e []int) bool {
	for _, a := range s {
			if equal(a, e) {
					return true
			}
	}
	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
			return false
	}
	for i, v := range a {
			if v != b[i] {
					return false
			}
	}
	return true
}