package main

import (
	"fmt"
	"github.com/dergeberl/aoc/utils"
	"os"
	"strconv"
	"strings"
)

type region struct {
	xMin, xMax int
	yMin, yMax int
	zMin, zMax int
	state      bool
}

type cubeRegions []region

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay22Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay22Part2(string(input)))
}

//SolveDay22Part1 returns number of on cubes within x=-50..50,y=-50..50,z=-50..50
func SolveDay22Part1(input string) int64 {
	c := parseInput(input)
	return c.countActive(countPart1)
}

//SolveDay22Part2 returns number of on cubes
func SolveDay22Part2(input string) int64 {
	c := parseInput(input)
	return c.countActive(countPart2)
}

//parseInput returns cubeRegions for a string input
func parseInput(input string) cubeRegions {
	lines, _ := utils.InputToSlice(input)
	c := make(cubeRegions, 0)
	for l := range lines {
		var state bool
		var line string
		if strings.HasPrefix(lines[l], "on") {
			state = true
			line = strings.TrimPrefix(lines[l], "on ")
		} else {
			line = strings.TrimPrefix(lines[l], "off ")
		}
		points := strings.Split(line, ",")
		var xMin, xMax int
		var yMin, yMax int
		var zMin, zMax int
		for i := range points {
			tmpPoint := strings.Split(points[i], "=")
			if len(tmpPoint) != 2 {
				panic("wrong input")
			}
			tempRange := strings.Split(tmpPoint[1], "..")
			if len(tempRange) != 2 {
				panic("wrong input")
			}
			minA, _ := strconv.Atoi(tempRange[0])
			maxA, _ := strconv.Atoi(tempRange[1])
			switch tmpPoint[0] {
			case "x":
				xMin, xMax = minA, maxA
			case "y":
				yMin, yMax = minA, maxA
			case "z":
				zMin, zMax = minA, maxA
			}
		}
		c = append(c, region{
			xMin:  xMin,
			xMax:  xMax,
			yMin:  yMin,
			yMax:  yMax,
			zMin:  zMin,
			zMax:  zMax,
			state: state,
		})
	}
	return c
}

//returns the number of active points after all rules, with a given count function
func (c cubeRegions) countActive(countFunc func(r region) int64) int64 {
	if c == nil {
		return 0
	}
	var regions cubeRegions

	for i := range c {
		var newRegions cubeRegions
		for cur := range regions {
			newRegions = append(newRegions, regions[cur].sub(c[i])...)
		}
		regions = newRegions
		if c[i].state {
			regions = append(regions, c[i])
		}
	}

	var sum int64
	for i := range regions {
		sum += countFunc(regions[i])
	}
	return sum
}

//sub returns a new cubeRegions without the region which is subtracted
func (r region) sub(sRegion region) cubeRegions {
	if sRegion.xMin > r.xMax ||
		sRegion.yMin > r.yMax ||
		sRegion.zMin > r.zMax ||
		sRegion.xMax < r.xMin ||
		sRegion.yMax < r.yMin ||
		sRegion.zMax < r.zMin {
		return cubeRegions{r}
	}
	if sRegion.xMin < r.xMin {
		sRegion.xMin = r.xMin
	}
	if sRegion.xMax > r.xMax {
		sRegion.xMax = r.xMax
	}
	if sRegion.yMin < r.yMin {
		sRegion.yMin = r.yMin
	}
	if sRegion.yMax > r.yMax {
		sRegion.yMax = r.yMax
	}
	if sRegion.zMin < r.zMin {
		sRegion.zMin = r.zMin
	}
	if sRegion.zMax > r.zMax {
		sRegion.zMax = r.zMax
	}
	newRegions := cubeRegions{r}
	sRegion.xMax++
	sRegion.yMax++
	sRegion.zMax++
	newRegions = newRegions.cutAll(&sRegion.xMin, nil, nil)
	newRegions = newRegions.cutAll(&sRegion.xMax, nil, nil)
	newRegions = newRegions.cutAll(nil, &sRegion.yMin, nil)
	newRegions = newRegions.cutAll(nil, &sRegion.yMax, nil)
	newRegions = newRegions.cutAll(nil, nil, &sRegion.zMin)
	newRegions = newRegions.cutAll(nil, nil, &sRegion.zMax)
	sRegion.xMax--
	sRegion.yMax--
	sRegion.zMax--
	for i := range newRegions {
		if newRegions[i].xMin == sRegion.xMin &&
			newRegions[i].xMax == sRegion.xMax &&
			newRegions[i].yMin == sRegion.yMin &&
			newRegions[i].yMax == sRegion.yMax &&
			newRegions[i].zMin == sRegion.zMin &&
			newRegions[i].zMax == sRegion.zMax {
			newRegions[i] = newRegions[len(newRegions)-1]
			newRegions = newRegions[:len(newRegions)-1]
			break
		}
	}
	return newRegions
}

