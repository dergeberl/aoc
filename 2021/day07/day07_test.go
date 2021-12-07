package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay07Part1(t *testing.T) {
	t.Run("Test SolveDay07Part1", func(t *testing.T) {
		got := SolveDay07Part1(`16,1,2,0,4,2,7,1,2,14`)
		expected := 37

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay07Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay07Part1(string(i))
	}
}

func TestSolveDay07Part2(t *testing.T) {
	t.Run("Test SolveDay07Part2", func(t *testing.T) {
		got := SolveDay07Part2(`16,1,2,0,4,2,7,1,2,14`)
		expected := 168
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay07Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay07Part2(string(i))
	}
}
