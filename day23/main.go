package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"time"

	"github.com/dubbe/advent-of-code-2020/helpers"
)


func main() {
	start := time.Now()
	lines, err := helpers.ReadLines("input")
	helpers.Check(err)
	fmt.Printf("result A: %v\n", a(lines, 100))
	elapsed := time.Since(start)
	fmt.Printf("result A took %s\n", elapsed)

	startB := time.Now()
	fmt.Printf("result B: %v\n", b2(lines, 10000000))
	elapsed = time.Since(startB)
	fmt.Printf("result B took %s\n", elapsed)
	
	elapsed = time.Since(start)
	fmt.Printf("Solution took %s±n", elapsed)
}

func a(lines []string, iterations int) string {
	cups := ring.New(len(lines[0]))
	cupsIndex := make([]*ring.Ring, len(lines[0]))

	i := 0
	for _, cup := range lines[0] {
		c, err := strconv.Atoi(string(cup))
		if err != nil {
			fmt.Println("ERRROR")
		} 
		cups.Value = c
		cupsIndex[c-1] = cups
		cups = cups.Next()
		i++
	}

	cups = playGame(cups, iterations, cupsIndex, i)
		
	returnRing := findRingFromValue(cups, 1)	

	ret := ""
	returnRing.Next()
	returnRing.Do(func(i interface{}) {
		if i.(int) == 1 {
			return
		}
		ret += fmt.Sprintf("%v", i)
	})

	return ret
}

func b2(lines []string, iterations int) int {
	numberOfCups := 1000000
	cups := ring.New(numberOfCups)
	cupsIndex := make([]*ring.Ring, numberOfCups)

	i:=0
	for _, cup := range lines[0] {
		c, err := strconv.Atoi(string(cup))
		if err != nil {
			fmt.Println("ERRROR")
		} else {
			cups.Value = c
			cupsIndex[c-1] = cups
			cups = cups.Next()
		}
		i++
	}

	maxNum := 0
	for i := ringMax(cups) + 1; i <= numberOfCups; i++ {
		cups.Value = i
		cupsIndex[i-1] = cups
		maxNum = i
		cups = cups.Next()
	}

	playGame(cups, iterations, cupsIndex, maxNum)

	returnRing := cupsIndex[0]	
	a := returnRing.Next()
	b := a.Next()

	return a.Value.(int) * b.Value.(int)
}

func playGame(cups *ring.Ring, iterations int, cupsIndex []*ring.Ring, max int) *ring.Ring {
	for i:=0; i<iterations; i++ {
		pickedUpCups := cups.Unlink(3)
		destination := cups.Value.(int) - 1
		
		if destination == 0 {
			destination = ringMax(cups)
		}

		for ringContains(pickedUpCups, destination) {
			destination--
			if destination == 0 {
				destination = max
			}
		}

		destinationRing := cupsIndex[destination-1]
		destinationRing.Link(pickedUpCups)

		cups = cups.Next()
	}

	return cups
}

func printRing(text string, r *ring.Ring) {
	fmt.Printf("%s: ", text)
	r.Do(func(x interface{}) {
		fmt.Printf("%v ", x)
	})
	fmt.Printf("\n")
}

func ringContains(data *ring.Ring, find int) bool {
	current := data
	for i := 0; i < current.Len(); i++ {
		if current.Value.(int) == find {
			return true
		}
		current = current.Next()
	}
	return false
}

func ringMax(r *ring.Ring) int {
	v := 0
	r.Do(func(i interface{}) {
		num, ok := i.(int)
		if ok {
			if(num > v) {
				v = num
			}
		}
	})
	return v
}

func findRingFromValue(r *ring.Ring, find int) *ring.Ring {
	
	current := r
	for i := 0; i < r.Len(); i++ {
		if current.Value.(int) == find {
			return current
		}
		current = current.Next()
	}
	return nil
}
