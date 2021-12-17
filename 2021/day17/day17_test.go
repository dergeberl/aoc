package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay17Part1(t *testing.T) {
	t.Run("Test SolveDay17Part1", func(t *testing.T) {
		got := SolveDay17Part1(`target area: x=20..30, y=-10..-5`)
		expected := 45

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay17Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay17Part1(string(i))
	}
}

func TestSolveDay17Part2(t *testing.T) {
	t.Run("Test SolveDay17Part2", func(t *testing.T) {
		got := SolveDay17Part2(`target area: x=20..30, y=-10..-5`)
		expected := 112
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay17Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay17Part2(string(i))
	}
}
