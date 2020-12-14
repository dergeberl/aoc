package main

import (
	"io/ioutil"
	"testing"
)

func TestSolveDay14Part1(t *testing.T) {
	t.Run("Test SolveDay14Part1", func(t *testing.T) {
		i := stringListToSlice(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)
		got := SolveDay14Part1(i)
		expected := int64(165)

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func BenchmarkSolveDay14Part1(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay14Part1(input)
	}
}

func TestSolveDay14Part2(t *testing.T) {
	t.Run("Test SolveDay14Part2", func(t *testing.T) {
		i := stringListToSlice(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)
		got := SolveDay14Part2(i)
		expected := int64(208)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test SolveDay14Part2", func(t *testing.T) {
		i := stringListToSlice(`mask = 000000000000000000000000000000X1001X
mem[26] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)
		got := SolveDay14Part2(i)
		expected := int64(208)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

}

func BenchmarkSolveDay14Part2(b *testing.B) {
	i, _ := ioutil.ReadFile("input.txt")
	input := stringListToSlice(string(i))
	for n := 0; n < b.N; n++ {
		_ = SolveDay14Part2(input)
	}
}

func TestGetMaskValues(t *testing.T) {
	t.Run("Test getMaskValues", func(t *testing.T) {
		i := "mem[7] = 101"
		gotAddress, gotValue := getMaskValues(i)
		expectedAddress, expectedValue := int64(7), int64(101)
		if gotAddress != expectedAddress || gotValue != expectedValue {
			t.Errorf("expected '%d', '%d' but got '%d', '%d'", expectedAddress, expectedValue, gotAddress, gotValue)
		}
	})
	t.Run("Test getMaskValues", func(t *testing.T) {
		i := "me] = 101"
		gotAddress, gotValue := getMaskValues(i)
		expectedAddress, expectedValue := int64(0), int64(0)
		if gotAddress != expectedAddress || gotValue != expectedValue {
			t.Errorf("expected '%d', '%d' but got '%d', '%d'", expectedAddress, expectedValue, gotAddress, gotValue)
		}
	})
	t.Run("Test getMaskValues", func(t *testing.T) {
		i := "mem[7] = a"
		gotAddress, gotValue := getMaskValues(i)
		expectedAddress, expectedValue := int64(0), int64(0)
		if gotAddress != expectedAddress || gotValue != expectedValue {
			t.Errorf("expected '%d', '%d' but got '%d', '%d'", expectedAddress, expectedValue, gotAddress, gotValue)
		}
	})
	t.Run("Test getMaskValues", func(t *testing.T) {
		i := "mem[7] is 101"
		gotAddress, gotValue := getMaskValues(i)
		expectedAddress, expectedValue := int64(0), int64(0)
		if gotAddress != expectedAddress || gotValue != expectedValue {
			t.Errorf("expected '%d', '%d' but got '%d', '%d'", expectedAddress, expectedValue, gotAddress, gotValue)
		}
	})
}
func TestApplyMaskOnValue(t *testing.T) {
	t.Run("Test applyMaskOnValue", func(t *testing.T) {
		got := applyMaskOnValue("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 11)
		expected := int64(73)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test applyMaskOnValue", func(t *testing.T) {
		got := applyMaskOnValue("XXXXXXXXXXXxXXXXXXXXXXXXXXXXX1XXXX0X", 11)
		expected := int64(0)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test applyMaskOnValue", func(t *testing.T) {
		got := applyMaskOnValue("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXX0X", 11)
		expected := int64(0)
		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func TestApplyMaskOnAddress(t *testing.T) {
	t.Run("Test applyMaskOnAddress", func(t *testing.T) {
		got := applyMaskOnAddress("000000000000000000000000000000X1001X", 42)
		expected := []int64{26, 27, 58, 59}
		if !equalInt64(got, expected) {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test applyMaskOnAddress", func(t *testing.T) {
		got := applyMaskOnAddress("00000000000000000000000000000000X0XX", 26)
		expected := []int64{16, 17, 18, 19, 24, 25, 26, 27}
		if !equalInt64(got, expected) {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test applyMaskOnAddress", func(t *testing.T) {
		got := applyMaskOnAddress("000000000000000000000000000000X1001", 42)
		expected := []int64{}
		if !equalInt64(got, expected) {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
	t.Run("Test applyMaskOnAddress", func(t *testing.T) {
		got := applyMaskOnAddress("000000000000000S0000000000000000X0XX", 42)
		expected := []int64{}
		if !equalInt64(got, expected) {
			t.Errorf("expected '%d' but got '%d'", expected, got)
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
