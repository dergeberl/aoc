package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type picture struct {
	tiles           []tile
	tilesOrder      []int
	possiblePicture []string
	size            int
}

type tile struct {
	id        int
	position  int
	rawTile   []string
	cleanTile []string
	outlines  outlines
}

type outlines struct {
	a []string
	b []string
	c []string
	d []string
}

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay20Part1(input))
	fmt.Printf("Part 2: %v\n", SolveDay20Part2(input))
}

//SolveDay20Part1 multiply all tile IDs from the edges and return the number
func SolveDay20Part1(input string) (s int) {
	var pic picture
	pic.init(input)
	pic.getTilesOrder()
	edges := pic.getEdges()
	s = 1
	for _, edgeId := range edges {
		s *= edgeId
	}
	return
}

//SolveDay20Part2 returns the number of waves (count of '#' except the '#' from a see monster)
func SolveDay20Part2(input string) (s int) {
	var pic picture
	pic.init(input)
	pic.getTilesOrder()
	pic.reorderTiles()

	for i, _ := range pic.tiles {
		pic.tiles[i].rotateCleanTile()
	}

	pic.generateAllPossiblePictures()

	var regex []*regexp.Regexp
	for i := 0; i < len(pic.tiles[0].cleanTile)*pic.size-19; i++ {
		r := regexp.MustCompile(fmt.Sprintf("(?m)((^.{%v}.{18}[#].*)\\n(.{%v}[#].{4}[#][#].{4}[#][#].{4}[#][#][#].*)\\n(.{%v}.{1}[#].{2}[#].{2}[#].{2}[#].{2}[#].{2}[#].*))", i, i, i))
		regex = append(regex, r)
	}

	var numberOfSeaMonsters int
	for _, pics := range pic.possiblePicture {
		for _, r := range regex {
			found := r.FindAllString(pics, -1)
			numberOfSeaMonsters += len(found)
		}
		if numberOfSeaMonsters != 0 {
			break
		}
	}

	var waveCount int
	for _, line := range strings.Split(pic.possiblePicture[0], "\n") {
		for _, char := range line {
			if char == '#' {
				waveCount++
			}
		}
	}
	waveCount -= numberOfSeaMonsters * 15
	return waveCount
}

//init generates the tiles in the picture form the input
func (p *picture) init(input string) {
	var tiles []tile
	for _, givenTiles := range strings.Split(input, "\n\n") {
		splitTile := strings.Split(givenTiles, "\n")
		var tileId int
		if strings.HasPrefix(splitTile[0], "Tile") {
			tileId, _ = strconv.Atoi(strings.Trim(splitTile[0], "Tile :"))
		}
		tiles = append(tiles, tile{
			id:      tileId,
			rawTile: splitTile[1:],
		})
	}
	if len(tiles)%int(math.Sqrt(float64(len(tiles)))) == 0 {
		p.tiles = tiles
		p.size = int(math.Sqrt(float64(len(tiles))))
	}
}

//getTilesOrder search for the right tiles order
func (p *picture) getTilesOrder() {
	tiles := p.tiles
	var tilesOrder []int
	for i := 0; i <= 7; i++ {
		for id, firstTile := range tiles {
			tiles[id].toPosition(i)
			tilesOrder = []int{}
			tilesOrder = append(tilesOrder, firstTile.id)
			for iLine := 0; iLine < p.size; iLine++ {
				for iColumn := 0; iColumn < p.size; iColumn++ {
					if iLine == 0 && iColumn == 0 {
						continue
					}
					if (p.size*iLine)+iColumn != len(tilesOrder) {
						break
					}
					for i, curTile := range tiles {
						tileNew := true
						for _, usedTile := range tilesOrder {
							if usedTile == curTile.id {
								tileNew = false
								break
							}
						}
						if !tileNew {
							continue
						}
						if iLine == 0 {
							//first line (check only left tile)
							lastTile := getTileById(tilesOrder[len(tilesOrder)-1], tiles)
							found := false
							for i := 0; i <= 7; i++ {
								curTile.toPosition(i)
								if compareLines(curTile.outlines.d, lastTile.outlines.b) {
									tilesOrder = append(tilesOrder, curTile.id)
									found = true
									break
								}
							}
							if found {
								tiles[i] = curTile
								break
							}
							curTile.toPosition(0)
						} else if iColumn == 0 {
							// not first line but first column (check only tile above)
							upperTile := getTileById(tilesOrder[len(tilesOrder)-p.size], tiles)
							found := false
							for i := 0; i <= 7; i++ {
								curTile.toPosition(i)
								if compareLines(curTile.outlines.a, upperTile.outlines.c) {
									tilesOrder = append(tilesOrder, curTile.id)
									found = true
									break
								}
							}
							if found {
								tiles[i] = curTile
								break
							}
							curTile.toPosition(0)
						} else {
							// not first line and not first column (check upper and left tile)
							upperTile := getTileById(tilesOrder[len(tilesOrder)-p.size], tiles)
							lastTile := getTileById(tilesOrder[len(tilesOrder)-1], tiles)
							found := false

							for i := 0; i <= 7; i++ {
								curTile.toPosition(i)
								if compareLines(curTile.outlines.a, upperTile.outlines.c) && compareLines(curTile.outlines.d, lastTile.outlines.b) {
									tilesOrder = append(tilesOrder, curTile.id)
									found = true
									break
								}
							}
							if found {
								tiles[i] = curTile
								break
							}
							curTile.generateOutlines()
						}

					}
				}

			}
			if len(tilesOrder) == len(tiles) {
				break

			}

		}
		if len(tilesOrder) == len(tiles) {
			break

		}
	}
	p.tilesOrder = tilesOrder
}

