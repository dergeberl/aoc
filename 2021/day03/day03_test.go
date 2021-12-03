package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay03Part1(t *testing.T) {
	t.Run("Test SolveDay03Part1", func(t *testing.T) {
		got := SolveDay03Part1(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)
		expected := 198

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay03Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay03Part1(string(i))
	}
}

func TestSolveDay03Part2(t *testing.T) {
	t.Run("Test SolveDay03Part2", func(t *testing.T) {
		got := SolveDay03Part2(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)
		expected := 230
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay03Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay03Part2(string(i))
	}
}