//cutAll cuts all cubeRegions at point and returns new cubeRegions
func (r cubeRegions) cutAll(x, y, z *int) cubeRegions {
	var newRegions cubeRegions
	for i := range r {
		newRegions = append(newRegions, r[i].cut(x, y, z)...)
	}
	return newRegions
}

//cut one region at point and returns new cubeRegions
func (r region) cut(x, y, z *int) cubeRegions {
	newRegions := make(cubeRegions, 0)
	var tempRegion1 region
	var tempRegion2 region
	if x != nil {
		tempRegion1 = region{
			xMin:  r.xMin,
			xMax:  *x - 1,
			yMin:  r.yMin,
			yMax:  r.yMax,
			zMin:  r.zMin,
			zMax:  r.zMax,
			state: true,
		}
		tempRegion2 = region{
			xMin:  *x,
			xMax:  r.xMax,
			yMin:  r.yMin,
			yMax:  r.yMax,
			zMin:  r.zMin,
			zMax:  r.zMax,
			state: true,
		}
		newRegions = append(newRegions, tempRegion1, tempRegion2)
	}
	if y != nil {
		tempRegion1 = region{
			xMin:  r.xMin,
			xMax:  r.xMax,
			yMin:  r.yMin,
			yMax:  *y - 1,
			zMin:  r.zMin,
			zMax:  r.zMax,
			state: true,
		}
		tempRegion2 = region{
			xMin:  r.xMin,
			xMax:  r.xMax,
			yMin:  *y,
			yMax:  r.yMax,
			zMin:  r.zMin,
			zMax:  r.zMax,
			state: true,
		}
		newRegions = append(newRegions, tempRegion1, tempRegion2)

	}

	if z != nil {
		tempRegion1 = region{
			xMin:  r.xMin,
			xMax:  r.xMax,
			yMin:  r.yMin,
			yMax:  r.yMax,
			zMin:  r.zMin,
			zMax:  *z - 1,
			state: true,
		}
		tempRegion2 = region{
			xMin:  r.xMin,
			xMax:  r.xMax,
			yMin:  r.yMin,
			yMax:  r.yMax,
			zMin:  *z,
			zMax:  r.zMax,
			state: true,
		}
		newRegions = append(newRegions, tempRegion1, tempRegion2)

	}
	if countPart2(tempRegion1) <= 0 || countPart2(tempRegion2) <= 0 {
		return []region{r}
	}
	return newRegions
}

//countPart1 reruns number of blocks only if they are within 50 cubes away from 0
func countPart1(r region) int64 {
	if r.xMax > 50 {
		return 0
	}
	if r.yMax > 50 {
		return 0
	}
	if r.zMax > 50 {
		return 0
	}

	if r.xMin < -50 {
		return 0
	}
	if r.yMin < -50 {
		return 0
	}
	if r.zMin < -50 {
		return 0
	}
	return int64((r.xMax - r.xMin + 1) * (r.yMax - r.yMin + 1) * (r.zMax - r.zMin + 1))
}

//countPart2 reruns number if blocks for a region
func countPart2(r region) int64 {
	return int64((r.xMax - r.xMin + 1) * (r.yMax - r.yMin + 1) * (r.zMax - r.zMin + 1))
}
