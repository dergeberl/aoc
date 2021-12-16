package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay15Part1(t *testing.T) {
	t.Run("Test SolveDay15Part1", func(t *testing.T) {
		got := SolveDay15Part1(`1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`)
		expected := 40

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay15Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay15Part1(string(i))
	}
}

func TestSolveDay15Part2(t *testing.T) {
	t.Run("Test SolveDay15Part2", func(t *testing.T) {
		got := SolveDay15Part2(`1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`)
		expected := 315
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay15Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay15Part2(string(i))
	}
}
