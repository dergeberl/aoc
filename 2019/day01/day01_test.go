package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay01Part1(t *testing.T) {
	t.Run("Test SolveDay01Part1", func(t *testing.T) {
		got := SolveDay01Part1(`14
1969
100756`)
		expected := 2 + 654 + 33583

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay01Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay01Part1(string(i))
	}
}

func TestSolveDay01Part2(t *testing.T) {
	t.Run("Test SolveDay01Part2", func(t *testing.T) {
		got := SolveDay01Part2(`100756`)
		expected := 50346
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay01Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay01Part2(string(i))
	}
}
