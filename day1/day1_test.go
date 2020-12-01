package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay1Part1(t *testing.T) {
	t.Run("Test SolveDay1Part1 with valid list", func(t *testing.T) {
		got := SolveDay1Part1([]int{1721, 979, 366, 299, 675, 1456})
		expected := 514579

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay1Part1 with invalid list", func(t *testing.T) {
		got := SolveDay1Part1([]int{1720, 979, 366, 299, 675, 1456})
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
func TestSolveDay1Part2(t *testing.T) {
	t.Run("Test SolveDay1Part2 with valid list", func(t *testing.T) {
		got := SolveDay1Part2([]int{1721, 979, 366, 299, 675, 1456})
		expected := 241861950

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay1Part2 with invalid list", func(t *testing.T) {
		got := SolveDay1Part2([]int{1720, 979, 365, 299, 675, 1456})
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay1Part1(b *testing.B) {
	// run the Fib function b.N times
	i, _ := ioutil.ReadFile("input.txt")
	input := listToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay1Part1(input)
	}
}

func BenchmarkSolveDay1Part2(b *testing.B) {
	// run the Fib function b.N times
	i, _ := ioutil.ReadFile("input.txt")
	input := listToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay1Part2(input)
	}
}

func TestListToSlice(t *testing.T) {

	t.Run("Test ListToSlice with valid list", func(t *testing.T) {
		s := `1721
979
366
299
675
1456`
		got := listToSlice(s)
		expected := []int{1721, 979, 366, 299, 675, 1456}

		if !equal(got, expected) {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test ListToSlice with invalid list", func(t *testing.T) {
		s := `1721
979
s366
299
675
1456`
		got := listToSlice(s)
		expected := []int{}

		if !equal(got, expected) {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func equal(a, b []int) bool {
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
