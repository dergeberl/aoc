package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay18Part1(t *testing.T) {
	t.Run("Test SolveDay18Part1", func(t *testing.T) {
		i := stringListToSlice(`1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`)
		got := SolveDay18Part1(i)
		expected := 71 + 51 + 26 + 437 + 12240 + 13632

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay18Part1", func(t *testing.T) {
		i := stringListToSlice(`1 + 2 * 3 - 4`)
		got := SolveDay18Part1(i)
		expected := 5

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay18Part1", func(t *testing.T) {
		i := stringListToSlice(`9 - ( 2 + 1)`)
		got := SolveDay18Part1(i)
		expected := 6

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay18Part1", func(t *testing.T) {
		i := stringListToSlice(`9 - (( 2 + 1)`)
		got := SolveDay18Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay18Part1", func(t *testing.T) {
		i := stringListToSlice(`9 - ( s + 1)`)
		got := SolveDay18Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay18Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay18Part1(input)
	}
}

func TestSolveDay18Part2(t *testing.T) {
	t.Run("Test SolveDay18Part2", func(t *testing.T) {
		i := stringListToSlice(`1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`)
		got := SolveDay18Part2(i)
		expected := 231 + 51 + 46 + 1445 + 669060 + 23340
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay18Part2", func(t *testing.T) {
		i := stringListToSlice(`1 + 2 * 3 - 2`)
		got := SolveDay18Part2(i)
		expected := 3
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay18Part2", func(t *testing.T) {
		i := stringListToSlice(`9 - ( 2 + 1)`)
		got := SolveDay18Part2(i)
		expected := 6
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay18Part2", func(t *testing.T) {
		i := stringListToSlice(`(2 * 2) - (3 - 2)`)
		got := SolveDay18Part2(i)
		expected := 3
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay18Part2", func(t *testing.T) {
		i := stringListToSlice(`9 - ( 2 + s)`)
		got := SolveDay18Part2(i)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay18Part2", func(t *testing.T) {
		i := stringListToSlice(`9 - ( 2 + 1))`)
		got := SolveDay18Part2(i)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay18Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay18Part2(input)
	}
}

func TestStringListToSlice(t *testing.T) {
	t.Run("Test stringListToSlice with valid list", func(t *testing.T) {
		s := `string
string 2
string:3`
		got := stringListToSlice(s)
		expected := []string{"string", "string 2", "string:3"}
		if !equalString(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func equalString(a, b []string) bool {
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
