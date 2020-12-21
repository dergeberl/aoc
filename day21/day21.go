package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay21Part1(stringListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay21Part2(stringListToSlice(input)))
}

//SolveDay21Part1 counts the listed ingredient that are not a allergen
func SolveDay21Part1(input []string) (s int) {
	allergenTranslate, ingredientsList := readIngredientList(input)

	ingredientCache := make(map[string]bool)
	//count each non allergen ingredient
	for _, ingredient := range ingredientsList {
		if _, ok := ingredientCache[ingredient]; ok {
			if !ingredientCache[ingredient] {
				s++
			}
			continue
		}
		var found bool
		for _, allergen := range allergenTranslate {
			if ingredient == allergen {
				found = true
				ingredientCache[ingredient] = true
				break
			}
		}
		if !found {
			ingredientCache[ingredient] = false
			s++
		}
	}
	return s
}

//SolveDay21Part2 returns comma separated list of the allergens (sort alphabetically by the english name)
func SolveDay21Part2(input []string) (list string) {
	allergenTranslate, _ := readIngredientList(input)

	var keys []string
	for k := range allergenTranslate {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// generate list
	for _, key := range keys {
		list += allergenTranslate[key]
		list += ","
	}
	list = strings.TrimSuffix(list, ",")
	return
}

//readIngredientList reads the input an returns a dict of allergen [english]otherLang and a list of all ingredients
func readIngredientList(input []string) (allergenDictionary map[string]string, ingredientsList []string) {
	var allergenTranslateCandidates map[string][]string

	//read input
	ingredientsPerAllergen := make(map[string][]string)
	for _, line := range input {
		splitLine := strings.Split(line, " (contains ")

		for _, ingredient := range strings.Split(splitLine[0], " ") {
			ingredientsList = append(ingredientsList, ingredient)
		}

		allergensList := strings.Split(strings.Trim(splitLine[1], ")"), ", ")
		for _, allergen := range allergensList {
			ingredientsPerAllergen[allergen] = append(ingredientsPerAllergen[allergen], splitLine[0])
		}
	}

	//find the possible translation candidates for each allergen
	allergenTranslateCandidates = make(map[string][]string)
	for allergenEnglish, allergen := range ingredientsPerAllergen {
		tmp := make(map[string]int)
		for _, ingredientList := range allergen {
			for _, ingredient := range strings.Split(ingredientList, " ") {
				tmp[ingredient]++
			}
		}
		for ingredient, count := range tmp {
			if count == len(allergen) {
				allergenTranslateCandidates[allergenEnglish] = append(allergenTranslateCandidates[allergenEnglish], ingredient)
			}
		}
	}

	// resolve allergen translation
	allergenDictionary = make(map[string]string)
	allergenSize := len(allergenTranslateCandidates)
	for {
		for english, otherLang := range allergenTranslateCandidates {
			if len(otherLang) == 1 {
				allergenDictionary[english] = otherLang[0]
			}
		}
		if len(allergenDictionary) == allergenSize {
			break
		}
		newAllergenTranslateCandidates := make(map[string][]string)
		for english, otherLangList := range allergenTranslateCandidates {
			for _, otherLang := range otherLangList {
				found := false
				for _, alreadyFound := range allergenDictionary {
					if otherLang == alreadyFound {
						found = true
						break
					}
				}
				if !found {
					newAllergenTranslateCandidates[english] = append(newAllergenTranslateCandidates[english], otherLang)
				}
			}
		}
		allergenTranslateCandidates = make(map[string][]string)
		for i, val := range newAllergenTranslateCandidates {
			allergenTranslateCandidates[i] = val
		}
	}

	return allergenDictionary, ingredientsList
}

//Helper functions
//stringListToSlice converts the list of strings (each string one row) to a slice
func stringListToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}

//intListToSlice converts the list of numbers (each number one row) to a slice
func intListToSlice(list string) (i []int) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			return nil
		}
		i = append(i, lineInt)
	}
	return
}
