package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay14Part1(stringListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay14Part2(stringListToSlice(input)))
}

//SolveDay14Part1 returns the sum of all addresses after mask the values
func SolveDay14Part1(input []string) (sum int64) {
	var curMask string
	curValues := make(map[int64]int64)
	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			curMask = strings.TrimPrefix(line, "mask = ")
		} else {
			address, value := getMaskValues(line)
			curValues[address] = applyMaskOnValue(curMask, value)
		}
	}
	for _, value := range curValues {
		sum += value
	}
	return
}

//SolveDay14Part2 returns the sum of all addresses after mask the addresses
func SolveDay14Part2(input []string) (sum int64) {
	var curMask string
	curValues := make(map[int64]int64)
	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			curMask = strings.TrimPrefix(line, "mask = ")
		} else {
			address, value := getMaskValues(line)
			for _, add := range applyMaskOnAddress(curMask, address) {
				curValues[add] = value
			}
		}
	}
	for _, value := range curValues {
		sum += value
	}
	return
}

//getMaskValues returns the mem address and the value
func getMaskValues(input string) (address, value int64) {
	split := strings.Split(input, " = ")
	if len(split) != 2 {
		return 0, 0
	}
	addressInt, err := strconv.Atoi(strings.Trim(split[0], "mem[]"))
	if err != nil {
		return 0, 0
	}
	valueInt, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0
	}
	return int64(addressInt), int64(valueInt)
}

//applyMaskOnValue returns the new value of the given value after apply the mask
func applyMaskOnValue(mask string, value int64) (newValue int64) {
	if len(mask) != 36 {
		return 0
	}
	curBit := int64(34359738368)
	for _, maskChar := range mask {
		newValue <<= 1
		switch maskChar {
		// 1
		case 49:
			newValue++
		// 0
		case 48:
		// X
		case 88:
			if (value & curBit) != 0 {
				newValue++
			}
		default:
			return 0
		}
		curBit /= 2
	}
	return newValue
}

//applyMaskOnAddress apply the mask to the address and returns a list of new addresses
func applyMaskOnAddress(mask string, address int64) (addressesInt []int64) {
	if len(mask) != 36 {
		return nil
	}
	//apply mask to address
	var maskedAddress string
	curBit := int64(34359738368)
	for _, maskChar := range mask {
		switch maskChar {
		// 1 == 49
		case 49:
			maskedAddress = maskedAddress + "1"
		// X == 88
		case 88:
			maskedAddress = maskedAddress + "X"
		// 0 == 48
		case 48:
			if (address & curBit) == 0 {
				maskedAddress = maskedAddress + "0"
			} else {
				maskedAddress = maskedAddress + "1"
			}
		default:
			return nil
		}
		curBit /= 2
	}
	//generate address list by inserting 0 and 1 for each X
	addressList := []string{maskedAddress}
	for i, bit := range mask {
		// 48 == 0 / 49 == 1
		if bit == 48 || bit == 49 {
			continue
		}
		var tempAddressList []string
		for _, runMask := range addressList {
			tempAddressList = append(tempAddressList, runMask[:i]+"0"+runMask[i+1:], runMask[:i]+"1"+runMask[i+1:])
		}
		addressList = tempAddressList
	}
	//calculate the numbers of the addresses
	for _, address := range addressList {
		tmpAddressInt, err := strconv.ParseInt(address, 2, 64)
		if err != nil {
			return nil
		}
		addressesInt = append(addressesInt, tmpAddressInt)
	}
	return
}

//Helper functions
//stringListToSlice converts the list of strings (each string one row) to a slice
func stringListToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}
