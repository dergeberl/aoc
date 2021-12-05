package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay05Part1(t *testing.T) {
	t.Run("Test SolveDay05Part1", func(t *testing.T) {
		got := SolveDay05Part1(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`)
		expected := 5

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay05Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay05Part1(string(i))
	}
}

func TestSolveDay05Part2(t *testing.T) {
	t.Run("Test SolveDay05Part2", func(t *testing.T) {
		got := SolveDay05Part2(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
906,28 -> 906,957`)
		expected := 12
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay05Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay05Part2(string(i))
	}
}
