package main

import (
	"os"
	"testing"
)

func TestSolveDay02Part1(t *testing.T) {
	t.Run("Test SolveDay02Part1", func(t *testing.T) {
		got := SolveDay02Part1(`A Y
B X
C Z`)
		expected := 15

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay02Part1(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay02Part1(string(i))
	}
}

func TestSolveDay02Part2(t *testing.T) {
	t.Run("Test SolveDay02Part2", func(t *testing.T) {
		got := SolveDay02Part2(`A Y
B X
C Z`)
		expected := 12
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay02Part2(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay02Part2(string(i))
	}
}
