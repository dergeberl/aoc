package utils

import "testing"

func TestSortString(t *testing.T) {
	t.Run("Test SortString with valid list", func(t *testing.T) {
		s := "azb"
		got := SortString(s)
		expected := "abz"
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
