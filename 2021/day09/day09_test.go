package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay09Part1(t *testing.T) {
	t.Run("Test SolveDay09Part1", func(t *testing.T) {
		got := SolveDay09Part1(`2199943210
3987894921
9856789892
8767896789
9899965678`)
		expected := 15

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay09Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay09Part1(string(i))
	}
}

func TestSolveDay09Part2(t *testing.T) {
	t.Run("Test SolveDay09Part2", func(t *testing.T) {
		got := SolveDay09Part2(`2199943210
3987894921
9856789892
8767896789
9899965678`)
		expected := 1134
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay09Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay09Part2(string(i))
	}
}
