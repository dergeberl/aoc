package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay14Part1(t *testing.T) {
	t.Run("Test SolveDay14Part1", func(t *testing.T) {
		got := SolveDay14Part1(`NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`)
		expected := 1588

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay14Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay14Part1(string(i))
	}
}

func TestSolveDay14Part2(t *testing.T) {
	t.Run("Test SolveDay14Part2", func(t *testing.T) {
		got := SolveDay14Part2(`NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`)
		expected := 2188189693529
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay14Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay14Part2(string(i))
	}
}
