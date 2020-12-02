package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay2Part1(t *testing.T) {
	t.Run("Test SolveDay2Part1 with valid list", func(t *testing.T) {
		got := SolveDay2Part1([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"})
		expected := 2

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay2Part1 with invalid list", func(t *testing.T) {
		got := SolveDay2Part1([]string{"a1-3 a: abcde", "a1-3 b: cdefg", "a2-9 c: ccccccccc"})
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay2Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := listToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay2Part1(input)
	}
}

func TestSolveDay2Part2(t *testing.T) {
	t.Run("Test SolveDay2Part2 with valid list", func(t *testing.T) {
		got := SolveDay2Part2([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"})
		expected := 1

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test SolveDay2Part2 with invalid list", func(t *testing.T) {
		got := SolveDay2Part2([]string{"a1-3 a: abcde", "1a-3 b: cdefg", "2a-9 c: ccccccccc"})
		expected := 0

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay2Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := listToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay2Part2(input)
	}
}
func TestGetPasswordValues(t *testing.T) {
	t.Run("Test getPasswordValues with invalid string", func(t *testing.T) {
		s := `1-a3 a: abcde`
		gotvmin, gotvmax, gotchar, gotpassword, goterr := getPasswordValues(s)
		expectedvmin, expectedvmax, expectedchar, expectedpassword, expectederr := 0, 0, "", "", invalidMinMaxError

		if gotvmin != expectedvmin || gotvmax != expectedvmax || gotchar != expectedchar || gotpassword != expectedpassword || goterr != expectederr {
			t.Errorf("expected '%v' '%v' '%v' '%v' '%v' but got '%v' '%v' '%v' '%v' '%v'", expectedvmin, expectedvmax, expectedchar, expectedpassword, expectederr, gotvmin, gotvmax, gotchar, gotpassword, goterr)
		}
	})
	t.Run("Test getPasswordValues with valid string", func(t *testing.T) {
		s := `1-3 a: abcde`
		gotvmin, gotvmax, gotchar, gotpassword, goterr := getPasswordValues(s)
		var expectederr error
		expectedvmin, expectedvmax, expectedchar, expectedpassword, expectederr := 1, 3, "a", "abcde", nil

		if gotvmin != expectedvmin || gotvmax != expectedvmax || gotchar != expectedchar || gotpassword != expectedpassword || goterr != expectederr {
			t.Errorf("expected '%v' '%v' '%v' '%v' '%v' but got '%v' '%v' '%v' '%v' '%v'", expectedvmin, expectedvmax, expectedchar, expectedpassword, expectederr, gotvmin, gotvmax, gotchar, gotpassword, goterr)
		}
	})

}

func BenchmarkGetPasswordValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _, _, _, _ = getPasswordValues(`1-3 a: abcde`)
	}
}

func TestListToSlice(t *testing.T) {
	t.Run("Test ListToSlice with valid list", func(t *testing.T) {
		s := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`
		got := listToSlice(s)
		expected := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}

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