//getEdges returns the tiles IDs from the 4 edges of the picture
func (p *picture) getEdges() []int {
	var edges []int
	edges = append(edges, p.tilesOrder[0])
	edges = append(edges, p.tilesOrder[p.size-1])
	edges = append(edges, p.tilesOrder[len(p.tilesOrder)-p.size])
	edges = append(edges, p.tilesOrder[len(p.tilesOrder)-1])
	return edges
}

//reorderTiles reorder the tiles to the found right order
func (p *picture) reorderTiles() {
	var newOrder []tile
	newOrder = make([]tile, len(p.tiles))

	if len(p.tilesOrder) != 0 {
		for i, id := range p.tilesOrder {
			for _, curTile := range p.tiles {
				if curTile.id == id {
					newOrder[i] = curTile
				}
			}
		}
	}
	p.tiles = make([]tile, len(newOrder))
	p.tiles = newOrder
}

//generateAllPossiblePictures generates the overall picture in all orientations
func (p *picture) generateAllPossiblePictures() {
	var completePicture, allPossiblePicture []string
	for iLines := 0; iLines < len(p.tiles); iLines = iLines + p.size {
		for iTiles := iLines; iTiles < iLines+p.size; iTiles++ {
			for tileLineI, tileLines := range p.tiles[iTiles].cleanTile {
				if iTiles%p.size == 0 {
					completePicture = append(completePicture, tileLines)
				} else {
					completePicture[tileLineI+(len(tileLines)*(iLines/p.size))] = completePicture[tileLineI+(len(tileLines)*(iLines/p.size))] + tileLines
				}
			}
		}
	}
	for i := 0; i <= 7; i++ {
		curPicture := positionTile(i, completePicture)
		var curPictureString string
		for _, line := range curPicture {
			curPictureString = curPictureString + line + "\n"
		}
		allPossiblePicture = append(allPossiblePicture, strings.TrimSuffix(curPictureString, "\n"))
	}
	p.possiblePicture = allPossiblePicture
}

//toPosition reorder the outlines to a set position (0 r0, 1 r90, 2 r180, 3 r270, 4 r0 mirror, 5 r90 mirror, 6 r180 mirror, 7 r270 mirror)
func (tile *tile) toPosition(pos int) {
	tile.generateOutlines()
	switch pos {
	default:
		break
	case 1, 2, 3:
		tile.rotateOutlines(pos, false)
	case 4, 5, 6, 7:
		tile.rotateOutlines(pos-4, true)
	}
	tile.position = pos
}

//generateOutlines generate the outlines from the rawTile
func (tile *tile) generateOutlines() {
	originalTile := tile.rawTile
	//a
	tile.outlines.a = strings.Split(originalTile[0], "")

	//c
	var cLine string
	for _, char := range originalTile[len(originalTile)-1] {
		cLine = string(char) + cLine
	}
	tile.outlines.c = strings.Split(cLine, "")

	// b && d
	var dLine string
	tile.outlines.b = []string{}
	for _, line := range originalTile {
		tile.outlines.b = append(tile.outlines.b, string(line[len(line)-1]))
		dLine = string(line[0]) + dLine
	}
	tile.outlines.d = strings.Split(dLine, "")
	tile.position = 0
}

//generateCleanTile remove the border from rawTile and save it to cleanTile
func (tile *tile) generateCleanTile() {
	var cleanTile []string
	for _, line := range tile.rawTile[1 : len(tile.rawTile)-1] {
		cleanTile = append(cleanTile, line[1:len(line)-1])
	}
	tile.cleanTile = cleanTile
}

//rotateCleanTile generates cleanTile and rotate it to  the right position
func (tile *tile) rotateCleanTile() {
	tile.generateCleanTile()
	var newTile []string
	newTile = positionTile(tile.position, tile.cleanTile)
	tile.cleanTile = make([]string, len(newTile))
	tile.cleanTile = newTile
}

