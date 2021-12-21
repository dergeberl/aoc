package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dergeberl/aoc/utils"
)

type player struct {
	position int
	score    int
}

type game struct {
	player1       player
	player2       player
	currentPlayer int
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay21Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay21Part2(string(input)))
}

//SolveDay21Part1 returns the number of rolled die multiplied by loosing player score
//This game is played the normal way
func SolveDay21Part1(input string) int {
	g := parseInput(input)

	var curNum, cur3Nums, dieCount int
	for {
		dieCount += 3
		cur3Nums, curNum = getNext3DiceRoles(curNum)

		if g.currentPlayer == 1 {
			g.player1.position += cur3Nums
			for g.player1.position > 10 {
				g.player1.position -= 10
			}
			g.player1.score += g.player1.position

			if g.player1.score >= 1000 {
				return (dieCount) * g.player2.score
			}
			g.currentPlayer = 2
			continue
		}

		g.player2.position += cur3Nums
		for g.player2.position > 10 {
			g.player2.position -= 10
		}

		g.player2.score += g.player2.position

		if g.player2.score >= 1000 {
			return (dieCount) * g.player1.score
		}
		g.currentPlayer = 1

	}
}

//SolveDay21Part2 returns the number of wins of the better performing input after play a strange universe round
func SolveDay21Part2(input string) int64 {
	g := parseInput(input)
	p1, p2 := g.playUniverseMode(nil)
	if p1 > p2 {
		return p1
	}
	return p2
}

//playUniverseMode returns the number of wins for player1 and player2 after all possible games are played
func (g game) playUniverseMode(cache *map[game][2]int64) (int64, int64) {
	if g.player1.score >= 21 {
		return 1, 0
	}
	if g.player2.score >= 21 {
		return 0, 1
	}
	if cache == nil {
		cache = &map[game][2]int64{}
	}

	// check cache
	if c, ok := (*cache)[g]; ok {
		return c[0], c[1]
	}

	var p1, p2 int64
	for i1 := 1; i1 <= 3; i1++ {
		for i2 := 1; i2 <= 3; i2++ {
			for i3 := 1; i3 <= 3; i3++ {
				newGame := g
				if g.currentPlayer == 1 {
					newGame.player1.position = g.player1.position + i1 + i2 + i3
					for newGame.player1.position > 10 {
						newGame.player1.position -= 10
					}
					newGame.player1.score = g.player1.score + newGame.player1.position
					newGame.currentPlayer = 2
					tmpP1, tmpP2 := newGame.playUniverseMode(cache)
					p1 += tmpP1
					p2 += tmpP2
					continue
				}
				newGame.player2.position = g.player2.position + i1 + i2 + i3
				for newGame.player2.position > 10 {
					newGame.player2.position -= 10
				}
				newGame.player2.score = g.player2.score + newGame.player2.position
				newGame.currentPlayer = 1
				tmpP1, tmpP2 := newGame.playUniverseMode(cache)
				p1 += tmpP1
				p2 += tmpP2

			}
		}
	}

	//add to cache
	(*cache)[g] = [2]int64{p1, p2}
	return p1, p2
}

//getNext3DiceRoles returns the next sum of the next 3 dices and the next dice start point
func getNext3DiceRoles(cur int) (next3numbers int, nextCur int) {
	cur++
	if cur > 100 {
		cur = 1
	}
	next3numbers = cur

	cur++
	if cur > 100 {
		cur = 1
	}
	next3numbers += cur

	cur++
	if cur > 100 {
		cur = 1
	}
	next3numbers += cur
	return next3numbers, cur
}

//parseInput returns a game for the input
func parseInput(input string) game {
	lines, _ := utils.InputToSlice(input)
	if len(lines) != 2 {
		panic("wrong input")
	}

	var g game

	for i := range lines {
		num := strings.Split(lines[i], ": ")
		if len(num) != 2 {
			panic("wrong input")
		}
		start, _ := strconv.Atoi(num[1])
		if strings.HasPrefix(lines[i], "Player 1") {
			g.player1 = player{
				position: start,
				score:    0,
			}
		}
		if strings.HasPrefix(lines[i], "Player 2") {
			g.player2 = player{
				position: start,
				score:    0,
			}
		}
	}
	g.currentPlayer = 1
	return g
}
