package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay22Part1(t *testing.T) {
	t.Run("Test SolveDay22Part1", func(t *testing.T) {
		i := `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`
		got := SolveDay22Part1(i)
		expected := 306

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay22Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay22Part1(input)
	}
}

func TestSolveDay22Part2(t *testing.T) {
	t.Run("Test SolveDay22Part2", func(t *testing.T) {
		i := `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`
		got := SolveDay22Part2(i)
		expected := 291
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay22Part2", func(t *testing.T) {
		i := `Player 1:
43
19

Player 2:
2
29
14`
		got := SolveDay22Part2(i)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay22Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay22Part2(input)
	}
}
