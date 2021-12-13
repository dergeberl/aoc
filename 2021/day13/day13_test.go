package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay13Part1(t *testing.T) {
	t.Run("Test SolveDay13Part1", func(t *testing.T) {
		got := SolveDay13Part1(`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`)
		expected := 17

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay13Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay13Part1(string(i))
	}
}

func TestSolveDay13Part2(t *testing.T) {
	t.Run("Test SolveDay13Part2", func(t *testing.T) {
		got := SolveDay13Part2(`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`)
		expected := `#####
#...#
#...#
#...#
#####
.....
.....
`
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func BenchmarkSolveDay13Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay13Part2(string(i))
	}
}
