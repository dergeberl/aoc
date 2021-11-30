package utils

import "testing"

func TestInputToSlice(t *testing.T) {
	t.Run("Test InputToSlice with valid list", func(t *testing.T) {
		s := `string
string 2
string:3`
		got, _ := InputToSlice(s)
		expected := []string{"string", "string 2", "string:3"}
		if !equalString(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func TestInputToIntSlice(t *testing.T) {
	t.Run("Test InputToIntSlice with valid list", func(t *testing.T) {
		s := `123
1234
000
00123`
		got, _ := InputToIntSlice(s)
		expected := []int{123, 1234, 0, 123}
		if !equalInt(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Test InputToIntSlice with invalid list", func(t *testing.T) {
		s := `123
1234a
000
00123`
		_, goterr := InputToIntSlice(s)
		if goterr == nil {
			t.Errorf("expected error but got none")
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