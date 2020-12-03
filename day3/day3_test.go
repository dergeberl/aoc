package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay3Part1(t *testing.T) {
	t.Run("Test SolveDay2Part1 with valid list", func(t *testing.T) {
		i := listToSlice(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)
		got := SolveDay3Part1(i, 3, 1)
		expected := 7

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay3Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := listToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay3Part1(input, 3, 1)
	}
}

func TestSolveDay3Part2(t *testing.T) {
	t.Run("Test SolveDay2Part2 with valid list", func(t *testing.T) {
		i := listToSlice(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)
		got := SolveDay3Part2(i)
		expected := 336
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay3Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := listToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay3Part2(input)
	}
}

func TestListToSlice(t *testing.T) {
	t.Run("Test ListToSlice with valid list", func(t *testing.T) {
		s := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
		got := listToSlice(s)
		expected := []string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#", ".#...##..#.", "..#.##.....", ".#.#.#....#", ".#........#", "#.##...#...", "#...##....#", ".#..#...#.#"}
		if !equal(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func equal(a, b []string) bool {
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
