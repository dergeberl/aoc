package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay21Part1(t *testing.T) {
	t.Run("Test SolveDay21Part1", func(t *testing.T) {
		got := SolveDay21Part1(`Player 1 starting position: 4
Player 2 starting position: 8`)
		expected := 739785

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay21Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay21Part1(string(i))
	}
}

func TestSolveDay21Part2(t *testing.T) {
	t.Run("Test SolveDay21Part2", func(t *testing.T) {
		got := SolveDay21Part2(`Player 1 starting position: 4
Player 2 starting position: 8`)
		expected := int64(444356092776315)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay21Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay21Part2(string(i))
	}
}
