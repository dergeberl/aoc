package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay25Part1(t *testing.T) {
	t.Run("Test SolveDay25Part1", func(t *testing.T) {
		got := SolveDay25Part1(`v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`)
		expected := 58

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay25Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay25Part1(string(i))
	}
}
