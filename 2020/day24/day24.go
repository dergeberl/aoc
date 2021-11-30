package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type coordinate struct {
	x int
	y int
	z int
}

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay24Part1(stringListToSlice(input)))
	fmt.Printf("Part 2: %v\n", SolveDay24Part2(stringListToSlice(input)))
}

//SolveDay24Part1 count the black tiles after the flips from the input
func SolveDay24Part1(input []string) (s int) {
	tiles := parseInput(input)

	for _, t := range tiles {
		if t {
			s++
		}
	}
	return
}

//SolveDay24Part2 count the black tiles after the flips with the rules after 100 days
func SolveDay24Part2(input []string) (s int) {
	tiles := parseInput(input)
	for i := 0; i < 100; i++ {
		makeBigger(tiles)
		newTiles := flipTiles(tiles)
		tiles = make(map[coordinate]bool, len(newTiles))
		for key, value := range newTiles {
			tiles[key] = value
		}
	}

	for _, t := range tiles {
		if t {
			s++
		}
	}
	return
}

//parseInput decodes the input list in a map
func parseInput(input []string) map[coordinate]bool {
	tiles := make(map[coordinate]bool)
	for _, line := range input {
		var x, y, z int
		for len(line) > 0 {
			if line[:1] == "e" {
				x++
				y--
				line = line[1:]
			} else if line[:1] == "w" {
				x--
				y++
				line = line[1:]
			} else if line[:2] == "se" {
				y--
				z++
				line = line[2:]
			} else if line[:2] == "sw" {
				z++
				x--
				line = line[2:]
			} else if line[:2] == "nw" {
				y++
				z--
				line = line[2:]
			} else if line[:2] == "ne" {
				z--
				x++
				line = line[2:]
			}
		}
		tiles[coordinate{x: x, y: y, z: z}] = !tiles[coordinate{x: x, y: y, z: z}]
	}

	return tiles
}

//flipTiles flip all tiles to the right color
func flipTiles(tiles map[coordinate]bool) map[coordinate]bool {
	tilesNew := make(map[coordinate]bool)
	for c, t := range tiles {
		var count int
		if tiles[coordinate{x: c.x + 1, y: c.y + 0, z: c.z - 1}] {
			count++
		}
		if tiles[coordinate{x: c.x + 1, y: c.y - 1, z: c.z + 0}] {
			count++
		}
		if tiles[coordinate{x: c.x + 0, y: c.y - 1, z: c.z + 1}] {
			count++
		}
		if tiles[coordinate{x: c.x - 1, y: c.y - 0, z: c.z + 1}] {
			count++
		}
		if tiles[coordinate{x: c.x - 1, y: c.y + 1, z: c.z + 0}] {
			count++
		}
		if tiles[coordinate{x: c.x + 0, y: c.y + 1, z: c.z - 1}] {
			count++
		}
		if t {
			if count == 0 || count > 2 {
				tilesNew[c] = false
			} else {
				tilesNew[c] = true
			}
		} else {
			if count == 2 {
				tilesNew[c] = true
			} else {
				tilesNew[c] = false
			}
		}
	}
	return tilesNew
}

//makeBigger increase the map around a black color (for the next day)
func makeBigger(tiles map[coordinate]bool) {
	for c, t := range tiles {
		if !t {
			continue
		}
		tiles[coordinate{x: c.x + 1, y: c.y + 0, z: c.z - 1}] = tiles[coordinate{x: c.x + 1, y: c.y + 0, z: c.z - 1}]
		tiles[coordinate{x: c.x + 1, y: c.y - 1, z: c.z + 0}] = tiles[coordinate{x: c.x + 1, y: c.y - 1, z: c.z + 0}]
		tiles[coordinate{x: c.x + 0, y: c.y - 1, z: c.z + 1}] = tiles[coordinate{x: c.x + 0, y: c.y - 1, z: c.z + 1}]
		tiles[coordinate{x: c.x - 1, y: c.y - 0, z: c.z + 1}] = tiles[coordinate{x: c.x - 1, y: c.y - 0, z: c.z + 1}]
		tiles[coordinate{x: c.x - 1, y: c.y + 1, z: c.z + 0}] = tiles[coordinate{x: c.x - 1, y: c.y + 1, z: c.z + 0}]
		tiles[coordinate{x: c.x + 0, y: c.y + 1, z: c.z - 1}] = tiles[coordinate{x: c.x + 0, y: c.y + 1, z: c.z - 1}]
	}
}

//Helper functions
//stringListToSlice converts the list of strings (each string one row) to a slice
func stringListToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}
