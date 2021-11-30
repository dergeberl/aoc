package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay13Part1(t *testing.T) {
	t.Run("Test SolveDay13Part1", func(t *testing.T) {
		i := stringListToSlice(`939
7,13,x,x,59,x,31,19`)
		got := SolveDay13Part1(i)
		expected := 295

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part1", func(t *testing.T) {
		i := stringListToSlice(`9s39
7,13,x,x,59,x,31,19`)
		got := SolveDay13Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part1", func(t *testing.T) {
		i := stringListToSlice(``)
		got := SolveDay13Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part1", func(t *testing.T) {
		i := stringListToSlice(`123
`)
		got := SolveDay13Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part1", func(t *testing.T) {
		i := stringListToSlice(`123
0,0,0`)
		got := SolveDay13Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay13Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay13Part1(input)
	}
}

func TestSolveDay13Part2(t *testing.T) {
	t.Run("Test SolveDay13Part2", func(t *testing.T) {
		i := stringListToSlice(`939
7,13,x,x,59,x,31,19`)
		got := SolveDay13Part2(i)
		expected := 1068781
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part2", func(t *testing.T) {
		i := stringListToSlice(`1234
17,x,13,19`)
		got := SolveDay13Part2(i)
		expected := 3417
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part2", func(t *testing.T) {
		i := stringListToSlice(`1234
67,7,59,61`)
		got := SolveDay13Part2(i)
		expected := 754018
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part2", func(t *testing.T) {
		i := stringListToSlice(`1234
67,x,7,59,61`)
		got := SolveDay13Part2(i)
		expected := 779210
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part2", func(t *testing.T) {
		i := stringListToSlice(`1234
67,7,x,59,61`)
		got := SolveDay13Part2(i)
		expected := 1261476
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part2", func(t *testing.T) {
		i := stringListToSlice(`1234
1789,37,47,1889`)
		got := SolveDay13Part2(i)
		expected := 1202161486
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part2", func(t *testing.T) {
		i := stringListToSlice(`1234`)
		got := SolveDay13Part2(i)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay13Part2", func(t *testing.T) {
		i := stringListToSlice(`1234
0,0,0`)
		got := SolveDay13Part2(i)
		expected := 1
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay13Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay13Part2(input)
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
