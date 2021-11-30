package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Refactor with struct
//password is a struct for the fields of the passwords
type password struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay4Part1(input))
	fmt.Printf("Part 2: %v\n", SolveDay4Part2(input))
	fmt.Printf("Part 1 refactor: %v\n", SolveDay4Part1r(input))
	fmt.Printf("Part 2 refactor: %v\n", SolveDay4Part2r(input))
}

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
		if checkRequired(cur) {
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
				if !checkValue(value[1], value[0]) {
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

//SolveDay4Part1r returns number of passwords that have all required fields
func SolveDay4Part1r(i string) (sum int) {
	return len(extractPassword(i))
}

//SolveDay4Part2r returns number of passwords that have all required fields and valid values
func SolveDay4Part2r(i string) (sum int) {
	return len(deleteInvalidPasswords(extractPassword(i)))
}

//extractPassword returns a password slice from a given list
func extractPassword(list string) (passwords []password) {
	for _, pw := range strings.Split(list, "\n\n") {
		currentPassword := password{}
		for _, passwordLine := range strings.Split(pw, "\n") {
			for _, passwordKeyValue := range strings.Split(passwordLine, " ") {
				value := strings.Split(passwordKeyValue, ":")
				switch value[0] {
				case "byr":
					byr, err := strconv.Atoi(value[1])
					if err != nil {
						continue
					}
					currentPassword.byr = byr
				case "iyr":
					iyr, err := strconv.Atoi(value[1])
					if err != nil {
						continue
					}
					currentPassword.iyr = iyr
				case "eyr":
					eyr, err := strconv.Atoi(value[1])
					if err != nil {
						continue
					}
					currentPassword.eyr = eyr
				case "hgt":
					currentPassword.hgt = value[1]
				case "hcl":
					currentPassword.hcl = value[1]
				case "ecl":
					currentPassword.ecl = value[1]
				case "pid":
					currentPassword.pid = value[1]
				case "cid":
					currentPassword.cid = value[1]
				}
			}
		}
		if checkPasswordRequiredFields(currentPassword) {
			passwords = append(passwords, currentPassword)
		}
	}
	return
}

//checkPasswordRequiredFields check a password if all required fields are given
func checkPasswordRequiredFields(pw password) bool {
	return pw.byr != 0 && pw.iyr != 0 && pw.eyr != 0 && pw.hgt != "" && pw.hcl != "" && pw.ecl != "" && pw.pid != ""
}

//deleteInvalidPasswords remove all invalid passwords from the slice
func deleteInvalidPasswords(passwords []password) (validPasswords []password) {
	for _, pw := range passwords {
		if pw.byr < 1920 || pw.byr > 2002 ||
			pw.iyr < 2010 || pw.iyr > 2020 ||
			pw.eyr < 2020 || pw.eyr > 2030 {
			continue
		}
		if strings.HasSuffix(pw.hgt, "in") {
			hgt, err := strconv.Atoi(strings.TrimSuffix(pw.hgt, "in"))
			if hgt < 59 || hgt > 76 || err != nil {
				continue
			}
		} else if strings.HasSuffix(pw.hgt, "cm") {
			hgt, err := strconv.Atoi(strings.TrimSuffix(pw.hgt, "cm"))
			if hgt < 150 || hgt > 193 || err != nil {
				continue
			}
		} else {
			continue
		}

		if len(pw.hcl) != 7 || !strings.HasPrefix(pw.hcl, "#") {
			continue
		}
		hcl := strings.TrimPrefix(pw.hcl, "#")
		hcl = strings.Trim(hcl, "abcdef1234567890")
		if hcl != "" {
			continue
		}
		if !(pw.ecl == "amb" ||
			pw.ecl == "blu" ||
			pw.ecl == "brn" ||
			pw.ecl == "gry" ||
			pw.ecl == "grn" ||
			pw.ecl == "hzl" ||
			pw.ecl == "oth") {
			continue
		}
		_, err := strconv.Atoi(pw.pid)
		if !(len(pw.pid) == 9 && err == nil) {
			continue
		}
		validPasswords = append(validPasswords, pw)
	}
	return
}
