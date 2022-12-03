package main

import (
	"os"
	"testing"
)

func TestSolveDay03Part1(t *testing.T) {
	t.Run("Test SolveDay03Part1", func(t *testing.T) {
		got := SolveDay03Part1(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)
		expected := 157

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay03Part1(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay03Part1(string(i))
	}
}

func TestSolveDay03Part2(t *testing.T) {
	t.Run("Test SolveDay03Part2", func(t *testing.T) {
		got := SolveDay03Part2(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)
		expected := 70
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay03Part2(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay03Part2(string(i))
	}
}
