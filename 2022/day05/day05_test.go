package main

import (
	"os"
	"testing"
)

func TestSolveDay05Part1(t *testing.T) {
	t.Run("Test SolveDay05Part1", func(t *testing.T) {
		got := SolveDay05Part1(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)
		expected := "CMZ"

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func BenchmarkSolveDay05Part1(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay05Part1(string(i))
	}
}

func TestSolveDay05Part2(t *testing.T) {
	t.Run("Test SolveDay05Part2", func(t *testing.T) {
		got := SolveDay05Part2(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)
		expected := "MCD"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func BenchmarkSolveDay05Part2(b *testing.B) {
	i, _ := os.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay05Part2(string(i))
	}
}
