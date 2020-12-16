package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay16Part1(t *testing.T) {
	t.Run("Test SolveDay16Part1", func(t *testing.T) {
		i := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`
		got := SolveDay16Part1(i)
		expected := 71

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay16Part1", func(t *testing.T) {
		i := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,x,47
40,4,50
55,2,20
38,6,12`
		got := SolveDay16Part1(i)
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay16Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay16Part1(input)
	}
}

func TestSolveDay16Part2(t *testing.T) {
	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		i := `departure class: 0-1 or 4-19
departure row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
100,100,100`
		got := SolveDay16Part2(i)
		expected := 12 * 11
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay16Part2", func(t *testing.T) {
		i := `departure class: 0-1 or 4-19
departure row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,x,18
15,1,5
5,14,9
100,100,100`
		got := SolveDay16Part2(i)
		expected := 0
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay16Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	for n := 0; n < b.N; n++ {
		_ = SolveDay16Part2(input)
	}
}
func TestGetTicketData(t *testing.T) {
	t.Run("Test getTicketData", func(t *testing.T) {
		i := `departure class: 01 or 4-19

your ticket:
11,12,13

nearby tickets:
3,9,18
`
		got1, got2, got3, got4 := getTicketData(i)

		if got1 != nil || got2 != nil || got3 != nil || got4 != nil {
			t.Errorf("expected '%v''%v''%v''%v' but got '%v''%v''%v''%v'", nil, nil, nil, nil, got1, got2, got3, got4)
		}
	})
	t.Run("Test getTicketData", func(t *testing.T) {
		i := `departure class: 01 or 4-19`
		got1, got2, got3, got4 := getTicketData(i)

		if got1 != nil || got2 != nil || got3 != nil || got4 != nil {
			t.Errorf("expected '%v''%v''%v''%v' but got '%v''%v''%v''%v'", nil, nil, nil, nil, got1, got2, got3, got4)
		}
	})
	t.Run("Test getTicketData", func(t *testing.T) {
		i := `departure class 01 or 4-19

your ticket:
11,12,13

nearby tickets:
3,9,18
`
		got1, got2, got3, got4 := getTicketData(i)

		if got1 != nil || got2 != nil || got3 != nil || got4 != nil {
			t.Errorf("expected '%v''%v''%v''%v' but got '%v''%v''%v''%v'", nil, nil, nil, nil, got1, got2, got3, got4)
		}
	})
	t.Run("Test getTicketData", func(t *testing.T) {
		i := `departure class: 0-a or 4-19

your ticket:
11,12,13

nearby tickets:
3,9,18
`
		got1, got2, got3, got4 := getTicketData(i)

		if got1 != nil || got2 != nil || got3 != nil || got4 != nil {
			t.Errorf("expected '%v''%v''%v''%v' but got '%v''%v''%v''%v'", nil, nil, nil, nil, got1, got2, got3, got4)
		}
	})
	t.Run("Test getTicketData", func(t *testing.T) {
		i := `departure class: a-1 or 4-19

your ticket:
11,12,13

nearby tickets:
3,9,18
`
		got1, got2, got3, got4 := getTicketData(i)

		if got1 != nil || got2 != nil || got3 != nil || got4 != nil {
			t.Errorf("expected '%v''%v''%v''%v' but got '%v''%v''%v''%v'", nil, nil, nil, nil, got1, got2, got3, got4)
		}
	})
	t.Run("Test getTicketData", func(t *testing.T) {
		i := `departure class: 0-3 or 4-19

your ticket:
11,x,13

nearby tickets:
3,9,18
`
		got1, got2, got3, got4 := getTicketData(i)

		if got1 != nil || got2 != nil || got3 != nil || got4 != nil {
			t.Errorf("expected '%v''%v''%v''%v' but got '%v''%v''%v''%v'", nil, nil, nil, nil, got1, got2, got3, got4)
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

func equalInt64(a, b []int64) bool {
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
