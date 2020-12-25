package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay17Part1(t *testing.T) {
	t.Run("Test SolveDay17Part1", func(t *testing.T) {
		i := stringListToSlice(`.#.
..#
###`)
		got := SolveDay17Part1(i)
		expected := 112

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay17Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay17Part1(input)
	}
}

func TestSolveDay17Part2(t *testing.T) {
	t.Run("Test SolveDay17Part2", func(t *testing.T) {
		i := stringListToSlice(`.#.
..#
###`)
		got := SolveDay17Part2(i)
		expected := 848
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay17Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay17Part2(input)
	}
}

func TestCountNeighbors(t *testing.T) {
	t.Run("Test countNeighbors", func(t *testing.T) {
		inputMap := map[coordinates]bool{
			{x: 1, y: 0, z: 0, w: 0}: true,
			{x: 2, y: 1, z: 0, w: 0}: true,
			{x: 0, y: 2, z: 0, w: 0}: true,
			{x: 1, y: 2, z: 0, w: 0}: true,
			{x: 2, y: 2, z: 0, w: 0}: true,
		}
		got := countNeighbors(inputMap, -1, 2, 2)
		expected := 3

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test countNeighbors", func(t *testing.T) {
		inputMap := map[coordinates]bool{
			{x: 1, y: 0, z: 0, w: 0}: true,
			{x: 2, y: 1, z: 0, w: 0}: true,
			{x: 0, y: 2, z: 0, w: 0}: true,
			{x: 1, y: 2, z: 0, w: 0}: true,
			{x: 2, y: 2, z: 0, w: 0}: true,
		}
		got := countNeighbors(inputMap, -1, 2, 2, 5, 5)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test countNeighbors", func(t *testing.T) {
		inputMap := map[coordinates]bool{
			{x: 1, y: 0, z: 0, w: 0}: true,
			{x: 2, y: 1, z: 0, w: 0}: true,
			{x: 0, y: 2, z: 0, w: 0}: true,
			{x: 1, y: 2, z: 0, w: 0}: true,
			{x: 2, y: 2, z: 0, w: 0}: true,
		}
		got := countNeighbors(inputMap, -1, 2, 2, 1)
		expected := 3

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
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
