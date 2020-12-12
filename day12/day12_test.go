package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay12Part1(t *testing.T) {
	t.Run("Test SolveDay12Part1", func(t *testing.T) {
		i := stringListToSlice(`F10
N3
F7
R90
R180
R180
L180
L180
S1
N1
W1
E1
F11
R90
F1
R180
F1
R90
F1
R180
F1`)
		got := SolveDay12Part1(i)
		expected := 25

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay12Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay12Part1(input)
	}
}

func TestSolveDay12Part2(t *testing.T) {
	t.Run("Test SolveDay12Part2", func(t *testing.T) {
		i := stringListToSlice(`F10
N3
F7
R90
R180
R180
L180
L180
S1
N1
W1
E1
F11
R90
F1
R180
F1
R90
F1
R180
F1`)
		got := SolveDay12Part2(i)
		expected := 286
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay12Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay12Part2(input)
	}
}

func TestSliceToStepSlice(t *testing.T) {
	t.Run("Test sliceToStepSlice with valid list", func(t *testing.T) {
		s := []string{"R1",
			"L1",
			"E12345"}
		got := sliceToStepSlice(s)
		expected := []step{{
			action: "R",
			number: 1,
		},
			{
				action: "L",
				number: 1,
			},
			{
				action: "E",
				number: 12345,
			}}
		if !equalSlice(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test sliceToStepSlice with valid list", func(t *testing.T) {
		s := []string{"R1",
			"L1",
			"E1234s5"}
		got := sliceToStepSlice(s)
		expected := []step{}
		if !equalSlice(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
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

func TestIntListToSlice(t *testing.T) {
	t.Run("Test intListToSlice with valid list", func(t *testing.T) {
		s := `123
1234
000
00123`
		got := intListToSlice(s)
		expected := []int{123,1234,0,123}
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

func equalSlice(a, b []step) bool {
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
