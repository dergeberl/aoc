package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay9Part1(t *testing.T) {
	t.Run("Test SolveDay9Part1", func(t *testing.T) {
		i := intListToSlice(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`)
		got := SolveDay9Part1(i, 5)
		expected := 127

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay9Part1 infinite loop", func(t *testing.T) {
		i := intListToSlice(`1
2
3`)
		got := SolveDay9Part1(i, 2)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay9Part1 invalid input", func(t *testing.T) {
		i := intListToSlice(`1`)
		got := SolveDay9Part1(i, 2)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay9Part1 invalid input", func(t *testing.T) {
		i := intListToSlice(`1
2`)
		got := SolveDay9Part1(i, 1)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay9Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := intListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay9Part1(input, 25)
	}
}

func TestSolveDay9Part2(t *testing.T) {
	t.Run("Test SolveDay9Part2", func(t *testing.T) {
		i := intListToSlice(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`)
		got := SolveDay9Part2(i, 5)
		expected := 62
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay9Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := intListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay9Part2(input, 5)
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
