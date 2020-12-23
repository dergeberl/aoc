package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay23Part1(t *testing.T) {
	t.Run("Test SolveDay23Part1", func(t *testing.T) {
		i := "389125467"
		got := SolveDay23Part1(i)
		expected := 67384529

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func BenchmarkSolveDay23Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay23Part1(input)
	}
}

func TestSolveDay23Part2(t *testing.T) {
	t.Run("Test SolveDay23Part2", func(t *testing.T) {
		i := "389125467"
		got := SolveDay23Part2(i)
		expected := 149245887792
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func BenchmarkSolveDay23Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay23Part2(input)
	}
}