//rotateOutlines rotate and/or mirror the outlines (rotate define the number of 90° turns)
func (tile *tile) rotateOutlines(rotate int, mirrored bool) {
	a := make([]string, len(tile.outlines.a))
	b := make([]string, len(tile.outlines.b))
	c := make([]string, len(tile.outlines.c))
	d := make([]string, len(tile.outlines.d))
	aTmp := make([]string, len(tile.outlines.a))
	bTmp := make([]string, len(tile.outlines.b))
	cTmp := make([]string, len(tile.outlines.c))
	dTmp := make([]string, len(tile.outlines.d))

	for i, val := range tile.outlines.a {
		a[i] = val
		aTmp[i] = val
	}
	for i, val := range tile.outlines.b {
		b[i] = val
		bTmp[i] = val
	}
	for i, val := range tile.outlines.c {
		c[i] = val
		cTmp[i] = val
	}
	for i, val := range tile.outlines.d {
		d[i] = val
		dTmp[i] = val
	}

	if mirrored {
		aTmp = make([]string, len(a))
		bTmp = make([]string, len(b))
		cTmp = make([]string, len(c))
		dTmp = make([]string, len(d))
		for i, b2 := range b {
			bTmp[len(b)-1-i] = b2
		}
		for i, d2 := range d {
			dTmp[len(d)-1-i] = d2
		}

		for i, a2 := range a {
			cTmp[len(c)-1-i] = a2
		}
		for i, c2 := range c {
			aTmp[len(a)-1-i] = c2
		}

		for i, val := range aTmp {
			a[i] = val
		}
		for i, val := range bTmp {
			b[i] = val
		}
		for i, val := range cTmp {
			c[i] = val
		}
		for i, val := range dTmp {
			d[i] = val
		}
	}

	for i := 0; i < rotate; i++ {
		for i, val := range aTmp {
			b[i] = val
		}
		for i, val := range bTmp {
			c[i] = val
		}
		for i, val := range cTmp {
			d[i] = val
		}
		for i, val := range dTmp {
			a[i] = val
		}
		aTmp = make([]string, len(tile.outlines.a))
		bTmp = make([]string, len(tile.outlines.b))
		cTmp = make([]string, len(tile.outlines.c))
		dTmp = make([]string, len(tile.outlines.d))
		for i, val := range a {
			aTmp[i] = val
		}
		for i, val := range b {
			bTmp[i] = val
		}
		for i, val := range c {
			cTmp[i] = val
		}
		for i, val := range d {
			dTmp[i] = val
		}
	}
	tile.outlines.a, tile.outlines.b, tile.outlines.c, tile.outlines.d = make([]string, len(a)), make([]string, len(a)), make([]string, len(a)), make([]string, len(a))
	tile.outlines.a, tile.outlines.b, tile.outlines.c, tile.outlines.d = a, b, c, d
}

//positionTile rotate and/or mirror to a set position (0 r0, 1 r90, 2 r180, 3 r270, 4 r0 mirror, 5 r90 mirror, 6 r180 mirror, 7 r270 mirror)
func positionTile(pos int, input []string) []string {
	var rotate int
	var mirrored bool
	switch pos {
	case 0:
		return input
	case 1, 2, 3:
		rotate = pos
		mirrored = false
	case 4, 5, 6, 7:
		rotate = pos
		mirrored = true
	}

	tmpTile := make([]string, len(input))
	cleanTile := make([]string, len(input))

	for i, val := range input {
		cleanTile[i] = val
	}

	if mirrored {
		tmpTile = mirrorStringSlice(cleanTile)
		cleanTile = make([]string, len(tmpTile))
		for i, val := range tmpTile {
			cleanTile[i] = val
		}
	}

	for i := 0; i < rotate; i++ {
		tmpTile = make([]string, len(input))
		tmpTile = rotateStringSlice(cleanTile)
		cleanTile = make([]string, len(tmpTile))
		for i, val := range tmpTile {
			cleanTile[i] = val
		}
	}

	return cleanTile
}

//mirrorStringSlice mirrors the given string slice from top to bottom
func mirrorStringSlice(input []string) []string {
	tmp := make([]string, len(input))
	for i, line := range input {
		tmp[len(input)-1-i] = line
	}
	return tmp
}

//rotateStringSlice rotate a given string slice to +90°
func rotateStringSlice(input []string) []string {
	tmp := make([]string, len(input))
	for _, line := range input {
		for charI, char := range line {
			tmp[charI] = string(char) + tmp[charI]
		}
	}
	return tmp
}

//compareLines compare two lines (the lines are reversed)
func compareLines(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, char := range a {
		if char != b[len(b)-1-i] {
			return false
		}
	}
	return true
}

//getTileById search for a tile in a slice of tiles by a tileID
func getTileById(id int, tiles []tile) tile {
	for _, t := range tiles {
		if t.id == id {
			return t
		}
	}
	return tile{}
}
