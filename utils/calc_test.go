package utils

import "testing"

func TestSumSlice(t *testing.T) {
	t.Run("Test SumSlice with valid list", func(t *testing.T) {
		s := []int{1, 2, 3}
		got := SumSlice(s)
		expected := 6
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
