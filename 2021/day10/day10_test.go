package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay10Part1(t *testing.T) {
	t.Run("Test SolveDay10Part1", func(t *testing.T) {
		got := SolveDay10Part1(`[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`)
		expected := 26397

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay10Part1", func(t *testing.T) {
		got := SolveDay10Part1(`<)`)
		expected := 3

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay10Part1", func(t *testing.T) {
		got := SolveDay10Part1(`<]`)
		expected := 57

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay10Part1", func(t *testing.T) {
		got := SolveDay10Part1(`<}`)
		expected := 1197

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay10Part1", func(t *testing.T) {
		got := SolveDay10Part1(`[>`)
		expected := 25137

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay10Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay10Part1(string(i))
	}
}

func TestSolveDay10Part2(t *testing.T) {
	t.Run("Test SolveDay10Part2", func(t *testing.T) {
		got := SolveDay10Part2(`[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`)
		expected := 288957
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay10Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		_ = SolveDay10Part2(string(i))
	}
}
