package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay8Part1(t *testing.T) {
	t.Run("Test SolveDay8Part1", func(t *testing.T) {
		i := stringListToSlice(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)
		got := SolveDay8Part1(i)
		expected := 5

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay8Part1", func(t *testing.T) {
		i := stringListToSlice(`acc a`)
		got := SolveDay8Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay8Part1", func(t *testing.T) {
		i := stringListToSlice(`jmp a`)
		got := SolveDay8Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay8Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay8Part1(input)
	}
}

func TestSolveDay8Part1r(t *testing.T) {
	t.Run("Test SolveDay8Part1r", func(t *testing.T) {
		i := stringListToSlice(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)
		got := SolveDay8Part1r(i)
		expected := 5

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay8Part1r(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay8Part1r(input)
	}
}

func TestSolveDay8Part2(t *testing.T) {
	t.Run("Test SolveDay8Part2", func(t *testing.T) {
		i := stringListToSlice(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)
		got := SolveDay8Part2(i)
		expected := 8
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay8Part2", func(t *testing.T) {
		i := stringListToSlice(`jmp a`)
		got := SolveDay8Part2(i)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay8Part2", func(t *testing.T) {
		i := stringListToSlice(`acc a`)
		got := SolveDay8Part2(i)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay8Part2", func(t *testing.T) {
		i := stringListToSlice(`nop a`)
		got := SolveDay8Part2(i)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay8Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay8Part2(input)
	}
}
func TestSolveDay8Part2r(t *testing.T) {
	t.Run("Test SolveDay8Part2r", func(t *testing.T) {
		i := stringListToSlice(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)
		got := SolveDay8Part2r(i)
		expected := 8
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay8Part2r(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay8Part2r(input)
	}
}

func TestGetTasks(t *testing.T) {
	t.Run("Test getTasks with valid list", func(t *testing.T) {
		s := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99`
		got := getTasks(stringListToSlice(s))
		expected := map[int]map[string]int{
			0: {"nop": 0},
			1: {"acc": 1},
			2: {"jmp": +4},
			3: {"acc": +3},
			4: {"jmp": -3},
			5: {"acc": -99},
		}
		for i, v := range expected {
			for i2, v2 := range v {
				if v2 != got[i][i2]{
					t.Errorf("expected '%v' but got '%v'", expected, got)
				}

			}

		}
	})
	t.Run("Test getTasks with invalid list", func(t *testing.T) {
		s := `nop +0
acc a
jmp +4
acc +3
jmp -3
acc -99`
		got := getTasks(stringListToSlice(s))
		if got != nil{
			t.Errorf("expected '%v' but got '%v'", nil, got)
		}
	})
}

func BenchmarkGetTasks(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = getTasks(input)
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
