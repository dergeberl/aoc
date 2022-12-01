package main

import (
	"os"
	"testing"
)

func TestSolveDay01Part1(t *testing.T) {
	t.Run("Test SolveDay01Part1", func(t *testing.T) {
		got := SolveDay01Part1(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
		expected := 24000

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay01Part1(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay01Part1(string(i))
	}
}

func TestSolveDay01Part2(t *testing.T) {
	t.Run("Test SolveDay01Part2", func(t *testing.T) {
		got := SolveDay01Part2(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
		expected := 45000
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay01Part2(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay01Part2(string(i))
	}
}
