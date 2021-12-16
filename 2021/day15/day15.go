package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"

	"github.com/dergeberl/aoc/utils"
)

type point struct {
	x, y int
}

type node struct {
	index          int
	p              point
	visited        bool
	connectedNodes []*node
	risk           int
	ownRisk        int
}

type nodes []*node

type card map[point]*int

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Part 1: %v\n", SolveDay15Part1(string(input)))
	fmt.Printf("Part 2: %v\n", SolveDay15Part2(string(input)))
}

//SolveDay15Part1 returns the path with the lowest risk for an input
func SolveDay15Part1(input string) int {
	c := parseInput(input, false)
	return c.getNodes().solve()
}

//SolveDay15Part2 returns the path with the lowest risk for an input which got resized
func SolveDay15Part2(input string) int {
	c := parseInput(input, true)
	return c.readBigCard().getNodes().solve()

}

//solve returns the risk level with the path with the lowest risk level
func (n nodes) solve() int {
	points := make(map[point]*node)
	for i := range n {
		points[n[i].p] = n[i]
	}
	points[point{}].risk = 0
	heap.Init(&n)

	// calculate first
	n.runNode(points)

	// calc rest break if next time the start field is used
	for !n.runNode(points) {
	}

	return points[n.getMaxPoint()].risk
}

//runNode checks the node with is nos visited with the lowest risk
//returns true if point is start
func (n *nodes) runNode(points map[point]*node) bool {
	p := heap.Pop(n).(*node)
	for _, connected := range p.connectedNodes {
		if points[connected.p].risk > p.risk+connected.ownRisk {
			points[connected.p].risk = p.risk + connected.ownRisk
			points[connected.p].visited = false
			heap.Fix(n, points[connected.p].index)
		}
	}
	p.visited = true
	heap.Fix(n, p.index)
	return p.p == point{}
}

//getMaxPoint returns the highest point
func (n nodes) getMaxPoint() point {
	var x, y int
	for i := range n {
		if n[i].p.x > x {
			x = n[i].p.x
		}
		if n[i].p.y > y {
			y = n[i].p.y
		}
	}
	return point{x: x, y: y}
}

//readBigCard returns a card witch is 5 times bigger on x and y axis
func (c card) readBigCard() card {
	var x, y int

	for p := range c {
		if p.x > x {
			x = p.x
		}
		if p.y > y {
			y = p.y
		}
	}

	newCard := make(map[point]*int)
	for p := range c {
		t1 := (*c[p] + 1)
		t2 := (*c[p] + 2)
		t3 := (*c[p] + 3)
		t4 := (*c[p] + 4)
		if t1 > 9 {
			t1 -= 9
		}
		if t2 > 9 {
			t2 -= 9
		}
		if t3 > 9 {
			t3 -= 9
		}
		if t4 > 9 {
			t4 -= 9
		}
		newCard[point{x: p.x, y: p.y}] = c[p]
		newCard[point{x: p.x + (1 * (x + 1)), y: p.y}] = &t1
		newCard[point{x: p.x + (2 * (x + 1)), y: p.y}] = &t2
		newCard[point{x: p.x + (3 * (x + 1)), y: p.y}] = &t3
		newCard[point{x: p.x + (4 * (x + 1)), y: p.y}] = &t4
	}
	c = newCard

	for p := range c {
		if p.x > x {
			x = p.x
		}
		if p.y > y {
			y = p.y
		}
	}

	newCard = make(map[point]*int)
	for p := range c {
		t1 := (*c[p] + 1)
		t2 := (*c[p] + 2)
		t3 := (*c[p] + 3)
		t4 := (*c[p] + 4)
		if t1 > 9 {
			t1 -= 9
		}
		if t2 > 9 {
			t2 -= 9
		}
		if t3 > 9 {
			t3 -= 9
		}
		if t4 > 9 {
			t4 -= 9
		}
		newCard[point{x: p.x, y: p.y}] = c[p]
		newCard[point{x: p.x, y: p.y + (1 * (y + 1))}] = &t1
		newCard[point{x: p.x, y: p.y + (2 * (y + 1))}] = &t2
		newCard[point{x: p.x, y: p.y + (3 * (y + 1))}] = &t3
		newCard[point{x: p.x, y: p.y + (4 * (y + 1))}] = &t4
	}
	c = newCard

	return c
}

//parseInput returns a card for an input
func parseInput(input string, part2 bool) card {
	line, _ := utils.InputToSlice(input)
	c := make(card)
	for y := range line {
		for x, v := range line[y] {
			number := int(v) - 48
			c[point{x: x, y: y}] = &number
		}
	}
	return c
}

//getNodes creates nodes from a card
func (c card) getNodes() nodes {
	tMap := make(map[point]*node)
	for p := range c {
		tMap[p] = &node{
			p:       p,
			visited: false,
			risk:    math.MaxInt,
			ownRisk: *c[p],
		}
	}

	n := make(nodes, 0)
	i := 0
	for p := range tMap {
		connections := make([]*node, 0)
		down := point{x: p.x, y: p.y + 1}
		if tMap[down] != nil {
			connections = append(connections, tMap[down])
		}
		up := point{x: p.x, y: p.y - 1}
		if tMap[up] != nil {
			connections = append(connections, tMap[up])
		}
		right := point{x: p.x + 1, y: p.y}
		if tMap[right] != nil {
			connections = append(connections, tMap[right])
		}
		left := point{x: p.x - 1, y: p.y}
		if tMap[left] != nil {
			connections = append(connections, tMap[left])
		}
		tMap[p].connectedNodes = connections
		tMap[p].index = i
		n = append(n, tMap[p])
		i++
	}
	return n
}

//functions for heap and sort

func (n nodes) Len() int { return len(n) }

func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
	n[i].index = i
	n[j].index = j
}

func (n nodes) Less(i, j int) bool {
	if n[i].visited && !n[j].visited {
		return false
	}
	if !n[i].visited && n[j].visited {
		return true
	}
	return n[i].risk < n[j].risk
}

func (no *nodes) Push(x interface{}) {
	n := len(*no)
	item := x.(*node)
	item.index = n
	*no = append(*no, item)
}

func (no *nodes) Pop() interface{} {
	n := len((*no))
	return (*no)[n-1]
}
