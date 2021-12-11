package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay11Part1(t *testing.T) {
	t.Run("Test SolveDay11Part1", func(t *testing.T) {
		got := SolveDay11Part1(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`)
		expected := 1656

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay11Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay11Part1(string(i))
	}
}

func TestSolveDay11Part2(t *testing.T) {
	t.Run("Test SolveDay11Part2", func(t *testing.T) {
		got := SolveDay11Part2(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`)
		expected := 195
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay11Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay11Part2(string(i))
	}
}
