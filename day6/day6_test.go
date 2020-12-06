package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay6Part1(t *testing.T) {
	t.Run("Test SolveDay6Part1", func(t *testing.T) {
		i := `abc

a
b
c

ab
ac

a
a
a
a

b`
		got := SolveDay6Part1(i)
		expected := 11

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay6Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay6Part1(input)
	}
}

func TestSolveDay6Part2(t *testing.T) {
	t.Run("Test SolveDay6Part2", func(t *testing.T) {
		i := `abc

a
b
c

ab
ac

a
a
a
a

b`
		got := SolveDay6Part2(i)
		expected := 6
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay6Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay6Part2(input)
	}
}

func TestDeleteDuplicates(t *testing.T) {
	t.Run("Test deleteDuplicates", func(t *testing.T) {
		s := `abcdef`
		got := deleteDuplicates(s)
		expected := "abcdef"
		if got != expected{
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test deleteDuplicates", func(t *testing.T) {
		s := `aaaaaaaabcdef`
		got := deleteDuplicates(s)
		expected := "abcdef"
		if got != expected{
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test deleteDuplicates", func(t *testing.T) {
		s := `aaaaaaaabcdaaaaaaefaabbbbfffbbbaa`
		got := deleteDuplicates(s)
		expected := "abcdef"
		if got != expected{
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
