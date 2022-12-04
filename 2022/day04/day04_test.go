package main

import (
	"os"
	"testing"
)

func TestSolveDay04Part1(t *testing.T) {
	t.Run("Test SolveDay04Part1", func(t *testing.T) {
		got := SolveDay04Part1(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)
		expected := 2

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay04Part1(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay04Part1(string(i))
	}
}

func TestSolveDay04Part2(t *testing.T) {
	t.Run("Test SolveDay04Part2", func(t *testing.T) {
		got := SolveDay04Part2(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
6-6,4-7
2-6,4-8`)
		expected := 5
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay04Part2(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay04Part2(string(i))
	}
}
