package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// checkRequired check if all required fields of a password are set
func checkRequired(m map[string]string) bool {
	return m["byr"] != "" && m["iyr"] != "" && m["eyr"] != "" && m["hgt"] != "" && m["hcl"] != "" && m["ecl"] != "" && m["pid"] != ""
}

// checkValue checks the password value
func checkValue(v string, p string) bool {
	validRegex := map[string]string{
		"byr": "^(19[2-9][0-9]|200[0-2])$",
		"iyr": "^(201[0-9]|2020)$",
		"eyr": "^(202[0-9]|2030)$",
		"hgt": "^(((59|6[0-9]|7[0-6])in)|((1[5-8][0-9]|19[0-3])cm))$",
		"hcl": "^#([0-9]|[a-f]){6}$",
		"ecl": "^(amb|blu|brn|gry|grn|hzl|oth)$",
		"pid": "^([0-9]){9}$",
		"cid": ".*",
	}
	if validRegex[p] == "" {
		return false
	}
	byr, err := regexp.MatchString(validRegex[p], v)
	return byr && err == nil
}

//SolveDay4Part1 returns number of passwords that have all required fields
func SolveDay4Part1(i string) (sum int) {
	for _, password := range strings.Split(i, "\n\n") {
		cur := make(map[string]string)
		for _, passwordLine := range strings.Split(password, "\n") {
			for _, passwordKeyValue := range strings.Split(passwordLine, " ") {
				value := strings.Split(passwordKeyValue, ":")
				cur[value[0]] = value[1]
			}
		}
		if checkRequired(cur){
			sum++
		}
	}
	return
}

//SolveDay4Part2 returns number of passwords that have all required fields and valid values
func SolveDay4Part2(i string) (sum int) {
	for _, password := range strings.Split(i, "\n\n") {
		invalid := false
		cur := make(map[string]string)
		for _, passwordLine := range strings.Split(password, "\n") {
			if invalid {
				break
			}
			for _, passwordKeyValue := range strings.Split(passwordLine, " ") {
				value := strings.Split(passwordKeyValue, ":")
				if !checkValue(value[1], value[0]){
					invalid = true
					break
				}
				cur[value[0]] = value[1]
			}
		}
		if !invalid && checkRequired(cur) {
			sum++
		}
	}
	return
}

func main() {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay4Part1(input))
	fmt.Printf("Part 2: %v\n", SolveDay4Part2(input))
}
