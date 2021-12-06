package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay06Part1(t *testing.T) {
	t.Run("Test SolveDay06Part1", func(t *testing.T) {
		got := SolveDay06Part1(`3,4,3,1,2`)
		expected := 5934

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay06Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay06Part1(string(i))
	}
}

func TestSolveDay06Part2(t *testing.T) {
	t.Run("Test SolveDay06Part2", func(t *testing.T) {
		got := SolveDay06Part2(`3,4,3,1,2`)
		expected := 26984457539
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay06Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay06Part2(string(i))
	}
}
