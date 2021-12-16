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

func TestProductSlice(t *testing.T) {
	t.Run("Test ProductSlice with valid list", func(t *testing.T) {
		s := []int{10, 10, 10}
		got := ProductSlice(s)
		expected := 10 * 10 * 10
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

	t.Run("Test ProductSlice with valid list", func(t *testing.T) {
		s := []int{}
		got := ProductSlice(s)
		expected := 1
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

	t.Run("Test ProductSlice with valid list", func(t *testing.T) {
		s := []int{80}
		got := ProductSlice(s)
		expected := 80
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func TestMaximumSliceSlice(t *testing.T) {
	t.Run("Test MaximumSlice with valid list", func(t *testing.T) {
		s := []int{10, 1, 5}
		got := MaximumSlice(s)
		expected := 10
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

	t.Run("Test MaximumSlice with valid list", func(t *testing.T) {
		s := []int{100, 123, 500, 2}
		got := MaximumSlice(s)
		expected := 500
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

	t.Run("Test MaximumSlice with valid list", func(t *testing.T) {
		s := []int{}
		got := MaximumSlice(s)
		expected := 0
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func TestMinimumSlice(t *testing.T) {
	t.Run("Test MaximumSlice with valid list", func(t *testing.T) {
		s := []int{10, 1, 5}
		got := MinimumSlice(s)
		expected := 1
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

	t.Run("Test MaximumSlice with valid list", func(t *testing.T) {
		s := []int{100, 123, 500, 66}
		got := MinimumSlice(s)
		expected := 66
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

	t.Run("Test MaximumSlice with valid list", func(t *testing.T) {
		s := []int{}
		got := MinimumSlice(s)
		expected := 0
		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
