package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type image map[point]*bool

type enhancementAlgorithm map[int]bool

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay20Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay20Part2(string(input)))
}

//SolveDay20Part1 returns the number of lit pixels after enhance 2 times
func SolveDay20Part1(input string) int {
	algorithm, img := parseInput(input)
	img = img.increase(3)
	img = img.enhance(algorithm)
	img = img.enhance(algorithm)
	img = img.decrease(1)
	return img.countLit()
}

//SolveDay20Part2 returns the number of lit pixels after enhance 50 times
func SolveDay20Part2(input string) int {
	algorithm, img := parseInput(input)
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			img = img.increase(4)
		}
		img = img.enhance(algorithm)
		img = img.decrease(1)
	}
	return img.countLit()
}

//enhance apply the enhancementAlgorithm to an image and returns the new image
func (i image) enhance(algorithm enhancementAlgorithm) image {
	newImage := make(image)
	for pixel := range i {
		x, y := pixel.x, pixel.y
		p := make([]bool, 9)
		p[0] = i.getPixel(point{x: x - 1, y: y - 1})
		p[1] = i.getPixel(point{x: x, y: y - 1})
		p[2] = i.getPixel(point{x: x + 1, y: y - 1})
		p[3] = i.getPixel(point{x: x - 1, y: y})
		p[4] = i.getPixel(point{x: x, y: y})
		p[5] = i.getPixel(point{x: x + 1, y: y})
		p[6] = i.getPixel(point{x: x - 1, y: y + 1})
		p[7] = i.getPixel(point{x: x, y: y + 1})
		p[8] = i.getPixel(point{x: x + 1, y: y + 1})
		var bits int
		for b := range p {
			bits = bits << 1
			if p[b] {
				bits++
			}
		}
		newPixel := algorithm[bits]
		newImage[pixel] = &newPixel
	}
	return newImage
}

//getPixel returns the status of a pixel (default false)
func (i image) getPixel(p point) bool {
	if i[p] != nil {
		return *i[p]
	}
	return false
}

//increase the image by x pixel on each site
func (i image) increase(by int) image {
	minX, minY, maxX, maxY := i.getSize()
	for y := minY - by; y <= maxY+by; y++ {
		for x := minX - by; x <= maxX+by; x++ {
			if i[point{x: x, y: y}] == nil {
				f := false
				i[point{x: x, y: y}] = &f
			}
		}
	}
	return i
}

//decrease the image by x pixel on each site
func (i image) decrease(by int) image {
	minX, minY, maxX, maxY := i.getSize()
	newImg := make(image)
	for y := minY + by; y <= maxY-by; y++ {
		for x := minX + by; x <= maxX-by; x++ {
			newImg[point{x: x, y: y}] = i[point{x: x, y: y}]
		}
	}
	return newImg
}

//getSize returns the size of the image (minX, minY, maxX, maxY)
func (i image) getSize() (int, int, int, int) {
	var maxX, maxY int
	minX := math.MaxInt
	minY := math.MaxInt
	for p := range i {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
	}
	return minX, minY, maxX, maxY
}

//print the image for debugging
func (i image) print() {
	var maxX, maxY int
	minX := math.MaxInt
	minY := math.MaxInt
	for p := range i {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
	}
	for y := minY; y <= maxY; y++ {

		for x := minX; x <= maxX; x++ {
			if i[point{x: x, y: y}] != nil && *i[point{x: x, y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

//countLit returns the number of lit pixels in the image
func (i image) countLit() int {
	var count int
	for p := range i {
		if i[p] != nil && *i[p] {
			count++
		}
	}
	return count
}

//parseInput returns the enhancementAlgorithm and image for an input
func parseInput(input string) (enhancementAlgorithm, image) {
	in := strings.Split(input, "\n\n")

	if len(in) != 2 {
		panic("wrong input")
	}
	enhancement := make(enhancementAlgorithm)
	for i, r := range in[0] {
		switch r {
		case '.':
			enhancement[i] = false
		case '#':
			enhancement[i] = true
		}
	}

	img := make(image)
	imgLines := strings.Split(in[1], "\n")
	for y := range imgLines {
		for x, r := range imgLines[y] {
			t := true
			f := false
			switch r {
			case '.':
				img[point{x: x, y: y}] = &f
			case '#':
				img[point{x: x, y: y}] = &t
			}
		}
	}
	return enhancement, img
}
