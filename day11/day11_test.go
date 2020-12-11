package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay11Part1(t *testing.T) {
	t.Run("Test SolveDay11Part1", func(t *testing.T) {
		i := stringListToSlice(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)
		got := SolveDay11Part1(i)
		expected := 37

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay11Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay11Part1(input)
	}
}

func TestSolveDay11Part2(t *testing.T) {
	t.Run("Test SolveDay11Part2", func(t *testing.T) {
		i := stringListToSlice(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)
		got := SolveDay11Part2(i)
		expected := 26
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func TestApplySeatRulesPart1(t *testing.T) {
	t.Run("Test applySeatRulesPart1", func(t *testing.T) {
		i := stringListToSlice(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)
		got := printMap(applySeatRulesPart1(sliceToMap(i)))
		expected := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

		if got != expected {
			t.Errorf("expected '%s' but got '%s'", expected, got)
		}
	})
	t.Run("Test applySeatRulesPart1", func(t *testing.T) {
		i := stringListToSlice(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`)
		got := printMap(applySeatRulesPart1(sliceToMap(i)))
		expected := `#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`

		if got != expected {
			t.Errorf("expected '%s' but got '%s'", expected, got)
		}
	})
}

func TestApplySeatRulesPart2(t *testing.T) {
	t.Run("Test applySeatRulesPart2", func(t *testing.T) {
		i := stringListToSlice(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)
		got := printMap(applySeatRulesPart2(sliceToMap(i)))
		expected := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

		if got != expected {
			t.Errorf("expected '%s' but got '%s'", expected, got)
		}
	})
	t.Run("Test applySeatRulesPart2", func(t *testing.T) {
		i := stringListToSlice(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`)
		got := printMap(applySeatRulesPart2(sliceToMap(i)))
		expected := `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`

		if got != expected {
			t.Errorf("expected '%s' but got '%s'", expected, got)
		}
	})
}


func TestConversion(t *testing.T) {
	t.Run("Test Conversion", func(t *testing.T) {
		input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
		i := stringListToSlice(input)
		got := printMap(sliceToMap(i))
		expected := input
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func TestCountSeats(t *testing.T) {
	t.Run("Test countSeats", func(t *testing.T) {
		input := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
		i := stringListToSlice(input)
		got := countSeats(sliceToMap(i))
		expected := 0
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test countSeats", func(t *testing.T) {
		input := `L.LL.LL.LL
#######.L#
L.L.#..L..
##LL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.L###LL.L
L.LLLLL.LL`
		i := stringListToSlice(input)
		got := countSeats(sliceToMap(i))
		expected := 14
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func BenchmarkSolveDay11Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay11Part2(input)
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
