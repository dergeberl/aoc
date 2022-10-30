package main

import (
	"os"
	"testing"
)

func TestSolveDay0Part1(t *testing.T) {
	t.Run("Test SolveDay0Part1", func(t *testing.T) {
		got := SolveDay0Part1(``)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay0Part1(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay0Part1(string(i))
	}
}

func TestSolveDay0Part2(t *testing.T) {
	t.Run("Test SolveDay0Part2", func(t *testing.T) {
		got := SolveDay0Part2(``)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay0Part2(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay0Part2(string(i))
	}
}
