package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay15Part1(t *testing.T) {
	t.Run("Test SolveDay15Part1", func(t *testing.T) {
		i := `0,3,6`
		got := SolveDay15Part1(i)
		expected := 436

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay15Part1", func(t *testing.T) {
		i := `1,3,2`
		got := SolveDay15Part1(i)
		expected := 1

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay15Part1", func(t *testing.T) {
		i := `2,1,3`
		got := SolveDay15Part1(i)
		expected := 10

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay15Part1", func(t *testing.T) {
		i := `1,2,3`
		got := SolveDay15Part1(i)
		expected := 27

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay15Part1", func(t *testing.T) {
		i := `2,3,1`
		got := SolveDay15Part1(i)
		expected := 78

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay15Part1", func(t *testing.T) {
		i := `3,2,1`
		got := SolveDay15Part1(i)
		expected := 438

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay15Part1", func(t *testing.T) {
		i := `3,1,2`
		got := SolveDay15Part1(i)
		expected := 1836

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay15Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay15Part1(input)
	}
}

func TestSolveDay15Part2(t *testing.T) {
	t.Run("Test SolveDay15Part2", func(t *testing.T) {
		i := `0,3,6`
		got := SolveDay15Part2(i)
		expected := 175594

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay15Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay15Part2(input)
	}
}

func TestSolvePlayElvesGame(t *testing.T) {
	t.Run("Test playElvesGame", func(t *testing.T) {
		i := `0,3,6`
		got := playElvesGame(i, 2020)
		expected := 436

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test playElvesGame", func(t *testing.T) {
		i := `0,3,6`
		got := playElvesGame(i, 2020)
		expected := 436

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test playElvesGame", func(t *testing.T) {
		i := `0,3,6`
		got := playElvesGame(i, 0)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test playElvesGame", func(t *testing.T) {
		i := `0,a,6`
		got := playElvesGame(i, 0)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkPlayElvesGame(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = playElvesGame("0,3,6", 10000)
	}
}

func TestSolvePlayElvesGameR(t *testing.T) {
	t.Run("Test playElvesGameR", func(t *testing.T) {
		i := `0,3,6`
		got := playElvesGameR(i, 2020)
		expected := 436

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test playElvesGameR", func(t *testing.T) {
		i := `0,3,6`
		got := playElvesGameR(i, 2020)
		expected := 436

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test playElvesGameR", func(t *testing.T) {
		i := `0,3,6`
		got := playElvesGameR(i, 0)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test playElvesGameR", func(t *testing.T) {
		i := `0,a,6`
		got := playElvesGameR(i, 0)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkPlayElvesGameR(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = playElvesGameR("0,3,6", 10000)
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
