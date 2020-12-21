package main

import (
	"fmt"
	"sort"
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

type foodListItem struct {
	ingredients []string
	allergens []string
}

func a(lines []string) int {
	foodList := []foodListItem{}
	ingredients := map[string]string{}
	allergens := []string{}

	for _, line := range lines {
		food := foodListItem{}

		l := strings.Split(line, " (contains ")

		for _, i := range strings.Split(l[0], " ") {
			ingredients[i] = ""
			food.ingredients = append(food.ingredients, i)
		}

		for _, a := range strings.Split(l[1], ", ") {
			allergens = append(allergens, strings.TrimRight(a, ")"))
			food.allergens = append(food.allergens, strings.TrimRight(a, ")"))
		}

		foodList = append(foodList, food)
	}

	allergens = removeDuplicateValues(allergens)

	for {
		allergen := allergens[0]
		possibleIngredients := []string{}

		for _, food := range foodList {

			if contains(food.allergens, allergen) {

					newPossibleIngredients := []string{}
					for _, i := range food.ingredients {
						if (len(possibleIngredients) == 0 || contains(possibleIngredients, i)) && ingredients[i] == "" {
							newPossibleIngredients = append(newPossibleIngredients, i)
						}
					}
					possibleIngredients = newPossibleIngredients
				}

		}

		copy(allergens[0:], allergens[1:]) 
		allergens = allergens[:len(allergens)-1]
		if len(possibleIngredients) == 1 {
			ingredients[possibleIngredients[0]] = allergen
			
		} else {
			allergens = append(allergens, allergen)
		}

		if len(allergens) == 0 {
			break;
		}
		
	}


	sum := 0

	for i, v := range ingredients {
		if v == "" {
			for _, f := range foodList {
				if contains(f.ingredients, i) {
					sum++
				}
			}
		}
	}


	


	return sum
}

func b(lines []string) string {

	foodList := []foodListItem{}
	ingredients := map[string]string{}
	allergens := []string{}

	for _, line := range lines {
		food := foodListItem{}

		l := strings.Split(line, " (contains ")

		for _, i := range strings.Split(l[0], " ") {
			ingredients[i] = ""
			food.ingredients = append(food.ingredients, i)
		}

		for _, a := range strings.Split(l[1], ", ") {
			allergens = append(allergens, strings.TrimRight(a, ")"))
			food.allergens = append(food.allergens, strings.TrimRight(a, ")"))
		}

		foodList = append(foodList, food)
	}

	allergens = removeDuplicateValues(allergens)

	for {
		allergen := allergens[0]
		possibleIngredients := []string{}

		for _, food := range foodList {

			if contains(food.allergens, allergen) {

					newPossibleIngredients := []string{}
					for _, i := range food.ingredients {
						if (len(possibleIngredients) == 0 || contains(possibleIngredients, i)) && ingredients[i] == "" {
							newPossibleIngredients = append(newPossibleIngredients, i)
						}
					}
					possibleIngredients = newPossibleIngredients
				}

		}

		copy(allergens[0:], allergens[1:]) 
		allergens = allergens[:len(allergens)-1]
		if len(possibleIngredients) == 1 {
			ingredients[possibleIngredients[0]] = allergen
			
		} else {
			allergens = append(allergens, allergen)
		}

		if len(allergens) == 0 {
			break;
		}
		
	}

	newIngredientsList := map[string]string {}
	keys := make([]string, 0, len(newIngredientsList))

	for k, v := range ingredients {
		if v != "" {
			newIngredientsList[v] = k
			keys = append(keys, v)
		}
	}

	sort.Strings(keys)

	ret := ""
	for _, k := range keys {
		ret += fmt.Sprintf("%s,", newIngredientsList[k])
	}

	return strings.TrimRight(ret, ",")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func removeDuplicateValues(strSlice []string) []string { 
	keys := make(map[string]bool) 
	list := []string{} 

	// If the key(values of the slice) is not equal 
	// to the already present value in new slice (list) 
	// then we append it. else we jump on another element. 
	for _, entry := range strSlice { 
			if _, value := keys[entry]; !value { 
					keys[entry] = true
					list = append(list, entry) 
			} 
	} 
	return list 
} 