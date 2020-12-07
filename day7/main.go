package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dubbe/advent-of-code-2020/helpers"
)


func main() {
	lines, err := helpers.ReadLines("input")
	helpers.Check(err)
	a(lines)
	b(lines)
}

func a(lines []string) int {
	//bag := "shiny gold"
	bagRules := map[string][]string{}
	re := regexp.MustCompile(`bags|bag|\.|(\d)`)
	for _, rule := range lines {
		rules := strings.Split(rule, "contain")
		parentBag := strings.TrimSpace(strings.ReplaceAll(rules[0], "bags" , ""))
		childBags := strings.Split(rules[1], ",")
		for i, childBag := range childBags {		
			childBags[i] = strings.TrimSpace(re.ReplaceAllString(childBag, ""))
		}
		bagRules[parentBag] = childBags
	}
	found := map[string]bool{}
	findBag(bagRules, "shiny gold", found)
	fmt.Printf("found %d bags \n", len(found))
	return len(found)
}

func findBag(bagRules map[string][]string, bagColor string, found map[string]bool) {
	for parentBag, childBags := range bagRules {
		for _, childBag := range childBags {
			if(childBag == bagColor) {
				found[parentBag] = true
				findBag(bagRules, parentBag, found)
			}
		} 
	}
}

func b(lines []string) int {
	bagRules := map[string][]map[string]int{}
	re := regexp.MustCompile(`bags|bag|\.|(\d)`)
	reAmount := regexp.MustCompile(`(\d)`)
	for _, rule := range lines {
		rules := strings.Split(rule, "contain")
		parentBag := strings.TrimSpace(strings.ReplaceAll(rules[0], "bags" , ""))
		childBags := strings.Split(rules[1], ",")
		childBagsRules := []map[string]int{}
		for _, childBag := range childBags {
			a := 0	
			amount := reAmount.FindStringSubmatch(childBag)
			if(len(amount) > 0) {
				a, _ = strconv.Atoi(amount[0])
			}
			childBagsRule := map[string]int{}
			childBagsRule[strings.TrimSpace(re.ReplaceAllString(childBag, ""))] = a
			childBagsRules = append(childBagsRules, childBagsRule)
		}
		bagRules[parentBag] = childBagsRules
	}
	count := 0
	countBags(bagRules, "shiny gold", &count)
	fmt.Printf("count %d bags \n", count)
	return count
}

func countBags(bagRules map[string][]map[string]int, bagColor string, count *int) {
	for parentBag, childBags := range bagRules {
		if(parentBag == bagColor) {
			for _, childBag := range childBags {
				for color, amount := range childBag {
					*count = *count + amount
					for i := 0; i < amount; i++ {
						countBags(bagRules, color, count)
					}
				}
			}
			
		}
	}
}
