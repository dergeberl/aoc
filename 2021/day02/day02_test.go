package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay02Part1(t *testing.T) {
	t.Run("Test SolveDay02Part1", func(t *testing.T) {
		got := SolveDay02Part1(`forward 5
down 5
forward 8
up 3
down 8
forward 2`)
		expected := 150

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay02Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay02Part1(string(i))
	}
}

func TestSolveDay02Part2(t *testing.T) {
	t.Run("Test SolveDay02Part2", func(t *testing.T) {
		got := SolveDay02Part2(`forward 5
down 5
forward 8
up 3
down 8
forward 2`)
		expected := 900
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay02Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay02Part2(string(i))
	}
}
