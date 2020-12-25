package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay21Part1(t *testing.T) {
	t.Run("Test SolveDay21Part1", func(t *testing.T) {
		i := stringListToSlice(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`)
		got := SolveDay21Part1(i)
		expected := 5

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay21Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay21Part1(input)
	}
}

func TestSolveDay21Part2(t *testing.T) {
	t.Run("Test SolveDay21Part2", func(t *testing.T) {
		i := stringListToSlice(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`)
		got := SolveDay21Part2(i)
		expected := "mxmxvkd,sqjhc,fvjkl"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func BenchmarkSolveDay21Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay21Part2(input)
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
