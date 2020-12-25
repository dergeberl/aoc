package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay25Part1(t *testing.T) {
	t.Run("Test SolveDay25Part1", func(t *testing.T) {
		i := intListToSlice(`5764801
17807724`)
		got := SolveDay25Part1(i)
		expected := 14897079

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay25Part1", func(t *testing.T) {
		i := intListToSlice(`17807724
5764801`)
		got := SolveDay25Part1(i)
		expected := 14897079

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay25Part1", func(t *testing.T) {
		i := intListToSlice(`5764801
17807724
123456`)
		got := SolveDay25Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay25Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := intListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay25Part1(input)
	}
}

func TestIntListToSlice(t *testing.T) {
	t.Run("Test intListToSlice with valid list", func(t *testing.T) {
		s := `123
1234
000
00123`
		got := intListToSlice(s)
		expected := []int{123, 1234, 0, 123}
		if !equalInt(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test intListToSlice with invalid list", func(t *testing.T) {
		s := `123
1234a
000
00123`
		got := intListToSlice(s)
		expected := []int{}
		if !equalInt(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func equalInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
