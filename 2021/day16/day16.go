package main

import (
	"encoding/hex"
	"fmt"
	"github.com/dergeberl/aoc/utils"
	"os"
	"strconv"
	"strings"
)

type bits []int

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay16Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay16Part2(string(input)))
}

//SolveDay16Part1 returns the sum of the versions of all packets
func SolveDay16Part1(input string) int {
	b := parseInput(input)
	versions, _, _ := b.parsePacket(0)
	return versions
}

//SolveDay16Part2 returns the calculated number for all packets
func SolveDay16Part2(input string) int {
	b := parseInput(input)
	_, num, _ := b.parsePacket(0)
	return num
}

//parsePacket parses a packet with the given start position and returns the version, calculated number and current position
func (p bits) parsePacket(position int) (int, int, int) {
	version := p[position : position+3].toInt()
	mode := p[position+3 : position+6].toInt()
	if mode == 4 {
		position += 6
		var num bits
		for {
			num = append(num, p[position+1:position+5]...)
			if p[position] == 0 {
				position += 5
				break
			}
			position += 5
		}
		return version, num.toInt(), position
	}
	if p[position+6] == 0 {
		return p.parsePacketOperatorSize(position, version, mode)

	}
	if p[position+6] == 1 {
		return p.packetOperatorCount(position, version, mode)
	}
	return version, 0, position
}

//packetOperatorSize parse operator packets with the length type bit site
func (p bits) parsePacketOperatorSize(position int, version int, mode int) (int, int, int) {
	var nums []int
	numberOfBits := p[position+7 : position+22].toInt()
	position += 22
	endpos := numberOfBits + position
	for position < endpos {
		tempVersion := 0
		num := 0
		tempVersion, num, position = p.parsePacket(position)
		nums = append(nums, num)
		version += tempVersion
	}
	return version, calcNumbersByMode(mode, nums), position
}

//packetOperatorCount parse operator packets with the length type packet count
func (p bits) packetOperatorCount(position int, version int, mode int) (int, int, int) {
	var nums []int
	numberOfPackets := p[position+7 : position+18].toInt()
	position += 18
	for i := 0; i < numberOfPackets; i++ {
		tempVersion := 0
		num := 0
		tempVersion, num, position = p.parsePacket(position)
		nums = append(nums, num)
		version += tempVersion
	}
	return version, calcNumbersByMode(mode, nums), position
}

//calcNumbersByMode returns the calculated number for an operator mode and the given numbers
func calcNumbersByMode(mode int, numbers []int) int {
	switch mode {
	case 0:
		return utils.SumSlice(numbers)
	case 1:
		return utils.ProductSlice(numbers)
	case 2:
		return utils.MinimumSlice(numbers)
	case 3:
		return utils.MaximumSlice(numbers)
	case 5:
		if len(numbers) != 2 {
			panic("wrong input")
		}
		if numbers[0] > numbers[1] {
			return 1
		}
		return 0
	case 6:
		if len(numbers) != 2 {
			panic("wrong input")
		}
		if numbers[0] < numbers[1] {
			return 1
		}
		return 0
	case 7:
		if len(numbers) != 2 {
			panic("wrong input")
		}
		if numbers[0] == numbers[1] {
			return 1
		}
		return 0
	}
	return 0
}

//toInt returns the int for the bits
func (p bits) toInt() int {
	var s int
	for i := range p {
		s = s << 1
		s += p[i]
	}
	return s
}

func parseInput(input string) bits {
	input = strings.TrimSuffix(input, "\n")
	inputHex, _ := hex.DecodeString(input)
	var inputBits string
	for i := range inputHex {
		inputBits += fmt.Sprintf("%08v", strconv.FormatInt(int64(inputHex[i]), 2))
	}
	var b bits
	for _, r := range inputBits {
		b = append(b, int(r)-48)
	}
	return b
}
